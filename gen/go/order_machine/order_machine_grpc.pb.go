// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: order_machine/order_machine.proto

package order_machine

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

// OrderMachineServiceClient is the client API for OrderMachineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderMachineServiceClient interface {
	// 创建订单
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	// 订单列表
	ListOrder(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderResponse, error)
	// 订单状态流转
	SwitchOrderState(ctx context.Context, in *SwitchOrderStateRequest, opts ...grpc.CallOption) (*SwitchOrderStateResponse, error)
}

type orderMachineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderMachineServiceClient(cc grpc.ClientConnInterface) OrderMachineServiceClient {
	return &orderMachineServiceClient{cc}
}

func (c *orderMachineServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, "/order_machine.OrderMachineService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderMachineServiceClient) ListOrder(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderResponse, error) {
	out := new(ListOrderResponse)
	err := c.cc.Invoke(ctx, "/order_machine.OrderMachineService/ListOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderMachineServiceClient) SwitchOrderState(ctx context.Context, in *SwitchOrderStateRequest, opts ...grpc.CallOption) (*SwitchOrderStateResponse, error) {
	out := new(SwitchOrderStateResponse)
	err := c.cc.Invoke(ctx, "/order_machine.OrderMachineService/SwitchOrderState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderMachineServiceServer is the server API for OrderMachineService service.
// All implementations must embed UnimplementedOrderMachineServiceServer
// for forward compatibility
type OrderMachineServiceServer interface {
	// 创建订单
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	// 订单列表
	ListOrder(context.Context, *ListOrderRequest) (*ListOrderResponse, error)
	// 订单状态流转
	SwitchOrderState(context.Context, *SwitchOrderStateRequest) (*SwitchOrderStateResponse, error)
	mustEmbedUnimplementedOrderMachineServiceServer()
}

// UnimplementedOrderMachineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderMachineServiceServer struct {
}

func (UnimplementedOrderMachineServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderMachineServiceServer) ListOrder(context.Context, *ListOrderRequest) (*ListOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrder not implemented")
}
func (UnimplementedOrderMachineServiceServer) SwitchOrderState(context.Context, *SwitchOrderStateRequest) (*SwitchOrderStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwitchOrderState not implemented")
}
func (UnimplementedOrderMachineServiceServer) mustEmbedUnimplementedOrderMachineServiceServer() {}

// UnsafeOrderMachineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderMachineServiceServer will
// result in compilation errors.
type UnsafeOrderMachineServiceServer interface {
	mustEmbedUnimplementedOrderMachineServiceServer()
}

func RegisterOrderMachineServiceServer(s grpc.ServiceRegistrar, srv OrderMachineServiceServer) {
	s.RegisterService(&OrderMachineService_ServiceDesc, srv)
}

func _OrderMachineService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderMachineServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_machine.OrderMachineService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderMachineServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderMachineService_ListOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderMachineServiceServer).ListOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_machine.OrderMachineService/ListOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderMachineServiceServer).ListOrder(ctx, req.(*ListOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderMachineService_SwitchOrderState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwitchOrderStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderMachineServiceServer).SwitchOrderState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_machine.OrderMachineService/SwitchOrderState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderMachineServiceServer).SwitchOrderState(ctx, req.(*SwitchOrderStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderMachineService_ServiceDesc is the grpc.ServiceDesc for OrderMachineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderMachineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order_machine.OrderMachineService",
	HandlerType: (*OrderMachineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderMachineService_CreateOrder_Handler,
		},
		{
			MethodName: "ListOrder",
			Handler:    _OrderMachineService_ListOrder_Handler,
		},
		{
			MethodName: "SwitchOrderState",
			Handler:    _OrderMachineService_SwitchOrderState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order_machine/order_machine.proto",
}