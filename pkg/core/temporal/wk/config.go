package wk

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/worker"
)

type WorkflowPanicPolicy worker.WorkflowPanicPolicy

var (
	workflowPanicPolicyBlock = []byte("block")
	workflowPanicPolicyFail  = []byte("fail")

	ErrUnknownWorkflowPanicPolicy = errors.New("unknown workflow panic policy")
)

func (w *WorkflowPanicPolicy) UnmarshalText(text []byte) error {
	switch {
	case bytes.Equal(text, workflowPanicPolicyBlock):
		*w = WorkflowPanicPolicy(worker.BlockWorkflow)
	case bytes.Equal(text, workflowPanicPolicyFail):
		*w = WorkflowPanicPolicy(worker.FailWorkflow)
	default:
		return fmt.Errorf("%w: %s", ErrUnknownWorkflowPanicPolicy, text)
	}
	return nil
}

type config struct {
	name                       string
	taskQueue                  string
	buildID                    string
	useBuildIDForVersioning    bool
	backgroundAcitivityContext context.Context
	interceptors               []interceptor.WorkerInterceptor
	onFatalError               func(error)
	client                     client.Client
}

func defaultOpts() config {
	return config{
		name: "Temporal Worker Server",
	}
}

type Option func(*config)

func WithName(name string) Option {
	return func(c *config) {
		c.name = name
	}
}

func WithTaskQueue(taskQueue string) Option {
	return func(c *config) {
		c.taskQueue = taskQueue
	}
}

func WithBuildID(buildID string) Option {
	return func(c *config) {
		c.buildID = buildID
		c.useBuildIDForVersioning = true
	}
}

func WithBackgroundActivityContext(ctx context.Context) Option {
	return func(c *config) {
		c.backgroundAcitivityContext = ctx
	}
}

func WithInterceptors(interceptors ...interceptor.WorkerInterceptor) Option {
	return func(c *config) {
		c.interceptors = interceptors
	}
}

func WithOnFatalError(onFatalError func(error)) Option {
	return func(c *config) {
		c.onFatalError = onFatalError
	}
}

func WithClient(client client.Client) Option {
	return func(c *config) {
		c.client = client
	}
}
