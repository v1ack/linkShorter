// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package linkShorter

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

// ShorterClient is the client API for Shorter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShorterClient interface {
	Create(ctx context.Context, in *CreateLinkRequest, opts ...grpc.CallOption) (*CreateLinkResponse, error)
	Get(ctx context.Context, in *GetLinkRequest, opts ...grpc.CallOption) (*GetLinkResponse, error)
}

type shorterClient struct {
	cc grpc.ClientConnInterface
}

func NewShorterClient(cc grpc.ClientConnInterface) ShorterClient {
	return &shorterClient{cc}
}

func (c *shorterClient) Create(ctx context.Context, in *CreateLinkRequest, opts ...grpc.CallOption) (*CreateLinkResponse, error) {
	out := new(CreateLinkResponse)
	err := c.cc.Invoke(ctx, "/shorter.Shorter/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shorterClient) Get(ctx context.Context, in *GetLinkRequest, opts ...grpc.CallOption) (*GetLinkResponse, error) {
	out := new(GetLinkResponse)
	err := c.cc.Invoke(ctx, "/shorter.Shorter/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShorterServer is the server API for Shorter service.
// All implementations should embed UnimplementedShorterServer
// for forward compatibility
type ShorterServer interface {
	Create(context.Context, *CreateLinkRequest) (*CreateLinkResponse, error)
	Get(context.Context, *GetLinkRequest) (*GetLinkResponse, error)
}

// UnimplementedShorterServer should be embedded to have forward compatible implementations.
type UnimplementedShorterServer struct {
}

func (UnimplementedShorterServer) Create(context.Context, *CreateLinkRequest) (*CreateLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedShorterServer) Get(context.Context, *GetLinkRequest) (*GetLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeShorterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShorterServer will
// result in compilation errors.
type UnsafeShorterServer interface {
	mustEmbedUnimplementedShorterServer()
}

func RegisterShorterServer(s grpc.ServiceRegistrar, srv ShorterServer) {
	s.RegisterService(&Shorter_ServiceDesc, srv)
}

func _Shorter_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShorterServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shorter.Shorter/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShorterServer).Create(ctx, req.(*CreateLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shorter_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShorterServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shorter.Shorter/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShorterServer).Get(ctx, req.(*GetLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Shorter_ServiceDesc is the grpc.ServiceDesc for Shorter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shorter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shorter.Shorter",
	HandlerType: (*ShorterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Shorter_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Shorter_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shorter.proto",
}
