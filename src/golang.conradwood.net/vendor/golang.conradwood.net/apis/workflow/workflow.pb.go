// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/workflow/workflow.proto
// DO NOT EDIT!

/*
Package workflow is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/workflow/workflow.proto

It has these top-level messages:
	InstanceID
	Instance
	CreateInstanceRequest
	SubmitFormRequest
	Workflow
	WorkFlowVersion
	FormInstance
	FormID
	Form
	FormField
	FormValues
	FormValue
*/
package workflow

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

type InstanceID struct {
	ID uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *InstanceID) Reset()                    { *m = InstanceID{} }
func (m *InstanceID) String() string            { return proto.CompactTextString(m) }
func (*InstanceID) ProtoMessage()               {}
func (*InstanceID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InstanceID) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type Instance struct {
	ID         uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	WorkFlowID uint64 `protobuf:"varint,2,opt,name=WorkFlowID" json:"WorkFlowID,omitempty"`
	Creator    string `protobuf:"bytes,3,opt,name=Creator" json:"Creator,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Instance) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Instance) GetWorkFlowID() uint64 {
	if m != nil {
		return m.WorkFlowID
	}
	return 0
}

func (m *Instance) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type CreateInstanceRequest struct {
	WorkFlowID uint64 `protobuf:"varint,1,opt,name=WorkFlowID" json:"WorkFlowID,omitempty"`
}

func (m *CreateInstanceRequest) Reset()                    { *m = CreateInstanceRequest{} }
func (m *CreateInstanceRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateInstanceRequest) ProtoMessage()               {}
func (*CreateInstanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateInstanceRequest) GetWorkFlowID() uint64 {
	if m != nil {
		return m.WorkFlowID
	}
	return 0
}

type SubmitFormRequest struct {
	WorkFlowID uint64      `protobuf:"varint,1,opt,name=WorkFlowID" json:"WorkFlowID,omitempty"`
	Values     *FormValues `protobuf:"bytes,2,opt,name=Values" json:"Values,omitempty"`
}

func (m *SubmitFormRequest) Reset()                    { *m = SubmitFormRequest{} }
func (m *SubmitFormRequest) String() string            { return proto.CompactTextString(m) }
func (*SubmitFormRequest) ProtoMessage()               {}
func (*SubmitFormRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SubmitFormRequest) GetWorkFlowID() uint64 {
	if m != nil {
		return m.WorkFlowID
	}
	return 0
}

func (m *SubmitFormRequest) GetValues() *FormValues {
	if m != nil {
		return m.Values
	}
	return nil
}

type Workflow struct {
	ID uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *Workflow) Reset()                    { *m = Workflow{} }
func (m *Workflow) String() string            { return proto.CompactTextString(m) }
func (*Workflow) ProtoMessage()               {}
func (*Workflow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Workflow) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type WorkFlowVersion struct {
	ID         uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	WorkFlowID uint64 `protobuf:"varint,2,opt,name=WorkFlowID" json:"WorkFlowID,omitempty"`
}

func (m *WorkFlowVersion) Reset()                    { *m = WorkFlowVersion{} }
func (m *WorkFlowVersion) String() string            { return proto.CompactTextString(m) }
func (*WorkFlowVersion) ProtoMessage()               {}
func (*WorkFlowVersion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WorkFlowVersion) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *WorkFlowVersion) GetWorkFlowID() uint64 {
	if m != nil {
		return m.WorkFlowID
	}
	return 0
}

type FormInstance struct {
	Form   *Form       `protobuf:"bytes,1,opt,name=Form" json:"Form,omitempty"`
	Values *FormValues `protobuf:"bytes,2,opt,name=Values" json:"Values,omitempty"`
}

func (m *FormInstance) Reset()                    { *m = FormInstance{} }
func (m *FormInstance) String() string            { return proto.CompactTextString(m) }
func (*FormInstance) ProtoMessage()               {}
func (*FormInstance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FormInstance) GetForm() *Form {
	if m != nil {
		return m.Form
	}
	return nil
}

func (m *FormInstance) GetValues() *FormValues {
	if m != nil {
		return m.Values
	}
	return nil
}

type FormID struct {
	ID uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *FormID) Reset()                    { *m = FormID{} }
func (m *FormID) String() string            { return proto.CompactTextString(m) }
func (*FormID) ProtoMessage()               {}
func (*FormID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FormID) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type Form struct {
	ID     uint64       `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Fields []*FormField `protobuf:"bytes,2,rep,name=Fields" json:"Fields,omitempty"`
}

func (m *Form) Reset()                    { *m = Form{} }
func (m *Form) String() string            { return proto.CompactTextString(m) }
func (*Form) ProtoMessage()               {}
func (*Form) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Form) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Form) GetFields() []*FormField {
	if m != nil {
		return m.Fields
	}
	return nil
}

