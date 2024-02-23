package workflow

import (
	"time"

	"go.temporal.io/sdk/temporal"
)

var WorkflowConfigs = struct {
	RetryPolicy *temporal.RetryPolicy
}{
	RetryPolicy: &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    3,
	},
}
