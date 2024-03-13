package main

import (
	"log"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/workers"
	"github.com/anhgeeky/go-temporal-labs/core/broker/kafka"
	"github.com/anhgeeky/go-temporal-labs/core/configs"
	"github.com/anhgeeky/go-temporal-labs/core/temporal/wk"
	"github.com/spf13/viper"

	"go.temporal.io/sdk/client"
)

func main() {
	// ======================= CONFIG =======================
	_, b, _, _ := runtime.Caller(0)
	filePath := filepath.Join(filepath.Dir(b), "../..", ".env")
	configs.LoadConfig(filePath)

	externalCfg := &config.ExternalConfig{}
	err := viper.Unmarshal(externalCfg)
	if err != nil {
		log.Fatalln("Could not load `ExternalConfig` configuration", err)
	}

	temporalCfg := &config.TemporalConfig{}
	err = viper.Unmarshal(temporalCfg)
	if err != nil {
		log.Fatalln("Could not load `TemporalConfig` configuration", err)
	}
	// ======================= CONFIG =======================

	// ======================= TEMPORAL =======================
	c, err := client.NewLazyClient(client.Options{
		HostPort:  temporalCfg.TemporalHost,
		Namespace: temporalCfg.TemporalNamespace,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	// ======================= TEMPORAL =======================

	// ======================= BROKER =======================
	kafkaCfg := &config.KafkaConfig{}
	err = viper.Unmarshal(kafkaCfg)
	if err != nil {
		log.Fatalln("Could not load `KafkaConfig` configuration", err)
	}
	bk := kafka.ConnectBrokerKafka(kafkaCfg.Brokers)
	// ======================= BROKER =======================

	workerName := "Versioning Worker"

	taskQueue := config.TaskQueues.TRANSFER_QUEUE
	wg := sync.WaitGroup{}
	// // ======================= WORKER 1 =======================
	// w1, _ := wk.NewWorker(workers.TransferWorkerV1{Broker: bk, Config: *externalCfg},
	// 	wk.WithName(workerName),
	// 	wk.WithClient(c),
	// 	wk.WithTaskQueue(taskQueue),
	// 	wk.WithBuildID(config.VERSION_1_0),
	// )
	// w1.RunWithGroup(&wg)
	// // ======================= WORKER 2 =======================
	// w2, _ := wk.NewWorker(workers.TransferWorkerV2{Broker: bk, Config: *externalCfg},
	// 	wk.WithName(workerName),
	// 	wk.WithClient(c),
	// 	wk.WithTaskQueue(taskQueue),
	// 	wk.WithBuildID(config.VERSION_2_0),
	// )
	// w2.RunWithGroup(&wg)
	// ======================= WORKER 3 =======================
	w3, _ := wk.NewWorker(workers.TransferWorkerV3{Broker: bk, Config: *externalCfg},
		wk.WithName(workerName),
		wk.WithClient(c),
		wk.WithTaskQueue(taskQueue),
		wk.WithBuildID(config.VERSION_3_0),
	)
	w3.RunWithGroup(&wg)
	// ======================= WORKER 4 =======================
	w4, _ := wk.NewWorker(workers.TransferWorkerV4{Broker: bk, Config: *externalCfg},
		wk.WithName(workerName),
		wk.WithClient(c),
		wk.WithTaskQueue(taskQueue),
		wk.WithBuildID(config.VERSION_4_0),
	)
	w4.RunWithGroup(&wg)
	// Auto update latest worker BuildIds
	wk.UpdateLatestWorkerBuildIDs(c, &wg, config.TaskQueues.TRANSFER_QUEUE, config.VERSION_1_0, config.VERSION_4_0)
	wg.Wait()
}
