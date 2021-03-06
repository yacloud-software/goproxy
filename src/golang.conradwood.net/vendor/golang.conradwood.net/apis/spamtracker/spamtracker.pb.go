// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/spamtracker/spamtracker.proto
// DO NOT EDIT!

/*
Package spamtracker is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/spamtracker/spamtracker.proto

It has these top-level messages:
	PingResponse
	AddKeywordRequest
	Keyword
	KeywordGroup
*/
package spamtracker

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

type AppliesTo int32

const (
	AppliesTo_Undefined AppliesTo = 0
	AppliesTo_Subject   AppliesTo = 1
	AppliesTo_Body      AppliesTo = 2
	AppliesTo_From      AppliesTo = 3
)

var AppliesTo_name = map[int32]string{
	0: "Undefined",
	1: "Subject",
	2: "Body",
	3: "From",
}
var AppliesTo_value = map[string]int32{
	"Undefined": 0,
	"Subject":   1,
	"Body":      2,
	"From":      3,
}

func (x AppliesTo) String() string {
	return proto.EnumName(AppliesTo_name, int32(x))
}
func (AppliesTo) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

// *all* the keywords must be found in an email to be spam detected
type AddKeywordRequest struct {
	AppliesTo []AppliesTo `protobuf:"varint,1,rep,packed,name=AppliesTo,enum=spamtracker.AppliesTo" json:"AppliesTo,omitempty"`
	Keywords  []string    `protobuf:"bytes,2,rep,name=Keywords" json:"Keywords,omitempty"`
}

func (m *AddKeywordRequest) Reset()                    { *m = AddKeywordRequest{} }
func (m *AddKeywordRequest) String() string            { return proto.CompactTextString(m) }
func (*AddKeywordRequest) ProtoMessage()               {}
func (*AddKeywordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddKeywordRequest) GetAppliesTo() []AppliesTo {
	if m != nil {
		return m.AppliesTo
	}
	return nil
}

func (m *AddKeywordRequest) GetKeywords() []string {
	if m != nil {
		return m.Keywords
	}
	return nil
}

type Keyword struct {
	ID           uint64        `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Word         string        `protobuf:"bytes,2,opt,name=Word" json:"Word,omitempty"`
	KeywordGroup *KeywordGroup `protobuf:"bytes,3,opt,name=KeywordGroup" json:"KeywordGroup,omitempty"`
}

func (m *Keyword) Reset()                    { *m = Keyword{} }
func (m *Keyword) String() string            { return proto.CompactTextString(m) }
func (*Keyword) ProtoMessage()               {}
func (*Keyword) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Keyword) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Keyword) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func (m *Keyword) GetKeywordGroup() *KeywordGroup {
	if m != nil {
		return m.KeywordGroup
	}
	return nil
}

type KeywordGroup struct {
	ID        uint64    `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	AppliesTo AppliesTo `protobuf:"varint,2,opt,name=AppliesTo,enum=spamtracker.AppliesTo" json:"AppliesTo,omitempty"`
}

func (m *KeywordGroup) Reset()                    { *m = KeywordGroup{} }
func (m *KeywordGroup) String() string            { return proto.CompactTextString(m) }
func (*KeywordGroup) ProtoMessage()               {}
func (*KeywordGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *KeywordGroup) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *KeywordGroup) GetAppliesTo() AppliesTo {
	if m != nil {
		return m.AppliesTo
	}
	return AppliesTo_Undefined
}

