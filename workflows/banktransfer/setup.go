package banktransfer

import (
	"github.com/anhgeeky/go-temporal-labs/banktransfer/activities"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/workflows"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

// Transfer workflow V1
func RegisterTransferWorkflowV1(w worker.Registry, transferActivity activities.TransferActivity) {
	w.RegisterWorkflowWithOptions(workflows.TransferWorkflowV1, workflow.RegisterOptions{Name: config.Workflows.TransferWorkflow})
	w.RegisterActivity(transferActivity.CheckBalance)
	w.RegisterActivity(transferActivity.CreateOTP)
	w.RegisterActivity(transferActivity.CreateTransaction)
}

// Transfer workflow V2
func RegisterTransferWorkflowV2(w worker.Registry, transferActivity activities.TransferActivity) {
	w.RegisterWorkflowWithOptions(workflows.TransferWorkflowV2, workflow.RegisterOptions{Name: config.Workflows.TransferWorkflow})
	w.RegisterActivity(transferActivity.CheckBalance)
	w.RegisterActivity(transferActivity.CreateOTP)
	w.RegisterActivity(transferActivity.CreateTransaction)
	w.RegisterActivity(transferActivity.NewActivityForV2)
}

// Transfer workflow V3
func RegisterTransferWorkflowV3(w worker.Registry, transferActivity activities.TransferActivity) {
	w.RegisterWorkflowWithOptions(workflows.TransferWorkflowV3, workflow.RegisterOptions{Name: config.Workflows.TransferWorkflow})
	w.RegisterActivity(transferActivity.CheckBalance)
	w.RegisterActivity(transferActivity.CreateOTP)
	w.RegisterActivity(transferActivity.CreateTransaction)
	w.RegisterActivity(transferActivity.NewActivityForV3)
}

// Transfer workflow V4
func RegisterTransferWorkflowV4(w worker.Registry, transferActivity activities.TransferActivity) {
	w.RegisterWorkflowWithOptions(workflows.TransferWorkflowV4, workflow.RegisterOptions{Name: config.Workflows.TransferWorkflow})
	w.RegisterActivity(transferActivity.CheckBalance)
	w.RegisterActivity(transferActivity.CreateOTP)
	w.RegisterActivity(transferActivity.CreateTransaction)
	w.RegisterActivity(transferActivity.NewActivityForV4)
}
