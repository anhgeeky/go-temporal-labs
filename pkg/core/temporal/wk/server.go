package wk

import (
	"context"
	"errors"
	"log"
	"sync"

	"go.temporal.io/sdk/worker"
)

var (
	ErrClientRequired = errors.New("client required: use WithClient to set the client")
)

// Worker is a services.Server that is able to initialize and manage the temporal Worker together with the
type Worker struct {
	name string
	w    worker.Worker
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
		name: opts.name,
		w:    w,
	}, nil
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
		err := w.w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalf("Unable to start %s worker: %v", w.name, err)
		}
	}()

	return nil
}
