// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bootstrapv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BootstrapClient is the client API for Bootstrap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BootstrapClient interface {
	// Get trust anchor ARN
	GetTrustAnchorARN(ctx context.Context, in *GetTrustAnchorARNRequest, opts ...grpc.CallOption) (*GetTrustAnchorARNResponse, error)
}

type bootstrapClient struct {
	cc grpc.ClientConnInterface
}

func NewBootstrapClient(cc grpc.ClientConnInterface) BootstrapClient {
	return &bootstrapClient{cc}
}

func (c *bootstrapClient) GetTrustAnchorARN(ctx context.Context, in *GetTrustAnchorARNRequest, opts ...grpc.CallOption) (*GetTrustAnchorARNResponse, error) {
	out := new(GetTrustAnchorARNResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.bootstrap.v1.Bootstrap/GetTrustAnchorARN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BootstrapServer is the server API for Bootstrap service.
// All implementations must embed UnimplementedBootstrapServer
// for forward compatibility
type BootstrapServer interface {
	// Get trust anchor ARN
	GetTrustAnchorARN(context.Context, *GetTrustAnchorARNRequest) (*GetTrustAnchorARNResponse, error)
	mustEmbedUnimplementedBootstrapServer()
}

// UnimplementedBootstrapServer must be embedded to have forward compatible implementations.
type UnimplementedBootstrapServer struct {
}

func (UnimplementedBootstrapServer) GetTrustAnchorARN(context.Context, *GetTrustAnchorARNRequest) (*GetTrustAnchorARNResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrustAnchorARN not implemented")
}
func (UnimplementedBootstrapServer) mustEmbedUnimplementedBootstrapServer() {}

// UnsafeBootstrapServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BootstrapServer will
// result in compilation errors.
type UnsafeBootstrapServer interface {
	mustEmbedUnimplementedBootstrapServer()
}

func RegisterBootstrapServer(s grpc.ServiceRegistrar, srv BootstrapServer) {
	s.RegisterService(&_Bootstrap_serviceDesc, srv)
}

func _Bootstrap_GetTrustAnchorARN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrustAnchorARNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootstrapServer).GetTrustAnchorARN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.bootstrap.v1.Bootstrap/GetTrustAnchorARN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootstrapServer).GetTrustAnchorARN(ctx, req.(*GetTrustAnchorARNRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bootstrap_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.api.server.bootstrap.v1.Bootstrap",
	HandlerType: (*BootstrapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTrustAnchorARN",
			Handler:    _Bootstrap_GetTrustAnchorARN_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spire/api/server/bootstrap/v1/bootstrap.proto",
}
