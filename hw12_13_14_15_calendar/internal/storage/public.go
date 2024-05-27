package storage

import (
	"context"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/config"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage/event"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage/memcache"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage/sqlstorage"
)

type Storage interface {
	Service
	event.Events
}

type Service interface {
	Connect(ctx context.Context, connect string) error
	Close(ctx context.Context) error
}

func New(config *config.DatabaseConfig) Storage {
	if config.EnableInMemory {
		storeService := memcache.New()
		return &storeService
	}

	storeService := sqlstorage.New()
	return &storeService
}
