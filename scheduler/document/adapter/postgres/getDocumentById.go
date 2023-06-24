package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/renatospaka/scheduler/scheduler/document/core/entity"
)

func (d *DocumentRepository) getDocumentById(ctx context.Context, id int) (*entity.Document, error) {
	log.Println("scheduler.document.adapter.postgres.GetDocument")

	query := `
	SELECT d.name, d.is_active
	FROM "Document" d 
	WHERE d.id = $1`
	stmt, err := d.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows := stmt.QueryRow(id)
	if err != nil {
		return nil, err
	}

	var (
		name       sql.NullString
		active     sql.NullBool
	)
	err = rows.Scan(&name, &active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	worker, err := entity.NewDocument(id, active.Bool, name.String)
	if err != nil {
		return nil, err
	}
	return worker, nil
}
