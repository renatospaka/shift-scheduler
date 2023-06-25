package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
)

func (w *WorkerRepository) getWorkerById(ctx context.Context, id int) (*entity.Worker, error) {
	log.Println("scheduler.worker.adapter.postgres.GetWorker")

	query := `
	SELECT w.name, w.profession, w.is_active
	FROM "Worker" w 
	WHERE w.id = $1`
	stmt, err := w.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows := stmt.QueryRow(id)
	var (
		name       sql.NullString
		profession sql.NullString
		active     sql.NullBool
	)
	err = rows.Scan(&name, &profession, &active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	worker, err := entity.NewWorker(id, active.Bool, name.String, profession.String)
	if err != nil {
		return nil, err
	}
	return worker, nil
}
