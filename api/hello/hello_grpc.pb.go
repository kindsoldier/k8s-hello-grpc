// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hello

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloClient interface {
	Install(ctx context.Context, in *InstallRequest, opts ...grpc.CallOption) (Hello_InstallClient, error)
}

type helloClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloClient(cc grpc.ClientConnInterface) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Install(ctx context.Context, in *InstallRequest, opts ...grpc.CallOption) (Hello_InstallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Hello_serviceDesc.Streams[0], "/services.Hello/Install", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloInstallClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hello_InstallClient interface {
	Recv() (*InstallResult, error)
	grpc.ClientStream
}

type helloInstallClient struct {
	grpc.ClientStream
}

func (x *helloInstallClient) Recv() (*InstallResult, error) {
	m := new(InstallResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServer is the server API for Hello service.
// All implementations must embed UnimplementedHelloServer
// for forward compatibility
type HelloServer interface {
	Install(*InstallRequest, Hello_InstallServer) error
	mustEmbedUnimplementedHelloServer()
}

// UnimplementedHelloServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (UnimplementedHelloServer) Install(*InstallRequest, Hello_InstallServer) error {
	return status.Errorf(codes.Unimplemented, "method Install not implemented")
}
func (UnimplementedHelloServer) mustEmbedUnimplementedHelloServer() {}

// UnsafeHelloServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServer will
// result in compilation errors.
type UnsafeHelloServer interface {
	mustEmbedUnimplementedHelloServer()
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Install_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InstallRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServer).Install(m, &helloInstallServer{stream})
}

type Hello_InstallServer interface {
	Send(*InstallResult) error
	grpc.ServerStream
}

type helloInstallServer struct {
	grpc.ServerStream
}

func (x *helloInstallServer) Send(m *InstallResult) error {
	return x.ServerStream.SendMsg(m)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Install",
			Handler:       _Hello_Install_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hello.proto",
}
