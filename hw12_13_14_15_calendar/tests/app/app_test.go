package app

import (
	"context"
	"testing"
	"time"

	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/cmd/config"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/app"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/logger"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage/event"
	"github.com/stretchr/testify/require"
)

func TestAppBase(t *testing.T) {
	storeConfig := config.DatabaseConfig{EnableInMemory: true}
	logConfig := config.LoggerConfig{EnableFileLog: false, Level: "INFO"}
	sLog, _ := logger.New(&logConfig)

	notifyDur := 1 * time.Hour
	initDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	baseEvent := &event.Event{
		ID:           1,
		Title:        "test",
		Start:        initDate.Add(1 * time.Hour),
		Stop:         initDate.Add(2 * time.Hour),
		Description:  "description",
		UserID:       1,
		Notification: &notifyDur,
	}

	t.Run("Create event", func(t *testing.T) {
		storageService := storage.New(&storeConfig)
		appService := app.New(storageService, sLog)
		ctx := context.Background()

		id, err := appService.CreateEvent(ctx, *baseEvent)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		events, err := appService.ListDay(ctx, initDate)
		require.NoError(t, err)
		require.Equal(t, 1, len(events))
	})

	t.Run("Update event", func(t *testing.T) {
		storageService := storage.New(&storeConfig)
		appService := app.New(storageService, sLog)
		ctx := context.Background()

		id, err := appService.CreateEvent(ctx, *baseEvent)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		updatedEvent := event.Event{
			ID:           baseEvent.ID,
			Title:        "New title",
			Start:        baseEvent.Stop.Add(12 * time.Minute),
			Stop:         baseEvent.Stop.Add(23 * time.Minute),
			Description:  baseEvent.Description,
			UserID:       baseEvent.UserID,
			Notification: baseEvent.Notification,
		}

		err = appService.UpdateEvent(ctx, id, updatedEvent)
		require.NoError(t, err)

		events, err := appService.ListDay(ctx, initDate)
		require.NoError(t, err)
		require.Equal(t, 1, len(events))
	})

	t.Run("Delete event", func(t *testing.T) {
		storageService := storage.New(&storeConfig)
		appService := app.New(storageService, sLog)
		ctx := context.Background()

		id, err := appService.CreateEvent(ctx, *baseEvent)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		err = appService.Delete(ctx, id)
		require.NoError(t, err)

		events, err := appService.ListDay(ctx, initDate)
		require.NoError(t, err)
		require.Equal(t, 0, len(events))
	})
}

func TestAppFails(t *testing.T) {
	storeConfig := config.DatabaseConfig{EnableInMemory: true}
	logConfig := config.LoggerConfig{EnableFileLog: false, Level: "INFO"}
	sLog, _ := logger.New(&logConfig)

	notifyDur := 1 * time.Hour
	initDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	baseEvent := &event.Event{
		ID:           1,
		Title:        "test",
		Start:        initDate.Add(1 * time.Hour),
		Stop:         initDate.Add(2 * time.Hour),
		Description:  "description",
		UserID:       1,
		Notification: &notifyDur,
	}

	t.Run("Create event without title", func(t *testing.T) {
		storageService := storage.New(&storeConfig)
		appService := app.New(storageService, sLog)
		ctx := context.Background()

		ev := event.Event{
			ID:           1,
			Title:        "",
			Start:        initDate.Add(1 * time.Hour),
			Stop:         initDate.Add(2 * time.Hour),
			Description:  "description",
			UserID:       1,
			Notification: &notifyDur,
		}

		id, err := appService.CreateEvent(ctx, ev)
		require.Error(t, err)
		require.Equal(t, -1, id)
	})

	t.Run("Create event with same id", func(t *testing.T) {
		storageService := storage.New(&storeConfig)
		appService := app.New(storageService, sLog)
		ctx := context.Background()

		id, err := appService.CreateEvent(ctx, *baseEvent)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		_, err = appService.CreateEvent(ctx, *baseEvent)
		require.Error(t, err, "expected error to add event dupl")
	})

	t.Run("Update non existing event", func(t *testing.T) {
		storageService := storage.New(&storeConfig)
		appService := app.New(storageService, sLog)
		ctx := context.Background()

		id, err := appService.CreateEvent(ctx, *baseEvent)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		err = appService.UpdateEvent(ctx, id+1, *baseEvent)
		require.Error(t, err, "expected error to update non existing event")
	})

	t.Run("Delete non existing event", func(t *testing.T) {
		storageService := storage.New(&storeConfig)
		appService := app.New(storageService, sLog)
		ctx := context.Background()

		id, err := appService.CreateEvent(ctx, *baseEvent)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		err = appService.Delete(ctx, id+1)
		require.Error(t, err, "expected error to delete non existing event")
	})

	t.Run("Incorrect time to update event", func(t *testing.T) {
		storageService := storage.New(&storeConfig)
		appService := app.New(storageService, sLog)
		ctx := context.Background()

		id, err := appService.CreateEvent(ctx, *baseEvent)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		updatedEvent := event.Event{
			ID:           id,
			Title:        "New title",
			Start:        baseEvent.Stop.Add(12 * time.Hour),
			Stop:         baseEvent.Stop,
			Description:  "New description",
			UserID:       1,
			Notification: &notifyDur,
		}

		err = appService.UpdateEvent(ctx, id, updatedEvent)
		require.Error(t, err, "incorrect time to update existing event")
	})

	t.Run("Incorrect time to create event", func(t *testing.T) {
		storageService := storage.New(&storeConfig)
		appService := app.New(storageService, sLog)
		ctx := context.Background()

		baseEvent.Start = baseEvent.Stop.Add(1 * time.Minute)
		_, err := appService.CreateEvent(ctx, *baseEvent)
		require.Error(t, err, "incorrect time to add event")
	})
}
