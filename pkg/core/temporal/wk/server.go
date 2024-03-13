package wk

import (
	"context"
	"errors"
	"log"
	"sync"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

var (
	ErrClientRequired = errors.New("client required: use WithClient to set the client")
)

// Registerer is the entity that registers workflows and activities.
type Registerer interface {
	Register(worker.Registry)
}

// Worker is a services.Server that is able to initialize and manage the temporal Worker together with the
type Worker struct {
	name    string
	buildID string
	w       worker.Worker
}

// NewWorker implements
func NewWorker(registerer Registerer, options ...Option) (Worker, error) {
	opts := defaultOpts()
	for _, opt := range options {
		opt(&opts)
	}
	if opts.client == nil {
		return Worker{}, ErrClientRequired
	}
	w := worker.New(opts.client, opts.taskQueue, worker.Options{
		BackgroundActivityContext: opts.backgroundAcitivityContext,
		Interceptors:              opts.interceptors,
		OnFatalError:              opts.onFatalError,
		BuildID:                   opts.buildID,
		UseBuildIDForVersioning:   opts.useBuildIDForVersioning,
	})
	registerer.Register(w)
	return Worker{
		name:    opts.name,
		buildID: opts.buildID,
		w:       w,
	}, nil
}

// RunAsNewWorkerVersioning implements
func RunAsNewWorkerVersioning(c client.Client, wg *sync.WaitGroup, name, taskQueue, buildID string, registerer Registerer, options ...Option) (Worker, error) {
	opts := defaultOpts()
	for _, opt := range options {
		opt(&opts)
	}
	if c == nil {
		return Worker{}, ErrClientRequired
	}
	w := worker.New(c, taskQueue, worker.Options{
		BackgroundActivityContext: opts.backgroundAcitivityContext,
		Interceptors:              opts.interceptors,
		OnFatalError:              opts.onFatalError,
		BuildID:                   buildID,
		UseBuildIDForVersioning:   true,
	})
	registerer.Register(w)
	newWorker := Worker{
		name:    name,
		buildID: buildID,
		w:       w,
	}

	newWorker.RunWithGroup(wg)

	return newWorker, nil
}

func (w *Worker) Name() string {
	return w.name
}

func (w *Worker) Listen(_ context.Context) error {
	return w.w.Start()
}

func (w *Worker) Run(_ context.Context) error {
	return w.w.Run(worker.InterruptCh())
}

func (w *Worker) Close(_ context.Context) error {
	w.w.Stop()
	return nil
}

func (w *Worker) RunWithGroup(wg *sync.WaitGroup) error {
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("%s worker buildID started: %v", w.name, w.buildID)
		err := w.w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalf("Unable to start %s worker, buildID %s: %v", w.name, w.buildID, err)
		}
	}()

	return nil
}
