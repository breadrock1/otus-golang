package memcache

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage/event"
)

type MemStorage struct {
	mu     *sync.Mutex
	lastID int
	cache  map[int]*event.Event
}

func New() MemStorage {
	store := MemStorage{
		mu:     &sync.Mutex{},
		lastID: 0,
		cache:  make(map[int]*event.Event),
	}

	return store
}

func (s *MemStorage) Connect(_ context.Context, _ string) error {
	return nil
}

func (s *MemStorage) Close(_ context.Context) error {
	return nil
}

func (s *MemStorage) Create(_ context.Context, ev event.Event) (int, error) {
	if !s.isTimeCorrect(&ev) {
		return -1, errors.New("time not correct")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if ev.ID == -1 {
		ev.ID = len(s.cache)
	}

	if _, ok := s.cache[ev.ID]; ok && ev.ID != 0 {
		msg := fmt.Sprintf("event %d already exists", ev.ID)
		return ev.ID, errors.New(msg)
	}

	s.cache[ev.ID] = &ev

	return ev.ID, nil
}

func (s *MemStorage) Update(_ context.Context, id int, newEvent event.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isTimeCorrect(&newEvent) {
		return errors.New("time not correct")
	}

	if _, ok := s.cache[id]; !ok {
		return fmt.Errorf("non exist: %d", id)
	}

	newEvent.ID = id
	s.cache[id] = &newEvent
	return nil
}

func (s *MemStorage) Delete(_ context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.cache[id]; !ok {
		return errors.New("event with passed id does not exist")
	}

	delete(s.cache, id)
	return nil
}

func (s *MemStorage) DeleteAll(_ context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cache = make(map[int]*event.Event)
	return nil
}

func (s *MemStorage) ListAll(_ context.Context) ([]event.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	events := make([]event.Event, 0)
	for _, ev := range s.cache {
		events = append(events, *ev)
	}

	return events, nil
}

func (s *MemStorage) ListDay(_ context.Context, date time.Time) ([]event.Event, error) {
	start := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	stop := start.Add(24 * time.Hour)
	return s.consumeFilter(start, stop), nil
}

func (s *MemStorage) ListWeek(_ context.Context, date time.Time) ([]event.Event, error) {
	start := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	stop := start.AddDate(0, 0, 7)
	return s.consumeFilter(start, stop), nil
}

func (s *MemStorage) ListMonth(_ context.Context, date time.Time) ([]event.Event, error) {
	start := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	stop := start.AddDate(0, 1, 0)
	return s.consumeFilter(start, stop), nil
}

func (s *MemStorage) GetEventsByNotifier(_ context.Context, start time.Time, end time.Time) ([]event.Event, error) {
	events := make([]event.Event, 0)
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, ev := range s.cache {
		notifyHours := ev.Notification.Hours()
		notifyTime := ev.Start.Add(time.Hour * time.Duration(notifyHours))
		if notifyHours > 0 && notifyTime.After(start) && notifyTime.Before(end) {
			events = append(events, *ev)
		}
	}

	return events, nil
}

func (s *MemStorage) RemoveAfter(_ context.Context, time time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for k, ev := range s.cache {
		if ev.Start.After(time) {
			delete(s.cache, k)
		}
	}

	return nil
}

func (s *MemStorage) IsTimeBusy(_ context.Context, ev event.Event) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, currEvent := range s.cache {
		isSameEvent := currEvent.ID == ev.ID
		isSameUser := currEvent.UserID == ev.UserID
		if isSameEvent && isSameUser {
			return true, nil
		}
	}

	return false, nil
}

func (s *MemStorage) consumeFilter(start, stop time.Time) []event.Event {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]event.Event, 0)
	for _, ev := range s.cache {
		isEqualStart := ev.Start.Equal(start)
		isRangeTrue := ev.Start.After(start) && ev.Start.Before(stop)
		if isEqualStart || isRangeTrue {
			result = append(result, *ev)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Start.Before(result[j].Start)
	})

	return result
}

func (s *MemStorage) isTimeCorrect(ev *event.Event) bool {
	if ev.Start.After(ev.Stop) {
		return false
	}

	if ev.Start.Before(time.Now()) {
		return false
	}

	return true
}
