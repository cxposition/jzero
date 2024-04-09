// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: credential.proto

package credentialpb

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

const (
	Credential_CredentialVersion_FullMethodName = "/credentialpb.credential/CredentialVersion"
)

// CredentialClient is the client API for Credential service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CredentialClient interface {
	CredentialVersion(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CredentialVersionResponse, error)
}

type credentialClient struct {
	cc grpc.ClientConnInterface
}

func NewCredentialClient(cc grpc.ClientConnInterface) CredentialClient {
	return &credentialClient{cc}
}

func (c *credentialClient) CredentialVersion(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CredentialVersionResponse, error) {
	out := new(CredentialVersionResponse)
	err := c.cc.Invoke(ctx, Credential_CredentialVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialServer is the server API for Credential service.
// All implementations must embed UnimplementedCredentialServer
// for forward compatibility
type CredentialServer interface {
	CredentialVersion(context.Context, *Empty) (*CredentialVersionResponse, error)
	mustEmbedUnimplementedCredentialServer()
}

// UnimplementedCredentialServer must be embedded to have forward compatible implementations.
type UnimplementedCredentialServer struct {
}

func (UnimplementedCredentialServer) CredentialVersion(context.Context, *Empty) (*CredentialVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CredentialVersion not implemented")
}
func (UnimplementedCredentialServer) mustEmbedUnimplementedCredentialServer() {}

// UnsafeCredentialServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CredentialServer will
// result in compilation errors.
type UnsafeCredentialServer interface {
	mustEmbedUnimplementedCredentialServer()
}

func RegisterCredentialServer(s grpc.ServiceRegistrar, srv CredentialServer) {
	s.RegisterService(&Credential_ServiceDesc, srv)
}

func _Credential_CredentialVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServer).CredentialVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Credential_CredentialVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServer).CredentialVersion(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Credential_ServiceDesc is the grpc.ServiceDesc for Credential service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Credential_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "credentialpb.credential",
	HandlerType: (*CredentialServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CredentialVersion",
			Handler:    _Credential_CredentialVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "credential.proto",
}

const (
	Credentialv2_CredentialVersion_FullMethodName = "/credentialpb.credentialv2/CredentialVersion"
)

// Credentialv2Client is the client API for Credentialv2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Credentialv2Client interface {
	CredentialVersion(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CredentialVersionResponse, error)
}

type credentialv2Client struct {
	cc grpc.ClientConnInterface
}

func NewCredentialv2Client(cc grpc.ClientConnInterface) Credentialv2Client {
	return &credentialv2Client{cc}
}

func (c *credentialv2Client) CredentialVersion(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CredentialVersionResponse, error) {
	out := new(CredentialVersionResponse)
	err := c.cc.Invoke(ctx, Credentialv2_CredentialVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Credentialv2Server is the server API for Credentialv2 service.
// All implementations must embed UnimplementedCredentialv2Server
// for forward compatibility
type Credentialv2Server interface {
	CredentialVersion(context.Context, *Empty) (*CredentialVersionResponse, error)
	mustEmbedUnimplementedCredentialv2Server()
}

// UnimplementedCredentialv2Server must be embedded to have forward compatible implementations.
type UnimplementedCredentialv2Server struct {
}

func (UnimplementedCredentialv2Server) CredentialVersion(context.Context, *Empty) (*CredentialVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CredentialVersion not implemented")
}
func (UnimplementedCredentialv2Server) mustEmbedUnimplementedCredentialv2Server() {}

// UnsafeCredentialv2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Credentialv2Server will
// result in compilation errors.
type UnsafeCredentialv2Server interface {
	mustEmbedUnimplementedCredentialv2Server()
}

func RegisterCredentialv2Server(s grpc.ServiceRegistrar, srv Credentialv2Server) {
	s.RegisterService(&Credentialv2_ServiceDesc, srv)
}

func _Credentialv2_CredentialVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Credentialv2Server).CredentialVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Credentialv2_CredentialVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Credentialv2Server).CredentialVersion(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Credentialv2_ServiceDesc is the grpc.ServiceDesc for Credentialv2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Credentialv2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "credentialpb.credentialv2",
	HandlerType: (*Credentialv2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CredentialVersion",
			Handler:    _Credentialv2_CredentialVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "credential.proto",
}
