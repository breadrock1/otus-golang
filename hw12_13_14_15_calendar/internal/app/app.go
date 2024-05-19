package app

import (
	"context"
	"errors"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/logger"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage"
	"time"
)

var EmptyTitleError = errors.New("empty title")
var EventIsBusyError = errors.New("event is busy")
var StartMoteThanNow = errors.New("event start time more than now")

type App struct {
	logger  *logger.Logger
	storage storage.Storage
}

func New(storage storage.Storage, logger *logger.Logger) *App {
	return &App{
		logger:  logger,
		storage: storage,
	}
}

func (a *App) CreateEvent(ctx context.Context, event storage.Event) (int, error) {
	if event.Title == "" {
		return -1, EmptyTitleError
	}

	if event.Start.After(event.Stop) {
		event.Start, event.Stop = event.Stop, event.Start
	}

	if !time.Now().After(event.Start) {
		return -1, StartMoteThanNow
	}

	isBusy, err := a.storage.IsTimeBusy(ctx, event.UserID, event.Start, event.Stop, 0)
	if err != nil {
		return -1, err
	}

	if isBusy {
		return -1, EventIsBusyError
	}

	return a.storage.Create(ctx, event)
}

func (a *App) UpdateEvent(ctx context.Context, id int, event storage.Event) error {
	if event.Title == "" {
		return EmptyTitleError
	}

	if event.Start.After(event.Stop) {
		event.Start, event.Stop = event.Stop, event.Start
	}

	if !time.Now().After(event.Start) {
		return StartMoteThanNow
	}

	isBusy, err := a.storage.IsTimeBusy(ctx, event.UserID, event.Start, event.Stop, id)
	if err != nil {
		return err
	}

	if isBusy {
		return EventIsBusyError
	}

	return a.storage.Update(ctx, id, event)
}

func (a *App) Delete(ctx context.Context, id int) error {
	return a.storage.Delete(ctx, id)
}

func (a *App) DeleteAll(ctx context.Context) error {
	return a.storage.DeleteAll(ctx)
}

func (a *App) ListAll(ctx context.Context) ([]storage.Event, error) {
	return a.storage.ListAll(ctx)
}

func (a *App) ListDay(ctx context.Context, date time.Time) ([]storage.Event, error) {
	return a.storage.ListDay(ctx, date)
}

func (a *App) ListWeek(ctx context.Context, date time.Time) ([]storage.Event, error) {
	return a.storage.ListWeek(ctx, date)
}

func (a *App) ListMonth(ctx context.Context, date time.Time) ([]storage.Event, error) {
	return a.storage.ListMonth(ctx, date)
}
