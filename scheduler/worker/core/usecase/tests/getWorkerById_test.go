package tests

import (
	"testing"

	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
	"github.com/renatospaka/scheduler/scheduler/worker/core/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetWorkerById(t *testing.T) {
	// worker is not being mocked
	worker, _ := entity.NewWorker(1, true, "Worker 1", "teacher")
	
	// repo is mocked
	repo := new(workerRepositoryMock)
	repo.On("GetWorker", 1).Return(worker, nil)

	// usecases uses the mocked repo
	usecases := usecase.NewWorkerUsecase(repo, nil)
	in := usecase.GetWorkerByIdInputDto{ID: 1}
	returnedWorker, err := usecases.GetWorkerById(in)
	assert.Nil(t, err)
	assert.NotNil(t, returnedWorker)
	assert.Equal(t, worker.ID(), returnedWorker.ID)
	assert.Equal(t, worker.Name(), returnedWorker.Name)
	assert.Equal(t, worker.IsActive(), returnedWorker.IsActive)

	repo.AssertExpectations(t)
}
