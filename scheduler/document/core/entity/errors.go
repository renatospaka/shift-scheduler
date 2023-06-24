package entity

import "errors"

var (
	ErrDocumentIdNotFound          = errors.New("documento id not found")
	ErrDocumentIdRequired          = errors.New("documento id is required")
	ErrDocumentIdInvalid           = errors.New("invalid documento id")
	ErrDocumentIdMustBeNumeric     = errors.New("documento id must be a number")
	ErrDocumentIDIsMissing         = errors.New("documento id is missing")
	ErrDocumentNameIsMissing       = errors.New("documento name is missing")
)
