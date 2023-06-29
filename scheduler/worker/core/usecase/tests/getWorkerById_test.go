package tests

import (
	"testing"

	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
	"github.com/renatospaka/scheduler/scheduler/worker/core/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetWorkerById(t *testing.T) {
	// using Worker Repository Mock to inject into the usecase
	repo := new(workerRepositoryMock)
	usecases := usecase.NewWorkerUsecase(repo)

	// execute the method as if the repo was the real one
	in := usecase.GetWorkerByIdInputDto{ID: 1}
	worker, err := usecases.GetWorkerById(in)
	assert.Nil(t, err)
	assert.NotNil(t, worker)
	assert.Equal(t, 1, worker.ID)
	assert.Equal(t, "Worker 1", worker.Name)
	assert.Equal(t, "teacher", worker.Profession)
	assert.Equal(t, true, worker.IsActive)
	assert.Equal(t, "success", worker.StandardStatusOutputDto.Status)
}

func TestGetWorkerById_NotFound(t *testing.T) {
	// using Worker Repository Mock to inject into the usecase
	repo := new(workerRepositoryMock)
	usecases := usecase.NewWorkerUsecase(repo)

	// execute the method as if the repo was the real one
	in := usecase.GetWorkerByIdInputDto{ID: 0}
	worker, err := usecases.GetWorkerById(in)
	assert.NotNil(t, err)
	assert.NotNil(t, worker)
	assert.Equal(t, 0, worker.ID)
	assert.Equal(t, "", worker.Name)
	assert.Equal(t, "", worker.Profession)
	assert.ErrorIs(t, err, entity.ErrWorkerIdInvalid)
}
