// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/promutils/promutils.proto
// DO NOT EDIT!

/*
Package promutils is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/promutils/promutils.proto

It has these top-level messages:
	TextQuery
	ParsedQuery
*/
package promutils

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// a 'string', such as max(foo_metric{instance="bar"})
type TextQuery struct {
	Text string `protobuf:"bytes,1,opt,name=Text" json:"Text,omitempty"`
}

func (m *TextQuery) Reset()                    { *m = TextQuery{} }
func (m *TextQuery) String() string            { return proto.CompactTextString(m) }
func (*TextQuery) ProtoMessage()               {}
func (*TextQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TextQuery) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

// a ParsedQuery can be interrogated and examined
type ParsedQuery struct {
}

func (m *ParsedQuery) Reset()                    { *m = ParsedQuery{} }
func (m *ParsedQuery) String() string            { return proto.CompactTextString(m) }
func (*ParsedQuery) ProtoMessage()               {}
func (*ParsedQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*TextQuery)(nil), "promutils.TextQuery")
	proto.RegisterType((*ParsedQuery)(nil), "promutils.ParsedQuery")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PromUtils service

type PromUtilsClient interface {
	// Query a textual query and turn it into something useful
	ParseQuery(ctx context.Context, in *TextQuery, opts ...grpc.CallOption) (*ParsedQuery, error)
}

type promUtilsClient struct {
	cc *grpc.ClientConn
}

func NewPromUtilsClient(cc *grpc.ClientConn) PromUtilsClient {
	return &promUtilsClient{cc}
}

func (c *promUtilsClient) ParseQuery(ctx context.Context, in *TextQuery, opts ...grpc.CallOption) (*ParsedQuery, error) {
	out := new(ParsedQuery)
	err := grpc.Invoke(ctx, "/promutils.PromUtils/ParseQuery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PromUtils service

type PromUtilsServer interface {
	// Query a textual query and turn it into something useful
	ParseQuery(context.Context, *TextQuery) (*ParsedQuery, error)
}

func RegisterPromUtilsServer(s *grpc.Server, srv PromUtilsServer) {
	s.RegisterService(&_PromUtils_serviceDesc, srv)
}

func _PromUtils_ParseQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromUtilsServer).ParseQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/promutils.PromUtils/ParseQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromUtilsServer).ParseQuery(ctx, req.(*TextQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _PromUtils_serviceDesc = grpc.ServiceDesc{
	ServiceName: "promutils.PromUtils",
	HandlerType: (*PromUtilsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParseQuery",
			Handler:    _PromUtils_ParseQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/promutils/promutils.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/promutils/promutils.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0xcf, 0x49,
	0xcc, 0x4b, 0xd7, 0x4b, 0xce, 0xcf, 0x2b, 0x4a, 0x4c, 0x29, 0xcf, 0xcf, 0x4f, 0xd1, 0xcb, 0x4b,
	0x2d, 0xd1, 0x4f, 0x2c, 0xc8, 0x2c, 0xd6, 0x2f, 0x28, 0xca, 0xcf, 0x2d, 0x2d, 0xc9, 0xcc, 0x41,
	0x62, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x71, 0xc2, 0x05, 0x94, 0xe4, 0xb9, 0x38, 0x43,
	0x52, 0x2b, 0x4a, 0x02, 0x4b, 0x53, 0x8b, 0x2a, 0x85, 0x84, 0xb8, 0x58, 0x40, 0x1c, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x89, 0x97, 0x8b, 0x3b, 0x20, 0xb1, 0xa8, 0x38, 0x35,
	0x05, 0xac, 0xc4, 0xc8, 0x9d, 0x8b, 0x33, 0xa0, 0x28, 0x3f, 0x37, 0x14, 0xa4, 0x59, 0xc8, 0x8a,
	0x8b, 0x0b, 0x2c, 0x07, 0xd1, 0x2d, 0xa2, 0x87, 0xb0, 0x07, 0x6e, 0xa6, 0x94, 0x18, 0x92, 0x28,
	0x92, 0x41, 0x4e, 0x6a, 0x5c, 0x2a, 0x79, 0xa9, 0x25, 0xc8, 0x0e, 0x87, 0x7a, 0x05, 0xe4, 0x76,
	0x84, 0x9e, 0x24, 0x36, 0xb0, 0x93, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0xe3, 0x44,
	0xb0, 0xea, 0x00, 0x00, 0x00,
}
