package internalhttp

import (
	"github.com/labstack/echo/v4"
	"time"

	log "github.com/sirupsen/logrus"
)

func (s *Server) CustomLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		res := next(c)

		ip := c.RealIP()
		userAgent := c.Request().Header.Get("user-agent")
		log.WithField("IP", ip).
			WithField("Method", c.Request().Method).
			WithField("URL", c.Request().URL).
			WithField("HTTP version", c.Request().Proto).
			WithField("User-Agent", userAgent).
			WithField("Latency", time.Since(start)).
			Info("http request processed")

		return res
	}
}
