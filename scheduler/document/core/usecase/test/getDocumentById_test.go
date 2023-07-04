package test

import (
	"testing"

	"github.com/renatospaka/scheduler/scheduler/document/core/entity"
	"github.com/renatospaka/scheduler/scheduler/document/core/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetDocumentById(t *testing.T) {
	// using Document Repository Mock to inject into the usecase
	repo := new(documentRepositoryMock)
	usecases := usecase.NewDocumentUsecase(repo)

	// execute the method as if the repo was the real one
	in := usecase.GetDocumentByIdInputDto{ID: 1}
	document, err := usecases.GetDocumentById(in)
	assert.Nil(t, err)
	assert.NotNil(t, document)
	assert.Equal(t, 1, document.ID)
	assert.Equal(t, "Document 1", document.Name)
	assert.Equal(t, true, document.IsActive)
	assert.Equal(t, "success", document.StandardStatusOutputDto.Status)
}

func TestGetDocumentById_NotFound(t *testing.T) {
	// using Document Repository Mock to inject into the usecase
	repo := new(documentRepositoryMock)
	usecases := usecase.NewDocumentUsecase(repo)

	// execute the method as if the repo was the real one
	in := usecase.GetDocumentByIdInputDto{ID: 0}
	document, err := usecases.GetDocumentById(in)
	assert.NotNil(t, err)
	assert.NotNil(t, document)
	assert.Equal(t, 0, document.ID)
	assert.Equal(t, "", document.Name)
	assert.ErrorIs(t, err, entity.ErrDocumentIdInvalid)
}
