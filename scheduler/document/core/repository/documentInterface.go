package repository

import (
	"context"

	"github.com/renatospaka/scheduler/scheduler/document/core/entity"
)

type DocumentInterface interface {
	SetContext(context.Context)
	GetWorker(id int) (*entity.Document, error) 
}
