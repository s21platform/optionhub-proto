// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: optionhub.v1.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	OptionhubService_SetAttributeByID_FullMethodName = "/OptionhubService/SetAttributeByID"
)

// OptionhubServiceClient is the client API for OptionhubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OptionhubServiceClient interface {
	SetAttributeByID(ctx context.Context, in *SetAttributeByIdIn, opts ...grpc.CallOption) (*empty.Empty, error)
}

type optionhubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOptionhubServiceClient(cc grpc.ClientConnInterface) OptionhubServiceClient {
	return &optionhubServiceClient{cc}
}

func (c *optionhubServiceClient) SetAttributeByID(ctx context.Context, in *SetAttributeByIdIn, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, OptionhubService_SetAttributeByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OptionhubServiceServer is the server API for OptionhubService service.
// All implementations must embed UnimplementedOptionhubServiceServer
// for forward compatibility.
type OptionhubServiceServer interface {
	SetAttributeByID(context.Context, *SetAttributeByIdIn) (*empty.Empty, error)
	mustEmbedUnimplementedOptionhubServiceServer()
}

// UnimplementedOptionhubServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOptionhubServiceServer struct{}

func (UnimplementedOptionhubServiceServer) SetAttributeByID(context.Context, *SetAttributeByIdIn) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAttributeByID not implemented")
}
func (UnimplementedOptionhubServiceServer) mustEmbedUnimplementedOptionhubServiceServer() {}
func (UnimplementedOptionhubServiceServer) testEmbeddedByValue()                          {}

// UnsafeOptionhubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OptionhubServiceServer will
// result in compilation errors.
type UnsafeOptionhubServiceServer interface {
	mustEmbedUnimplementedOptionhubServiceServer()
}

func RegisterOptionhubServiceServer(s grpc.ServiceRegistrar, srv OptionhubServiceServer) {
	// If the following call pancis, it indicates UnimplementedOptionhubServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OptionhubService_ServiceDesc, srv)
}

func _OptionhubService_SetAttributeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAttributeByIdIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OptionhubServiceServer).SetAttributeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OptionhubService_SetAttributeByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OptionhubServiceServer).SetAttributeByID(ctx, req.(*SetAttributeByIdIn))
	}
	return interceptor(ctx, in, info, handler)
}

// OptionhubService_ServiceDesc is the grpc.ServiceDesc for OptionhubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OptionhubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OptionhubService",
	HandlerType: (*OptionhubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetAttributeByID",
			Handler:    _OptionhubService_SetAttributeByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "optionhub.v1.proto",
}