// currently we only do string fields, so it is really simple
type FormField struct {
	Key         string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=DisplayName" json:"DisplayName,omitempty"`
}

func (m *FormField) Reset()                    { *m = FormField{} }
func (m *FormField) String() string            { return proto.CompactTextString(m) }
func (*FormField) ProtoMessage()               {}
func (*FormField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *FormField) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *FormField) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

// very simple, currently we only do string fields, so really, these are key values
type FormValues struct {
	Values map[string]*FormValue `protobuf:"bytes,1,rep,name=Values" json:"Values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FormValues) Reset()                    { *m = FormValues{} }
func (m *FormValues) String() string            { return proto.CompactTextString(m) }
func (*FormValues) ProtoMessage()               {}
func (*FormValues) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *FormValues) GetValues() map[string]*FormValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type FormValue struct {
	Key    string   `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=Values" json:"Values,omitempty"`
}

func (m *FormValue) Reset()                    { *m = FormValue{} }
func (m *FormValue) String() string            { return proto.CompactTextString(m) }
func (*FormValue) ProtoMessage()               {}
func (*FormValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *FormValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *FormValue) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*InstanceID)(nil), "workflow.InstanceID")
	proto.RegisterType((*Instance)(nil), "workflow.Instance")
	proto.RegisterType((*CreateInstanceRequest)(nil), "workflow.CreateInstanceRequest")
	proto.RegisterType((*SubmitFormRequest)(nil), "workflow.SubmitFormRequest")
	proto.RegisterType((*Workflow)(nil), "workflow.Workflow")
	proto.RegisterType((*WorkFlowVersion)(nil), "workflow.WorkFlowVersion")
	proto.RegisterType((*FormInstance)(nil), "workflow.FormInstance")
	proto.RegisterType((*FormID)(nil), "workflow.FormID")
	proto.RegisterType((*Form)(nil), "workflow.Form")
	proto.RegisterType((*FormField)(nil), "workflow.FormField")
	proto.RegisterType((*FormValues)(nil), "workflow.FormValues")
	proto.RegisterType((*FormValue)(nil), "workflow.FormValue")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WorkFlow service

type WorkFlowClient interface {
	// create an instance of a workflow (latest version)
	CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*Instance, error)
	// return the form for a given workflow instance and (calling) user
	GetFormForMe(ctx context.Context, in *InstanceID, opts ...grpc.CallOption) (*Form, error)
	// submit the form for a given instance
	SubmitForm(ctx context.Context, in *SubmitFormRequest, opts ...grpc.CallOption) (*common.Void, error)
}

type workFlowClient struct {
	cc *grpc.ClientConn
}

func NewWorkFlowClient(cc *grpc.ClientConn) WorkFlowClient {
	return &workFlowClient{cc}
}

func (c *workFlowClient) CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*Instance, error) {
	out := new(Instance)
	err := grpc.Invoke(ctx, "/workflow.WorkFlow/CreateInstance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workFlowClient) GetFormForMe(ctx context.Context, in *InstanceID, opts ...grpc.CallOption) (*Form, error) {
	out := new(Form)
	err := grpc.Invoke(ctx, "/workflow.WorkFlow/GetFormForMe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workFlowClient) SubmitForm(ctx context.Context, in *SubmitFormRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/workflow.WorkFlow/SubmitForm", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorkFlow service

type WorkFlowServer interface {
	// create an instance of a workflow (latest version)
	CreateInstance(context.Context, *CreateInstanceRequest) (*Instance, error)
	// return the form for a given workflow instance and (calling) user
	GetFormForMe(context.Context, *InstanceID) (*Form, error)
	// submit the form for a given instance
	SubmitForm(context.Context, *SubmitFormRequest) (*common.Void, error)
}

func RegisterWorkFlowServer(s *grpc.Server, srv WorkFlowServer) {
	s.RegisterService(&_WorkFlow_serviceDesc, srv)
}

func _WorkFlow_CreateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkFlowServer).CreateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.WorkFlow/CreateInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkFlowServer).CreateInstance(ctx, req.(*CreateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkFlow_GetFormForMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstanceID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkFlowServer).GetFormForMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.WorkFlow/GetFormForMe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkFlowServer).GetFormForMe(ctx, req.(*InstanceID))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkFlow_SubmitForm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitFormRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkFlowServer).SubmitForm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.WorkFlow/SubmitForm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkFlowServer).SubmitForm(ctx, req.(*SubmitFormRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkFlow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "workflow.WorkFlow",
	HandlerType: (*WorkFlowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInstance",
			Handler:    _WorkFlow_CreateInstance_Handler,
		},
		{
			MethodName: "GetFormForMe",
			Handler:    _WorkFlow_GetFormForMe_Handler,
		},
		{
			MethodName: "SubmitForm",
			Handler:    _WorkFlow_SubmitForm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/workflow/workflow.proto",
}

// Client API for FormRenderer service

type FormRendererClient interface {
	// save a new form (or update one)
	SaveForm(ctx context.Context, in *Form, opts ...grpc.CallOption) (*Form, error)
	// get a form by ID
	GetFormInstance(ctx context.Context, in *FormID, opts ...grpc.CallOption) (*FormInstance, error)
}

type formRendererClient struct {
	cc *grpc.ClientConn
}

func NewFormRendererClient(cc *grpc.ClientConn) FormRendererClient {
	return &formRendererClient{cc}
}

func (c *formRendererClient) SaveForm(ctx context.Context, in *Form, opts ...grpc.CallOption) (*Form, error) {
	out := new(Form)
	err := grpc.Invoke(ctx, "/workflow.FormRenderer/SaveForm", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formRendererClient) GetFormInstance(ctx context.Context, in *FormID, opts ...grpc.CallOption) (*FormInstance, error) {
	out := new(FormInstance)
	err := grpc.Invoke(ctx, "/workflow.FormRenderer/GetFormInstance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FormRenderer service

type FormRendererServer interface {
	// save a new form (or update one)
	SaveForm(context.Context, *Form) (*Form, error)
	// get a form by ID
	GetFormInstance(context.Context, *FormID) (*FormInstance, error)
}

func RegisterFormRendererServer(s *grpc.Server, srv FormRendererServer) {
	s.RegisterService(&_FormRenderer_serviceDesc, srv)
}

func _FormRenderer_SaveForm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Form)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormRendererServer).SaveForm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.FormRenderer/SaveForm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormRendererServer).SaveForm(ctx, req.(*Form))
	}
	return interceptor(ctx, in, info, handler)
}

