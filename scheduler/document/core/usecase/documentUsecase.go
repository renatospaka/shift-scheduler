package usecase

import (
	"context"

	"github.com/renatospaka/scheduler/scheduler/document/core/repository"
)

type DocumentUsecase struct {
	ctx        context.Context
	repo       repository.DocumentInterface
	// dispatcher *event.EventDispatcher
}

func NewDocumentUsecase(repo repository.DocumentInterface) *DocumentUsecase {
	return &DocumentUsecase{
		repo: repo,
	}
}

func (u *DocumentUsecase) GetDocumentById(in GetDocumentByIdInputDto) (GetDocumentByIdOutputDto, error) {
	return u.getDocumentById(in)
}

func (u *DocumentUsecase) SetContext(ctx context.Context) {
	u.ctx = ctx
	u.repo.SetContext(ctx)
}

func (u *DocumentUsecase) Context() context.Context {
	return u.ctx
}
