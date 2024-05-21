// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: calendar.proto

package grpcserv

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalendarClient interface {
	Create(ctx context.Context, in *Event, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *Event, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteEvent, opts ...grpc.CallOption) (*DeleteResponse, error)
	ListPerDay(ctx context.Context, in *ListPerDatetime, opts ...grpc.CallOption) (*ListEventsResponse, error)
	ListPerWeek(ctx context.Context, in *ListPerDatetime, opts ...grpc.CallOption) (*ListEventsResponse, error)
	ListPerMonth(ctx context.Context, in *ListPerDatetime, opts ...grpc.CallOption) (*ListEventsResponse, error)
}

type calendarClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarClient(cc grpc.ClientConnInterface) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) Create(ctx context.Context, in *Event, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) Update(ctx context.Context, in *Event, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) Delete(ctx context.Context, in *DeleteEvent, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) ListPerDay(ctx context.Context, in *ListPerDatetime, opts ...grpc.CallOption) (*ListEventsResponse, error) {
	out := new(ListEventsResponse)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/ListPerDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) ListPerWeek(ctx context.Context, in *ListPerDatetime, opts ...grpc.CallOption) (*ListEventsResponse, error) {
	out := new(ListEventsResponse)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/ListPerWeek", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) ListPerMonth(ctx context.Context, in *ListPerDatetime, opts ...grpc.CallOption) (*ListEventsResponse, error) {
	out := new(ListEventsResponse)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/ListPerMonth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
// All implementations must embed UnimplementedCalendarServer
// for forward compatibility
type CalendarServer interface {
	Create(context.Context, *Event) (*CreateResponse, error)
	Update(context.Context, *Event) (*UpdateResponse, error)
	Delete(context.Context, *DeleteEvent) (*DeleteResponse, error)
	ListPerDay(context.Context, *ListPerDatetime) (*ListEventsResponse, error)
	ListPerWeek(context.Context, *ListPerDatetime) (*ListEventsResponse, error)
	ListPerMonth(context.Context, *ListPerDatetime) (*ListEventsResponse, error)
	mustEmbedUnimplementedCalendarServer()
}

// UnimplementedCalendarServer must be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (UnimplementedCalendarServer) Create(context.Context, *Event) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCalendarServer) Update(context.Context, *Event) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCalendarServer) Delete(context.Context, *DeleteEvent) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCalendarServer) ListPerDay(context.Context, *ListPerDatetime) (*ListEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPerDay not implemented")
}
func (UnimplementedCalendarServer) ListPerWeek(context.Context, *ListPerDatetime) (*ListEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPerWeek not implemented")
}
func (UnimplementedCalendarServer) ListPerMonth(context.Context, *ListPerDatetime) (*ListEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPerMonth not implemented")
}
func (UnimplementedCalendarServer) mustEmbedUnimplementedCalendarServer() {}

// UnsafeCalendarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalendarServer will
// result in compilation errors.
type UnsafeCalendarServer interface {
	mustEmbedUnimplementedCalendarServer()
}

func RegisterCalendarServer(s grpc.ServiceRegistrar, srv CalendarServer) {
	s.RegisterService(&Calendar_ServiceDesc, srv)
}

func _Calendar_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).Create(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).Update(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).Delete(ctx, req.(*DeleteEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_ListPerDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPerDatetime)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).ListPerDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/ListPerDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).ListPerDay(ctx, req.(*ListPerDatetime))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_ListPerWeek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPerDatetime)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).ListPerWeek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/ListPerWeek",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).ListPerWeek(ctx, req.(*ListPerDatetime))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_ListPerMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPerDatetime)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).ListPerMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/ListPerMonth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).ListPerMonth(ctx, req.(*ListPerDatetime))
	}
	return interceptor(ctx, in, info, handler)
}

// Calendar_ServiceDesc is the grpc.ServiceDesc for Calendar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calendar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calendar.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Calendar_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Calendar_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Calendar_Delete_Handler,
		},
		{
			MethodName: "ListPerDay",
			Handler:    _Calendar_ListPerDay_Handler,
		},
		{
			MethodName: "ListPerWeek",
			Handler:    _Calendar_ListPerWeek_Handler,
		},
		{
			MethodName: "ListPerMonth",
			Handler:    _Calendar_ListPerMonth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calendar.proto",
}