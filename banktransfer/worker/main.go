package main

import (
	"log"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/activities"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/configs"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/workflows"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	c, err := client.NewLazyClient(client.Options{
		HostPort: config.TEMPORAL_CLUSTER_HOST,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	w := worker.New(c, configs.Workflows.TRANSFER, worker.Options{})

	// Transfer workflow
	transferActivity := &activities.TransferActivity{}
	w.RegisterActivity(transferActivity.CreateTransfer)
	w.RegisterActivity(transferActivity.SendTransferNotification)
	w.RegisterWorkflow(workflows.TransferWorkflow)

	// Verify workflow
	// verifyActivity := &activities.VerifyActivity{}
	// w.RegisterActivity(verifyActivity.VerifyOtp)
	// w.RegisterWorkflow(workflows.VerifyOtpWorkflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
