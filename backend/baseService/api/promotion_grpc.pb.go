// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.1
// source: promotion.proto

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
	PromotionService_PromotionCreate_FullMethodName         = "/api.PromotionService/PromotionCreate"
	PromotionService_PromotionQuery_FullMethodName          = "/api.PromotionService/PromotionQuery"
	PromotionService_PromotionUpdate_FullMethodName         = "/api.PromotionService/PromotionUpdate"
	PromotionService_PromotionDelete_FullMethodName         = "/api.PromotionService/PromotionDelete"
	PromotionService_CreateSpecificPromotion_FullMethodName = "/api.PromotionService/CreateSpecificPromotion"
	PromotionService_QuerySpecificPromotion_FullMethodName  = "/api.PromotionService/QuerySpecificPromotion"
	PromotionService_Calculate_FullMethodName               = "/api.PromotionService/Calculate"
)

// PromotionServiceClient is the client API for PromotionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PromotionServiceClient interface {
	PromotionCreate(ctx context.Context, in *PromotionCreateRequest, opts ...grpc.CallOption) (*PromotionCreateResponse, error)
	PromotionQuery(ctx context.Context, in *PromotionQueryRequest, opts ...grpc.CallOption) (*PromotionQueryResponse, error)
	PromotionUpdate(ctx context.Context, in *PromotionUpdateRequest, opts ...grpc.CallOption) (*PromotionUpdateResponse, error)
	PromotionDelete(ctx context.Context, in *PromotionDeleteRequest, opts ...grpc.CallOption) (*PromotionDeleteResponse, error)
	CreateSpecificPromotion(ctx context.Context, in *CreateSpecificPromotionRequest, opts ...grpc.CallOption) (*CreateSpecificPromotionResponse, error)
	QuerySpecificPromotion(ctx context.Context, in *QuerySpecificPromotionRequest, opts ...grpc.CallOption) (*QuerySpecificPromotionResponse, error)
	Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
}

type promotionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPromotionServiceClient(cc grpc.ClientConnInterface) PromotionServiceClient {
	return &promotionServiceClient{cc}
}

func (c *promotionServiceClient) PromotionCreate(ctx context.Context, in *PromotionCreateRequest, opts ...grpc.CallOption) (*PromotionCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PromotionCreateResponse)
	err := c.cc.Invoke(ctx, PromotionService_PromotionCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) PromotionQuery(ctx context.Context, in *PromotionQueryRequest, opts ...grpc.CallOption) (*PromotionQueryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PromotionQueryResponse)
	err := c.cc.Invoke(ctx, PromotionService_PromotionQuery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) PromotionUpdate(ctx context.Context, in *PromotionUpdateRequest, opts ...grpc.CallOption) (*PromotionUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PromotionUpdateResponse)
	err := c.cc.Invoke(ctx, PromotionService_PromotionUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) PromotionDelete(ctx context.Context, in *PromotionDeleteRequest, opts ...grpc.CallOption) (*PromotionDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PromotionDeleteResponse)
	err := c.cc.Invoke(ctx, PromotionService_PromotionDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) CreateSpecificPromotion(ctx context.Context, in *CreateSpecificPromotionRequest, opts ...grpc.CallOption) (*CreateSpecificPromotionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSpecificPromotionResponse)
	err := c.cc.Invoke(ctx, PromotionService_CreateSpecificPromotion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) QuerySpecificPromotion(ctx context.Context, in *QuerySpecificPromotionRequest, opts ...grpc.CallOption) (*QuerySpecificPromotionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QuerySpecificPromotionResponse)
	err := c.cc.Invoke(ctx, PromotionService_QuerySpecificPromotion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, PromotionService_Calculate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromotionServiceServer is the server API for PromotionService service.
// All implementations should embed UnimplementedPromotionServiceServer
// for forward compatibility.
type PromotionServiceServer interface {
	PromotionCreate(context.Context, *PromotionCreateRequest) (*PromotionCreateResponse, error)
	PromotionQuery(context.Context, *PromotionQueryRequest) (*PromotionQueryResponse, error)
	PromotionUpdate(context.Context, *PromotionUpdateRequest) (*PromotionUpdateResponse, error)
	PromotionDelete(context.Context, *PromotionDeleteRequest) (*PromotionDeleteResponse, error)
	CreateSpecificPromotion(context.Context, *CreateSpecificPromotionRequest) (*CreateSpecificPromotionResponse, error)
	QuerySpecificPromotion(context.Context, *QuerySpecificPromotionRequest) (*QuerySpecificPromotionResponse, error)
	Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error)
}

// UnimplementedPromotionServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPromotionServiceServer struct{}

func (UnimplementedPromotionServiceServer) PromotionCreate(context.Context, *PromotionCreateRequest) (*PromotionCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromotionCreate not implemented")
}
func (UnimplementedPromotionServiceServer) PromotionQuery(context.Context, *PromotionQueryRequest) (*PromotionQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromotionQuery not implemented")
}
func (UnimplementedPromotionServiceServer) PromotionUpdate(context.Context, *PromotionUpdateRequest) (*PromotionUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromotionUpdate not implemented")
}
func (UnimplementedPromotionServiceServer) PromotionDelete(context.Context, *PromotionDeleteRequest) (*PromotionDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromotionDelete not implemented")
}
func (UnimplementedPromotionServiceServer) CreateSpecificPromotion(context.Context, *CreateSpecificPromotionRequest) (*CreateSpecificPromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSpecificPromotion not implemented")
}
func (UnimplementedPromotionServiceServer) QuerySpecificPromotion(context.Context, *QuerySpecificPromotionRequest) (*QuerySpecificPromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySpecificPromotion not implemented")
}
func (UnimplementedPromotionServiceServer) Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedPromotionServiceServer) testEmbeddedByValue() {}

// UnsafePromotionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PromotionServiceServer will
// result in compilation errors.
type UnsafePromotionServiceServer interface {
	mustEmbedUnimplementedPromotionServiceServer()
}

func RegisterPromotionServiceServer(s grpc.ServiceRegistrar, srv PromotionServiceServer) {
	// If the following call pancis, it indicates UnimplementedPromotionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PromotionService_ServiceDesc, srv)
}

func _PromotionService_PromotionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromotionCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).PromotionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromotionService_PromotionCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).PromotionCreate(ctx, req.(*PromotionCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_PromotionQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromotionQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).PromotionQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromotionService_PromotionQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).PromotionQuery(ctx, req.(*PromotionQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_PromotionUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromotionUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).PromotionUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromotionService_PromotionUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).PromotionUpdate(ctx, req.(*PromotionUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_PromotionDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromotionDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).PromotionDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromotionService_PromotionDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).PromotionDelete(ctx, req.(*PromotionDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_CreateSpecificPromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSpecificPromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).CreateSpecificPromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromotionService_CreateSpecificPromotion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).CreateSpecificPromotion(ctx, req.(*CreateSpecificPromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_QuerySpecificPromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySpecificPromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).QuerySpecificPromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromotionService_QuerySpecificPromotion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).QuerySpecificPromotion(ctx, req.(*QuerySpecificPromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromotionService_Calculate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).Calculate(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PromotionService_ServiceDesc is the grpc.ServiceDesc for PromotionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PromotionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.PromotionService",
	HandlerType: (*PromotionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PromotionCreate",
			Handler:    _PromotionService_PromotionCreate_Handler,
		},
		{
			MethodName: "PromotionQuery",
			Handler:    _PromotionService_PromotionQuery_Handler,
		},
		{
			MethodName: "PromotionUpdate",
			Handler:    _PromotionService_PromotionUpdate_Handler,
		},
		{
			MethodName: "PromotionDelete",
			Handler:    _PromotionService_PromotionDelete_Handler,
		},
		{
			MethodName: "CreateSpecificPromotion",
			Handler:    _PromotionService_CreateSpecificPromotion_Handler,
		},
		{
			MethodName: "QuerySpecificPromotion",
			Handler:    _PromotionService_QuerySpecificPromotion_Handler,
		},
		{
			MethodName: "Calculate",
			Handler:    _PromotionService_Calculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "promotion.proto",
}
