package grpcserv

import "github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/app"

type Service struct {
	app *app.App
}

func NewUserServer(app *app.App) *Service {
	return &Service{app: app}
}
