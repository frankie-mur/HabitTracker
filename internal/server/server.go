package server

import (
	"fmt"
	"net"
	"strconv"

	"github.com/frankie-mur/habit_tracker/api"
	"github.com/frankie-mur/habit_tracker/log"
	"google.golang.org/grpc"
)

// Server is the implementation of the gRPC server.
type Server struct {
	api.UnimplementedHabitsServer
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

// ListenAndServe starts listening to the port and serving requests.
func (s *Server) ListenAndServe(port int) error {
	const addr = "127.0.0.1"

	listener, err := net.Listen("tcp", net.JoinHostPort(addr, strconv.Itoa(port)))
	if err != nil {
		return fmt.Errorf("unable to listen to tcp port %d: %w", port, err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterHabitsServer(grpcServer, s)

	s.lgr.Logf("starting server on port %d\n", port)

	err = grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("error while listening: %w", err)
	}

	//Stop or GracefulStop was called, no reason to be alarmed.
	return nil

}
