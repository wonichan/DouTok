// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.1
// source: trade.proto

package api

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
	TradeService_QueryTradeOrder_FullMethodName        = "/api.TradeService/QueryTradeOrder"
	TradeService_TradeCreate_FullMethodName            = "/api.TradeService/TradeCreate"
	TradeService_MergeTrade_FullMethodName             = "/api.TradeService/MergeTrade"
	TradeService_UpdateExtendInfo_FullMethodName       = "/api.TradeService/UpdateExtendInfo"
	TradeService_UpdateTradeOrderStatus_FullMethodName = "/api.TradeService/UpdateTradeOrderStatus"
	TradeService_UpdateDeliveryInfo_FullMethodName     = "/api.TradeService/UpdateDeliveryInfo"
	TradeService_AddSubOrder_FullMethodName            = "/api.TradeService/AddSubOrder"
)

// TradeServiceClient is the client API for TradeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TradeServiceClient interface {
	QueryTradeOrder(ctx context.Context, in *QueryTradeOrderRequest, opts ...grpc.CallOption) (*QueryTradeOrderResponse, error)
	TradeCreate(ctx context.Context, in *TradeCreateRequest, opts ...grpc.CallOption) (*TradeCreateResponse, error)
	MergeTrade(ctx context.Context, in *MergeTradeRequest, opts ...grpc.CallOption) (*MergeTradeResponse, error)
	UpdateExtendInfo(ctx context.Context, in *UpdateExtendInfoRequest, opts ...grpc.CallOption) (*UpdateExtendInfoResponse, error)
	UpdateTradeOrderStatus(ctx context.Context, in *UpdateTradeOrderStatusRequest, opts ...grpc.CallOption) (*UpdateTradeOrderStatusResponse, error)
	UpdateDeliveryInfo(ctx context.Context, in *UpdateDeliveryInfoRequest, opts ...grpc.CallOption) (*UpdateDeliveryInfoResponse, error)
	AddSubOrder(ctx context.Context, in *AddSubOrderRequest, opts ...grpc.CallOption) (*AddSubOrderResponse, error)
}

type tradeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTradeServiceClient(cc grpc.ClientConnInterface) TradeServiceClient {
	return &tradeServiceClient{cc}
}

