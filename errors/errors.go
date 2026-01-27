// Package errors wraps error handling with traceback and public error formatting
package errors

import (
	"github.com/cockroachdb/errors"
)

// Error creation functions
var (
	WithStack = errors.WithStack
	Wrapf     = errors.Wrapf
	New       = errors.New
	Newf      = errors.Newf
)

// Error inspection functions
var (
	Is = errors.Is
	As = errors.As
)
