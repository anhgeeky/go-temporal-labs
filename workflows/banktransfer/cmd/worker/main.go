package main

import (
	"log"
	"path/filepath"
	"runtime"
	"sync"

	tranFlow "github.com/anhgeeky/go-temporal-labs/banktransfer"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"github.com/anhgeeky/go-temporal-labs/core/broker/kafka"
	"github.com/anhgeeky/go-temporal-labs/core/configs"
	notiFlow "github.com/anhgeeky/go-temporal-labs/notification"
	"github.com/spf13/viper"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
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

	taskQueue := config.TaskQueues.TRANSFER_QUEUE

	wg := sync.WaitGroup{}
	// createAndRunWorker(c, taskQueue, config.VERSION_1_0, &wg, externalCfg, bk)
	// createAndRunWorker(c, taskQueue, config.VERSION_2_0, &wg, externalCfg, bk)
	createAndRunWorker(c, taskQueue, config.VERSION_3_0, &wg, externalCfg, bk)
	createAndRunWorker(c, taskQueue, config.VERSION_4_0, &wg, externalCfg, bk)
	wg.Wait()
}

func createAndRunWorker(c client.Client, taskQueue, buildID string, wg *sync.WaitGroup, externalCfg *config.ExternalConfig, bk broker.Broker) {
	log.Println("Start worker: ", taskQueue, "Build ID:", buildID)
	w := worker.New(c, taskQueue, worker.Options{
		// Both of these options must be set to opt into the feature
		BuildID:                 buildID,
		UseBuildIDForVersioning: true,
	})

	switch buildID {
	case config.VERSION_1_0:
		tranFlow.SetupBankTransferWorkflow(w, externalCfg, bk)
	case config.VERSION_2_0:
		tranFlow.SetupBankTransferWorkflowV2(w, externalCfg, bk)
	case config.VERSION_3_0:
		tranFlow.SetupBankTransferWorkflowV3(w, externalCfg, bk)
	case config.VERSION_4_0:
		tranFlow.SetupBankTransferWorkflowV4(w, externalCfg, bk)
	}
	notiFlow.SetupNotificationWorkflow(w)

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalf("Unable to start %s worker: %v", buildID, err)
		}
	}()
}
