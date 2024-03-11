package main

import (
	"log"
	"path/filepath"
	"runtime"
	"sync"

	tranPkg "github.com/anhgeeky/go-temporal-labs/banktransfer"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	tranFlow "github.com/anhgeeky/go-temporal-labs/banktransfer/workflows"
	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"github.com/anhgeeky/go-temporal-labs/core/broker/kafka"
	"github.com/anhgeeky/go-temporal-labs/core/configs"
	notiPkg "github.com/anhgeeky/go-temporal-labs/notification"
	notiWorkflow "github.com/anhgeeky/go-temporal-labs/notification/workflows"
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
	createAndRunWorker(c, taskQueue, config.VERSION_1_0, &wg, externalCfg, bk)
	createAndRunWorker(c, taskQueue, config.VERSION_2_0, &wg, externalCfg, bk)
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

	// TODO: Implement Worker for build ID

	tranPkg.SetupBankTransferWorkflow(w, tranFlow.TransferWorkflow, externalCfg, bk)
	tranPkg.SetupBankTransferWorkflowV2(w, tranFlow.TransferWorkflowV2, externalCfg, bk)
	tranPkg.SetupBankTransferWorkflowV3(w, tranFlow.TransferWorkflowV3, externalCfg, bk)
	tranPkg.SetupBankTransferWorkflowV4(w, tranFlow.TransferWorkflowV4, externalCfg, bk)
	notiPkg.SetupNotificationWorkflow(w, notiWorkflow.NotificationWorkflow)

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalf("Unable to start %s worker: %v", buildID, err)
		}
	}()
}
