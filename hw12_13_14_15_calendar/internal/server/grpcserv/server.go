package grpcserv

import (
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/app"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/logger"
	"google.golang.org/grpc"
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
	server.Init()
	return server
}

func (s *Server) Init() {
	//lis, err := net.Listen("tcp", s.address)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var opts []grpc.ServerOption
	//grpcServer := grpc.NewServer(opts...)
	//userGrpcServer := NewUserServer(s.app)
	//
	//s.server = grpc.NewServer(opts...)
	//s.server.RegisterService(NewService(s.app))

}
