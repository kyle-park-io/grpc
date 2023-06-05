// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: order_management.proto

package bidirectionstreaming

import (
	context "context"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderManagementClient is the client API for OrderManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderManagementClient interface {
	ProcessOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_ProcessOrdersClient, error)
}

type orderManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderManagementClient(cc grpc.ClientConnInterface) OrderManagementClient {
	return &orderManagementClient{cc}
}

func (c *orderManagementClient) ProcessOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_ProcessOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderManagement_ServiceDesc.Streams[0], "/ecommerce.OrderManagement/processOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementProcessOrdersClient{stream}
	return x, nil
}

type OrderManagement_ProcessOrdersClient interface {
	Send(*wrappers.StringValue) error
	Recv() (*CombinedShipment, error)
	grpc.ClientStream
}

type orderManagementProcessOrdersClient struct {
	grpc.ClientStream
}

func (x *orderManagementProcessOrdersClient) Send(m *wrappers.StringValue) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderManagementProcessOrdersClient) Recv() (*CombinedShipment, error) {
	m := new(CombinedShipment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderManagementServer is the server API for OrderManagement service.
// All implementations should embed UnimplementedOrderManagementServer
// for forward compatibility
type OrderManagementServer interface {
	ProcessOrders(OrderManagement_ProcessOrdersServer) error
}

// UnimplementedOrderManagementServer should be embedded to have forward compatible implementations.
type UnimplementedOrderManagementServer struct {
}

func (UnimplementedOrderManagementServer) ProcessOrders(OrderManagement_ProcessOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessOrders not implemented")
}

// UnsafeOrderManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderManagementServer will
// result in compilation errors.
type UnsafeOrderManagementServer interface {
	mustEmbedUnimplementedOrderManagementServer()
}

func RegisterOrderManagementServer(s grpc.ServiceRegistrar, srv OrderManagementServer) {
	s.RegisterService(&OrderManagement_ServiceDesc, srv)
}

func _OrderManagement_ProcessOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderManagementServer).ProcessOrders(&orderManagementProcessOrdersServer{stream})
}

type OrderManagement_ProcessOrdersServer interface {
	Send(*CombinedShipment) error
	Recv() (*wrappers.StringValue, error)
	grpc.ServerStream
}

type orderManagementProcessOrdersServer struct {
	grpc.ServerStream
}

func (x *orderManagementProcessOrdersServer) Send(m *CombinedShipment) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderManagementProcessOrdersServer) Recv() (*wrappers.StringValue, error) {
	m := new(wrappers.StringValue)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderManagement_ServiceDesc is the grpc.ServiceDesc for OrderManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.OrderManagement",
	HandlerType: (*OrderManagementServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "processOrders",
			Handler:       _OrderManagement_ProcessOrders_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "order_management.proto",
}
