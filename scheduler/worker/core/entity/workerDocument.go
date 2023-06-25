package entity

import "github.com/renatospaka/scheduler/scheduler/document/core/entity"

type WorkerDocument struct {
	id       int
	document *entity.Document
	isValid  bool
}

func NewWorkerDocument(id int, docto *entity.Document) (*WorkerDocument, error) {
	d := &WorkerDocument{
		id:       id,
		document: docto,
	}

	err := d.Validate()
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (d *WorkerDocument) ID() int {
	return d.id
}

func (d *WorkerDocument) Document() *entity.Document {
	return d.document
}

func (d *WorkerDocument) Validate() error {
	d.isValid = false

	if d.id <= 0 {
		return ErrWorkerDocumentIdInvalid
	}

	err := d.document.Validate()
	if err != nil {
		return err
	}

	d.isValid = true
	return nil
}
