package main

import (
	"log"

	noti "github.com/anhgeeky/go-temporal-labs/notification"
	"github.com/anhgeeky/go-temporal-labs/notification/config"

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
	w := worker.New(c, config.TaskQueues.NOTIFICATION_QUEUE, worker.Options{})

	noti.SetupNotificationWorkflow(w)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
