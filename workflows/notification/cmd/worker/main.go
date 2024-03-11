package main

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/anhgeeky/go-temporal-labs/core/configs"
	notiPkg "github.com/anhgeeky/go-temporal-labs/notification"
	"github.com/anhgeeky/go-temporal-labs/notification/config"
	notiWorkflow "github.com/anhgeeky/go-temporal-labs/notification/workflows"
	"github.com/spf13/viper"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
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

	c, err := client.NewLazyClient(client.Options{
		HostPort:  temporalCfg.TemporalHost,
		Namespace: temporalCfg.TemporalNamespace,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	w := worker.New(c, config.TaskQueues.NOTIFICATION_QUEUE, worker.Options{})

	notiPkg.SetupNotificationWorkflow(w, notiWorkflow.NotificationWorkflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
