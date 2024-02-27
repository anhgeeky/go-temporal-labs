package main

import (
	"log"
	"path/filepath"
	"runtime"

	tranFlow "github.com/anhgeeky/go-temporal-labs/banktransfer"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/core/configs"
	notiFlow "github.com/anhgeeky/go-temporal-labs/notification"
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
	w := worker.New(c, config.TaskQueues.BANK_TRANSFER_QUEUE, worker.Options{})

	tranFlow.SetupBankTransferWorkflow(w, cfg)
	notiFlow.SetupNotificationWorkflow(w, cfg.NotificationHost)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
