package main

import (
	"log"

	app "github.com/anhgeeky/go-temporal-labs/bank-transfer"
	"github.com/anhgeeky/go-temporal-labs/bank-transfer/config"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	c, err := client.NewLazyClient(client.Options{
		HostPort: config.TemporalHost,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	w := worker.New(c, app.Workflows.BANK_TRANSFER, worker.Options{})

	a := &app.Activities{}

	w.RegisterActivity(a.CreateTransfer)
	w.RegisterActivity(a.SendTransferNotification)

	w.RegisterWorkflow(app.TransferWorkflow)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
