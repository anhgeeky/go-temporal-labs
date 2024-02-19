package main

import (
	"log"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/activities"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/utils"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/workflows"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	c, err := client.NewLazyClient(client.Options{
		HostPort: config.TEMPORAL_HOST,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	w := worker.New(c, utils.Workflows.BANK_TRANSFER, worker.Options{})

	a := &activities.TransferActivity{}

	w.RegisterActivity(a.CreateTransfer)
	w.RegisterActivity(a.SendTransferNotification)
	w.RegisterWorkflow(workflows.TransferWorkflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
