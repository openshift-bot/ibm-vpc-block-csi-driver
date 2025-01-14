// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package provider

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// APIKeyProviderClient is the client API for APIKeyProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIKeyProviderClient interface {
	// Get container API key
	GetContainerAPIKey(ctx context.Context, in *Cipher, opts ...grpc.CallOption) (*APIKey, error)
	// Get VPC API key
	GetVPCAPIKey(ctx context.Context, in *Cipher, opts ...grpc.CallOption) (*APIKey, error)
}

type aPIKeyProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIKeyProviderClient(cc grpc.ClientConnInterface) APIKeyProviderClient {
	return &aPIKeyProviderClient{cc}
}

func (c *aPIKeyProviderClient) GetContainerAPIKey(ctx context.Context, in *Cipher, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/provider.APIKeyProvider/GetContainerAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyProviderClient) GetVPCAPIKey(ctx context.Context, in *Cipher, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/provider.APIKeyProvider/GetVPCAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIKeyProviderServer is the server API for APIKeyProvider service.
// All implementations must embed UnimplementedAPIKeyProviderServer
// for forward compatibility
type APIKeyProviderServer interface {
	// Get container API key
	GetContainerAPIKey(context.Context, *Cipher) (*APIKey, error)
	// Get VPC API key
	GetVPCAPIKey(context.Context, *Cipher) (*APIKey, error)
	mustEmbedUnimplementedAPIKeyProviderServer()
}

// UnimplementedAPIKeyProviderServer must be embedded to have forward compatible implementations.
type UnimplementedAPIKeyProviderServer struct {
}

func (UnimplementedAPIKeyProviderServer) GetContainerAPIKey(context.Context, *Cipher) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainerAPIKey not implemented")
}
func (UnimplementedAPIKeyProviderServer) GetVPCAPIKey(context.Context, *Cipher) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVPCAPIKey not implemented")
}
func (UnimplementedAPIKeyProviderServer) mustEmbedUnimplementedAPIKeyProviderServer() {}

// UnsafeAPIKeyProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIKeyProviderServer will
// result in compilation errors.
type UnsafeAPIKeyProviderServer interface {
	mustEmbedUnimplementedAPIKeyProviderServer()
}

func RegisterAPIKeyProviderServer(s *grpc.Server, srv APIKeyProviderServer) {
	s.RegisterService(&_APIKeyProvider_serviceDesc, srv)
}

func _APIKeyProvider_GetContainerAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cipher)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyProviderServer).GetContainerAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provider.APIKeyProvider/GetContainerAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyProviderServer).GetContainerAPIKey(ctx, req.(*Cipher))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKeyProvider_GetVPCAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cipher)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyProviderServer).GetVPCAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provider.APIKeyProvider/GetVPCAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyProviderServer).GetVPCAPIKey(ctx, req.(*Cipher))
	}
	return interceptor(ctx, in, info, handler)
}

var _APIKeyProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "provider.APIKeyProvider",
	HandlerType: (*APIKeyProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContainerAPIKey",
			Handler:    _APIKeyProvider_GetContainerAPIKey_Handler,
		},
		{
			MethodName: "GetVPCAPIKey",
			Handler:    _APIKeyProvider_GetVPCAPIKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider/provider.proto",
}
