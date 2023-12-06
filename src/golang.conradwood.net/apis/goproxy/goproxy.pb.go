// Code generated by protoc-gen-go.
// source: protos/golang.conradwood.net/apis/goproxy/goproxy.proto
// DO NOT EDIT!

/*
Package goproxy is a generated protocol buffer package.

It is generated from these files:
	protos/golang.conradwood.net/apis/goproxy/goproxy.proto

It has these top-level messages:
	Config
	UpStreamProxy
	ArtefactDef
	CachedModule
	ModuleInfoRequest
	ModuleInfo
	VersionInfo
	GetPathRequest
	BinData
	Override
*/
package goproxy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "golang.conradwood.net/apis/common"
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

type MODULETYPE int32

const (
	MODULETYPE_UNKNOWN        MODULETYPE = 0
	MODULETYPE_PROTO          MODULETYPE = 2
	MODULETYPE_ARTEFACT       MODULETYPE = 3
	MODULETYPE_EXTERNALMODULE MODULETYPE = 4
	MODULETYPE_UPSTREAMPROXY  MODULETYPE = 5
)

var MODULETYPE_name = map[int32]string{
	0: "UNKNOWN",
	2: "PROTO",
	3: "ARTEFACT",
	4: "EXTERNALMODULE",
	5: "UPSTREAMPROXY",
}
var MODULETYPE_value = map[string]int32{
	"UNKNOWN":        0,
	"PROTO":          2,
	"ARTEFACT":       3,
	"EXTERNALMODULE": 4,
	"UPSTREAMPROXY":  5,
}

