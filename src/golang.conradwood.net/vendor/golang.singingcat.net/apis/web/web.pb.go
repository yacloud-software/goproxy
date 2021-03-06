// Code generated by protoc-gen-go.
// source: golang.singingcat.net/apis/web/web.proto
// DO NOT EDIT!

/*
Package web is a generated protocol buffer package.

It is generated from these files:
	golang.singingcat.net/apis/web/web.proto

It has these top-level messages:
	SensorSaveRequest
	ColourRequest
	ColourResponse
*/
package web

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import singingcat "golang.singingcat.net/apis/singingcat"
import sensors "golang.singingcat.net/apis/sensors"

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

type SensorSaveRequest struct {
	Node       uint32               `protobuf:"varint,1,opt,name=Node" json:"Node,omitempty"`
	SensorName string               `protobuf:"bytes,2,opt,name=SensorName" json:"SensorName,omitempty"`
	RAWValue   *sensors.SensorValue `protobuf:"bytes,3,opt,name=RAWValue" json:"RAWValue,omitempty"`
}

func (m *SensorSaveRequest) Reset()                    { *m = SensorSaveRequest{} }
func (m *SensorSaveRequest) String() string            { return proto.CompactTextString(m) }
func (*SensorSaveRequest) ProtoMessage()               {}
func (*SensorSaveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SensorSaveRequest) GetNode() uint32 {
	if m != nil {
		return m.Node
	}
	return 0
}

func (m *SensorSaveRequest) GetSensorName() string {
	if m != nil {
		return m.SensorName
	}
	return ""
}

func (m *SensorSaveRequest) GetRAWValue() *sensors.SensorValue {
	if m != nil {
		return m.RAWValue
	}
	return nil
}

type ColourRequest struct {
	Text string `protobuf:"bytes,1,opt,name=Text" json:"Text,omitempty"`
}

