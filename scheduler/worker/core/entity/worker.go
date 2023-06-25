package entity

import (
	"log"
	"strings"

	"github.com/renatospaka/scheduler/scheduler/document/core/entity"
)

type Worker struct {
	id         int
	isActive   bool
	isValid    bool
	name       string
	profession string
	documents  []*WorkerDocument
}

func NewWorker(id int, isActive bool, name, profession string) (*Worker, error) {
	log.Println("initiating worker entity")

	w := &Worker{
		id:         id,
		isActive:   isActive,
		isValid:    false,
		name:       name,
		profession: profession,
	}
	w.documents = make([]*WorkerDocument, 0)
	
	err := w.Validate()
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (w *Worker) ID() int {
	return w.id
}

func (w *Worker) ChangeName(name string) error {
	current := w.name
	w.name = name

	err := w.Validate()
	if err != nil {
		w.name = strings.TrimSpace(current)
		return err
	}
	return nil
}

func (w *Worker) Name() string {
	return w.name
}

func (w *Worker) ChangeProfession(profession string) error {
	current := w.profession
	w.profession = profession

	err := w.Validate()
	if err != nil {
		w.profession = strings.TrimSpace(current)
		return err
	}
	return nil
}

func (w *Worker) Profession() string {
	return w.profession
}

func (w *Worker) ChangeActive(active bool) {
	w.isActive = active
}

func (w *Worker) IsActive() bool {
	return w.isActive
}

func (w *Worker) AddDocument(docto *entity.Document) error {
	err := docto.Validate()
	if err != nil {
		return err
	}

	if w.documents == nil {
		w.documents = make([]*WorkerDocument, 0)
	}
	id := len(w.documents) + 1

	document := &WorkerDocument{
		id:       id,
		document: docto,
	}
	w.documents = append(w.documents, document)
	return nil
}

func (w *Worker) Documents() []*WorkerDocument {
	return w.documents
}

func (w *Worker) WorkerDocument(id int) *WorkerDocument {
	if id <= 0 {
		return nil
	}


	// small slice
	if id <= 50 {
		for ix, docto := range w.documents {
			if docto.id == id {
				return w.documents[ix] 
			}
		}
	} 

	// big slice
	if id > 50 {
		left, right := 0, len(w.documents) - 1
		for left < right {
			if w.documents[left].id == id {
				return w.documents[left]
			}
			if w.documents[right].id == id {
				return w.documents[right]
			}

			left++
			right--
			if left == right && w.documents[right].id == id {
				return w.documents[right]
			}
		}
	}
	return nil
}

func (w *Worker) IsValid() bool {
	return w.isValid
}

func (w *Worker) Validate() error {
	w.isValid = false

	if w.id <= 0 {
		return ErrWorkerIdInvalid
	}

	if strings.TrimSpace(w.name) == "" {
		return ErrWorkerNameIsMissing
	}

	if strings.TrimSpace(w.profession) == "" {
		return ErrWorkerProfessionIsMissing
	}

	err := w.validateDocuments()
	if err != nil {
		return err
	}

	w.isValid = true
	return nil
}

func (w *Worker) validateDocuments() error {
	if  len(w.documents) == 0 {
		return nil
	}

	for _, docto := range w.documents {
		err := docto.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
