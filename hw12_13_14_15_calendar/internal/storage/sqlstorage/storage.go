package sqlstorage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage/event"
)

type SQLStorage struct {
	db *sql.DB
}

func New() SQLStorage {
	return SQLStorage{}
}

func (s *SQLStorage) Connect(ctx context.Context, connect string) error {
	db, err := sql.Open("pgx", connect)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	s.db = db
	return s.db.PingContext(ctx)
}

func (s *SQLStorage) Close(_ context.Context) error {
	return s.db.Close()
}

func (s *SQLStorage) Create(ctx context.Context, ev event.Event) (int, error) {
	query := `
		INSERT INTO event (title, start, stop, description, user_id, notification)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING event_id;
	`

	args := []interface{}{ev.Title, ev.Start, ev.Stop, ev.Description, ev.UserID, ev.Notification}
	var id int
	err := s.db.QueryRowContext(ctx, query, args...).Scan(id)
	if err != nil {
		return 0, fmt.Errorf("db exec: %w", err)
	}
	return id, nil
}

func (s *SQLStorage) Update(ctx context.Context, id int, ev event.Event) error {
	query := `
		UPDATE event
		SET title = $1,
		    start = $2,
		    stop = $3,
		    description = $4,
		    notification = $5
		WHERE event_id = $6;
	`

	args := []interface{}{ev.Title, ev.Start, ev.Stop, ev.Description, ev.Notification, id}
	result, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		_ = fmt.Errorf("exec: %w", err)
		return err
	}

	if count, err := result.RowsAffected(); err != nil || count < 1 {
		_ = fmt.Errorf("affected: %w", err)
		return err
	}

	return nil
}

func (s *SQLStorage) Delete(ctx context.Context, id int) error {
	query := `
		DELETE FROM event WHERE event_id = $1;
	`

	args := []interface{}{id}
	if _, err := s.db.ExecContext(ctx, query, args...); err != nil {
		_ = fmt.Errorf("exec: %w", err)
		return err
	}
	return nil
}

func (s *SQLStorage) DeleteAll(ctx context.Context) error {
	query := `
		TRUNCATE TABLE event RESTART IDENTITY;
	`

	args := []interface{}{}
	if _, err := s.db.ExecContext(ctx, query, args...); err != nil {
		_ = fmt.Errorf("exec: %w", err)
		return err
	}
	return nil
}

func (s *SQLStorage) ListAll(ctx context.Context) ([]event.Event, error) {
	query := `
		SELECT *
		FROM event
		ORDER BY start;
	`

	return s.extractList(ctx, query)
}

func (s *SQLStorage) ListDay(ctx context.Context, date time.Time) ([]event.Event, error) {
	year, month, day := date.Date()

	query := `
		SELECT event_id, title, start, stop, description, user_id, notification
		FROM event
		WHERE extract(year from start) = $1
			AND extract(month from start) = $2
		    AND extract(day from start) = $3
		ORDER BY start;
	`

	return s.extractList(ctx, query, year, month, day)
}

func (s *SQLStorage) ListWeek(ctx context.Context, date time.Time) ([]event.Event, error) {
	year, week := date.ISOWeek()

	query := `
		SELECT event_id, title, start, stop, description, user_id, notification
		FROM event
		WHERE extract(isoyear from start) = $1
	     AND extract(week from start) = $2
		ORDER BY start;
	`

	return s.extractList(ctx, query, year, week)
}

func (s *SQLStorage) ListMonth(ctx context.Context, date time.Time) ([]event.Event, error) {
	year, month, _ := date.Date()

	query := `
		SELECT event_id, title, start, stop, description, user_id, notification
		FROM event
		WHERE extract(year from start) = $1
		  	AND extract(month from start) = $2
		ORDER BY start;
	`

	return s.extractList(ctx, query, year, month)
}

func (s *SQLStorage) GetEventsByNotifier(ctx context.Context, start time.Time, end time.Time) ([]event.Event, error) {
	query := `
		SELECT event_id, title, start, stop, description, user_id, notification
		FROM event
		WHERE extract(day FROM notification) > 0
		  AND (start - (interval '1' day * notification))>=$1
		  AND (start - (interval '1' day * notification))<=$2
	`

	return s.extractList(ctx, query, start, end)
}

func (s *SQLStorage) RemoveAfter(_ context.Context, time time.Time) error {
	query := `
		DELETE FROM Events WHERE EXTRACT(day FROM start) < $1
	`

	_, err := s.db.Exec(query, time)
	return err
}

func (s *SQLStorage) IsTimeBusy(ctx context.Context, ev event.Event) (bool, error) {
	query := `
		SELECT Count(*) AS count
		FROM event
		WHERE user_id = $1
		  	AND event_id != $2;
	`

	var count int
	err := s.db.QueryRowContext(ctx, query, ev.UserID, ev.ID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("db query: %w", err)
	}

	return count > 0, nil
}

func (s *SQLStorage) extractList(
	ctx context.Context,
	query string,
	args ...interface{},
) (events []event.Event, errEvent error) {
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("db query: %w", err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			errEvent = err
		}
	}()

	for rows.Next() {
		var ev event.Event
		var notification sql.NullInt64
		err := rows.Scan(
			&ev.ID,
			&ev.Title,
			&ev.Start,
			&ev.Stop,
			&ev.Description,
			&ev.UserID,
			&notification,
		)
		if err != nil {
			errEvent = fmt.Errorf("db scan: %w", err)
			return
		}

		if notification.Valid {
			ev.Notification = (*time.Duration)(&notification.Int64)
		}

		events = append(events, ev)
	}

	if err := rows.Err(); err != nil {
		errEvent = fmt.Errorf("db rows: %w", err)
		return
	}

	return
}
