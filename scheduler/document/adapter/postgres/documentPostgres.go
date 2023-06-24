package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/renatospaka/scheduler/scheduler/document/core/entity"
)

type DocumentRepository struct {
	DB  *sql.DB
	ctx context.Context
}

func NewDocumentRepository(db *sql.DB) *DocumentRepository {
	log.Println("initiating document repository")

	return &DocumentRepository{
		DB: db,
	}
}

func (d *DocumentRepository) GetDocument(id int) (*entity.Document, error) {
	return d.getDocumentById(d.ctx, id)
}

func (d *DocumentRepository) SetContext(ctx context.Context) {
	d.ctx = ctx
}
