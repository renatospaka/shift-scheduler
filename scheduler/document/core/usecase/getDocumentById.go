package usecase

import (
	"log"

	"github.com/renatospaka/scheduler/core/dto"
	"github.com/renatospaka/scheduler/scheduler/document/core/entity"
	pkgController "github.com/renatospaka/scheduler/adapter/web/controller"
)

type GetDocumentByIdInputDto struct {
	ID int `json:"document_id"`
}

type GetDocumentByIdOutputDto struct {
	ID                          int    `json:"document_id"`
	Name                        string `json:"name"`
	IsActive                    bool   `json:"is_active"`
	dto.StandardStatusOutputDto `json:"request"`
}

func (u *DocumentUsecase) getDocumentById(in GetDocumentByIdInputDto) (GetDocumentByIdOutputDto, error) {
	log.Println("scheduler.document.core.usecase.GetDocument")

	if in.ID <= 0 {
		return GetDocumentByIdOutputDto{}, entity.ErrDocumentIdInvalid
	}
	docto, err := u.repo.GetDocument(in.ID)
	if err != nil {
		return GetDocumentByIdOutputDto{}, err
	}

	out := GetDocumentByIdOutputDto{
		ID:                      docto.ID(),
		Name:                    docto.Name(),
		IsActive:                docto.IsActive(),
		StandardStatusOutputDto: dto.StandardStatusOutputDto{
			Status: pkgController.REQUEST_SUCCESS,
		},
	}

	return out , nil
}
