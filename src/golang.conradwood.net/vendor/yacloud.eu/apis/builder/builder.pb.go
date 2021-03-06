// Code generated by protoc-gen-go.
// source: yacloud.eu/apis/builder/builder.proto
// DO NOT EDIT!

/*
Package builder is a generated protocol buffer package.

It is generated from these files:
	yacloud.eu/apis/builder/builder.proto

It has these top-level messages:
	Link
	ServeResponse
*/
package builder

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import h2gproxy "golang.conradwood.net/apis/h2gproxy"

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

type Link struct {
	ID   string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
}

func (m *Link) Reset()                    { *m = Link{} }
func (m *Link) String() string            { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()               {}
func (*Link) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Link) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Link) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ServeResponse struct {
	MimeType string `protobuf:"bytes,1,opt,name=MimeType" json:"MimeType,omitempty"`
	Body     []byte `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
	// links for the menu to serve
	Links []*Link `protobuf:"bytes,3,rep,name=Links" json:"Links,omitempty"`
	// if true, won't add header/footer
	ServeAsIs bool   `protobuf:"varint,4,opt,name=ServeAsIs" json:"ServeAsIs,omitempty"`
	Title     string `protobuf:"bytes,5,opt,name=Title" json:"Title,omitempty"`
}

func (m *ServeResponse) Reset()                    { *m = ServeResponse{} }
func (m *ServeResponse) String() string            { return proto.CompactTextString(m) }
func (*ServeResponse) ProtoMessage()               {}
func (*ServeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServeResponse) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *ServeResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *ServeResponse) GetLinks() []*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *ServeResponse) GetServeAsIs() bool {
	if m != nil {
		return m.ServeAsIs
	}
	return false
}

func (m *ServeResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func init() {
	proto.RegisterType((*Link)(nil), "builder.Link")
	proto.RegisterType((*ServeResponse)(nil), "builder.ServeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Builder service

type BuilderClient interface {
	ServeHTML(ctx context.Context, in *h2gproxy.ServeRequest, opts ...grpc.CallOption) (*ServeResponse, error)
}

type builderClient struct {
	cc *grpc.ClientConn
}

func NewBuilderClient(cc *grpc.ClientConn) BuilderClient {
	return &builderClient{cc}
}

func (c *builderClient) ServeHTML(ctx context.Context, in *h2gproxy.ServeRequest, opts ...grpc.CallOption) (*ServeResponse, error) {
	out := new(ServeResponse)
	err := grpc.Invoke(ctx, "/builder.Builder/ServeHTML", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Builder service

type BuilderServer interface {
	ServeHTML(context.Context, *h2gproxy.ServeRequest) (*ServeResponse, error)
}

func RegisterBuilderServer(s *grpc.Server, srv BuilderServer) {
	s.RegisterService(&_Builder_serviceDesc, srv)
}

func _Builder_ServeHTML_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(h2gproxy.ServeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServer).ServeHTML(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/builder.Builder/ServeHTML",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServer).ServeHTML(ctx, req.(*h2gproxy.ServeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Builder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "builder.Builder",
	HandlerType: (*BuilderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServeHTML",
			Handler:    _Builder_ServeHTML_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yacloud.eu/apis/builder/builder.proto",
}

func init() { proto.RegisterFile("yacloud.eu/apis/builder/builder.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x50, 0xc1, 0x4e, 0x83, 0x40,
	0x14, 0x0c, 0x14, 0x6c, 0xfb, 0xb4, 0x1e, 0x36, 0xa6, 0x22, 0xf1, 0x40, 0x6a, 0x4c, 0x88, 0x87,
	0x6d, 0x82, 0x47, 0x4f, 0x92, 0xc6, 0x48, 0xd2, 0x7a, 0x40, 0x7e, 0x80, 0x96, 0x17, 0x24, 0x52,
	0x1e, 0xb2, 0xa0, 0xf2, 0x27, 0x7e, 0xae, 0x61, 0x17, 0x30, 0x9e, 0xf6, 0xbd, 0xd9, 0x9d, 0x99,
	0x9d, 0x81, 0xdb, 0x36, 0x3e, 0xe4, 0xd4, 0x24, 0x1c, 0x9b, 0x75, 0x5c, 0x66, 0x62, 0xbd, 0x6f,
	0xb2, 0x3c, 0xc1, 0x6a, 0x38, 0x79, 0x59, 0x51, 0x4d, 0x6c, 0xda, 0xaf, 0xb6, 0x97, 0x52, 0x1e,
	0x17, 0x29, 0x3f, 0x50, 0x51, 0xc5, 0xc9, 0x17, 0x51, 0xc2, 0x0b, 0xac, 0x15, 0xf5, 0xcd, 0x4b,
	0xcb, 0x8a, 0xbe, 0xdb, 0x71, 0x50, 0xe4, 0xd5, 0x1d, 0x18, 0xdb, 0xac, 0x78, 0x67, 0xe7, 0xa0,
	0x07, 0x1b, 0x4b, 0x73, 0x34, 0x77, 0x1e, 0xea, 0xc1, 0x86, 0x31, 0x30, 0x5e, 0xe2, 0x23, 0x5a,
	0xba, 0x44, 0xe4, 0xbc, 0xfa, 0xd1, 0x60, 0xf1, 0x8a, 0xd5, 0x27, 0x86, 0x28, 0x4a, 0x2a, 0x04,
	0x32, 0x1b, 0x66, 0xbb, 0xec, 0x88, 0x51, 0x5b, 0x62, 0xcf, 0x1d, 0xf7, 0x4e, 0xc1, 0xa7, 0xa4,
	0x95, 0x0a, 0x67, 0xa1, 0x9c, 0xd9, 0x0d, 0x98, 0x9d, 0x9b, 0xb0, 0x26, 0xce, 0xc4, 0x3d, 0xf5,
	0x16, 0x7c, 0x48, 0xd2, 0xa1, 0xa1, 0xba, 0x63, 0xd7, 0x30, 0x97, 0x2e, 0x8f, 0x22, 0x10, 0x96,
	0xe1, 0x68, 0xee, 0x2c, 0xfc, 0x03, 0xd8, 0x05, 0x98, 0x51, 0x56, 0xe7, 0x68, 0x99, 0xd2, 0x4f,
	0x2d, 0xde, 0x13, 0x4c, 0x7d, 0x25, 0xc5, 0x1e, 0x7a, 0xfa, 0x73, 0xb4, 0xdb, 0xb2, 0x25, 0x1f,
	0xf3, 0xf6, 0x3f, 0xff, 0x68, 0x50, 0xd4, 0xf6, 0x72, 0x74, 0xfe, 0x17, 0xc8, 0xbf, 0x82, 0x4b,
	0x6c, 0xf8, 0xd0, 0x7b, 0xd7, 0xdc, 0xf0, 0x70, 0x7f, 0x22, 0x0b, 0xbb, 0xff, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x15, 0x68, 0x8b, 0xbd, 0x96, 0x01, 0x00, 0x00,
}
