package tests

import (
	"context"
	"fmt"

	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
	"github.com/stretchr/testify/mock"
)


type workerRepositoryMock struct {
	mock.Mock
}

func (w *workerRepositoryMock) GetWorker(id int) (*entity.Worker, error) {
	wo, _ := entity.NewWorker(1, true, "Worker 1", "teacher")
	return wo, nil
}

func (w *workerRepositoryMock) GetWorkerWithDocuments(id int) (*entity.Worker, error) {
	return nil, nil
}

func (w *workerRepositoryMock) SetContext(ctx context.Context) {
	fmt.Println("Setting the Context")
}
