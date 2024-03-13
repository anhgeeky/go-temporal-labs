package banktransfer

import (
	"github.com/anhgeeky/go-temporal-labs/banktransfer/activities"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/workflows"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

// Transfer workflow V1
func TransferWorkflowRegisterV1(w worker.Registry, transferActivity activities.TransferActivity) {
	w.RegisterWorkflowWithOptions(workflows.TransferWorkflow, workflow.RegisterOptions{Name: config.Workflows.TransferWorkflow})
	w.RegisterActivity(transferActivity.CheckBalance)
	w.RegisterActivity(transferActivity.CreateOTP)
	w.RegisterActivity(transferActivity.CreateTransaction)
}

// Transfer workflow V2
func TransferWorkflowRegisterV2(w worker.Registry, transferActivity activities.TransferActivity) {
	w.RegisterWorkflowWithOptions(workflows.TransferWorkflowV2, workflow.RegisterOptions{Name: config.Workflows.TransferWorkflow})
	w.RegisterActivity(transferActivity.CheckBalance)
	w.RegisterActivity(transferActivity.CreateOTP)
	w.RegisterActivity(transferActivity.CreateTransaction)
	w.RegisterActivity(transferActivity.NewActivityForV2)
}

// Transfer workflow V3
func TransferWorkflowRegisterV3(w worker.Registry, transferActivity activities.TransferActivity) {
	w.RegisterWorkflowWithOptions(workflows.TransferWorkflowV3, workflow.RegisterOptions{Name: config.Workflows.TransferWorkflow})
	w.RegisterActivity(transferActivity.CheckBalance)
	w.RegisterActivity(transferActivity.CreateOTP)
	w.RegisterActivity(transferActivity.CreateTransaction)
	w.RegisterActivity(transferActivity.NewActivityForV3)
}

// Transfer workflow V4
func TransferWorkflowRegisterV4(w worker.Registry, transferActivity activities.TransferActivity) {
	w.RegisterWorkflowWithOptions(workflows.TransferWorkflowV4, workflow.RegisterOptions{Name: config.Workflows.TransferWorkflow})
	w.RegisterActivity(transferActivity.CheckBalance)
	w.RegisterActivity(transferActivity.CreateOTP)
	w.RegisterActivity(transferActivity.CreateTransaction)
	w.RegisterActivity(transferActivity.NewActivityForV4)
}
