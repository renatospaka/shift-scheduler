package repository

import "github.com/renatospaka/scheduler/scheduler/worker/core/entity"

type WorkerInterface interface {
	GetWorker(id int) (*entity.Worker, error) 
}
