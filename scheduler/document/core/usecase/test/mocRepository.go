package test

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/renatospaka/scheduler/scheduler/document/core/entity"
)

// Mocked Postgress repository
// NOT TO BE CALLED under no circumstances out of the testing process
type documentRepositoryMock struct {
	db *sql.DB
}

// Mocked method GetDocument of Document's Repository
// NOT TO BE CALLED under no circumstances out of the testing process
func (d *documentRepositoryMock) GetDocument(id int) (*entity.Document, error) {
	if id == 0 {
		return nil, nil
	}

	document, _ := entity.NewDocument(1, true, "Document 1")
	return document, nil
}

// Mocked method SetContext of Document's Repository
// NOT TO BE CALLED under no circumstances out of the testing process
func (d *documentRepositoryMock) SetContext(ctx context.Context) {
	fmt.Println("Setting the Context")
}