func _FormRenderer_GetFormInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FormID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormRendererServer).GetFormInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.FormRenderer/GetFormInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormRendererServer).GetFormInstance(ctx, req.(*FormID))
	}
	return interceptor(ctx, in, info, handler)
}

var _FormRenderer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "workflow.FormRenderer",
	HandlerType: (*FormRendererServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveForm",
			Handler:    _FormRenderer_SaveForm_Handler,
		},
		{
			MethodName: "GetFormInstance",
			Handler:    _FormRenderer_GetFormInstance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/workflow/workflow.proto",
}

func init() { proto.RegisterFile("golang.conradwood.net/apis/workflow/workflow.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x93, 0x12, 0x92, 0x71, 0x95, 0x96, 0x6d, 0xa9, 0x2c, 0x53, 0x95, 0x68, 0x4f, 0xe1,
	0x47, 0xae, 0x64, 0x40, 0xad, 0xe0, 0x80, 0xa0, 0x6e, 0x90, 0x85, 0xe8, 0x61, 0x8b, 0xd2, 0x2b,
	0xdb, 0x78, 0xa9, 0xac, 0xd8, 0xbb, 0x61, 0xed, 0xd4, 0xca, 0x8b, 0xf0, 0x3c, 0x3c, 0x1a, 0xf2,
	0xfa, 0xdf, 0x8e, 0x10, 0x3d, 0x79, 0x77, 0xe6, 0x9b, 0x6f, 0xbe, 0x99, 0x1d, 0x0f, 0xd8, 0x77,
	0x22, 0xa0, 0xfc, 0xce, 0x5a, 0x08, 0x2e, 0xa9, 0x97, 0x08, 0xe1, 0x59, 0x9c, 0xc5, 0xa7, 0x74,
	0xe5, 0x47, 0xa7, 0x89, 0x90, 0xcb, 0x9f, 0x81, 0x48, 0xca, 0x83, 0xb5, 0x92, 0x22, 0x16, 0x68,
	0x58, 0xdc, 0x4d, 0xeb, 0x1f, 0xd1, 0x0b, 0x11, 0x86, 0x82, 0xe7, 0x9f, 0x2c, 0x12, 0x1f, 0x03,
	0xb8, 0x3c, 0x8a, 0x29, 0x5f, 0x30, 0xd7, 0x41, 0x63, 0xe8, 0xb9, 0x8e, 0xa1, 0x4d, 0xb4, 0xe9,
	0x0e, 0xe9, 0xb9, 0x0e, 0xfe, 0x0e, 0xc3, 0xc2, 0xdb, 0xf6, 0xa1, 0x13, 0x80, 0x1b, 0x21, 0x97,
	0xb3, 0x40, 0x24, 0xae, 0x63, 0xf4, 0x94, 0xbd, 0x66, 0x41, 0x06, 0x3c, 0xbe, 0x90, 0x8c, 0xc6,
	0x42, 0x1a, 0xfd, 0x89, 0x36, 0x1d, 0x91, 0xe2, 0x8a, 0xcf, 0xe0, 0xa9, 0x3a, 0xb2, 0x82, 0x9b,
	0xb0, 0x5f, 0x6b, 0x16, 0xc5, 0x2d, 0x4a, 0xad, 0x4d, 0x89, 0x29, 0x3c, 0xb9, 0x5e, 0xdf, 0x86,
	0x7e, 0x3c, 0x13, 0x32, 0xfc, 0xcf, 0x20, 0xf4, 0x1a, 0x06, 0x73, 0x1a, 0xac, 0x59, 0xa4, 0x34,
	0xea, 0xf6, 0xa1, 0x55, 0x36, 0x2f, 0xa5, 0xc9, 0x7c, 0x24, 0xc7, 0x60, 0x13, 0x86, 0x37, 0xb9,
	0xbb, 0xd3, 0x8d, 0x4f, 0xb0, 0x57, 0xf0, 0xce, 0x99, 0x8c, 0x7c, 0xc1, 0x1f, 0xda, 0x14, 0xfc,
	0x03, 0x76, 0xd3, 0xa4, 0x65, 0x53, 0x31, 0xec, 0xa4, 0x77, 0xc5, 0xa0, 0xdb, 0xe3, 0xa6, 0x34,
	0xa2, 0x7c, 0x0f, 0x2c, 0xc0, 0x80, 0x81, 0xca, 0xd0, 0x7d, 0xcc, 0x8b, 0x2c, 0x57, 0x47, 0xf3,
	0x2b, 0x18, 0xcc, 0x7c, 0x16, 0x78, 0x29, 0x7f, 0x7f, 0xaa, 0xdb, 0x07, 0x4d, 0x7e, 0xe5, 0x23,
	0x39, 0x04, 0x7f, 0x84, 0x51, 0x69, 0x44, 0xfb, 0xd0, 0xff, 0xca, 0x36, 0x8a, 0x6a, 0x44, 0xd2,
	0x23, 0x9a, 0x80, 0xee, 0xf8, 0xd1, 0x2a, 0xa0, 0x9b, 0x2b, 0x1a, 0x32, 0x25, 0x78, 0x44, 0xea,
	0x26, 0xfc, 0x5b, 0x03, 0xa8, 0x64, 0xa3, 0xf3, 0xb2, 0x38, 0x4d, 0x25, 0x9f, 0x6c, 0x2b, 0xce,
	0xca, 0x3e, 0x97, 0x3c, 0x96, 0x9b, 0xa2, 0x50, 0xf3, 0x0a, 0xf4, 0x9a, 0x39, 0xd5, 0xb2, 0xac,
	0xb4, 0x2c, 0xd9, 0x06, 0xbd, 0x80, 0x47, 0xf7, 0x29, 0x20, 0x6f, 0xdb, 0xc1, 0x16, 0x66, 0x92,
	0x21, 0xde, 0xf7, 0xce, 0x35, 0xfc, 0x2e, 0xab, 0x4c, 0xd9, 0xb7, 0x54, 0x76, 0x54, 0x7b, 0x85,
	0xfe, 0x74, 0x54, 0xc8, 0xb0, 0xff, 0x68, 0xd9, 0xc4, 0xa4, 0x0f, 0x8c, 0x2e, 0x61, 0xdc, 0x9c,
	0x6c, 0xf4, 0xbc, 0xca, 0xba, 0x75, 0xe6, 0x4d, 0x54, 0x01, 0xca, 0xa0, 0xb7, 0xb0, 0xfb, 0x85,
	0xa9, 0x21, 0x9f, 0x09, 0xf9, 0x8d, 0xa1, 0xc3, 0x2e, 0xc6, 0x75, 0xcc, 0xd6, 0xb4, 0xa0, 0x33,
	0x80, 0xea, 0xef, 0x40, 0xcf, 0x2a, 0x6f, 0xe7, 0x9f, 0x31, 0x77, 0xad, 0x7c, 0x09, 0xcc, 0x85,
	0xef, 0xd9, 0x49, 0x36, 0x94, 0x84, 0x71, 0x8f, 0x49, 0x26, 0xd1, 0x4b, 0x18, 0x5e, 0xd3, 0x7b,
	0x96, 0x0d, 0x4b, 0x33, 0x49, 0x27, 0xe9, 0x07, 0xd8, 0xcb, 0xa5, 0x96, 0xea, 0xf7, 0x9b, 0x10,
	0xd7, 0x31, 0x8f, 0x5a, 0x96, 0x1c, 0xf9, 0xf9, 0x04, 0x8e, 0x39, 0x8b, 0xeb, 0xbb, 0x2a, 0xdd,
	0x53, 0x25, 0xf8, 0x76, 0xa0, 0x76, 0xd4, 0x9b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x41, 0x7c,
	0x48, 0xfa, 0x13, 0x05, 0x00, 0x00,
}