func init() {
	proto.RegisterType((*PingResponse)(nil), "spamtracker.PingResponse")
	proto.RegisterType((*AddKeywordRequest)(nil), "spamtracker.AddKeywordRequest")
	proto.RegisterType((*Keyword)(nil), "spamtracker.Keyword")
	proto.RegisterType((*KeywordGroup)(nil), "spamtracker.KeywordGroup")
	proto.RegisterEnum("spamtracker.AppliesTo", AppliesTo_name, AppliesTo_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SpamTracker service

type SpamTrackerClient interface {
	// trigger a rewrite of the file (also happens automatically)
	RewriteFile(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*common.Void, error)
	// add keywords
	AddKeywords(ctx context.Context, in *AddKeywordRequest, opts ...grpc.CallOption) (*common.Void, error)
}

type spamTrackerClient struct {
	cc *grpc.ClientConn
}

func NewSpamTrackerClient(cc *grpc.ClientConn) SpamTrackerClient {
	return &spamTrackerClient{cc}
}

func (c *spamTrackerClient) RewriteFile(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/spamtracker.SpamTracker/RewriteFile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spamTrackerClient) AddKeywords(ctx context.Context, in *AddKeywordRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/spamtracker.SpamTracker/AddKeywords", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SpamTracker service

type SpamTrackerServer interface {
	// trigger a rewrite of the file (also happens automatically)
	RewriteFile(context.Context, *common.Void) (*common.Void, error)
	// add keywords
	AddKeywords(context.Context, *AddKeywordRequest) (*common.Void, error)
}

func RegisterSpamTrackerServer(s *grpc.Server, srv SpamTrackerServer) {
	s.RegisterService(&_SpamTracker_serviceDesc, srv)
}

func _SpamTracker_RewriteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpamTrackerServer).RewriteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spamtracker.SpamTracker/RewriteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpamTrackerServer).RewriteFile(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpamTracker_AddKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddKeywordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpamTrackerServer).AddKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spamtracker.SpamTracker/AddKeywords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpamTrackerServer).AddKeywords(ctx, req.(*AddKeywordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpamTracker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spamtracker.SpamTracker",
	HandlerType: (*SpamTrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RewriteFile",
			Handler:    _SpamTracker_RewriteFile_Handler,
		},
		{
			MethodName: "AddKeywords",
			Handler:    _SpamTracker_AddKeywords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/spamtracker/spamtracker.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/spamtracker/spamtracker.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x4e, 0xd4, 0x36, 0xe3, 0x50, 0xc2, 0x1e, 0x90, 0x13, 0x09, 0x30, 0x3e, 0x85, 0x1e,
	0x5c, 0x29, 0x54, 0x08, 0xa9, 0x17, 0x6a, 0x55, 0x45, 0x15, 0x17, 0xe4, 0x16, 0xb8, 0x70, 0xc0,
	0xf5, 0x0e, 0xd6, 0xd2, 0x78, 0x67, 0xd9, 0x5d, 0x2b, 0xea, 0xb5, 0x57, 0xfe, 0x80, 0x23, 0x5f,
	0x8a, 0xec, 0xb8, 0xc9, 0x1a, 0x44, 0x4f, 0x7e, 0xcf, 0x33, 0xfb, 0xde, 0xdb, 0x9d, 0x81, 0x37,
	0x25, 0x2d, 0x73, 0x59, 0x26, 0x05, 0x49, 0x9d, 0xf3, 0x15, 0x11, 0x4f, 0x24, 0xda, 0xc3, 0x5c,
	0x09, 0x73, 0x68, 0x54, 0x5e, 0x59, 0x9d, 0x17, 0xd7, 0xa8, 0x5d, 0x9c, 0x28, 0x4d, 0x96, 0x58,
	0xe0, 0xfc, 0x9a, 0x25, 0xf7, 0xc8, 0x14, 0x54, 0x55, 0x24, 0xbb, 0xcf, 0xfa, 0x70, 0x7c, 0x00,
	0xe3, 0x0f, 0x42, 0x96, 0x19, 0x1a, 0x45, 0xd2, 0x20, 0x9b, 0xc1, 0xde, 0x1d, 0x0e, 0xbd, 0xc8,
	0x9b, 0x8f, 0xb2, 0x0d, 0x8f, 0x11, 0x1e, 0x9f, 0x70, 0xfe, 0x1e, 0x6f, 0x56, 0xa4, 0x79, 0x86,
	0x3f, 0x6a, 0x34, 0x96, 0x1d, 0xc1, 0xe8, 0x44, 0xa9, 0xa5, 0x40, 0x73, 0x49, 0xa1, 0x17, 0x0d,
	0xe6, 0xfb, 0x8b, 0x27, 0x89, 0x1b, 0x72, 0x53, 0xcd, 0xb6, 0x8d, 0x8d, 0x4d, 0xa7, 0x63, 0x42,
	0x3f, 0x1a, 0x34, 0x36, 0x77, 0x3c, 0xfe, 0xe9, 0xc1, 0x6e, 0x47, 0xd8, 0x3e, 0xf8, 0xe7, 0xa7,
	0x6d, 0x90, 0x61, 0xe6, 0x9f, 0x9f, 0x32, 0x06, 0xc3, 0xcf, 0xa4, 0x79, 0xe8, 0xb7, 0xd1, 0x5a,
	0xcc, 0xbe, 0xc0, 0xb8, 0x6b, 0x7f, 0xa7, 0xa9, 0x56, 0xe1, 0x20, 0xf2, 0xe6, 0xc1, 0x62, 0xda,
	0x0b, 0xe1, 0x36, 0xa4, 0x4f, 0x7f, 0xdd, 0x4e, 0x77, 0x6a, 0x21, 0xed, 0xeb, 0xa3, 0xdf, 0xb7,
	0xd3, 0x47, 0xd7, 0xeb, 0x5a, 0xd9, 0xd4, 0x12, 0xc1, 0xb3, 0x9e, 0x5a, 0xfc, 0xb5, 0xaf, 0xfe,
	0x4f, 0xa2, 0xb7, 0xee, 0xfd, 0x9b, 0x58, 0xff, 0xbd, 0x7f, 0x0a, 0x5b, 0x5f, 0xe7, 0x2d, 0x0e,
	0x8e, 0x1d, 0x05, 0xf6, 0x10, 0x46, 0x1f, 0x25, 0xc7, 0x6f, 0x42, 0x22, 0x9f, 0x3c, 0x60, 0x01,
	0xec, 0x5e, 0xd4, 0x57, 0xdf, 0xb1, 0xb0, 0x13, 0x8f, 0xed, 0xc1, 0x30, 0x25, 0x7e, 0x33, 0xf1,
	0x1b, 0x74, 0xa6, 0xa9, 0x9a, 0x0c, 0x16, 0x35, 0x04, 0x17, 0x2a, 0xaf, 0x2e, 0xd7, 0x66, 0xec,
	0x25, 0x04, 0x19, 0xae, 0xb4, 0xb0, 0x78, 0x26, 0x96, 0xc8, 0xc6, 0x49, 0x37, 0xec, 0x4f, 0x24,
	0xf8, 0xac, 0xc7, 0xd8, 0x31, 0x04, 0xdb, 0x69, 0x1a, 0xf6, 0xac, 0x1f, 0xfa, 0xef, 0x39, 0xf7,
	0x0f, 0xa7, 0x2f, 0xe0, 0xb9, 0x44, 0xeb, 0x6e, 0x59, 0xb3, 0x61, 0xae, 0xc4, 0xd5, 0x4e, 0xbb,
	0x60, 0xaf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x95, 0xaa, 0xb7, 0xd6, 0xd9, 0x02, 0x00, 0x00,
}
