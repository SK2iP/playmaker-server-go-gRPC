// version 1.6

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: service.proto

package service_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Game_GetPlayerActions_FullMethodName     = "/protos.Game/GetPlayerActions"
	Game_GetCoachActions_FullMethodName      = "/protos.Game/GetCoachActions"
	Game_GetTrainerActions_FullMethodName    = "/protos.Game/GetTrainerActions"
	Game_SendInitMessage_FullMethodName      = "/protos.Game/SendInitMessage"
	Game_SendServerParams_FullMethodName     = "/protos.Game/SendServerParams"
	Game_SendPlayerParams_FullMethodName     = "/protos.Game/SendPlayerParams"
	Game_SendPlayerType_FullMethodName       = "/protos.Game/SendPlayerType"
	Game_Register_FullMethodName             = "/protos.Game/Register"
	Game_SendByeCommand_FullMethodName       = "/protos.Game/SendByeCommand"
	Game_GetBestPlannerAction_FullMethodName = "/protos.Game/GetBestPlannerAction"
)

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameClient interface {
	GetPlayerActions(ctx context.Context, in *State, opts ...grpc.CallOption) (*PlayerActions, error)
	GetCoachActions(ctx context.Context, in *State, opts ...grpc.CallOption) (*CoachActions, error)
	GetTrainerActions(ctx context.Context, in *State, opts ...grpc.CallOption) (*TrainerActions, error)
	SendInitMessage(ctx context.Context, in *InitMessage, opts ...grpc.CallOption) (*Empty, error)
	SendServerParams(ctx context.Context, in *ServerParam, opts ...grpc.CallOption) (*Empty, error)
	SendPlayerParams(ctx context.Context, in *PlayerParam, opts ...grpc.CallOption) (*Empty, error)
	SendPlayerType(ctx context.Context, in *PlayerType, opts ...grpc.CallOption) (*Empty, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	SendByeCommand(ctx context.Context, in *RegisterResponse, opts ...grpc.CallOption) (*Empty, error)
	GetBestPlannerAction(ctx context.Context, in *BestPlannerActionRequest, opts ...grpc.CallOption) (*BestPlannerActionResponse, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) GetPlayerActions(ctx context.Context, in *State, opts ...grpc.CallOption) (*PlayerActions, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlayerActions)
	err := c.cc.Invoke(ctx, Game_GetPlayerActions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GetCoachActions(ctx context.Context, in *State, opts ...grpc.CallOption) (*CoachActions, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CoachActions)
	err := c.cc.Invoke(ctx, Game_GetCoachActions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GetTrainerActions(ctx context.Context, in *State, opts ...grpc.CallOption) (*TrainerActions, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TrainerActions)
	err := c.cc.Invoke(ctx, Game_GetTrainerActions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) SendInitMessage(ctx context.Context, in *InitMessage, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Game_SendInitMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) SendServerParams(ctx context.Context, in *ServerParam, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Game_SendServerParams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) SendPlayerParams(ctx context.Context, in *PlayerParam, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Game_SendPlayerParams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) SendPlayerType(ctx context.Context, in *PlayerType, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Game_SendPlayerType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, Game_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) SendByeCommand(ctx context.Context, in *RegisterResponse, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Game_SendByeCommand_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GetBestPlannerAction(ctx context.Context, in *BestPlannerActionRequest, opts ...grpc.CallOption) (*BestPlannerActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BestPlannerActionResponse)
	err := c.cc.Invoke(ctx, Game_GetBestPlannerAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServer is the server API for Game service.
// All implementations must embed UnimplementedGameServer
// for forward compatibility.
type GameServer interface {
	GetPlayerActions(context.Context, *State) (*PlayerActions, error)
	GetCoachActions(context.Context, *State) (*CoachActions, error)
	GetTrainerActions(context.Context, *State) (*TrainerActions, error)
	SendInitMessage(context.Context, *InitMessage) (*Empty, error)
	SendServerParams(context.Context, *ServerParam) (*Empty, error)
	SendPlayerParams(context.Context, *PlayerParam) (*Empty, error)
	SendPlayerType(context.Context, *PlayerType) (*Empty, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	SendByeCommand(context.Context, *RegisterResponse) (*Empty, error)
	GetBestPlannerAction(context.Context, *BestPlannerActionRequest) (*BestPlannerActionResponse, error)
	mustEmbedUnimplementedGameServer()
}

// UnimplementedGameServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGameServer struct{}

func (UnimplementedGameServer) GetPlayerActions(context.Context, *State) (*PlayerActions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerActions not implemented")
}
func (UnimplementedGameServer) GetCoachActions(context.Context, *State) (*CoachActions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoachActions not implemented")
}
func (UnimplementedGameServer) GetTrainerActions(context.Context, *State) (*TrainerActions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrainerActions not implemented")
}
func (UnimplementedGameServer) SendInitMessage(context.Context, *InitMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendInitMessage not implemented")
}
func (UnimplementedGameServer) SendServerParams(context.Context, *ServerParam) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendServerParams not implemented")
}
func (UnimplementedGameServer) SendPlayerParams(context.Context, *PlayerParam) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPlayerParams not implemented")
}
func (UnimplementedGameServer) SendPlayerType(context.Context, *PlayerType) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPlayerType not implemented")
}
func (UnimplementedGameServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedGameServer) SendByeCommand(context.Context, *RegisterResponse) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendByeCommand not implemented")
}
func (UnimplementedGameServer) GetBestPlannerAction(context.Context, *BestPlannerActionRequest) (*BestPlannerActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBestPlannerAction not implemented")
}
func (UnimplementedGameServer) mustEmbedUnimplementedGameServer() {}
func (UnimplementedGameServer) testEmbeddedByValue()              {}

// UnsafeGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServer will
// result in compilation errors.
type UnsafeGameServer interface {
	mustEmbedUnimplementedGameServer()
}

func RegisterGameServer(s grpc.ServiceRegistrar, srv GameServer) {
	// If the following call pancis, it indicates UnimplementedGameServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Game_ServiceDesc, srv)
}

func _Game_GetPlayerActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetPlayerActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_GetPlayerActions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetPlayerActions(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GetCoachActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetCoachActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_GetCoachActions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetCoachActions(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GetTrainerActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetTrainerActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_GetTrainerActions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetTrainerActions(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_SendInitMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).SendInitMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_SendInitMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).SendInitMessage(ctx, req.(*InitMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_SendServerParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).SendServerParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_SendServerParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).SendServerParams(ctx, req.(*ServerParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_SendPlayerParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).SendPlayerParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_SendPlayerParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).SendPlayerParams(ctx, req.(*PlayerParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_SendPlayerType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).SendPlayerType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_SendPlayerType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).SendPlayerType(ctx, req.(*PlayerType))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_SendByeCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).SendByeCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_SendByeCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).SendByeCommand(ctx, req.(*RegisterResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GetBestPlannerAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BestPlannerActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetBestPlannerAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_GetBestPlannerAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetBestPlannerAction(ctx, req.(*BestPlannerActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Game_ServiceDesc is the grpc.ServiceDesc for Game service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Game_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayerActions",
			Handler:    _Game_GetPlayerActions_Handler,
		},
		{
			MethodName: "GetCoachActions",
			Handler:    _Game_GetCoachActions_Handler,
		},
		{
			MethodName: "GetTrainerActions",
			Handler:    _Game_GetTrainerActions_Handler,
		},
		{
			MethodName: "SendInitMessage",
			Handler:    _Game_SendInitMessage_Handler,
		},
		{
			MethodName: "SendServerParams",
			Handler:    _Game_SendServerParams_Handler,
		},
		{
			MethodName: "SendPlayerParams",
			Handler:    _Game_SendPlayerParams_Handler,
		},
		{
			MethodName: "SendPlayerType",
			Handler:    _Game_SendPlayerType_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Game_Register_Handler,
		},
		{
			MethodName: "SendByeCommand",
			Handler:    _Game_SendByeCommand_Handler,
		},
		{
			MethodName: "GetBestPlannerAction",
			Handler:    _Game_GetBestPlannerAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}