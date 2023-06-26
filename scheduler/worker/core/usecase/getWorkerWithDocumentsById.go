package usecase

import (
	"log"

	pkgController "github.com/renatospaka/scheduler/adapter/web/controller"
	"github.com/renatospaka/scheduler/core/dto"
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
	ID       int                        `json:"worker_document_id"`
	Document GetWorkerDocumentOutputDto `json:"document"`
}

type GetWorkerDocumentOutputDto struct {
	ID       int    `json:"document_id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
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

	out := GetWorkerWithDocumentsByIdOutputDto{
		ID:        worker.ID(),
		Name:      worker.Name(),
		IsActive:  worker.IsActive(),
		Documents: []GetWorkerDocumentsOutputDto{},
		StandardStatusOutputDto: dto.StandardStatusOutputDto{
			Status: pkgController.REQUEST_SUCCESS,
		},
	}

	for _, docto := range worker.Documents() {
		inner := docto.Document()
		document := GetWorkerDocumentsOutputDto{
			ID: docto.ID(),
			Document: GetWorkerDocumentOutputDto{
				ID:       inner.ID(),
				Name:     inner.Name(),
				IsActive: inner.IsActive(),
			},
		}

		out.Documents = append(out.Documents, document)
	}

	return out, nil
}
