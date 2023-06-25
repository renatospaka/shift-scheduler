package usecase

import (
	"log"

	"github.com/renatospaka/scheduler/core/dto"
	doctoUsecase "github.com/renatospaka/scheduler/scheduler/document/core/usecase"
	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
)

type GetWorkerWithDocumentsByIdInputDto struct {
	ID int `json:"worker_id"`
}

type GetWorkerWithDocumentsByIdOutputDto struct {
	ID                          int                           `json:"worker_id"`
	Name                        string                        `json:"name"`
	IsActive                    bool                          `json:"is_active"`
	Documents                   []GetWorkerDocumentsOutputDto `json:"documents"`
	dto.StandardStatusOutputDto `json:"request"`
}

type GetWorkerDocumentsOutputDto struct {
	ID       int                                   `json:"worker_document_id"`
	Document doctoUsecase.GetDocumentByIdOutputDto `json:"document"`
}

func (u *WorkerUsecase) getWorkerWithDocumentsById(in GetWorkerWithDocumentsByIdInputDto) (GetWorkerWithDocumentsByIdOutputDto, error) {
	log.Println("scheduler.worker.core.usecase.GetWorkerWithDocumentsById")

	if in.ID <= 0 {
		return GetWorkerWithDocumentsByIdOutputDto{}, entity.ErrWorkerIdInvalid
	}
	worker, err := u.repo.GetWorkerWithDocuments(in.ID)
	if err != nil {
		return GetWorkerWithDocumentsByIdOutputDto{}, err
	}
	if worker == nil {
		return GetWorkerWithDocumentsByIdOutputDto{}, entity.ErrWorkerIdNotFound
	}

	return GetWorkerWithDocumentsByIdOutputDto{}, nil
}
