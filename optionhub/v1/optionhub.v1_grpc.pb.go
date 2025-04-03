// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: optionhub.v1.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	OptionhubServiceV1_SetAttributeByID_FullMethodName   = "/OptionhubServiceV1/SetAttributeByID"
	OptionhubServiceV1_GetOptionRequests_FullMethodName  = "/OptionhubServiceV1/GetOptionRequests"
	OptionhubServiceV1_GetAttributeValues_FullMethodName = "/OptionhubServiceV1/GetAttributeValues"
)

// OptionhubServiceV1Client is the client API for OptionhubServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OptionhubServiceV1Client interface {
	SetAttributeByID(ctx context.Context, in *SetAttributeByIdIn, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetOptionRequests(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetOptionRequestsOut, error)
	GetAttributeValues(ctx context.Context, in *GetAttributeValuesIn, opts ...grpc.CallOption) (*GetAttributeValuesOut, error)
}

type optionhubServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewOptionhubServiceV1Client(cc grpc.ClientConnInterface) OptionhubServiceV1Client {
	return &optionhubServiceV1Client{cc}
}

func (c *optionhubServiceV1Client) SetAttributeByID(ctx context.Context, in *SetAttributeByIdIn, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, OptionhubServiceV1_SetAttributeByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *optionhubServiceV1Client) GetOptionRequests(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetOptionRequestsOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOptionRequestsOut)
	err := c.cc.Invoke(ctx, OptionhubServiceV1_GetOptionRequests_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *optionhubServiceV1Client) GetAttributeValues(ctx context.Context, in *GetAttributeValuesIn, opts ...grpc.CallOption) (*GetAttributeValuesOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAttributeValuesOut)
	err := c.cc.Invoke(ctx, OptionhubServiceV1_GetAttributeValues_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OptionhubServiceV1Server is the server API for OptionhubServiceV1 service.
// All implementations must embed UnimplementedOptionhubServiceV1Server
// for forward compatibility.
type OptionhubServiceV1Server interface {
	SetAttributeByID(context.Context, *SetAttributeByIdIn) (*emptypb.Empty, error)
	GetOptionRequests(context.Context, *emptypb.Empty) (*GetOptionRequestsOut, error)
	GetAttributeValues(context.Context, *GetAttributeValuesIn) (*GetAttributeValuesOut, error)
	mustEmbedUnimplementedOptionhubServiceV1Server()
}

// UnimplementedOptionhubServiceV1Server must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOptionhubServiceV1Server struct{}

func (UnimplementedOptionhubServiceV1Server) SetAttributeByID(context.Context, *SetAttributeByIdIn) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAttributeByID not implemented")
}
func (UnimplementedOptionhubServiceV1Server) GetOptionRequests(context.Context, *emptypb.Empty) (*GetOptionRequestsOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOptionRequests not implemented")
}
func (UnimplementedOptionhubServiceV1Server) GetAttributeValues(context.Context, *GetAttributeValuesIn) (*GetAttributeValuesOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttributeValues not implemented")
}
func (UnimplementedOptionhubServiceV1Server) mustEmbedUnimplementedOptionhubServiceV1Server() {}
func (UnimplementedOptionhubServiceV1Server) testEmbeddedByValue()                            {}

// UnsafeOptionhubServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OptionhubServiceV1Server will
// result in compilation errors.
type UnsafeOptionhubServiceV1Server interface {
	mustEmbedUnimplementedOptionhubServiceV1Server()
}

func RegisterOptionhubServiceV1Server(s grpc.ServiceRegistrar, srv OptionhubServiceV1Server) {
	// If the following call pancis, it indicates UnimplementedOptionhubServiceV1Server was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OptionhubServiceV1_ServiceDesc, srv)
}

func _OptionhubServiceV1_SetAttributeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAttributeByIdIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OptionhubServiceV1Server).SetAttributeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OptionhubServiceV1_SetAttributeByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OptionhubServiceV1Server).SetAttributeByID(ctx, req.(*SetAttributeByIdIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _OptionhubServiceV1_GetOptionRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OptionhubServiceV1Server).GetOptionRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OptionhubServiceV1_GetOptionRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OptionhubServiceV1Server).GetOptionRequests(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OptionhubServiceV1_GetAttributeValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttributeValuesIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OptionhubServiceV1Server).GetAttributeValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OptionhubServiceV1_GetAttributeValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OptionhubServiceV1Server).GetAttributeValues(ctx, req.(*GetAttributeValuesIn))
	}
	return interceptor(ctx, in, info, handler)
}

// OptionhubServiceV1_ServiceDesc is the grpc.ServiceDesc for OptionhubServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OptionhubServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OptionhubServiceV1",
	HandlerType: (*OptionhubServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetAttributeByID",
			Handler:    _OptionhubServiceV1_SetAttributeByID_Handler,
		},
		{
			MethodName: "GetOptionRequests",
			Handler:    _OptionhubServiceV1_GetOptionRequests_Handler,
		},
		{
			MethodName: "GetAttributeValues",
			Handler:    _OptionhubServiceV1_GetAttributeValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "optionhub.v1.proto",
}
