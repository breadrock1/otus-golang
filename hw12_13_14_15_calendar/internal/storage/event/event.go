package event

import (
	"context"
	"time"
)

type Event struct {
	ID           int
	Title        string         `json:"title" example:"Alarm"`
	Start        time.Time      `json:"start" example:"2024-05-10T10:07:35Z"`
	Stop         time.Time      `json:"stop" example:"2024-05-11T10:07:35Z"`
	Description  string         `json:"description" example:"Alarm to wake up"`
	UserID       int            `json:"user" example:"1"`
	Notification *time.Duration `json:"notification" example:"10"`
}

type Events interface {
	Create(ctx context.Context, event Event) (int, error)
	Update(ctx context.Context, id int, change Event) error
	Delete(ctx context.Context, id int) error
	DeleteAll(ctx context.Context) error
	ListAll(ctx context.Context) ([]Event, error)
	ListDay(ctx context.Context, date time.Time) ([]Event, error)
	ListWeek(ctx context.Context, date time.Time) ([]Event, error)
	ListMonth(ctx context.Context, date time.Time) ([]Event, error)
	GetEventsByNotifier(_ context.Context, start time.Time, end time.Time) ([]Event, error)
	RemoveAfter(_ context.Context, time time.Time) error
	IsTimeBusy(ctx context.Context, event Event) (bool, error)
}
