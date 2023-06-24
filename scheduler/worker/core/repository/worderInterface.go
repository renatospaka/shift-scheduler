package repository

import (
	"context"

	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
)

type WorkerInterface interface {
	SetContext(context.Context)
	GetWorker(id int) (*entity.Worker, error) 
}
