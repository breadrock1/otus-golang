package memcache

import (
	"context"
	"fmt"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage"
	"sort"
	"sync"
	"time"
)

type memStorage struct {
	mu     *sync.Mutex
	lastID int
	cache  map[int]*storage.Event
}

func New() storage.Storage {
	store := memStorage{
		mu:     &sync.Mutex{},
		lastID: 0,
		cache:  make(map[int]*storage.Event),
	}

	return &store
}

func (s *memStorage) Connect(_ context.Context, _ string) error {
	return nil
}

func (s *memStorage) Close(_ context.Context) error {
	return nil
}

func (s *memStorage) Create(_ context.Context, event storage.Event) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.lastID + 1
	s.cache[id] = &storage.Event{
		ID:           id,
		Title:        event.Title,
		Start:        event.Start,
		Stop:         event.Stop,
		Description:  event.Description,
		UserID:       event.UserID,
		Notification: event.Notification,
	}

	return id, nil
}

func (s *memStorage) Update(_ context.Context, id int, newEvent storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.cache[id]; !ok {
		return fmt.Errorf("non exist: %d", id)
	}

	newEvent.ID = id
	s.cache[id] = &newEvent
	return nil
}

func (s *memStorage) Delete(_ context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.cache, id)
	return nil
}

func (s *memStorage) DeleteAll(_ context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cache = make(map[int]*storage.Event)
	return nil
}

func (s *memStorage) ListAll(_ context.Context) ([]storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]storage.Event, 0, len(s.cache))
	for _, event := range s.cache {
		result = append(result, *event)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Start.Before(result[j].Start)
	})

	return result, nil
}

func (s *memStorage) ListDay(_ context.Context, date time.Time) ([]storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []storage.Event
	year, month, day := date.Date()
	for _, event := range s.cache {
		eventYear, eventMonth, eventDay := event.Start.Date()
		if eventYear == year && eventMonth == month && eventDay == day {
			result = append(result, *event)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Start.Before(result[j].Start)
	})

	return result, nil
}

func (s *memStorage) ListWeek(_ context.Context, date time.Time) ([]storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []storage.Event
	year, week := date.ISOWeek()
	for _, event := range s.cache {
		eventYear, eventWeek := event.Start.ISOWeek()
		if eventYear == year && eventWeek == week {
			result = append(result, *event)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Start.Before(result[j].Start)
	})

	return result, nil
}

func (s *memStorage) ListMonth(_ context.Context, date time.Time) ([]storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []storage.Event
	year, month, _ := date.Date()
	for _, event := range s.cache {
		eventYear, eventMonth, _ := event.Start.Date()
		if eventYear == year && eventMonth == month {
			result = append(result, *event)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Start.Before(result[j].Start)
	})

	return result, nil
}

func (s *memStorage) IsTimeBusy(_ context.Context, userID int, start, stop time.Time, excludeID int) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, event := range s.cache {
		filterByUser := event.UserID == userID
		isExcluded := event.ID != excludeID
		isStopAfter := event.Stop.After(start)
		isStartBefore := event.Start.Before(stop)
		if filterByUser && isExcluded && isStopAfter && isStartBefore {
			return true, nil
		}
	}

	return false, nil
}
