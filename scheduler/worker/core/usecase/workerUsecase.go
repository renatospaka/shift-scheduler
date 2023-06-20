package usecase

import (
	"log"

	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
	"github.com/renatospaka/scheduler/scheduler/worker/core/repository"
)

type WorkerUsecase struct {
	repo 	repository.WorkerInterface
}

func NewWorkerUsecase(repo repository.WorkerInterface) *WorkerUsecase {
	log.Println("initiating worker usecases")
	
	return &WorkerUsecase{
		repo: repo,
	}
}

func (u *WorkerUsecase) GetWorker(id int) (*entity.Worker, error) {
	log.Println("scheduler.worker.core.usecase.GetWorker")
	return nil, nil
}