// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: grpc-ldap.proto

package v1alpha1

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

// SimpleLDAPServiceClient is the client API for SimpleLDAPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleLDAPServiceClient interface {
	SayHi(ctx context.Context, in *SayHiRequest, opts ...grpc.CallOption) (*SayHiResponse, error)
}

type simpleLDAPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleLDAPServiceClient(cc grpc.ClientConnInterface) SimpleLDAPServiceClient {
	return &simpleLDAPServiceClient{cc}
}

func (c *simpleLDAPServiceClient) SayHi(ctx context.Context, in *SayHiRequest, opts ...grpc.CallOption) (*SayHiResponse, error) {
	out := new(SayHiResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.SimpleLDAPService/SayHi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleLDAPServiceServer is the server API for SimpleLDAPService service.
// All implementations must embed UnimplementedSimpleLDAPServiceServer
// for forward compatibility
type SimpleLDAPServiceServer interface {
	SayHi(context.Context, *SayHiRequest) (*SayHiResponse, error)
	mustEmbedUnimplementedSimpleLDAPServiceServer()
}

// UnimplementedSimpleLDAPServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSimpleLDAPServiceServer struct {
}

func (UnimplementedSimpleLDAPServiceServer) SayHi(context.Context, *SayHiRequest) (*SayHiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHi not implemented")
}
func (UnimplementedSimpleLDAPServiceServer) mustEmbedUnimplementedSimpleLDAPServiceServer() {}

// UnsafeSimpleLDAPServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleLDAPServiceServer will
// result in compilation errors.
type UnsafeSimpleLDAPServiceServer interface {
	mustEmbedUnimplementedSimpleLDAPServiceServer()
}

func RegisterSimpleLDAPServiceServer(s grpc.ServiceRegistrar, srv SimpleLDAPServiceServer) {
	s.RegisterService(&SimpleLDAPService_ServiceDesc, srv)
}

func _SimpleLDAPService_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleLDAPServiceServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.SimpleLDAPService/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleLDAPServiceServer).SayHi(ctx, req.(*SayHiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SimpleLDAPService_ServiceDesc is the grpc.ServiceDesc for SimpleLDAPService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimpleLDAPService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.SimpleLDAPService",
	HandlerType: (*SimpleLDAPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _SimpleLDAPService_SayHi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc-ldap.proto",
}
