package internal

import (
	"github.com/frankie-mur/habit_tracker/log"
)

// Server is the implementation of the gRPC server.
type Server struct {
	lgr log.Logger
}

func New(lgr log.Logger) *Server {
	return &Server{
		lgr: lgr,
	}
}

type Logger interface {
	Logf(format string, args ...any)
}
