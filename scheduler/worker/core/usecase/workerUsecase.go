package usecase

import (
	"context"
	"log"

	eventAdapter "github.com/renatospaka/scheduler/adapter/event"
	"github.com/renatospaka/scheduler/scheduler/worker/core/event"
	"github.com/renatospaka/scheduler/scheduler/worker/core/repository"
)

type WorkerUsecase struct {
	ctx        context.Context
	repo       repository.WorkerInterface
	dispatcher *eventAdapter.EventDispatcher
}

// Start an instance of the usecases of the worker domain
func NewWorkerUsecase(repo repository.WorkerInterface, dispatcher *eventAdapter.EventDispatcher) *WorkerUsecase {
	log.Println("initiating worker usecases")

	usecases := &WorkerUsecase{
		repo:       repo,
		dispatcher: dispatcher,
	}
	usecases.dispatcher.AddListener(event.EVENT_WORKER_GOTTEN, event.NewWorkerGottenListener())

	return usecases
}

// Factory to instantiate GetWorkerById usecase
func (u *WorkerUsecase) GetWorkerById(in GetWorkerByIdInputDto) (GetWorkerByIdOutputDto, error) {
	return u.getWorkerById(in)
}

// Factory to instantiate GetWorkerWithDocumentsById usecase
func (u *WorkerUsecase) GetWorkerWithDocumentsById(in GetWorkerWithDocumentsByIdInputDto) (GetWorkerWithDocumentsByIdOutputDto, error) {
	return u.getWorkerWithDocumentsById(in)
}

// Set the internal context of the usecase, overriding if any
func (u *WorkerUsecase) SetContext(ctx context.Context) {
	u.ctx = ctx
	u.repo.SetContext(ctx)
}

// Retrieve the internal context of the usecase
func (u *WorkerUsecase) Context() context.Context {
	return u.ctx
}
