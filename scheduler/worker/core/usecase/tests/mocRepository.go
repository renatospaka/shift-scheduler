package tests

import (
	"context"
	"fmt"

	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
	"github.com/stretchr/testify/mock"
)

// Mocked Postgress repository
// NOT TO BE CALLED under no circumstances out of the testing process
type workerRepositoryMock struct {
	mock.Mock
}

// Mocked method GetWorker of Worker's Repository
// NOT TO BE CALLED under no circumstances out of the testing process
func (w *workerRepositoryMock) GetWorker(id int) (*entity.Worker, error) {
	wo, _ := entity.NewWorker(1, true, "Worker 1", "teacher")
	return wo, nil
}

// Mocked method GetWorkerWithDocuments of Worker's Repository
// NOT TO BE CALLED under no circumstances out of the testing process
func (w *workerRepositoryMock) GetWorkerWithDocuments(id int) (*entity.Worker, error) {
	return nil, nil
}

// Mocked method GetWorkerWithDocuments of Worker's Repository
// NOT TO BE CALLED under no circumstances out of the testing process
func (w *workerRepositoryMock) SetContext(ctx context.Context) {
	fmt.Println("Setting the Context")
}
