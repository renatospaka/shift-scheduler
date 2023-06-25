package usecase

import (
	"context"
	"log"

	"github.com/renatospaka/scheduler/adapter/event"
	eventWorker "github.com/renatospaka/scheduler/scheduler/worker/core/event"
	"github.com/renatospaka/scheduler/scheduler/worker/core/repository"
)

type WorkerUsecase struct {
	ctx        context.Context
	repo       repository.WorkerInterface
	dispatcher *event.EventDispatcher
}

func NewWorkerUsecase(repo repository.WorkerInterface, dispatcher *event.EventDispatcher) *WorkerUsecase {
	log.Println("initiating worker usecases")

	usecases := &WorkerUsecase{
		repo:       repo,
		dispatcher: dispatcher,
	}
	usecases.dispatcher.AddListener(eventWorker.EVENT_WORKER_GOTTEN, eventWorker.NewWorkerGottenListener())

	return usecases
}

func (u *WorkerUsecase) GetWorkerById(in GetWorkerByIdInputDto) (GetWorkerByIdOutputDto, error) {
	return u.getWorkerById(in)
}

func (u *WorkerUsecase) GetWorkerWithDocumentsById(in GetWorkerWithDocumentsByIdInputDto) (GetWorkerWithDocumentsByIdOutputDto, error) {
	return u.getWorkerWithDocumentsById(in)
}

func (u *WorkerUsecase) SetContext(ctx context.Context) {
	u.ctx = ctx
	u.repo.SetContext(ctx)
}

func (u *WorkerUsecase) Context() context.Context {
	return u.ctx
}
