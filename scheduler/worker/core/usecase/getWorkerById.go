package usecase

import (
	"log"

	"github.com/renatospaka/scheduler/core/dto"
	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
	eventWorker "github.com/renatospaka/scheduler/scheduler/worker/core/event"
	pkgController "github.com/renatospaka/scheduler/adapter/web/controller"
)

type GetWorkerByIdInputDto struct {
	ID int `json:"worker_id"`
}

type GetWorkerByIdOutputDto struct {
	ID                          int    `json:"worker_id"`
	Name                        string `json:"name"`
	IsActive                    bool   `json:"is_active"`
	dto.StandardStatusOutputDto `json:"request"`
}

func (u *WorkerUsecase) getWorkerById(in GetWorkerByIdInputDto) (GetWorkerByIdOutputDto, error) {
	log.Println("scheduler.worker.core.usecase.GetWorker")

	if in.ID <= 0 {
		return GetWorkerByIdOutputDto{}, entity.ErrWorkerIdInvalid
	}
	worker, err := u.repo.GetWorker(in.ID)
	if err != nil {
		return GetWorkerByIdOutputDto{}, err
	}
	if worker == nil {
		return GetWorkerByIdOutputDto{}, entity.ErrWorkerIdNotFound
	}

	out := GetWorkerByIdOutputDto{
		ID:                      worker.ID(),
		Name:                    worker.Name(),
		IsActive:                worker.IsActive(),
		StandardStatusOutputDto: dto.StandardStatusOutputDto{
			Status: pkgController.REQUEST_SUCCESS,
		},
	}

	workerGotten := eventWorker.NewWorkerGottenEvent("getting worker ID #" + string(rune(in.ID)))
	u.dispatcher.Dispatch(workerGotten)

	return out, nil
}
