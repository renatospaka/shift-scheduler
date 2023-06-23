package usecase

import (
	"log"

	"github.com/renatospaka/scheduler/adapter/event"
	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
	eventWorker "github.com/renatospaka/scheduler/scheduler/worker/core/event"
	"github.com/renatospaka/scheduler/scheduler/worker/core/repository"
)

type WorkerUsecase struct {
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

func (u *WorkerUsecase) GetWorker(id int) (*entity.Worker, error) {
	log.Printf("try to get worker #%d\n", id)

	return u.getWorker(id)
}
