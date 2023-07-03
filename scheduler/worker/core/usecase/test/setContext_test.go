package test

import (
	"context"
	"testing"

	"github.com/renatospaka/scheduler/scheduler/worker/core/usecase"
	"github.com/stretchr/testify/assert"
)

func TestSetContext(t *testing.T) {
	// using Worker Repository Mock to inject into the usecase
	repo := new(workerRepositoryMock)
	usecases := usecase.NewWorkerUsecase(repo)

	usecases.SetContext(context.Background())
	ctx := usecases.Context()
	assert.NotNil(t, ctx)
	assert.IsType(t, context.Background(), ctx)
}
