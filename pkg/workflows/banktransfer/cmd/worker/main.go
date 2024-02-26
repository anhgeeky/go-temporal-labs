package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	tranFlow "github.com/anhgeeky/go-temporal-labs/banktransfer"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/core/configs"
	notiFlow "github.com/anhgeeky/go-temporal-labs/notification"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	filePath := filepath.Join(filepath.Dir(b), "../..", ".env")
	fmt.Println("File Path", filePath)
	configs.LoadConfig(filePath)

	log.Println("TEMPORAL_CLUSTER_HOST", config.TEMPORAL_CLUSTER_HOST)

	c, err := client.NewLazyClient(client.Options{
		HostPort: config.TEMPORAL_CLUSTER_HOST,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	w := worker.New(c, config.TaskQueues.BANK_TRANSFER_QUEUE, worker.Options{})

	tranFlow.SetupBankTransferWorkflow(w)
	notiFlow.SetupNotificationWorkflow(w)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
