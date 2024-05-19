package internalhttp

import (
	"context"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/app"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	address string
	app     *app.App
	logger  *logger.Logger
	server  *echo.Echo
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
	s.server = echo.New()

	s.server.Use(middleware.CORS())
	s.server.Use(middleware.Logger())

	s.CreateEventsGroup()

	s.server.GET("/swagger/*", echoSwagger.WrapHandler)
}

func (s *Server) Start(_ context.Context) error {
	return s.server.Start(s.address)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
