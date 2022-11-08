package util

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

//var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
var Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
	Level(zerolog.TraceLevel).
	With().
	Timestamp().
	Caller().
	Logger()
