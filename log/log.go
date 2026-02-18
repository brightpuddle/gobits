// Package logger provides a simple wrapper around zerolog.
package log

import (
	"io"
	"os"
	"runtime"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

// Logger aliases the zerolog.Logger
type Logger = zerolog.Logger

type Config struct {
	File string
}

var (
	DefaultConfig = Config{
		File: "out.log",
	}
	logger = New(DefaultConfig)

	// Convenience shortcuts for logging levels
	Debug = logger.Debug
	Info  = logger.Info
	Warn  = logger.Warn
	Error = logger.Error
	Fatal = logger.Fatal
	Panic = logger.Panic
	With  = logger.With
)

var (
	// Convenience shortcut for setting logging level
	DebugLevel = zerolog.DebugLevel
	InfoLevel  = zerolog.InfoLevel
	WarnLevel  = zerolog.WarnLevel
	ErrorLevel = zerolog.ErrorLevel
	SetLevel   = zerolog.SetGlobalLevel
)

// New creates a new multi-level logger
func New(cfg Config) Logger {
	if testing.Testing() {
		return zerolog.Nop()
	}

	// If filename is specified, open file and assume file logging
	file, err := os.Create(cfg.File)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create log file")
	}

	// Levels default to zero, i.e. debug
	return zerolog.New(

		zerolog.ConsoleWriter{
			Out:     io.MultiWriter(os.Stderr, file),
			NoColor: runtime.GOOS == "windows",
		}).With().Timestamp().Logger()
}

func init() {
	// defaults
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.DurationFieldInteger = true
}
