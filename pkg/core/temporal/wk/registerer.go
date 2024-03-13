package wk

import (
	"go.temporal.io/sdk/worker"
)

// Registerer is the entity that registers workflows and activities.
type Registerer interface {
	Register(worker.Registry)
}
