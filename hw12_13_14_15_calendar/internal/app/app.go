package app

import (
	"context"
	"errors"
	"time"

	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/logger"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage/event"
)

var (
	ErrEmptyTitle        = errors.New("empty title")
	ErrEventIsBusy       = errors.New("event is busy")
	ErrStartMoreThanNow  = errors.New("event start time more than now")
	ErrStartMoreThanStop = errors.New("event start time more than stop time")
)

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

func (a *App) CreateEvent(ctx context.Context, ev event.Event) (int, error) {
	if ev.Title == "" {
		return -1, ErrEmptyTitle
	}

	if ev.Start.After(ev.Stop) {
		return -1, ErrStartMoreThanStop
	}

	if time.Now().After(ev.Start) {
		return -1, ErrStartMoreThanNow
	}

	isBusy, err := a.storage.IsTimeBusy(ctx, ev)
	if err != nil {
		return -1, err
	}

	if isBusy {
		return -1, ErrEventIsBusy
	}

	return a.storage.Create(ctx, ev)
}

func (a *App) UpdateEvent(ctx context.Context, id int, ev event.Event) error {
	if ev.Title == "" {
		return ErrEmptyTitle
	}

	if ev.Start.After(ev.Stop) {
		return ErrStartMoreThanStop
	}

	if time.Now().After(ev.Start) {
		return ErrStartMoreThanNow
	}

	return a.storage.Update(ctx, id, ev)
}

func (a *App) Delete(ctx context.Context, id int) error {
	return a.storage.Delete(ctx, id)
}

func (a *App) DeleteAll(ctx context.Context) error {
	return a.storage.DeleteAll(ctx)
}

func (a *App) ListAll(ctx context.Context) ([]event.Event, error) {
	return a.storage.ListAll(ctx)
}

func (a *App) ListDay(ctx context.Context, date time.Time) ([]event.Event, error) {
	return a.storage.ListDay(ctx, date)
}

func (a *App) ListWeek(ctx context.Context, date time.Time) ([]event.Event, error) {
	return a.storage.ListWeek(ctx, date)
}

func (a *App) ListMonth(ctx context.Context, date time.Time) ([]event.Event, error) {
	return a.storage.ListMonth(ctx, date)
}
