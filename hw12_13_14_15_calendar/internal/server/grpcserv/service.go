package grpcserv

import (
	"context"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/app"
	"github.com/breadrock1/otus-golang/hw12_13_14_15_calendar/internal/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type Service struct {
	UnimplementedCalendarServer
	app app.App
}

func NewService(app app.App) Service {
	return Service{app: app}
}

func (s Service) Create(ctx context.Context, event *Event) (*CreateResponse, error) {
	storageEvent := s.intoEvent(event)
	id, err := s.app.CreateEvent(ctx, storageEvent)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &CreateResponse{Id: int32(id)}, nil
}

func (s Service) Update(ctx context.Context, event *Event) (*UpdateResponse, error) {
	storageEvent := s.intoEvent(event)
	if err := s.app.UpdateEvent(ctx, int(event.Id), storageEvent); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &UpdateResponse{}, nil
}

func (s Service) Delete(ctx context.Context, event *DeleteEvent) (*DeleteResponse, error) {
	if err := s.app.Delete(ctx, int(event.Id)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &DeleteResponse{}, nil
}

func (s Service) ListPerDay(ctx context.Context, datetime *ListPerDatetime) (*ListEventsResponse, error) {
	events, err := s.app.ListDay(ctx, datetime.Datetime.AsTime())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &ListEventsResponse{Events: s.fromEvent(events)}, nil
}

func (s Service) ListPerWeek(ctx context.Context, datetime *ListPerDatetime) (*ListEventsResponse, error) {
	events, err := s.app.ListWeek(ctx, datetime.Datetime.AsTime())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &ListEventsResponse{Events: s.fromEvent(events)}, nil
}

func (s Service) ListPerMonth(ctx context.Context, datetime *ListPerDatetime) (*ListEventsResponse, error) {
	events, err := s.app.ListMonth(ctx, datetime.Datetime.AsTime())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &ListEventsResponse{Events: s.fromEvent(events)}, nil
}

func (s Service) mustEmbedUnimplementedCalendarServer() {
	log.Println("Unimplemented")
}

func (s Service) intoEvent(event *Event) storage.Event {
	notification := event.Notification.AsDuration()
	return storage.Event{
		ID:           int(event.Id),
		Title:        event.Title,
		Start:        event.Start.AsTime(),
		Stop:         event.Stop.AsTime(),
		Description:  event.Description,
		UserID:       int(event.UserId),
		Notification: &notification,
	}
}

func (s Service) fromEvent(events []storage.Event) []*Event {
	allEvents := make([]*Event, 0)
	for _, event := range events {
		grpcEvent := &Event{
			Id:           int32(event.ID),
			UserId:       int32(event.UserID),
			Notification: durationpb.New(*event.Notification),
			Start:        timestamppb.New(event.Start),
			Stop:         timestamppb.New(event.Stop),
			Title:        event.Title,
			Description:  event.Description,
		}
		allEvents = append(allEvents, grpcEvent)
	}
	return allEvents
}
