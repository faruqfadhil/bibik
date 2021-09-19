package error

import "errors"

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrDataEmpty   = errors.New("data empty")
	ErrReadFile    = errors.New("internal error: failed to read file")
	ErrWriteFile   = errors.New("internal error: failed to write to file")
	ErrMarshal     = errors.New("internal error: failed to marshal")
	ErrUnmarshal   = errors.New("internal error: failed to unmarshall")
	ErrCreateFile  = errors.New("internal error: failed to create new file")
)
