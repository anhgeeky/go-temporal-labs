package workers

import "go.temporal.io/sdk/worker"

type TransferWorkerV1 struct {
}

func (r TransferWorkerV1) Register(register worker.Registry) {
	// register.RegisterWorkflow(sample.Workflow)
	// register.RegisterActivity(sample.Activity)
}

func NewWorker() {

}
