package test

import (
	"testing"

	"github.com/renatospaka/scheduler/scheduler/document/core/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewDocument(t *testing.T) {
	docto, err := entity.NewDocument(1, true, "Document 1")
	assert.Nil(t, err)
	assert.NotNil(t, docto)
	assert.Equal(t, "Document 1", docto.Name())
	assert.True(t, docto.IsActive())
	assert.True(t, docto.IsValid())
}

func TestNewDocumentInvalidId(t *testing.T) {
	docto, err := entity.NewDocument(0, true, "Document 1")
	assert.NotNil(t, err)
	assert.Nil(t, docto)
	assert.Error(t, err, entity.ErrDocumentIDIsMissing)
}

func TestNewDocumentInvalidName(t *testing.T) {
	docto, err := entity.NewDocument(1, true, "")
	assert.NotNil(t, err)
	assert.Nil(t, docto)
	assert.Error(t, err, entity.ErrDocumentNameIsMissing)
}

func TestChangeName(t *testing.T) {
	docto, err := entity.NewDocument(1, true, "Document 1")
	assert.Nil(t, err)
	assert.NotNil(t, docto)

	err = docto.ChangeName("Document 1 Smith")
	assert.Nil(t, err)
	assert.Equal(t, "Document 1 Smith", docto.Name())

	err = docto.ChangeName("")
	assert.NotNil(t, err)
	assert.Error(t, err, entity.ErrDocumentNameIsMissing)
}

func TestChangeActive(t *testing.T) {
	docto, err := entity.NewDocument(1, true, "Document 1")
	assert.Nil(t, err)
	assert.NotNil(t, docto)

	docto.ChangeActive(false)
	assert.False(t, docto.IsActive())

	docto.ChangeActive(true)
	assert.True(t, docto.IsActive())
}
