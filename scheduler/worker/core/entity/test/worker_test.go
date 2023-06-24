package test

import (
	"testing"

	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
	doctoEntity "github.com/renatospaka/scheduler/scheduler/document/core/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewWorker(t *testing.T) {
	worker, err := entity.NewWorker(1, true, "Worker 1", "professor")
	assert.Nil(t, err)
	assert.NotNil(t, worker)
	assert.Equal(t, "Worker 1", worker.Name())
	assert.Equal(t, "professor", worker.Profession())
	assert.True(t, worker.IsActive())
	assert.True(t, worker.IsValid())
}

func TestNewWorkerInvalidId(t *testing.T) {
	worker, err := entity.NewWorker(0, true, "Worker 1", "professor")
	assert.NotNil(t, err)
	assert.Nil(t, worker)
	assert.Error(t, err, entity.ErrWorkerIDIsMissing)
}

func TestNewWorkerInvalidName(t *testing.T) {
	worker, err := entity.NewWorker(1, true, "", "professor")
	assert.NotNil(t, err)
	assert.Nil(t, worker)
	assert.Error(t, err, entity.ErrWorkerNameIsMissing)
}

func TestNewWorkerInvalidProfession(t *testing.T) {
	worker, err := entity.NewWorker(1, true, "Worker 1", "")
	assert.NotNil(t, err)
	assert.Nil(t, worker)
	assert.Error(t, err, entity.ErrWorkerProfessionIsMissing)
}

func TestChangeName(t *testing.T) {
	worker, err := entity.NewWorker(1, true, "Worker 1", "professor")
	assert.Nil(t, err)
	assert.NotNil(t, worker)

	err = worker.ChangeName("Worker 1 Smith")
	assert.Nil(t, err)
	assert.Equal(t, "Worker 1 Smith", worker.Name())

	err = worker.ChangeName("")
	assert.NotNil(t, err)
	assert.Error(t, err, entity.ErrWorkerNameIsMissing)
}

func TestChangeProfession(t *testing.T) {
	worker, err := entity.NewWorker(1, true, "Worker 1", "professor")
	assert.Nil(t, err)
	assert.NotNil(t, worker)

	err = worker.ChangeProfession("teacher")
	assert.Nil(t, err)
	assert.Equal(t, "teacher", worker.Profession())

	err = worker.ChangeProfession("")
	assert.NotNil(t, err)
	assert.Error(t, err, entity.ErrWorkerProfessionIsMissing)
}

func TestChangeActive(t *testing.T) {
	worker, err := entity.NewWorker(1, true, "Worker 1", "professor")
	assert.Nil(t, err)
	assert.NotNil(t, worker)

	worker.ChangeActive(false)
	assert.False(t, worker.IsActive())

	worker.ChangeActive(true)
	assert.True(t, worker.IsActive())
}

func TestAddWorkerDocuments(t *testing.T) {
	worker, err := entity.NewWorker(1, true, "Worker 1", "professor")
	assert.Nil(t, err)
	assert.NotNil(t, worker)
	assert.True(t, worker.IsValid())

	docto, _ := doctoEntity.NewDocument(1, false, "Docto 1")
	err = worker.AddDocument(docto)
	assert.Nil(t, err)
	assert.True(t, worker.IsValid())
	assert.Equal(t, 1, len(worker.Documents()))

	doctos := worker.Documents()
	assert.NotNil(t, doctos)
	assert.Equal(t, 1, len(doctos))

	d := worker.WorkerDocument(1)
	assert.NotNil(t, d)
	assert.Equal(t, 1, d.ID())
	assert.Equal(t, "Docto 1", d.Document().Name())
	assert.NotEqual(t, "Docto 2", d.Document().Name())
	assert.False(t, d.Document().IsActive())
}

func TestAddDocument(t *testing.T) {
	worker, err := entity.NewWorker(1, true, "Worker 1", "professor")
	assert.Nil(t, err)
	assert.NotNil(t, worker)
	assert.True(t, worker.IsValid())

	docto1, _ := doctoEntity.NewDocument(1, false, "Docto 1")
	docto2, _ := doctoEntity.NewDocument(2, true, "Docto 2")
	docto3, _ := doctoEntity.NewDocument(3, true, "Docto 3")
	worker.AddDocument(docto1)
	worker.AddDocument(docto2)
	worker.AddDocument(docto3)
	assert.Equal(t, 3, len(worker.Documents()))

	doctos := worker.Documents()
	assert.NotNil(t, doctos)
	assert.Equal(t, 3, len(doctos))

	d := worker.WorkerDocument(3)
	assert.NotNil(t, d)
	assert.Equal(t, 3, d.ID())
	assert.Equal(t, "Docto 3", d.Document().Name())
	assert.True(t, d.Document().IsActive())
}
