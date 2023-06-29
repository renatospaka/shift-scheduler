package tests

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
)

// Mocked Postgress repository
// NOT TO BE CALLED under no circumstances out of the testing process
type workerRepositoryMock struct {
	db *sql.DB
}

// Mocked method GetWorker of Worker's Repository
// NOT TO BE CALLED under no circumstances out of the testing process
func (w *workerRepositoryMock) GetWorker(id int) (*entity.Worker, error) {
	if id == 0 {
		return nil, nil
	}
	
	worker, _ := entity.NewWorker(1, true, "Worker 1", "teacher")
	return worker, nil
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
