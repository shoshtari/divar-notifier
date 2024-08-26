package pkg

import "github.com/pkg/errors"

var (
	ErrInternal = errors.New("internal error")
	ErrCanceled = errors.New("user cacneled")
)
