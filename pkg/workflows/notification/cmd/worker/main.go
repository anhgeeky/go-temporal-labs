package main

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/anhgeeky/go-temporal-labs/core/configs"
	notiFlow "github.com/anhgeeky/go-temporal-labs/notification"
	"github.com/anhgeeky/go-temporal-labs/notification/config"
	"github.com/spf13/viper"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	filePath := filepath.Join(filepath.Dir(b), "../..", ".env")
	configs.LoadConfig(filePath)

	cfg := &config.ExternalConfigs{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		log.Fatalln("Could not load configuration", err)
	}

	log.Println("TemporalClusterHost", cfg.TemporalClusterHost)
	log.Println("TemporalClusterNamespace", cfg.TemporalClusterNamespace)

	c, err := client.NewLazyClient(client.Options{
		HostPort:  cfg.TemporalClusterHost,
		Namespace: cfg.TemporalClusterNamespace,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	w := worker.New(c, config.TaskQueues.NOTIFICATION_QUEUE, worker.Options{})

	notiFlow.SetupNotificationWorkflow(w, cfg.NotificationHost)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
