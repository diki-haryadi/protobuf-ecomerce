// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tokopedia/oauth2/v1/password.proto

package oauth2v1

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

// PasswordServiceClient is the client API for PasswordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PasswordServiceClient interface {
	PasswordGrant(ctx context.Context, in *PasswordGrantRequest, opts ...grpc.CallOption) (*PasswordGrantResponse, error)
}

type passwordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPasswordServiceClient(cc grpc.ClientConnInterface) PasswordServiceClient {
	return &passwordServiceClient{cc}
}

func (c *passwordServiceClient) PasswordGrant(ctx context.Context, in *PasswordGrantRequest, opts ...grpc.CallOption) (*PasswordGrantResponse, error) {
	out := new(PasswordGrantResponse)
	err := c.cc.Invoke(ctx, "/tokopedia.oauth2.v1.PasswordService/PasswordGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PasswordServiceServer is the server API for PasswordService service.
// All implementations should embed UnimplementedPasswordServiceServer
// for forward compatibility
type PasswordServiceServer interface {
	PasswordGrant(context.Context, *PasswordGrantRequest) (*PasswordGrantResponse, error)
}

// UnimplementedPasswordServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPasswordServiceServer struct {
}

func (UnimplementedPasswordServiceServer) PasswordGrant(context.Context, *PasswordGrantRequest) (*PasswordGrantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PasswordGrant not implemented")
}

// UnsafePasswordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PasswordServiceServer will
// result in compilation errors.
type UnsafePasswordServiceServer interface {
	mustEmbedUnimplementedPasswordServiceServer()
}

func RegisterPasswordServiceServer(s grpc.ServiceRegistrar, srv PasswordServiceServer) {
	s.RegisterService(&PasswordService_ServiceDesc, srv)
}

func _PasswordService_PasswordGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordServiceServer).PasswordGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokopedia.oauth2.v1.PasswordService/PasswordGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordServiceServer).PasswordGrant(ctx, req.(*PasswordGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PasswordService_ServiceDesc is the grpc.ServiceDesc for PasswordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PasswordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tokopedia.oauth2.v1.PasswordService",
	HandlerType: (*PasswordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PasswordGrant",
			Handler:    _PasswordService_PasswordGrant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tokopedia/oauth2/v1/password.proto",
}