package main

import (
	"context"
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

var (
	VERSION_1_0 = "1.0"
	VERSION_2_0 = "2.0"
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
	createAndRunWorker(c, taskQueue, VERSION_1_0, &wg, externalCfg, bk)
	createAndRunWorker(c, taskQueue, VERSION_2_0, &wg, externalCfg, bk)
	wg.Wait()

	err = updateLatestWorkerBuildId(c, taskQueue, VERSION_1_0, VERSION_2_0)
	if err != nil {
		log.Fatalln("Update latest worker build failure", err)
	}
}

// FAQ: https://docs.temporal.io/dev-guide/go/versioning
func updateLatestWorkerBuildId(c client.Client, taskQueue, compatibleBuildID, latestBuildID string) error {
	ctx := context.Background()
	// Now, let's update the task queue with a new compatible version:
	err := c.UpdateWorkerBuildIdCompatibility(ctx, &client.UpdateWorkerBuildIdCompatibilityOptions{
		TaskQueue: taskQueue,
		Operation: &client.BuildIDOpAddNewCompatibleVersion{
			BuildID:                   compatibleBuildID,
			ExistingCompatibleBuildID: latestBuildID,
		},
	})
	if err != nil {
		log.Fatalln("Unable to update build id compatability", err)
	}
	return err
}

func createAndRunWorker(c client.Client, taskQueue, buildID string, wg *sync.WaitGroup, externalCfg *config.ExternalConfig, bk broker.Broker) {
	w := worker.New(c, taskQueue, worker.Options{
		// Both of these options must be set to opt into the feature
		BuildID:                 buildID,
		UseBuildIDForVersioning: true,
	})

	switch buildID {
	case VERSION_1_0:
		tranFlow.SetupBankTransferWorkflow(w, externalCfg, bk)
	case VERSION_2_0:
		tranFlow.SetupBankTransferWorkflowV2(w, externalCfg, bk)
	}
	notiFlow.SetupNotificationWorkflow(w, externalCfg.NotificationHost)

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalf("Unable to start %s worker: %v", buildID, err)
		}
	}()
}
