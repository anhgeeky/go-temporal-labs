package main

import (
	"log"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/activities"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/configs"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/workflows"
	notiActivities "github.com/anhgeeky/go-temporal-labs/notification/activities"
	notiWorkflows "github.com/anhgeeky/go-temporal-labs/notification/workflows"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	c, err := client.NewLazyClient(client.Options{
		HostPort: configs.TEMPORAL_CLUSTER_HOST,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	w := worker.New(c, configs.TaskQueues.BANK_TRANSFER_QUEUE, worker.Options{})

	// Transfer workflow
	transferActivity := &activities.TransferActivity{}
	w.RegisterActivity(transferActivity.CreateTransfer)
	w.RegisterActivity(transferActivity.CheckBalance)
	w.RegisterActivity(transferActivity.CheckTargetAccount)
	w.RegisterActivity(transferActivity.CreateTransferTransaction)
	w.RegisterActivity(transferActivity.WriteCreditAccount)
	w.RegisterActivity(transferActivity.WriteDebitAccount)
	w.RegisterWorkflow(workflows.TransferWorkflow)

	// Notification workflow
	notificationActivity := &notiActivities.NotificationActivity{}
	w.RegisterActivity(notificationActivity.GetDeviceToken)
	w.RegisterActivity(notificationActivity.PushSMS)
	w.RegisterActivity(notificationActivity.PushNotification)
	w.RegisterActivity(notificationActivity.PushInternalApp)
	w.RegisterWorkflow(notiWorkflows.NotificationWorkflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
