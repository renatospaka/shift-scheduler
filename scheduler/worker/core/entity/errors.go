package entity

import "errors"

var (
	ErrWorkerIdNotFound          = errors.New("worker id not found")
	ErrWorkerIdRequired          = errors.New("worker id is required")
	ErrWorkerIdInvalid           = errors.New("invalid worker id")
	ErrWorkerIdMustBeNumeric     = errors.New("worker id must be a number")
	ErrWorkerIDIsMissing         = errors.New("worker id is missing")
	ErrWorkerNameIsMissing       = errors.New("worker name is missing")
	ErrWorkerProfessionIsMissing = errors.New("worker profession is missing")

	ErrWorkerDocumentIdInvalid     = errors.New("invalid worker x document id")
	ErrWorkerDocumentAlreadyLinked = errors.New("worker already linked to document")
)