func (m *ColourRequest) Reset()                    { *m = ColourRequest{} }
func (m *ColourRequest) String() string            { return proto.CompactTextString(m) }
func (*ColourRequest) ProtoMessage()               {}
func (*ColourRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ColourRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type ColourResponse struct {
	RGB        uint64 `protobuf:"varint,1,opt,name=RGB" json:"RGB,omitempty"`
	Brightness uint32 `protobuf:"varint,2,opt,name=Brightness" json:"Brightness,omitempty"`
}

func (m *ColourResponse) Reset()                    { *m = ColourResponse{} }
func (m *ColourResponse) String() string            { return proto.CompactTextString(m) }
func (*ColourResponse) ProtoMessage()               {}
func (*ColourResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ColourResponse) GetRGB() uint64 {
	if m != nil {
		return m.RGB
	}
	return 0
}

func (m *ColourResponse) GetBrightness() uint32 {
	if m != nil {
		return m.Brightness
	}
	return 0
}

func init() {
	proto.RegisterType((*SensorSaveRequest)(nil), "web.SensorSaveRequest")
	proto.RegisterType((*ColourRequest)(nil), "web.ColourRequest")
	proto.RegisterType((*ColourResponse)(nil), "web.ColourResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Web service

type WebClient interface {
	Ping(ctx context.Context, in *singingcat.Void, opts ...grpc.CallOption) (*singingcat.Void, error)
	// process & save the value
	ProcessSensorValue(ctx context.Context, in *SensorSaveRequest, opts ...grpc.CallOption) (*sensors.SensorRef, error)
}

type webClient struct {
	cc *grpc.ClientConn
}

func NewWebClient(cc *grpc.ClientConn) WebClient {
	return &webClient{cc}
}

func (c *webClient) Ping(ctx context.Context, in *singingcat.Void, opts ...grpc.CallOption) (*singingcat.Void, error) {
	out := new(singingcat.Void)
	err := grpc.Invoke(ctx, "/web.Web/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) ProcessSensorValue(ctx context.Context, in *SensorSaveRequest, opts ...grpc.CallOption) (*sensors.SensorRef, error) {
	out := new(sensors.SensorRef)
	err := grpc.Invoke(ctx, "/web.Web/ProcessSensorValue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Web service

type WebServer interface {
	Ping(context.Context, *singingcat.Void) (*singingcat.Void, error)
	// process & save the value
	ProcessSensorValue(context.Context, *SensorSaveRequest) (*sensors.SensorRef, error)
}

func RegisterWebServer(s *grpc.Server, srv WebServer) {
	s.RegisterService(&_Web_serviceDesc, srv)
}

func _Web_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(singingcat.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Web/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).Ping(ctx, req.(*singingcat.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_ProcessSensorValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SensorSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).ProcessSensorValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Web/ProcessSensorValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).ProcessSensorValue(ctx, req.(*SensorSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Web_serviceDesc = grpc.ServiceDesc{
	ServiceName: "web.Web",
	HandlerType: (*WebServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Web_Ping_Handler,
		},
		{
			MethodName: "ProcessSensorValue",
			Handler:    _Web_ProcessSensorValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.singingcat.net/apis/web/web.proto",
}

func init() { proto.RegisterFile("golang.singingcat.net/apis/web/web.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x65, 0xdd, 0x22, 0x76, 0xa4, 0x52, 0x07, 0x91, 0xb2, 0x87, 0xb2, 0xd4, 0xcb, 0xe2, 0x61,
	0x5b, 0x2a, 0x78, 0xd6, 0xf5, 0xe0, 0xad, 0x94, 0x54, 0xda, 0x73, 0xb6, 0x1d, 0x63, 0xa0, 0x26,
	0x75, 0x93, 0xb5, 0xfa, 0xf7, 0x92, 0xac, 0xad, 0xc1, 0x42, 0x0f, 0x21, 0x8f, 0x97, 0x37, 0x33,
	0x2f, 0xf3, 0x20, 0x13, 0x7a, 0xcd, 0x95, 0xc8, 0x8d, 0x54, 0x42, 0x2a, 0xb1, 0xe4, 0x36, 0x57,
	0x64, 0x87, 0x7c, 0x23, 0xcd, 0x70, 0x4b, 0xa5, 0x3b, 0xf9, 0xa6, 0xd2, 0x56, 0x63, 0xbc, 0xa5,
	0x32, 0xb9, 0x3f, 0x22, 0xff, 0xe3, 0x02, 0xd8, 0x14, 0x27, 0xa3, 0x63, 0x75, 0xa4, 0x8c, 0xae,
	0xf6, 0x77, 0x53, 0x31, 0xf8, 0x86, 0xcb, 0x99, 0x27, 0x66, 0xfc, 0x93, 0x18, 0x7d, 0xd4, 0x64,
	0x2c, 0x22, 0xb4, 0x26, 0x7a, 0x45, 0xbd, 0x28, 0x8d, 0xb2, 0x0e, 0xf3, 0x18, 0xfb, 0x00, 0x8d,
	0x70, 0xc2, 0xdf, 0xa9, 0x77, 0x92, 0x46, 0x59, 0x9b, 0x05, 0x0c, 0x8e, 0xe0, 0x8c, 0x3d, 0x2e,
	0xe6, 0x7c, 0x5d, 0x53, 0x2f, 0x4e, 0xa3, 0xec, 0x7c, 0x7c, 0x95, 0xef, 0x46, 0x35, 0x32, 0xff,
	0xc6, 0xf6, 0xaa, 0xc1, 0x0d, 0x74, 0x9e, 0xf4, 0x5a, 0xd7, 0x55, 0x30, 0xf6, 0x85, 0xbe, 0xac,
	0x1f, 0xdb, 0x66, 0x1e, 0x0f, 0x0a, 0xb8, 0xd8, 0x89, 0xcc, 0x46, 0x2b, 0x43, 0xd8, 0x85, 0x98,
	0x3d, 0x17, 0x5e, 0xd4, 0x62, 0x0e, 0x3a, 0x6b, 0x45, 0x25, 0xc5, 0x9b, 0x55, 0x64, 0x8c, 0xb7,
	0xd6, 0x61, 0x01, 0x33, 0x36, 0x10, 0x2f, 0xa8, 0xc4, 0x5b, 0x68, 0x4d, 0xa5, 0x12, 0xd8, 0x0d,
	0xd7, 0x33, 0xd7, 0x72, 0x95, 0x1c, 0x30, 0xf8, 0x00, 0x38, 0xad, 0xf4, 0x92, 0x8c, 0x09, 0xbc,
	0xe3, 0x75, 0xee, 0x72, 0x3a, 0xd8, 0x57, 0x82, 0xff, 0x7e, 0xca, 0xe8, 0xb5, 0x48, 0xa1, 0xaf,
	0xc8, 0x86, 0x8d, 0x7f, 0xb3, 0x71, 0x61, 0xb8, 0x3e, 0xe5, 0xa9, 0x4f, 0xe0, 0xee, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x48, 0xd4, 0x37, 0x5c, 0x1c, 0x02, 0x00, 0x00,
}
