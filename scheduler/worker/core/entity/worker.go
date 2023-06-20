package entity

import (
	"errors"
	"log"
	"strings"
)

type Worker struct {
	id         int
	isActive   bool
	isValid    bool
	name       string
	profession string
}

func NewWorker(id int, isActive bool, name, profession string) (*Worker, error ) {
	log.Println("initiating worker entity")
	
	worker := &Worker{
		id:         id,
		isActive:   isActive,
		isValid:    false,
		name:       name,
		profession: profession,
	}
	err := worker.Validate()

	if err != nil {
		return nil, err
	}
	return worker, nil
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

func (w *Worker) IsValid() bool {
	return w.isValid
}

func (w *Worker) Validate() error {
	w.isValid = false

	if w.id <= 0 {
		return ErrWorkerIDIsMissing
	}

	if strings.TrimSpace(w.name) == "" {
		return ErrWorkerNameIsMissing
	}

	if strings.TrimSpace(w.profession) == "" {
		return ErrWorkerProfessionIsMissing
	}

	w.isValid = true
	return nil
}

var (
	ErrWorkerIDIsMissing = errors.New("id is missing")
	ErrWorkerNameIsMissing = errors.New("name is missing")
	ErrWorkerProfessionIsMissing = errors.New("profession is missing")
)
