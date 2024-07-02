// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: chat.proto

// protobuf config

package chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Chat_SendMsgTest_FullMethodName = "/kyle_chat.Chat/SendMsgTest"
	Chat_SendMsg_FullMethodName     = "/kyle_chat.Chat/SendMsg"
	Chat_GetMsgTest_FullMethodName  = "/kyle_chat.Chat/GetMsgTest"
	Chat_GetMsg_FullMethodName      = "/kyle_chat.Chat/GetMsg"
)

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	SendMsgTest(ctx context.Context, in *TestMsg, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	SendMsg(ctx context.Context, in *ChatMsg, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	GetMsgTest(ctx context.Context, in *GetTestMsg, opts ...grpc.CallOption) (*TestMsg, error)
	GetMsg(ctx context.Context, in *GetChatMsg, opts ...grpc.CallOption) (*ChatMsg, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) SendMsgTest(ctx context.Context, in *TestMsg, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, Chat_SendMsgTest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SendMsg(ctx context.Context, in *ChatMsg, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, Chat_SendMsg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetMsgTest(ctx context.Context, in *GetTestMsg, opts ...grpc.CallOption) (*TestMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestMsg)
	err := c.cc.Invoke(ctx, Chat_GetMsgTest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetMsg(ctx context.Context, in *GetChatMsg, opts ...grpc.CallOption) (*ChatMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatMsg)
	err := c.cc.Invoke(ctx, Chat_GetMsg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations should embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	SendMsgTest(context.Context, *TestMsg) (*wrapperspb.StringValue, error)
	SendMsg(context.Context, *ChatMsg) (*wrapperspb.StringValue, error)
	GetMsgTest(context.Context, *GetTestMsg) (*TestMsg, error)
	GetMsg(context.Context, *GetChatMsg) (*ChatMsg, error)
}

// UnimplementedChatServer should be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) SendMsgTest(context.Context, *TestMsg) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsgTest not implemented")
}
func (UnimplementedChatServer) SendMsg(context.Context, *ChatMsg) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (UnimplementedChatServer) GetMsgTest(context.Context, *GetTestMsg) (*TestMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMsgTest not implemented")
}
func (UnimplementedChatServer) GetMsg(context.Context, *GetChatMsg) (*ChatMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMsg not implemented")
}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_SendMsgTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SendMsgTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_SendMsgTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SendMsgTest(ctx, req.(*TestMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_SendMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SendMsg(ctx, req.(*ChatMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetMsgTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetMsgTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_GetMsgTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetMsgTest(ctx, req.(*GetTestMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_GetMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetMsg(ctx, req.(*GetChatMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kyle_chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsgTest",
			Handler:    _Chat_SendMsgTest_Handler,
		},
		{
			MethodName: "SendMsg",
			Handler:    _Chat_SendMsg_Handler,
		},
		{
			MethodName: "GetMsgTest",
			Handler:    _Chat_GetMsgTest_Handler,
		},
		{
			MethodName: "GetMsg",
			Handler:    _Chat_GetMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}