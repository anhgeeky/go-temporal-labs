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

	taskQueue := config.TaskQueues.TRANSFER_QUEUE
	// workflowName := config.Workflows.TransferWorkflow

	// transferActivity := &activities.TransferActivity{
	// 	Broker: bk,
	// 	AccountService: account.AccountService{
	// 		Host: externalCfg.AccountHost,
	// 	},
	// 	MoneyTransferService: moneytransfer.MoneyTransferService{
	// 		Host: externalCfg.MoneyTransferHost,
	// 	},
	// }

	wg := sync.WaitGroup{}
	// ctx := context.Background()
	w, _ := wk.NewWorker(workers.TransferWorkerV1{
		Broker: bk,
		Config: *externalCfg,
	}, wk.PlatformConfig{},
		wk.WithClient(c),
		wk.WithTaskQueue(taskQueue),
		wk.WithBuildID(config.VERSION_1_0),
	)
	w.RunWithGroup(&wg)

	// createAndRunWorker(c, taskQueue, config.VERSION_1_0, &wg, externalCfg, bk)
	// createAndRunWorker(c, taskQueue, config.VERSION_2_0, &wg, externalCfg, bk)
	// createAndRunWorker(c, taskQueue, config.VERSION_3_0, &wg, externalCfg, bk)
	// createAndRunWorker(c, taskQueue, config.VERSION_4_0, &wg, externalCfg, bk)
	// wk.CreateNewWorker[messages.Transfer](
	// 	c, &wg, workflowName, taskQueue, config.VERSION_1_0,
	// 	// Register workflows
	// 	tranFlow.TransferWorkflow,
	// 	transferActivity.CheckBalance,
	// 	transferActivity.CreateOTP,
	// 	transferActivity.CreateTransaction,
	// )
	wg.Wait()
}

// func createAndRunWorker(c client.Client, taskQueue, buildID string, wg *sync.WaitGroup, externalCfg *config.ExternalConfig, bk broker.Broker) {
// 	log.Println("Start worker: ", taskQueue, "Build ID:", buildID)
// 	w := worker.New(c, taskQueue, worker.Options{
// 		// Both of these options must be set to opt into the feature
// 		BuildID:                 buildID,
// 		UseBuildIDForVersioning: true,
// 	})

// 	// TODO: Implement Worker for build ID

// 	// tranPkg.SetupBankTransferWorkflow(w, tranFlow.TransferWorkflow, externalCfg, bk)
// 	// tranPkg.SetupBankTransferWorkflowV2(w, tranFlow.TransferWorkflowV2, externalCfg, bk)
// 	// tranPkg.SetupBankTransferWorkflowV3(w, tranFlow.TransferWorkflowV3, externalCfg, bk)
// 	// tranPkg.SetupBankTransferWorkflowV4(w, tranFlow.TransferWorkflowV4, externalCfg, bk)
// 	// notiPkg.SetupNotificationWorkflow(w, notiWorkflow.NotificationWorkflow)

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		err := w.Run(worker.InterruptCh())
// 		if err != nil {
// 			log.Fatalf("Unable to start %s worker: %v", buildID, err)
// 		}
// 	}()
// }
