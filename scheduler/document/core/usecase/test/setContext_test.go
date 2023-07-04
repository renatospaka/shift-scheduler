package test

import (
	"context"
	"testing"

	"github.com/renatospaka/scheduler/scheduler/document/core/usecase"
	"github.com/stretchr/testify/assert"
)

func TestSetContext(t *testing.T) {
	// using Document Repository Mock to inject into the usecase
	repo := new(documentRepositoryMock)
	usecases := usecase.NewDocumentUsecase(repo)

	usecases.SetContext(context.Background())
	ctx := usecases.Context()
	assert.NotNil(t, ctx)
	assert.IsType(t, context.Background(), ctx)
}
