package grpcserv

import (
	"context"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/app"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/logger"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	address string
	app     *app.App
	logger  *logger.Logger
	server  *grpc.Server
}

func NewServer(address string, app *app.App, logger *logger.Logger) *Server {
	server := &Server{
		address: address,
		app:     app,
		logger:  logger,
	}
	return server
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal(err)
	}

	s.server = grpc.NewServer()
	service := NewService(*s.app)
	RegisterCalendarServer(s.server, service)
	return s.server.Serve(lis)
}

func (s *Server) Stop(_ context.Context) error {
	s.server.GracefulStop()
	return nil
}
