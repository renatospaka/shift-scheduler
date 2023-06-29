package usecase

import (
	"log"

	pkgController "github.com/renatospaka/scheduler/adapter/web/controller"
	"github.com/renatospaka/scheduler/core/dto"
	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
	// "github.com/renatospaka/scheduler/scheduler/worker/core/event"
)

type GetWorkerByIdInputDto struct {
	ID int `json:"worker_id"`
}

type GetWorkerByIdOutputDto struct {
	ID                          int    `json:"worker_id"`
	Name                        string `json:"name"`
	Profession                  string `json:"profession"`
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
		return GetWorkerByIdOutputDto{
			StandardStatusOutputDto: dto.StandardStatusOutputDto{ Status: pkgController.REQUEST_FAILURE,
			},
		}, err
	}
	if worker == nil {
		return GetWorkerByIdOutputDto{
			StandardStatusOutputDto: dto.StandardStatusOutputDto{ Status: pkgController.REQUEST_FAILURE,
			},
		}, entity.ErrWorkerIdNotFound
	}

	out := GetWorkerByIdOutputDto{
		ID:         worker.ID(),
		Name:       worker.Name(),
		Profession: worker.Profession(),
		IsActive:   worker.IsActive(),
		StandardStatusOutputDto: dto.StandardStatusOutputDto{
			Status: pkgController.REQUEST_SUCCESS,
		},
	}

	// No need for an event dispatcher at this moment
	// workerGotten := event.NewWorkerGottenEvent("getting worker ID #" + string(rune(in.ID)))
	// u.dispatcher.Dispatch(workerGotten)

	return out, nil
}
