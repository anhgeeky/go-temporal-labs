package temporal

import (
	"log"
	"sync"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

// Create a new worker
// func CreateNewWorker(c client.Client, wg *sync.WaitGroup, workflowFunc func(ctx workflow.Context) error, workflowName, taskQueue, buildID string, acts ...interface{}) {
// 	log.Println("Start worker: ", taskQueue, "Build ID:", buildID)
// 	w := worker.New(c, taskQueue, worker.Options{
// 		BuildID:                 buildID,
// 		UseBuildIDForVersioning: true,
// 	})

// 	w.RegisterWorkflowWithOptions(workflowFunc, workflow.RegisterOptions{Name: workflowName})
// 	for act := range acts {
// 		w.RegisterActivity(act)
// 	}

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		err := w.Run(worker.InterruptCh())
// 		if err != nil {
// 			log.Fatalf("Unable to start %s worker: %v", buildID, err)
// 		}
// 	}()
// }

func CreateNewWorker[T any](c client.Client, wg *sync.WaitGroup, workflowName, taskQueue, buildID string, workflowFunc func(ctx workflow.Context, params ...T) error, acts ...interface{}) {
	log.Println("Start worker: ", taskQueue, "Build ID:", buildID)
	w := worker.New(c, taskQueue, worker.Options{
		BuildID:                 buildID,
		UseBuildIDForVersioning: true,
	})

	w.RegisterWorkflowWithOptions(workflowFunc, workflow.RegisterOptions{Name: workflowName})
	for act := range acts {
		w.RegisterActivity(act)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalf("Unable to start %s worker: %v", buildID, err)
		}
	}()
}
