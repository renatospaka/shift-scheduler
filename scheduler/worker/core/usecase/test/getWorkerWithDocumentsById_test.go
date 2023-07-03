package test

import (
	"testing"

	"github.com/renatospaka/scheduler/scheduler/worker/core/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetWorkerWithDocumentsById(t *testing.T) {
	// using Worker Repository Mock to inject into the usecase
	repo := new(workerRepositoryMock)
	usecases := usecase.NewWorkerUsecase(repo)

	// execute the method as if the repo was the real one
	in := usecase.GetWorkerWithDocumentsByIdInputDto{ID: 1}
	worker, err := usecases.GetWorkerWithDocumentsById(in)
	assert.Nil(t, err)
	assert.NotNil(t, worker)
	assert.Equal(t, 1, worker.ID)
	assert.Equal(t, "Worker 1", worker.Name)
	assert.Equal(t, "teacher", worker.Profession)
	assert.Equal(t, true, worker.IsActive)
	assert.Equal(t, 2, len(worker.Documents))
	assert.Equal(t, 2312, worker.Documents[1].Document.ID)
	assert.Equal(t, "document 2312", worker.Documents[1].Document.Name)
	assert.Equal(t, "success", worker.StandardStatusOutputDto.Status)
}

func TestGetWorkerWithDocumentsById_NotFound(t *testing.T) {
	// using Worker Repository Mock to inject into the usecase
	repo := new(workerRepositoryMock)
	usecases := usecase.NewWorkerUsecase(repo)

	// execute the method as if the repo was the real one
	in := usecase.GetWorkerWithDocumentsByIdInputDto{ID: 1}
	worker, err := usecases.GetWorkerWithDocumentsById(in)
	assert.Nil(t, err)
	assert.NotNil(t, worker)
	assert.Equal(t, 1, worker.ID)
	assert.Equal(t, "Worker 1", worker.Name)
	assert.Equal(t, "teacher", worker.Profession)
	assert.Equal(t, true, worker.IsActive)
	assert.Equal(t, 2, len(worker.Documents))
	assert.Equal(t, 2312, worker.Documents[1].Document.ID)
	assert.Equal(t, "document 2312", worker.Documents[1].Document.Name)
	assert.Equal(t, "success", worker.StandardStatusOutputDto.Status)
}
