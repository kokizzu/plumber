// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/dproxy.proto

package services

import (
	context "context"
	fmt "fmt"
	events "github.com/batchcorp/collector-schemas/build/go/protos/events"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("services/dproxy.proto", fileDescriptor_b975708b01663117) }

var fileDescriptor_b975708b01663117 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x87, 0x59, 0x90, 0x2a, 0x41, 0x04, 0x53, 0xbd, 0xac, 0xbe, 0x81, 0x98, 0xf1, 0xcf, 0x03,
	0x88, 0xda, 0x8b, 0xb7, 0x22, 0xe2, 0xc1, 0x5b, 0x77, 0x3a, 0x34, 0x0b, 0x69, 0x66, 0xcd, 0x4c,
	0x8a, 0x3e, 0x97, 0x2f, 0x28, 0xbb, 0x21, 0xa0, 0xe2, 0x71, 0x7e, 0xdf, 0xc7, 0x07, 0x89, 0x39,
	0x15, 0x4a, 0xbb, 0x1e, 0x49, 0x60, 0x3d, 0x24, 0xfe, 0xf8, 0x74, 0x43, 0x62, 0x65, 0x7b, 0x50,
	0xe7, 0x76, 0x4e, 0x3b, 0x8a, 0x2a, 0xa0, 0x39, 0x46, 0x0a, 0x05, 0xdf, 0x7c, 0x35, 0x66, 0xb6,
	0x58, 0x8e, 0xbe, 0xbd, 0x30, 0xfb, 0x8f, 0x1c, 0x23, 0xa1, 0xda, 0x23, 0x57, 0x5c, 0xf7, 0x32,
	0xb9, 0xed, 0x9f, 0xfb, 0xaa, 0xb1, 0xd7, 0x66, 0xef, 0x3e, 0xab, 0xb7, 0xf3, 0x4a, 0xc6, 0xeb,
	0x99, 0xde, 0x33, 0x89, 0xb6, 0x27, 0xbf, 0x47, 0x19, 0x38, 0x0a, 0xd9, 0x27, 0x73, 0xb8, 0xa0,
	0x40, 0x4a, 0x25, 0x62, 0xcf, 0xaa, 0xf5, 0x73, 0xad, 0x89, 0xf3, 0xff, 0x61, 0x49, 0x3d, 0xbc,
	0x9a, 0x63, 0xf1, 0xae, 0x5b, 0x29, 0x7a, 0x57, 0xdf, 0xb7, 0x6c, 0xde, 0xee, 0x36, 0xbd, 0xfa,
	0xdc, 0x39, 0xe4, 0x2d, 0x4c, 0x10, 0x39, 0x0d, 0x80, 0x1c, 0x02, 0xa1, 0x72, 0xba, 0x14, 0xf4,
	0xb4, 0x5d, 0x09, 0x74, 0xb9, 0x0f, 0x6b, 0xd8, 0x30, 0x4c, 0x7f, 0x20, 0x50, 0x13, 0xdd, 0x6c,
	0x1a, 0x6e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x78, 0x42, 0x05, 0xf9, 0x4c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DProxyClient is the client API for DProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DProxyClient interface {
	Connect(ctx context.Context, in *events.Tunnel, opts ...grpc.CallOption) (DProxy_ConnectClient, error)
	Auth(ctx context.Context, in *events.AuthRequest, opts ...grpc.CallOption) (*events.AuthResponse, error)
	DeleteTunnel(ctx context.Context, in *events.DeleteTunnelRequest, opts ...grpc.CallOption) (*events.DeleteTunnelResponse, error)
}

type dProxyClient struct {
	cc *grpc.ClientConn
}

func NewDProxyClient(cc *grpc.ClientConn) DProxyClient {
	return &dProxyClient{cc}
}

func (c *dProxyClient) Connect(ctx context.Context, in *events.Tunnel, opts ...grpc.CallOption) (DProxy_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DProxy_serviceDesc.Streams[0], "/services.DProxy/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &dProxyConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DProxy_ConnectClient interface {
	Recv() (*events.Tunnel, error)
	grpc.ClientStream
}

type dProxyConnectClient struct {
	grpc.ClientStream
}

func (x *dProxyConnectClient) Recv() (*events.Tunnel, error) {
	m := new(events.Tunnel)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dProxyClient) Auth(ctx context.Context, in *events.AuthRequest, opts ...grpc.CallOption) (*events.AuthResponse, error) {
	out := new(events.AuthResponse)
	err := c.cc.Invoke(ctx, "/services.DProxy/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dProxyClient) DeleteTunnel(ctx context.Context, in *events.DeleteTunnelRequest, opts ...grpc.CallOption) (*events.DeleteTunnelResponse, error) {
	out := new(events.DeleteTunnelResponse)
	err := c.cc.Invoke(ctx, "/services.DProxy/DeleteTunnel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DProxyServer is the server API for DProxy service.
type DProxyServer interface {
	Connect(*events.Tunnel, DProxy_ConnectServer) error
	Auth(context.Context, *events.AuthRequest) (*events.AuthResponse, error)
	DeleteTunnel(context.Context, *events.DeleteTunnelRequest) (*events.DeleteTunnelResponse, error)
}

// UnimplementedDProxyServer can be embedded to have forward compatible implementations.
type UnimplementedDProxyServer struct {
}

func (*UnimplementedDProxyServer) Connect(req *events.Tunnel, srv DProxy_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedDProxyServer) Auth(ctx context.Context, req *events.AuthRequest) (*events.AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (*UnimplementedDProxyServer) DeleteTunnel(ctx context.Context, req *events.DeleteTunnelRequest) (*events.DeleteTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTunnel not implemented")
}

func RegisterDProxyServer(s *grpc.Server, srv DProxyServer) {
	s.RegisterService(&_DProxy_serviceDesc, srv)
}

func _DProxy_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(events.Tunnel)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DProxyServer).Connect(m, &dProxyConnectServer{stream})
}

type DProxy_ConnectServer interface {
	Send(*events.Tunnel) error
	grpc.ServerStream
}

type dProxyConnectServer struct {
	grpc.ServerStream
}

func (x *dProxyConnectServer) Send(m *events.Tunnel) error {
	return x.ServerStream.SendMsg(m)
}

func _DProxy_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DProxyServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.DProxy/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DProxyServer).Auth(ctx, req.(*events.AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DProxy_DeleteTunnel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.DeleteTunnelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DProxyServer).DeleteTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.DProxy/DeleteTunnel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DProxyServer).DeleteTunnel(ctx, req.(*events.DeleteTunnelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.DProxy",
	HandlerType: (*DProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _DProxy_Auth_Handler,
		},
		{
			MethodName: "DeleteTunnel",
			Handler:    _DProxy_DeleteTunnel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _DProxy_Connect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services/dproxy.proto",
}
