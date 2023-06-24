package entity

import (
	"log"
	"strings"
)

type Document struct {
	id         int
	isActive   bool
	isValid    bool
	name       string
}

func NewDocument(id int, isActive bool, name string) (*Document, error ) {
	log.Println("initiating document entity")
	
	docto := &Document{
		id:         id,
		isActive:   isActive,
		isValid:    false,
		name:       name,
	}
	err := docto.Validate()

	if err != nil {
		return nil, err
	}
	return docto, nil
}

func (d *Document) ID() int {
	return d.id
}

func (d *Document) ChangeName(name string) error {
	current := d.name
	d.name = name

	err := d.Validate()
	if err != nil {
		d.name = strings.TrimSpace(current)
		return err
	}
	return nil
}

func (d *Document) Name() string {
	return d.name
}

func (d *Document) ChangeActive(active bool) {
	d.isActive = active
}

func (d *Document) IsActive() bool {
	return d.isActive
}

func (d *Document) IsValid() bool {
	return d.isValid
}

func (d *Document) Validate() error {
	d.isValid = false

	if d.id <= 0 {
		return ErrDocumentIDIsMissing
	}

	if strings.TrimSpace(d.name) == "" {
		return ErrDocumentNameIsMissing
	}

	d.isValid = true
	return nil
}
