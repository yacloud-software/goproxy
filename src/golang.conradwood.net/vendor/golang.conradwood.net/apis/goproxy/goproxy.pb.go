// Code generated by protoc-gen-go.
// source: protos/golang.conradwood.net/apis/goproxy/goproxy.proto
// DO NOT EDIT!

/*
Package goproxy is a generated protocol buffer package.

It is generated from these files:
	protos/golang.conradwood.net/apis/goproxy/goproxy.proto

It has these top-level messages:
	PingResponse
*/
package goproxy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "golang.conradwood.net/apis/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// comment: message pingresponse
type PingResponse struct {
	// comment: field pingresponse.response
	Response string `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*PingResponse)(nil), "goproxy.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GoProxy service

type GoProxyClient interface {
	// comment: rpc ping
	Ping(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PingResponse, error)
}

type goProxyClient struct {
	cc *grpc.ClientConn
}

func NewGoProxyClient(cc *grpc.ClientConn) GoProxyClient {
	return &goProxyClient{cc}
}

func (c *goProxyClient) Ping(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/goproxy.GoProxy/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoProxy service

type GoProxyServer interface {
	// comment: rpc ping
	Ping(context.Context, *common.Void) (*PingResponse, error)
}

func RegisterGoProxyServer(s *grpc.Server, srv GoProxyServer) {
	s.RegisterService(&_GoProxy_serviceDesc, srv)
}

func _GoProxy_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoProxyServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goproxy.GoProxy/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoProxyServer).Ping(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goproxy.GoProxy",
	HandlerType: (*GoProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _GoProxy_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/golang.conradwood.net/apis/goproxy/goproxy.proto",
}

func init() {
	proto.RegisterFile("protos/golang.conradwood.net/apis/goproxy/goproxy.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x4b, 0xce, 0xcf, 0x2b, 0x4a, 0x4c,
	0x29, 0xcf, 0xcf, 0x4f, 0xd1, 0xcb, 0x4b, 0x2d, 0xd1, 0x4f, 0x2c, 0xc8, 0x04, 0x49, 0x15, 0x14,
	0xe5, 0x57, 0x54, 0xc2, 0x68, 0x3d, 0xb0, 0x0e, 0x21, 0x76, 0x28, 0x57, 0x4a, 0x0f, 0x8f, 0xd6,
	0xe4, 0xfc, 0xdc, 0xdc, 0xfc, 0x3c, 0x28, 0x05, 0xd1, 0xa8, 0xa4, 0xc5, 0xc5, 0x13, 0x90, 0x99,
	0x97, 0x1e, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc5, 0xc5, 0x01, 0x63, 0x4b,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x46, 0x66, 0x5c, 0xec, 0xee, 0xf9, 0x01, 0x20,
	0x6b, 0x84, 0xb4, 0xb9, 0x58, 0x40, 0xda, 0x84, 0x78, 0xf4, 0xa0, 0xa6, 0x85, 0xe5, 0x67, 0xa6,
	0x48, 0x89, 0xea, 0xc1, 0x5c, 0x85, 0x6c, 0xa6, 0x93, 0x2c, 0x97, 0x74, 0x5e, 0x6a, 0x09, 0xb2,
	0x93, 0x40, 0xce, 0x81, 0xa9, 0x4d, 0x62, 0x03, 0xbb, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x63, 0xf5, 0xda, 0x50, 0xfd, 0x00, 0x00, 0x00,
}
