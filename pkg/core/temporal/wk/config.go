package wk

import (
	"context"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/interceptor"
)

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
