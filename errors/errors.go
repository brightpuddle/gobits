// Package errors wraps error handling with traceback and public error formatting
package errors

import (
	"github.com/cockroachdb/errors"
)

var (
	WithStack = errors.WithStack
	Wrapf     = errors.Wrapf
	New       = errors.New
	Newf      = errors.Newf
)
