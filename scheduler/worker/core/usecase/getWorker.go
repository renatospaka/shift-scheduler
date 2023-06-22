package usecase

import (
	"log"

	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
)

func (u *WorkerUsecase) getWorker(id int) (*entity.Worker, error) {
	log.Println("scheduler.worker.core.usecase.GetWorker")
	return nil, nil
}