func (c *tradeServiceClient) QueryTradeOrder(ctx context.Context, in *QueryTradeOrderRequest, opts ...grpc.CallOption) (*QueryTradeOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryTradeOrderResponse)
	err := c.cc.Invoke(ctx, TradeService_QueryTradeOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) TradeCreate(ctx context.Context, in *TradeCreateRequest, opts ...grpc.CallOption) (*TradeCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TradeCreateResponse)
	err := c.cc.Invoke(ctx, TradeService_TradeCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) MergeTrade(ctx context.Context, in *MergeTradeRequest, opts ...grpc.CallOption) (*MergeTradeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MergeTradeResponse)
	err := c.cc.Invoke(ctx, TradeService_MergeTrade_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) UpdateExtendInfo(ctx context.Context, in *UpdateExtendInfoRequest, opts ...grpc.CallOption) (*UpdateExtendInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateExtendInfoResponse)
	err := c.cc.Invoke(ctx, TradeService_UpdateExtendInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) UpdateTradeOrderStatus(ctx context.Context, in *UpdateTradeOrderStatusRequest, opts ...grpc.CallOption) (*UpdateTradeOrderStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTradeOrderStatusResponse)
	err := c.cc.Invoke(ctx, TradeService_UpdateTradeOrderStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) UpdateDeliveryInfo(ctx context.Context, in *UpdateDeliveryInfoRequest, opts ...grpc.CallOption) (*UpdateDeliveryInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDeliveryInfoResponse)
	err := c.cc.Invoke(ctx, TradeService_UpdateDeliveryInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeServiceClient) AddSubOrder(ctx context.Context, in *AddSubOrderRequest, opts ...grpc.CallOption) (*AddSubOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddSubOrderResponse)
	err := c.cc.Invoke(ctx, TradeService_AddSubOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradeServiceServer is the server API for TradeService service.
// All implementations should embed UnimplementedTradeServiceServer
// for forward compatibility.
type TradeServiceServer interface {
	QueryTradeOrder(context.Context, *QueryTradeOrderRequest) (*QueryTradeOrderResponse, error)
	TradeCreate(context.Context, *TradeCreateRequest) (*TradeCreateResponse, error)
	MergeTrade(context.Context, *MergeTradeRequest) (*MergeTradeResponse, error)
	UpdateExtendInfo(context.Context, *UpdateExtendInfoRequest) (*UpdateExtendInfoResponse, error)
	UpdateTradeOrderStatus(context.Context, *UpdateTradeOrderStatusRequest) (*UpdateTradeOrderStatusResponse, error)
	UpdateDeliveryInfo(context.Context, *UpdateDeliveryInfoRequest) (*UpdateDeliveryInfoResponse, error)
	AddSubOrder(context.Context, *AddSubOrderRequest) (*AddSubOrderResponse, error)
}

// UnimplementedTradeServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTradeServiceServer struct{}

func (UnimplementedTradeServiceServer) QueryTradeOrder(context.Context, *QueryTradeOrderRequest) (*QueryTradeOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTradeOrder not implemented")
}
func (UnimplementedTradeServiceServer) TradeCreate(context.Context, *TradeCreateRequest) (*TradeCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeCreate not implemented")
}
func (UnimplementedTradeServiceServer) MergeTrade(context.Context, *MergeTradeRequest) (*MergeTradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MergeTrade not implemented")
}
func (UnimplementedTradeServiceServer) UpdateExtendInfo(context.Context, *UpdateExtendInfoRequest) (*UpdateExtendInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExtendInfo not implemented")
}
func (UnimplementedTradeServiceServer) UpdateTradeOrderStatus(context.Context, *UpdateTradeOrderStatusRequest) (*UpdateTradeOrderStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTradeOrderStatus not implemented")
}
func (UnimplementedTradeServiceServer) UpdateDeliveryInfo(context.Context, *UpdateDeliveryInfoRequest) (*UpdateDeliveryInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeliveryInfo not implemented")
}
func (UnimplementedTradeServiceServer) AddSubOrder(context.Context, *AddSubOrderRequest) (*AddSubOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubOrder not implemented")
}
func (UnimplementedTradeServiceServer) testEmbeddedByValue() {}

// UnsafeTradeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TradeServiceServer will
// result in compilation errors.
type UnsafeTradeServiceServer interface {
	mustEmbedUnimplementedTradeServiceServer()
}

func RegisterTradeServiceServer(s grpc.ServiceRegistrar, srv TradeServiceServer) {
	// If the following call pancis, it indicates UnimplementedTradeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TradeService_ServiceDesc, srv)
}

func _TradeService_QueryTradeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTradeOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).QueryTradeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeService_QueryTradeOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).QueryTradeOrder(ctx, req.(*QueryTradeOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_TradeCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).TradeCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeService_TradeCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).TradeCreate(ctx, req.(*TradeCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_MergeTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).MergeTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeService_MergeTrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).MergeTrade(ctx, req.(*MergeTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_UpdateExtendInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExtendInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).UpdateExtendInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeService_UpdateExtendInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).UpdateExtendInfo(ctx, req.(*UpdateExtendInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_UpdateTradeOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTradeOrderStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).UpdateTradeOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeService_UpdateTradeOrderStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).UpdateTradeOrderStatus(ctx, req.(*UpdateTradeOrderStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_UpdateDeliveryInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeliveryInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).UpdateDeliveryInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeService_UpdateDeliveryInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).UpdateDeliveryInfo(ctx, req.(*UpdateDeliveryInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeService_AddSubOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSubOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).AddSubOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeService_AddSubOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).AddSubOrder(ctx, req.(*AddSubOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TradeService_ServiceDesc is the grpc.ServiceDesc for TradeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TradeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.TradeService",
	HandlerType: (*TradeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryTradeOrder",
			Handler:    _TradeService_QueryTradeOrder_Handler,
		},
		{
			MethodName: "TradeCreate",
			Handler:    _TradeService_TradeCreate_Handler,
		},
		{
			MethodName: "MergeTrade",
			Handler:    _TradeService_MergeTrade_Handler,
		},
		{
			MethodName: "UpdateExtendInfo",
			Handler:    _TradeService_UpdateExtendInfo_Handler,
		},
		{
			MethodName: "UpdateTradeOrderStatus",
			Handler:    _TradeService_UpdateTradeOrderStatus_Handler,
		},
		{
			MethodName: "UpdateDeliveryInfo",
			Handler:    _TradeService_UpdateDeliveryInfo_Handler,
		},
		{
			MethodName: "AddSubOrder",
			Handler:    _TradeService_AddSubOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trade.proto",
}
