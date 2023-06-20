package repository

import (
	"database/sql"
	"log"

	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
)

type WorkerRepository struct {
	DB *sql.DB
}

func NewWorkerRepository(db *sql.DB) *WorkerRepository {
	log.Println("initiating worker repository")
	
	return &WorkerRepository{
		DB: db,
	}
}

func (w *WorkerRepository) GetWorker(id int) (*entity.Worker, error) {
	log.Println("scheduler.worker.adapter.postgres.GetWorker")
	return nil, nil
}
