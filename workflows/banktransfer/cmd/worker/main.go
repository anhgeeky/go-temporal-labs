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
	wk.RunAsNewWorkerVersioning(
		c, &wg, workerName, taskQueue, config.VERSION_1_0,
		workers.TransferWorkerV1{Broker: bk, Config: *externalCfg},
	)
	// // ======================= WORKER 2 =======================
	wk.RunAsNewWorkerVersioning(
		c, &wg, workerName, taskQueue, config.VERSION_2_0,
		workers.TransferWorkerV2{Broker: bk, Config: *externalCfg},
	)
	// ======================= WORKER 3 =======================
	wk.RunAsNewWorkerVersioning(
		c, &wg, workerName, taskQueue, config.VERSION_3_0,
		workers.TransferWorkerV3{Broker: bk, Config: *externalCfg},
	)
	// ======================= WORKER 4 =======================
	wk.RunAsNewWorkerVersioning(
		c, &wg, workerName, taskQueue, config.VERSION_4_0,
		workers.TransferWorkerV4{Broker: bk, Config: *externalCfg},
	)
	// Auto update latest worker BuildIds
	wk.UpdateLatestWorkerBuildIDs(c, &wg, config.TaskQueues.TRANSFER_QUEUE, config.VERSION_1_0, config.VERSION_4_0)
	wg.Wait()
}
