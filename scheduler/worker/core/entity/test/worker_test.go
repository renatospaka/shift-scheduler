package test

import (
	"testing"

	"github.com/renatospaka/scheduler/scheduler/worker/core/entity"
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