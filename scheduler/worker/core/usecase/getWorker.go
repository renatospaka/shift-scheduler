package usecase

import (
	"log"

	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
	eventWorker "github.com/renatospaka/scheduler/scheduler/worker/core/event"
)

func (u *WorkerUsecase) getWorker(id int) (*entity.Worker, error) {
	log.Println("scheduler.worker.core.usecase.GetWorker")

	workerGotten := eventWorker.NewWorkerGottenEvent("getting worker ID #" + string(id))
	u.dispatcher.Dispatch(workerGotten)

	return nil, nil
}
