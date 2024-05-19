package internalhttp

import "github.com/labstack/echo/v4"

func (s *Server) CreateHealthcheckGroup() error {
	group := s.server.Group("/healthcheck")
	group.GET("/", s.HealthCheck)
	return nil
}

func (s *Server) HealthCheck(c echo.Context) error {
	return c.JSON(200, createStatusResponse(200, "Ok"))
}
