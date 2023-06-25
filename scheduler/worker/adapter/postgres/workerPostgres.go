package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
)

type WorkerRepository struct {
	DB  *sql.DB
	ctx context.Context
}

func NewWorkerRepository(db *sql.DB) *WorkerRepository {
	log.Println("initiating worker repository")

	return &WorkerRepository{
		DB: db,
	}
}

func (w *WorkerRepository) GetWorker(id int) (*entity.Worker, error) {
	return w.getWorkerById(w.ctx, id)
}

func (w *WorkerRepository) GetWorkerWithDocuments(id int) (*entity.Worker, error) {
	return w.getWorkerWithDocumentsById(w.ctx, id)
}

func (w *WorkerRepository) SetContext(ctx context.Context) {
	w.ctx = ctx
}
