package sqlstorage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage"
)

type sqlStorage struct {
	db *sql.DB
}

func New() storage.Storage {
	return &sqlStorage{}
}

func (s *sqlStorage) Connect(ctx context.Context, connect string) error {
	db, err := sql.Open("pgx", connect)
	if err != nil {
		log.Fatalln("djgbdfg")
		return err
	}
	s.db = db
	return s.db.PingContext(ctx)
}

func (s *sqlStorage) Close(_ context.Context) error {
	return s.db.Close()
}

func (s *sqlStorage) Create(ctx context.Context, event storage.Event) (int, error) {
	var query string
	var args []interface{}

	//query = `
	//	INSERT INTO event (title, start, stop, description, user_id, notification)
	//	VALUES($1, $2, $3, $4, $5, $6)
	//	RETURNING event_id;
	//`

	args = []interface{}{event.Title, event.Start, event.Stop, event.Description, event.UserID, event.Notification}
	var id int
	err := s.db.QueryRowContext(ctx, query, args...).Scan(id)
	if err != nil {
		return 0, fmt.Errorf("db exec: %w", err)
	}
	return id, nil
}

func (s *sqlStorage) Update(ctx context.Context, id int, event storage.Event) error {
	var query string
	var args []interface{}

	//query = `
	//	UPDATE event
	//	SET title = $1,
	//	    start = $2,
	//	    stop = $3,
	//	    description = $4,
	//	    notification = $5
	//	WHERE event_id = $6;
	//`

	args = []interface{}{event.Title, event.Start, event.Stop, event.Description, event.Notification, id}
	result, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		_ = fmt.Errorf("exec: %s", err)
		return err
	}

	if count, err := result.RowsAffected(); err != nil || count < 1 {
		_ = fmt.Errorf("affected: %s", err)
		return err
	}

	return nil
}

func (s *sqlStorage) Delete(ctx context.Context, id int) error {
	var query string
	var args []interface{}

	//query = `
	//	DELETE FROM event WHERE event_id = $1;
	//`

	args = []interface{}{id}
	if _, err := s.db.ExecContext(ctx, query, args...); err != nil {
		_ = fmt.Errorf("exec: %s", err)
		return err
	}
	return nil
}

func (s *sqlStorage) DeleteAll(ctx context.Context) error {
	var query string
	var args []interface{}

	//query = `
	//	TRUNCATE TABLE event RESTART IDENTITY;
	//`

	args = []interface{}{}
	if _, err := s.db.ExecContext(ctx, query, args...); err != nil {
		_ = fmt.Errorf("exec: %s", err)
		return err
	}
	return nil
}

func (s *sqlStorage) ListAll(ctx context.Context) ([]storage.Event, error) {
	var query string

	//query = `
	//	SELECT *
	//	FROM event
	//	ORDER BY start;
	//`

	return s.extractList(ctx, query)
}

func (s *sqlStorage) ListDay(ctx context.Context, date time.Time) ([]storage.Event, error) {
	var query string
	year, month, day := date.Date()

	//query := `
	//	SELECT event_id, title, start, stop, description, user_id, notification
	//	FROM event
	//	WHERE extract(year from start) = $1
	//		AND extract(month from start) = $2
	//	    AND extract(day from start) = $3
	//	ORDER BY start;
	//`

	return s.extractList(ctx, query, year, month, day)
}

func (s *sqlStorage) ListWeek(ctx context.Context, date time.Time) ([]storage.Event, error) {
	var query string
	year, week := date.ISOWeek()

	//query := `
	//	SELECT event_id, title, start, stop, description, user_id, notification
	//	FROM event
	//	WHERE extract(isoyear from start) = $1
	//      AND extract(week from start) = $2
	//	ORDER BY start;
	//`

	return s.extractList(ctx, query, year, week)
}

func (s *sqlStorage) ListMonth(ctx context.Context, date time.Time) ([]storage.Event, error) {
	var query string
	year, month, _ := date.Date()

	//query := `
	//	SELECT event_id, title, start, stop, description, user_id, notification
	//	FROM event
	//	WHERE extract(year from start) = $1
	//	  	AND extract(month from start) = $2
	//	ORDER BY start;
	//`

	return s.extractList(ctx, query, year, month)
}

func (s *sqlStorage) IsTimeBusy(ctx context.Context, userID int, start, stop time.Time, excludeID int) (bool, error) {
	var query string

	//query := `
	//	SELECT Count(*) AS count
	//	FROM event
	//	WHERE user_id = $1
	//		AND start < $2
	//	   	AND stop > $3
	//	  	AND event_id != $4;
	//`

	var count int
	err := s.db.QueryRowContext(ctx, query, userID, stop, start, excludeID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("db query: %w", err)
	}

	return count > 0, nil
}

func (s *sqlStorage) extractList(ctx context.Context, query string, args ...interface{}) (resultEvent []storage.Event, errEvent error) {
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
		var event storage.Event
		var notification sql.NullInt64
		err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Start,
			&event.Stop,
			&event.Description,
			&event.UserID,
			&notification,
		)

		if err != nil {
			errEvent = fmt.Errorf("db scan: %w", err)
			return
		}

		if notification.Valid {
			event.Notification = (*time.Duration)(&notification.Int64)
		}

		resultEvent = append(resultEvent, event)
	}

	if err := rows.Err(); err != nil {
		errEvent = fmt.Errorf("db rows: %w", err)
		return
	}

	return
}
