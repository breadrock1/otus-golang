package memcache

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/cmd/config"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage/event"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
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

	t.Run("Create new event", func(t *testing.T) {
		storeConfig := config.DatabaseConfig{EnableInMemory: true}
		storeService := storage.New(&storeConfig)
		ctx := context.Background()

		id, err := storeService.Create(ctx, *baseEvent)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		events, err := storeService.ListDay(ctx, initDate)
		require.NoError(t, err)
		require.Equal(t, 1, len(events))
	})

	t.Run("Update event", func(t *testing.T) {
		storeConfig := config.DatabaseConfig{EnableInMemory: true}
		storeService := storage.New(&storeConfig)
		ctx := context.Background()

		id, err := storeService.Create(ctx, *baseEvent)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		updatedEvent := event.Event{
			ID:           id + 1,
			Title:        "New title",
			Start:        baseEvent.Stop.Add(12 * time.Minute),
			Stop:         baseEvent.Stop.Add(23 * time.Minute),
			Description:  "New description",
			UserID:       1,
			Notification: &notifyDur,
		}

		err = storeService.Update(ctx, id, updatedEvent)
		require.NoError(t, err)

		events, err := storeService.ListDay(ctx, initDate)
		require.NoError(t, err)
		require.Equal(t, 1, len(events))
	})

	t.Run("Delete event", func(t *testing.T) {
		storeConfig := config.DatabaseConfig{EnableInMemory: true}
		storeService := storage.New(&storeConfig)
		ctx := context.Background()

		id, err := storeService.Create(ctx, *baseEvent)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		err = storeService.Delete(ctx, id)
		require.NoError(t, err)

		events, err := storeService.ListDay(ctx, initDate)
		require.NoError(t, err)
		require.Equal(t, 0, len(events))
	})

	t.Run("List of events", func(t *testing.T) {
		storeConfig := config.DatabaseConfig{EnableInMemory: true}
		storeService := storage.New(&storeConfig)
		ctx := context.Background()

		for i := 1; i < 100; i++ {
			baseEvent.ID = i
			baseEvent.Start = baseEvent.Start.AddDate(0, 0, 1)
			baseEvent.Stop = baseEvent.Stop.AddDate(0, 0, 1)

			id, err := storeService.Create(ctx, *baseEvent)
			require.NoError(t, err)
			require.NotEmpty(t, id)
		}

		checkDate := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)

		list, err := storeService.ListDay(ctx, checkDate)
		require.NoError(t, err)
		require.Equal(t, len(list), 1)

		checkDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
		list, err = storeService.ListWeek(ctx, checkDate)
		require.NoError(t, err)
		require.Equal(t, len(list), 6)

		list, err = storeService.ListMonth(ctx, checkDate)
		require.NoError(t, err)
		require.Equal(t, len(list), 30)

		offsetDate := initDate.AddDate(0, 1, 0)
		list, err = storeService.ListMonth(ctx, offsetDate)
		require.NoError(t, err)
		require.Equal(t, len(list), 28)
	})
}

func TestStorageErrors(t *testing.T) {
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

	t.Run("Add event with same id", func(t *testing.T) {
		storeConfig := config.DatabaseConfig{EnableInMemory: true}
		storeService := storage.New(&storeConfig)
		ctx := context.Background()

		id, err := storeService.Create(ctx, *baseEvent)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		_, err = storeService.Create(ctx, *baseEvent)
		require.Error(t, err, "expected error to add event dupl")
	})

	t.Run("Update non existing event", func(t *testing.T) {
		storeConfig := config.DatabaseConfig{EnableInMemory: true}
		storeService := storage.New(&storeConfig)
		ctx := context.Background()

		id, err := storeService.Create(ctx, *baseEvent)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		err = storeService.Update(ctx, id+1, *baseEvent)
		require.Error(t, err, "expected error to update non existing event")
	})

	t.Run("Delete non existing event", func(t *testing.T) {
		storeConfig := config.DatabaseConfig{EnableInMemory: true}
		storeService := storage.New(&storeConfig)
		ctx := context.Background()

		id, err := storeService.Create(ctx, *baseEvent)
		require.NoError(t, err)
		require.NotEmpty(t, id)

		err = storeService.Delete(ctx, id+1)
		require.Error(t, err, "expected error to delete non existing event")
	})

	t.Run("Incorrect time to update event", func(t *testing.T) {
		storeConfig := config.DatabaseConfig{EnableInMemory: true}
		storeService := storage.New(&storeConfig)
		ctx := context.Background()

		fmt.Printf("%v", baseEvent)
		id, err := storeService.Create(ctx, *baseEvent)
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

		err = storeService.Update(ctx, id, updatedEvent)
		require.Error(t, err, "incorrect time to update existing event")
	})

	t.Run("Incorrect time to create event", func(t *testing.T) {
		storeConfig := config.DatabaseConfig{EnableInMemory: true}
		storeService := storage.New(&storeConfig)
		ctx := context.Background()

		baseEvent.Start = baseEvent.Stop.Add(1 * time.Minute)
		_, err := storeService.Create(ctx, *baseEvent)
		require.Error(t, err, "incorrect time to add event")
	})
}

func TestStorageConcurrency(t *testing.T) {
	notifyDur := 1 * time.Hour
	initDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	t.Run("Add events by goroutines", func(t *testing.T) {
		storeConfig := config.DatabaseConfig{EnableInMemory: true}
		storeService := storage.New(&storeConfig)
		ctx := context.Background()

		var goCalls []func()
		for i := 0; i < 100; i++ {
			i := i
			test := func() {
				ev := event.Event{
					ID:           i,
					Title:        fmt.Sprintf("Title for event: %d", i),
					Start:        initDate.Add(1 * time.Hour),
					Stop:         initDate.Add(2 * time.Hour),
					Description:  "description",
					UserID:       1,
					Notification: &notifyDur,
				}

				_, err := storeService.Create(ctx, ev)
				require.NoError(t, err)
			}
			goCalls = append(goCalls, test)
		}

		wg := &sync.WaitGroup{}
		for _, goCall := range goCalls {
			wg.Add(1)
			go func() {
				defer wg.Done()
				goCall()
			}()
		}

		wg.Wait()

		allEvents, err := storeService.ListAll(ctx)
		require.NoError(t, err, "failed to get all events")
		require.Equal(t, 100, len(allEvents))
	})
}
