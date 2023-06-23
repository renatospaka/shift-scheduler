package test

import (
	"testing"

	"github.com/renatospaka/scheduler/core/id"
	"github.com/stretchr/testify/assert"
)

func TestNewId(t *testing.T) {
	id := id.NewID()
	assert.NotNil(t, id)
	assert.NotEmpty(t, id.String())
}

func TestParse(t *testing.T) {
	uuid := "f2cfe684-9c99-455a-a29d-656d0dd99784"
	id, err := id.Parse(uuid)
	assert.NotEmpty(t, id)
	assert.Equal(t, uuid, id.String())
	assert.Nil(t, err)
}

func TestParseFail(t *testing.T) {
	uuid := "not-uuid"
	id, err := id.Parse(uuid)
	assert.Empty(t, id)
	assert.NotEqual(t, uuid, id.String())
	assert.NotNil(t, err)
}
