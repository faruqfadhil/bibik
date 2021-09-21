package error

import "errors"

var (
	ErrKeyNotFound      = errors.New("key not found")
	ErrDataEmpty        = errors.New("data empty")
	ErrReadFile         = errors.New("failed to read file")
	ErrWriteFile        = errors.New("failed to write to file")
	ErrMarshal          = errors.New("failed to marshal")
	ErrUnmarshal        = errors.New("failed to unmarshall")
	ErrCreateFile       = errors.New("failed to create new file")
	ErrCreateDir        = errors.New("failed to create ne directory")
	ErrExecDirrNotFound = errors.New("dirr not found")
)
