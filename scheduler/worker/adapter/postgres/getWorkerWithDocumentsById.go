package repository

import (
	"context"
	"database/sql"
	"log"

	entityDocto "github.com/renatospaka/scheduler/scheduler/document/core/entity"
	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
)

func (w *WorkerRepository) getWorkerWithDocumentsById(ctx context.Context, id int) (*entity.Worker, error) {
	log.Println("scheduler.worker.adapter.postgres.GetWorkerWithDocuments")

	query := `
	SELECT w.name, 
				w.profession, 
				w.is_active, 
				dw.id AS document_worker_id,
				d.id AS document_id, 
				d.name AS document_name, 
				d.is_active AS document_is_active
	FROM "Worker" w 
	LEFT JOIN "DocumentWorker" dw ON dw.worker_id = w.id
	LEFT JOIN "Document" d ON d.id = dw.worker_id
	WHERE w.id = $1`
	stmt, err := w.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}

	var (
		count  int
		worker *entity.Worker
	)

	for rows.Next() {
		var (
			name             sql.NullString
			profession       sql.NullString
			active           sql.NullBool
			documentWorkerId sql.NullInt16
			documentId       sql.NullInt16
			documentName     sql.NullString
			documentActive   sql.NullBool
		)

		/* 
	SELECT w.name, 
				w.profession, 
				w.is_active, 
				dw.id AS document_worker_id,
				d.id AS document_id, 
				d.name AS document_name, 
				d.is_active AS document_is_active
		*/
		err = rows.Scan(
			&name,
			&profession,
			&active,
			&documentWorkerId,
			&documentId,
			&documentName,
			&documentActive,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
		}

		log.Println(id, active.Bool, name.String, profession.String, documentId.Int16, documentActive.Bool, documentName.String)

		if count < 1 {
			// mount the worker
			worker, err = entity.NewWorker(id, active.Bool, name.String, profession.String)
			if err != nil {
				return nil, err
			}
		}

		// mount the document to be linked to the worker
		docto, err := entityDocto.NewDocument(int(documentId.Int16), documentActive.Bool, documentName.String)
		// log.Printf("docto => .ID(): %d, .Name(): %s\n", docto.ID(), docto.Name())
		if err != nil {
			return nil, err
		}

		// link the document to the worker
		workerDocto, err := entity.NewWorkerDocument(int(documentWorkerId.Int16), docto)
		// log.Printf("workerDocto => .ID(): %d, .Document().ID(): %d, .Document().Name(): %s\n", workerDocto.ID(), workerDocto.Document().ID(), workerDocto.Document().Name())
		if err != nil {
			return nil, err
		}

		err = worker.LinkToDocument(workerDocto)
		if err != nil {
			return nil, err
		}
		count++
	}

	err = rows.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return worker, nil
}