func (x MODULETYPE) String() string {
	return proto.EnumName(MODULETYPE_name, int32(x))
}
func (MODULETYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Config struct {
	LocalHosts        []string         `protobuf:"bytes,1,rep,name=LocalHosts" json:"LocalHosts,omitempty"`
	ArtefactResolvers []*ArtefactDef   `protobuf:"bytes,2,rep,name=ArtefactResolvers" json:"ArtefactResolvers,omitempty"`
	GoGetProxy        string           `protobuf:"bytes,3,opt,name=GoGetProxy" json:"GoGetProxy,omitempty"`
	GoProxies         []*UpStreamProxy `protobuf:"bytes,4,rep,name=GoProxies" json:"GoProxies,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetLocalHosts() []string {
	if m != nil {
		return m.LocalHosts
	}
	return nil
}

func (m *Config) GetArtefactResolvers() []*ArtefactDef {
	if m != nil {
		return m.ArtefactResolvers
	}
	return nil
}

func (m *Config) GetGoGetProxy() string {
	if m != nil {
		return m.GoGetProxy
	}
	return ""
}

func (m *Config) GetGoProxies() []*UpStreamProxy {
	if m != nil {
		return m.GoProxies
	}
	return nil
}

type UpStreamProxy struct {
	Matcher  string `protobuf:"bytes,1,opt,name=Matcher" json:"Matcher,omitempty"`
	Proxy    string `protobuf:"bytes,2,opt,name=Proxy" json:"Proxy,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=Username" json:"Username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=Password" json:"Password,omitempty"`
	Token    string `protobuf:"bytes,5,opt,name=Token" json:"Token,omitempty"`
}

func (m *UpStreamProxy) Reset()                    { *m = UpStreamProxy{} }
func (m *UpStreamProxy) String() string            { return proto.CompactTextString(m) }
func (*UpStreamProxy) ProtoMessage()               {}
func (*UpStreamProxy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpStreamProxy) GetMatcher() string {
	if m != nil {
		return m.Matcher
	}
	return ""
}

func (m *UpStreamProxy) GetProxy() string {
	if m != nil {
		return m.Proxy
	}
	return ""
}

func (m *UpStreamProxy) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UpStreamProxy) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UpStreamProxy) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ArtefactDef struct {
	Path       string `protobuf:"bytes,1,opt,name=Path" json:"Path,omitempty"`
	ArtefactID uint64 `protobuf:"varint,2,opt,name=ArtefactID" json:"ArtefactID,omitempty"`
	Domain     string `protobuf:"bytes,3,opt,name=Domain" json:"Domain,omitempty"`
	Name       string `protobuf:"bytes,4,opt,name=Name" json:"Name,omitempty"`
}

func (m *ArtefactDef) Reset()                    { *m = ArtefactDef{} }
func (m *ArtefactDef) String() string            { return proto.CompactTextString(m) }
func (*ArtefactDef) ProtoMessage()               {}
func (*ArtefactDef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ArtefactDef) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ArtefactDef) GetArtefactID() uint64 {
	if m != nil {
		return m.ArtefactID
	}
	return 0
}

func (m *ArtefactDef) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *ArtefactDef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CachedModule struct {
	ID           uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Path         string `protobuf:"bytes,2,opt,name=Path" json:"Path,omitempty"`
	Version      string `protobuf:"bytes,3,opt,name=Version" json:"Version,omitempty"`
	Suffix       string `protobuf:"bytes,4,opt,name=Suffix" json:"Suffix,omitempty"`
	Key          string `protobuf:"bytes,5,opt,name=Key" json:"Key,omitempty"`
	Created      uint32 `protobuf:"varint,6,opt,name=Created" json:"Created,omitempty"`
	LastUsed     uint32 `protobuf:"varint,7,opt,name=LastUsed" json:"LastUsed,omitempty"`
	ToBeDeleted  bool   `protobuf:"varint,8,opt,name=ToBeDeleted" json:"ToBeDeleted,omitempty"`
	FailingSince uint32 `protobuf:"varint,9,opt,name=FailingSince" json:"FailingSince,omitempty"`
	FailCounter  uint32 `protobuf:"varint,10,opt,name=FailCounter" json:"FailCounter,omitempty"`
	LastFailed   uint32 `protobuf:"varint,11,opt,name=LastFailed" json:"LastFailed,omitempty"`
	PutFailed    bool   `protobuf:"varint,12,opt,name=PutFailed" json:"PutFailed,omitempty"`
	PutError     string `protobuf:"bytes,13,opt,name=PutError" json:"PutError,omitempty"`
	Size         uint64 `protobuf:"varint,14,opt,name=Size" json:"Size,omitempty"`
}

func (m *CachedModule) Reset()                    { *m = CachedModule{} }
func (m *CachedModule) String() string            { return proto.CompactTextString(m) }
func (*CachedModule) ProtoMessage()               {}
func (*CachedModule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CachedModule) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *CachedModule) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CachedModule) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CachedModule) GetSuffix() string {
	if m != nil {
		return m.Suffix
	}
	return ""
}

func (m *CachedModule) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CachedModule) GetCreated() uint32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *CachedModule) GetLastUsed() uint32 {
	if m != nil {
		return m.LastUsed
	}
	return 0
}

func (m *CachedModule) GetToBeDeleted() bool {
	if m != nil {
		return m.ToBeDeleted
	}
	return false
}

func (m *CachedModule) GetFailingSince() uint32 {
	if m != nil {
		return m.FailingSince
	}
	return 0
}

func (m *CachedModule) GetFailCounter() uint32 {
	if m != nil {
		return m.FailCounter
	}
	return 0
}

func (m *CachedModule) GetLastFailed() uint32 {
	if m != nil {
		return m.LastFailed
	}
	return 0
}

func (m *CachedModule) GetPutFailed() bool {
	if m != nil {
		return m.PutFailed
	}
	return false
}

func (m *CachedModule) GetPutError() string {
	if m != nil {
		return m.PutError
	}
	return ""
}

func (m *CachedModule) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ModuleInfoRequest struct {
	URL string `protobuf:"bytes,1,opt,name=URL" json:"URL,omitempty"`
}

func (m *ModuleInfoRequest) Reset()                    { *m = ModuleInfoRequest{} }
func (m *ModuleInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*ModuleInfoRequest) ProtoMessage()               {}
func (*ModuleInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ModuleInfoRequest) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

type ModuleInfo struct {
	ModuleType MODULETYPE `protobuf:"varint,1,opt,name=ModuleType,enum=goproxy.MODULETYPE" json:"ModuleType,omitempty"`
	Exists     bool       `protobuf:"varint,2,opt,name=Exists" json:"Exists,omitempty"`
}

func (m *ModuleInfo) Reset()                    { *m = ModuleInfo{} }
func (m *ModuleInfo) String() string            { return proto.CompactTextString(m) }
func (*ModuleInfo) ProtoMessage()               {}
func (*ModuleInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ModuleInfo) GetModuleType() MODULETYPE {
	if m != nil {
		return m.ModuleType
	}
	return MODULETYPE_UNKNOWN
}

func (m *ModuleInfo) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

type VersionInfo struct {
	Version     uint64 `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	BuildTime   uint32 `protobuf:"varint,2,opt,name=BuildTime" json:"BuildTime,omitempty"`
	VersionName string `protobuf:"bytes,3,opt,name=VersionName" json:"VersionName,omitempty"`
}

func (m *VersionInfo) Reset()                    { *m = VersionInfo{} }
func (m *VersionInfo) String() string            { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()               {}
func (*VersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *VersionInfo) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *VersionInfo) GetBuildTime() uint32 {
	if m != nil {
		return m.BuildTime
	}
	return 0
}

func (m *VersionInfo) GetVersionName() string {
	if m != nil {
		return m.VersionName
	}
	return ""
}

type GetPathRequest struct {
	Path string `protobuf:"bytes,1,opt,name=Path" json:"Path,omitempty"`
}

func (m *GetPathRequest) Reset()                    { *m = GetPathRequest{} }
func (m *GetPathRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPathRequest) ProtoMessage()               {}
func (*GetPathRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetPathRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type BinData struct {
	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *BinData) Reset()                    { *m = BinData{} }
func (m *BinData) String() string            { return proto.CompactTextString(m) }
func (*BinData) ProtoMessage()               {}
func (*BinData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *BinData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// given a certain package we might always only want to serve a given package
type Override struct {
	ID      uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Package string `protobuf:"bytes,2,opt,name=Package" json:"Package,omitempty"`
	List    []byte `protobuf:"bytes,3,opt,name=List,proto3" json:"List,omitempty"`
}

func (m *Override) Reset()                    { *m = Override{} }
func (m *Override) String() string            { return proto.CompactTextString(m) }
func (*Override) ProtoMessage()               {}
func (*Override) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Override) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Override) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *Override) GetList() []byte {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "goproxy.Config")
	proto.RegisterType((*UpStreamProxy)(nil), "goproxy.UpStreamProxy")
	proto.RegisterType((*ArtefactDef)(nil), "goproxy.ArtefactDef")
	proto.RegisterType((*CachedModule)(nil), "goproxy.CachedModule")
	proto.RegisterType((*ModuleInfoRequest)(nil), "goproxy.ModuleInfoRequest")
	proto.RegisterType((*ModuleInfo)(nil), "goproxy.ModuleInfo")
	proto.RegisterType((*VersionInfo)(nil), "goproxy.VersionInfo")
	proto.RegisterType((*GetPathRequest)(nil), "goproxy.GetPathRequest")
	proto.RegisterType((*BinData)(nil), "goproxy.BinData")
	proto.RegisterType((*Override)(nil), "goproxy.Override")
	proto.RegisterEnum("goproxy.MODULETYPE", MODULETYPE_name, MODULETYPE_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GoProxy service

type GoProxyClient interface {
	// h2gproxy endpoint
	StreamHTTP(ctx context.Context, in *h2gproxy.StreamRequest, opts ...grpc.CallOption) (GoProxy_StreamHTTPClient, error)
	AnalyseURL(ctx context.Context, in *ModuleInfoRequest, opts ...grpc.CallOption) (*ModuleInfo, error)
	// given a path, will download response
	GetPath(ctx context.Context, in *GetPathRequest, opts ...grpc.CallOption) (GoProxy_GetPathClient, error)
	AddOverride(ctx context.Context, in *Override, opts ...grpc.CallOption) (*common.Void, error)
}

type goProxyClient struct {
	cc *grpc.ClientConn
}

func NewGoProxyClient(cc *grpc.ClientConn) GoProxyClient {
	return &goProxyClient{cc}
}

func (c *goProxyClient) StreamHTTP(ctx context.Context, in *h2gproxy.StreamRequest, opts ...grpc.CallOption) (GoProxy_StreamHTTPClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GoProxy_serviceDesc.Streams[0], c.cc, "/goproxy.GoProxy/StreamHTTP", opts...)
	if err != nil {
		return nil, err
	}
	x := &goProxyStreamHTTPClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoProxy_StreamHTTPClient interface {
	Recv() (*h2gproxy.StreamDataResponse, error)
	grpc.ClientStream
}

type goProxyStreamHTTPClient struct {
	grpc.ClientStream
}

func (x *goProxyStreamHTTPClient) Recv() (*h2gproxy.StreamDataResponse, error) {
	m := new(h2gproxy.StreamDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *goProxyClient) AnalyseURL(ctx context.Context, in *ModuleInfoRequest, opts ...grpc.CallOption) (*ModuleInfo, error) {
	out := new(ModuleInfo)
	err := grpc.Invoke(ctx, "/goproxy.GoProxy/AnalyseURL", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goProxyClient) GetPath(ctx context.Context, in *GetPathRequest, opts ...grpc.CallOption) (GoProxy_GetPathClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GoProxy_serviceDesc.Streams[1], c.cc, "/goproxy.GoProxy/GetPath", opts...)
	if err != nil {
		return nil, err
	}
	x := &goProxyGetPathClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoProxy_GetPathClient interface {
	Recv() (*BinData, error)
	grpc.ClientStream
}

type goProxyGetPathClient struct {
	grpc.ClientStream
}

func (x *goProxyGetPathClient) Recv() (*BinData, error) {
	m := new(BinData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *goProxyClient) AddOverride(ctx context.Context, in *Override, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/goproxy.GoProxy/AddOverride", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoProxy service

type GoProxyServer interface {
	// h2gproxy endpoint
	StreamHTTP(*h2gproxy.StreamRequest, GoProxy_StreamHTTPServer) error
	AnalyseURL(context.Context, *ModuleInfoRequest) (*ModuleInfo, error)
	// given a path, will download response
	GetPath(*GetPathRequest, GoProxy_GetPathServer) error
	AddOverride(context.Context, *Override) (*common.Void, error)
}

func RegisterGoProxyServer(s *grpc.Server, srv GoProxyServer) {
	s.RegisterService(&_GoProxy_serviceDesc, srv)
}

func _GoProxy_StreamHTTP_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(h2gproxy.StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoProxyServer).StreamHTTP(m, &goProxyStreamHTTPServer{stream})
}

type GoProxy_StreamHTTPServer interface {
	Send(*h2gproxy.StreamDataResponse) error
	grpc.ServerStream
}

type goProxyStreamHTTPServer struct {
	grpc.ServerStream
}

func (x *goProxyStreamHTTPServer) Send(m *h2gproxy.StreamDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GoProxy_AnalyseURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModuleInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoProxyServer).AnalyseURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goproxy.GoProxy/AnalyseURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoProxyServer).AnalyseURL(ctx, req.(*ModuleInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoProxy_GetPath_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetPathRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoProxyServer).GetPath(m, &goProxyGetPathServer{stream})
}

type GoProxy_GetPathServer interface {
	Send(*BinData) error
	grpc.ServerStream
}

type goProxyGetPathServer struct {
	grpc.ServerStream
}

func (x *goProxyGetPathServer) Send(m *BinData) error {
	return x.ServerStream.SendMsg(m)
}

func _GoProxy_AddOverride_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Override)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoProxyServer).AddOverride(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goproxy.GoProxy/AddOverride",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoProxyServer).AddOverride(ctx, req.(*Override))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goproxy.GoProxy",
	HandlerType: (*GoProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnalyseURL",
			Handler:    _GoProxy_AnalyseURL_Handler,
		},
		{
			MethodName: "AddOverride",
			Handler:    _GoProxy_AddOverride_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamHTTP",
			Handler:       _GoProxy_StreamHTTP_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetPath",
			Handler:       _GoProxy_GetPath_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/golang.conradwood.net/apis/goproxy/goproxy.proto",
}

// Client API for GoProxyTestRunner service

type GoProxyTestRunnerClient interface {
	Trigger(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*common.Void, error)
}

type goProxyTestRunnerClient struct {
	cc *grpc.ClientConn
}

func NewGoProxyTestRunnerClient(cc *grpc.ClientConn) GoProxyTestRunnerClient {
	return &goProxyTestRunnerClient{cc}
}

func (c *goProxyTestRunnerClient) Trigger(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/goproxy.GoProxyTestRunner/Trigger", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoProxyTestRunner service

type GoProxyTestRunnerServer interface {
	Trigger(context.Context, *common.Void) (*common.Void, error)
}

func RegisterGoProxyTestRunnerServer(s *grpc.Server, srv GoProxyTestRunnerServer) {
	s.RegisterService(&_GoProxyTestRunner_serviceDesc, srv)
}

func _GoProxyTestRunner_Trigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoProxyTestRunnerServer).Trigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goproxy.GoProxyTestRunner/Trigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoProxyTestRunnerServer).Trigger(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoProxyTestRunner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goproxy.GoProxyTestRunner",
	HandlerType: (*GoProxyTestRunnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Trigger",
			Handler:    _GoProxyTestRunner_Trigger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/golang.conradwood.net/apis/goproxy/goproxy.proto",
}

func init() {
	proto.RegisterFile("protos/golang.conradwood.net/apis/goproxy/goproxy.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 907 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x55, 0xdf, 0x8f, 0xda, 0x46,
	0x10, 0xae, 0x81, 0x3b, 0x1f, 0xc3, 0x0f, 0xc1, 0x26, 0x4a, 0x2c, 0x7a, 0xad, 0x90, 0xd5, 0x48,
	0xa7, 0x3e, 0x90, 0x88, 0x54, 0xad, 0x54, 0xa9, 0x0f, 0xfc, 0xca, 0xdd, 0x29, 0x1c, 0x58, 0x8b,
	0x49, 0x73, 0x7d, 0xdb, 0xe2, 0x01, 0xac, 0x80, 0x97, 0xae, 0x97, 0xe4, 0xe8, 0x63, 0xdf, 0xfb,
	0x4f, 0xf5, 0x9f, 0xea, 0x6b, 0xb5, 0xeb, 0xb5, 0x31, 0x49, 0x94, 0x3e, 0x79, 0xbe, 0x6f, 0xbe,
	0x9d, 0x99, 0x9d, 0x59, 0xef, 0xc2, 0x4f, 0x3b, 0xc1, 0x25, 0x8f, 0x9f, 0xaf, 0xf8, 0x86, 0x45,
	0xab, 0xce, 0x82, 0x47, 0x82, 0x05, 0x1f, 0x38, 0x0f, 0x3a, 0x11, 0xca, 0xe7, 0x6c, 0x17, 0x2a,
	0xd7, 0x4e, 0xf0, 0x87, 0x43, 0xfa, 0xed, 0xe8, 0x15, 0xc4, 0x36, 0xb0, 0xd5, 0xf9, 0xc2, 0xd2,
	0x05, 0xdf, 0x6e, 0x79, 0x64, 0x3e, 0xc9, 0xc2, 0x56, 0xf7, 0x0b, 0xfa, 0x75, 0x77, 0x95, 0xe4,
	0x4a, 0x8d, 0x64, 0x8d, 0xfb, 0x8f, 0x05, 0xe7, 0x03, 0x1e, 0x2d, 0xc3, 0x15, 0xf9, 0x16, 0x60,
	0xcc, 0x17, 0x6c, 0x73, 0xc3, 0x63, 0x19, 0x3b, 0x56, 0xbb, 0x78, 0x55, 0xa6, 0x39, 0x86, 0xf4,
	0xa1, 0xd9, 0x13, 0x12, 0x97, 0x6c, 0x21, 0x29, 0xc6, 0x7c, 0xf3, 0x1e, 0x45, 0xec, 0x14, 0xda,
	0xc5, 0xab, 0x4a, 0xf7, 0x71, 0x27, 0xdd, 0x42, 0xaa, 0x18, 0xe2, 0x92, 0x7e, 0x2a, 0x57, 0x39,
	0xae, 0xf9, 0x35, 0x4a, 0x4f, 0x89, 0x9d, 0x62, 0xdb, 0x52, 0x39, 0x8e, 0x0c, 0xf9, 0x01, 0xca,
	0xd7, 0x5c, 0x99, 0x21, 0xc6, 0x4e, 0x49, 0xc7, 0x7e, 0x92, 0xc5, 0x9e, 0xef, 0x66, 0x52, 0x20,
	0xdb, 0x6a, 0x29, 0x3d, 0x0a, 0xdd, 0xbf, 0x2d, 0xa8, 0x9d, 0x38, 0x89, 0x03, 0xf6, 0x1d, 0x93,
	0x8b, 0x35, 0x0a, 0xc7, 0xd2, 0x49, 0x52, 0x48, 0x1e, 0xc3, 0x59, 0x92, 0xbc, 0xa0, 0xf9, 0x04,
	0x90, 0x16, 0x5c, 0xcc, 0x63, 0x14, 0x11, 0xdb, 0xa2, 0xa9, 0x2a, 0xc3, 0xca, 0xe7, 0xb1, 0x38,
	0xfe, 0xc0, 0x45, 0xe0, 0x94, 0x12, 0x5f, 0x8a, 0x55, 0x34, 0x9f, 0xbf, 0xc3, 0xc8, 0x39, 0x4b,
	0xa2, 0x69, 0xe0, 0x6e, 0xa1, 0x92, 0xeb, 0x03, 0x21, 0x50, 0xf2, 0x98, 0x5c, 0x9b, 0x4a, 0xb4,
	0xad, 0x1a, 0x91, 0x4a, 0x6e, 0x87, 0xba, 0x96, 0x12, 0xcd, 0x31, 0xe4, 0x09, 0x9c, 0x0f, 0xf9,
	0x96, 0x85, 0x91, 0x29, 0xc7, 0x20, 0x15, 0x6b, 0xa2, 0x8a, 0x4c, 0x0a, 0xd1, 0xb6, 0xfb, 0x57,
	0x11, 0xaa, 0x03, 0xb6, 0x58, 0x63, 0x70, 0xc7, 0x83, 0xfd, 0x06, 0x49, 0x1d, 0x0a, 0xb7, 0x43,
	0x9d, 0xae, 0x44, 0x0b, 0xb7, 0xc3, 0xac, 0x80, 0x42, 0xae, 0x00, 0x07, 0xec, 0x37, 0x28, 0xe2,
	0x90, 0xa7, 0x19, 0x52, 0xa8, 0x52, 0xcf, 0xf6, 0xcb, 0x65, 0xf8, 0x60, 0x92, 0x18, 0x44, 0x1a,
	0x50, 0x7c, 0x8d, 0x07, 0xb3, 0x53, 0x65, 0xaa, 0x18, 0x03, 0x81, 0x4c, 0x62, 0xe0, 0x9c, 0xb7,
	0xad, 0xab, 0x1a, 0x4d, 0xa1, 0xea, 0xd9, 0x98, 0xc5, 0x72, 0x1e, 0x63, 0xe0, 0xd8, 0xda, 0x95,
	0x61, 0xd2, 0x86, 0x8a, 0xcf, 0xfb, 0x38, 0xc4, 0x0d, 0xaa, 0x95, 0x17, 0x6d, 0xeb, 0xea, 0x82,
	0xe6, 0x29, 0xe2, 0x42, 0xf5, 0x15, 0x0b, 0x37, 0x61, 0xb4, 0x9a, 0x85, 0xd1, 0x02, 0x9d, 0xb2,
	0x8e, 0x70, 0xc2, 0xa9, 0x28, 0x0a, 0x0f, 0xf8, 0x3e, 0x92, 0x28, 0x1c, 0xd0, 0x92, 0x3c, 0xa5,
	0xcf, 0x33, 0x8b, 0xa5, 0xa2, 0x30, 0x70, 0x2a, 0x5a, 0x90, 0x63, 0xc8, 0x25, 0x94, 0xbd, 0x7d,
	0xea, 0xae, 0xea, 0x2a, 0x8e, 0x84, 0x9e, 0xfa, 0x5e, 0x8e, 0x84, 0xe0, 0xc2, 0xa9, 0x99, 0xa9,
	0x1b, 0xac, 0xfa, 0x39, 0x0b, 0xff, 0x44, 0xa7, 0xae, 0x3b, 0xac, 0x6d, 0xf7, 0x19, 0x34, 0x93,
	0xee, 0xdf, 0x46, 0x4b, 0x4e, 0xf1, 0x8f, 0x3d, 0xc6, 0x52, 0xb5, 0x6c, 0x4e, 0xc7, 0x66, 0xf0,
	0xca, 0x74, 0xef, 0x01, 0x8e, 0x32, 0xf2, 0x32, 0x45, 0xfe, 0x61, 0x87, 0x5a, 0x56, 0xef, 0x3e,
	0xca, 0xce, 0xfb, 0xdd, 0x74, 0x38, 0x1f, 0x8f, 0xfc, 0x7b, 0x6f, 0x44, 0x73, 0x32, 0x35, 0x9f,
	0xd1, 0x43, 0xa8, 0xfe, 0xd1, 0x82, 0x2e, 0xda, 0x20, 0x77, 0x05, 0x15, 0x33, 0x42, 0x1d, 0x3b,
	0x37, 0xe0, 0xe4, 0x24, 0x64, 0x03, 0xbe, 0x84, 0x72, 0x7f, 0x1f, 0x6e, 0x02, 0x3f, 0xdc, 0xa2,
	0x8e, 0x51, 0xa3, 0x47, 0x42, 0x35, 0xd6, 0x08, 0x27, 0xc7, 0xbf, 0x21, 0x4f, 0xb9, 0xdf, 0x41,
	0x5d, 0xfd, 0xb0, 0x4c, 0xae, 0xd3, 0x7d, 0x7e, 0xe6, 0x84, 0xbb, 0xdf, 0x80, 0xdd, 0x0f, 0xa3,
	0x21, 0x93, 0x4c, 0xb9, 0xd5, 0x57, 0xbb, 0xab, 0x54, 0xdb, 0xee, 0x0d, 0x5c, 0x4c, 0xdf, 0xa3,
	0x10, 0x61, 0xf0, 0xe9, 0x79, 0x75, 0xc0, 0xf6, 0xd8, 0xe2, 0x1d, 0x5b, 0xa1, 0x39, 0xb2, 0x29,
	0x54, 0x91, 0xc6, 0x61, 0x2c, 0x75, 0x55, 0x55, 0xaa, 0xed, 0xef, 0x55, 0x4b, 0xb3, 0x4e, 0x91,
	0x0a, 0xd8, 0xf3, 0xc9, 0xeb, 0xc9, 0xf4, 0xd7, 0x49, 0xe3, 0x2b, 0x52, 0x86, 0x33, 0x8f, 0x4e,
	0xfd, 0x69, 0xa3, 0x40, 0xaa, 0x70, 0xd1, 0xa3, 0xfe, 0xe8, 0x55, 0x6f, 0xe0, 0x37, 0x8a, 0x84,
	0x40, 0x7d, 0xf4, 0xd6, 0x1f, 0xd1, 0x49, 0x6f, 0x9c, 0xac, 0x6d, 0x94, 0x48, 0x13, 0x6a, 0x73,
	0x6f, 0xe6, 0xd3, 0x51, 0xef, 0xce, 0xa3, 0xd3, 0xb7, 0xf7, 0x8d, 0xb3, 0xee, 0xbf, 0x16, 0xd8,
	0xc9, 0x35, 0x73, 0x20, 0x23, 0x80, 0xe4, 0x86, 0xb9, 0xf1, 0x7d, 0x8f, 0x3c, 0xed, 0x64, 0x17,
	0x69, 0xc2, 0x9a, 0x56, 0xb4, 0x2e, 0x3f, 0x76, 0xa8, 0xdd, 0x52, 0x8c, 0x77, 0x3c, 0x8a, 0xf1,
	0x85, 0x45, 0x7e, 0x01, 0xe8, 0x45, 0x6c, 0x73, 0x88, 0x71, 0x4e, 0xc7, 0xa4, 0x75, 0x1c, 0xf6,
	0xc7, 0x87, 0xa7, 0xf5, 0xe8, 0x33, 0x3e, 0xf2, 0x23, 0xd8, 0xa6, 0xf7, 0xe4, 0x69, 0xe6, 0x3f,
	0x9d, 0x46, 0xab, 0x91, 0x39, 0xcc, 0x00, 0x5e, 0x58, 0xa4, 0x03, 0x95, 0x5e, 0x10, 0x64, 0x1d,
	0x6f, 0x66, 0x92, 0x94, 0x6a, 0x55, 0x3b, 0xe6, 0x31, 0x79, 0xc3, 0xc3, 0xa0, 0xfb, 0x33, 0x34,
	0xcd, 0xc6, 0x7d, 0x8c, 0x25, 0xdd, 0x47, 0x11, 0x0a, 0xf2, 0x0c, 0x6c, 0x5f, 0x84, 0xab, 0x15,
	0x0a, 0x72, 0xa2, 0x3e, 0x5d, 0xdb, 0xef, 0xc1, 0xd7, 0x11, 0xca, 0xfc, 0x33, 0xa4, 0x9e, 0xa0,
	0x34, 0xe3, 0x6f, 0xee, 0xff, 0xbf, 0x88, 0xbf, 0x9f, 0xeb, 0xd7, 0xe9, 0xe5, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x64, 0x34, 0x8c, 0x9f, 0x45, 0x07, 0x00, 0x00,
}

