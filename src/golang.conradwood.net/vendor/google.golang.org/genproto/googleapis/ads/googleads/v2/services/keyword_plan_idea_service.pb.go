// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/keyword_plan_idea_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [KeywordIdeaService.GenerateKeywordIdeas][].
type GenerateKeywordIdeasRequest struct {
	// The ID of the customer with the recommendation.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The resource name of the language to target.
	// Required
	Language *wrappers.StringValue `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	// The resource names of the location to target.
	// Max 10
	GeoTargetConstants []*wrappers.StringValue `protobuf:"bytes,8,rep,name=geo_target_constants,json=geoTargetConstants,proto3" json:"geo_target_constants,omitempty"`
	// Targeting network.
	KeywordPlanNetwork enums.KeywordPlanNetworkEnum_KeywordPlanNetwork `protobuf:"varint,9,opt,name=keyword_plan_network,json=keywordPlanNetwork,proto3,enum=google.ads.googleads.v2.enums.KeywordPlanNetworkEnum_KeywordPlanNetwork" json:"keyword_plan_network,omitempty"`
	// The type of seed to generate keyword ideas.
	//
	// Types that are valid to be assigned to Seed:
	//	*GenerateKeywordIdeasRequest_KeywordAndUrlSeed
	//	*GenerateKeywordIdeasRequest_KeywordSeed
	//	*GenerateKeywordIdeasRequest_UrlSeed
	Seed                 isGenerateKeywordIdeasRequest_Seed `protobuf_oneof:"seed"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *GenerateKeywordIdeasRequest) Reset()         { *m = GenerateKeywordIdeasRequest{} }
func (m *GenerateKeywordIdeasRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeasRequest) ProtoMessage()    {}
func (*GenerateKeywordIdeasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cad9891eddcc8ab, []int{0}
}

func (m *GenerateKeywordIdeasRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Marshal(b, m, deterministic)
}
func (m *GenerateKeywordIdeasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeasRequest.Merge(m, src)
}
func (m *GenerateKeywordIdeasRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Size(m)
}
func (m *GenerateKeywordIdeasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeasRequest proto.InternalMessageInfo

func (m *GenerateKeywordIdeasRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *GenerateKeywordIdeasRequest) GetLanguage() *wrappers.StringValue {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetGeoTargetConstants() []*wrappers.StringValue {
	if m != nil {
		return m.GeoTargetConstants
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordPlanNetwork() enums.KeywordPlanNetworkEnum_KeywordPlanNetwork {
	if m != nil {
		return m.KeywordPlanNetwork
	}
	return enums.KeywordPlanNetworkEnum_UNSPECIFIED
}

type isGenerateKeywordIdeasRequest_Seed interface {
	isGenerateKeywordIdeasRequest_Seed()
}

type GenerateKeywordIdeasRequest_KeywordAndUrlSeed struct {
	KeywordAndUrlSeed *KeywordAndUrlSeed `protobuf:"bytes,2,opt,name=keyword_and_url_seed,json=keywordAndUrlSeed,proto3,oneof"`
}

type GenerateKeywordIdeasRequest_KeywordSeed struct {
	KeywordSeed *KeywordSeed `protobuf:"bytes,3,opt,name=keyword_seed,json=keywordSeed,proto3,oneof"`
}

type GenerateKeywordIdeasRequest_UrlSeed struct {
	UrlSeed *UrlSeed `protobuf:"bytes,5,opt,name=url_seed,json=urlSeed,proto3,oneof"`
}

func (*GenerateKeywordIdeasRequest_KeywordAndUrlSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (*GenerateKeywordIdeasRequest_KeywordSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (*GenerateKeywordIdeasRequest_UrlSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (m *GenerateKeywordIdeasRequest) GetSeed() isGenerateKeywordIdeasRequest_Seed {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordAndUrlSeed() *KeywordAndUrlSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_KeywordAndUrlSeed); ok {
		return x.KeywordAndUrlSeed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordSeed() *KeywordSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_KeywordSeed); ok {
		return x.KeywordSeed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetUrlSeed() *UrlSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_UrlSeed); ok {
		return x.UrlSeed
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GenerateKeywordIdeasRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GenerateKeywordIdeasRequest_KeywordAndUrlSeed)(nil),
		(*GenerateKeywordIdeasRequest_KeywordSeed)(nil),
		(*GenerateKeywordIdeasRequest_UrlSeed)(nil),
	}
}

// Keyword And Url Seed
type KeywordAndUrlSeed struct {
	// The URL to crawl in order to generate keyword ideas.
	Url *wrappers.StringValue `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Requires at least one keyword.
	Keywords             []*wrappers.StringValue `protobuf:"bytes,2,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordAndUrlSeed) Reset()         { *m = KeywordAndUrlSeed{} }
func (m *KeywordAndUrlSeed) String() string { return proto.CompactTextString(m) }
func (*KeywordAndUrlSeed) ProtoMessage()    {}
func (*KeywordAndUrlSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cad9891eddcc8ab, []int{1}
}

func (m *KeywordAndUrlSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordAndUrlSeed.Unmarshal(m, b)
}
func (m *KeywordAndUrlSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordAndUrlSeed.Marshal(b, m, deterministic)
}
func (m *KeywordAndUrlSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordAndUrlSeed.Merge(m, src)
}
func (m *KeywordAndUrlSeed) XXX_Size() int {
	return xxx_messageInfo_KeywordAndUrlSeed.Size(m)
}
func (m *KeywordAndUrlSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordAndUrlSeed.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordAndUrlSeed proto.InternalMessageInfo

func (m *KeywordAndUrlSeed) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *KeywordAndUrlSeed) GetKeywords() []*wrappers.StringValue {
	if m != nil {
		return m.Keywords
	}
	return nil
}

// Keyword Seed
type KeywordSeed struct {
	// Requires at least one keyword.
	Keywords             []*wrappers.StringValue `protobuf:"bytes,1,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordSeed) Reset()         { *m = KeywordSeed{} }
func (m *KeywordSeed) String() string { return proto.CompactTextString(m) }
func (*KeywordSeed) ProtoMessage()    {}
func (*KeywordSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cad9891eddcc8ab, []int{2}
}

func (m *KeywordSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordSeed.Unmarshal(m, b)
}
func (m *KeywordSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordSeed.Marshal(b, m, deterministic)
}
func (m *KeywordSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordSeed.Merge(m, src)
}
func (m *KeywordSeed) XXX_Size() int {
	return xxx_messageInfo_KeywordSeed.Size(m)
}
func (m *KeywordSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordSeed.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordSeed proto.InternalMessageInfo

func (m *KeywordSeed) GetKeywords() []*wrappers.StringValue {
	if m != nil {
		return m.Keywords
	}
	return nil
}

// Url Seed
type UrlSeed struct {
	// The URL to crawl in order to generate keyword ideas.
	Url                  *wrappers.StringValue `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UrlSeed) Reset()         { *m = UrlSeed{} }
func (m *UrlSeed) String() string { return proto.CompactTextString(m) }
func (*UrlSeed) ProtoMessage()    {}
func (*UrlSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cad9891eddcc8ab, []int{3}
}

func (m *UrlSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UrlSeed.Unmarshal(m, b)
}
func (m *UrlSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UrlSeed.Marshal(b, m, deterministic)
}
func (m *UrlSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UrlSeed.Merge(m, src)
}
func (m *UrlSeed) XXX_Size() int {
	return xxx_messageInfo_UrlSeed.Size(m)
}
func (m *UrlSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_UrlSeed.DiscardUnknown(m)
}

var xxx_messageInfo_UrlSeed proto.InternalMessageInfo

func (m *UrlSeed) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

// Response message for [KeywordIdeaService.GenerateKeywordIdeas][].
type GenerateKeywordIdeaResponse struct {
	// Results of generating keyword ideas.
	Results              []*GenerateKeywordIdeaResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GenerateKeywordIdeaResponse) Reset()         { *m = GenerateKeywordIdeaResponse{} }
func (m *GenerateKeywordIdeaResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeaResponse) ProtoMessage()    {}
func (*GenerateKeywordIdeaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cad9891eddcc8ab, []int{4}
}

func (m *GenerateKeywordIdeaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Marshal(b, m, deterministic)
}
func (m *GenerateKeywordIdeaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeaResponse.Merge(m, src)
}
func (m *GenerateKeywordIdeaResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Size(m)
}
func (m *GenerateKeywordIdeaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeaResponse proto.InternalMessageInfo

func (m *GenerateKeywordIdeaResponse) GetResults() []*GenerateKeywordIdeaResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result of generating keyword ideas.
type GenerateKeywordIdeaResult struct {
	// Text of the keyword idea.
	// As in Keyword Plan historical metrics, this text may not be an actual
	// keyword, but the canonical form of multiple keywords.
	// See KeywordPlanKeywordHistoricalMetrics message in KeywordPlanService.
	Text *wrappers.StringValue `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// The historical metrics for the keyword
	KeywordIdeaMetrics   *common.KeywordPlanHistoricalMetrics `protobuf:"bytes,3,opt,name=keyword_idea_metrics,json=keywordIdeaMetrics,proto3" json:"keyword_idea_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *GenerateKeywordIdeaResult) Reset()         { *m = GenerateKeywordIdeaResult{} }
func (m *GenerateKeywordIdeaResult) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeaResult) ProtoMessage()    {}
func (*GenerateKeywordIdeaResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cad9891eddcc8ab, []int{5}
}

func (m *GenerateKeywordIdeaResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeaResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Marshal(b, m, deterministic)
}
func (m *GenerateKeywordIdeaResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeaResult.Merge(m, src)
}
func (m *GenerateKeywordIdeaResult) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Size(m)
}
func (m *GenerateKeywordIdeaResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeaResult.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeaResult proto.InternalMessageInfo

func (m *GenerateKeywordIdeaResult) GetText() *wrappers.StringValue {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *GenerateKeywordIdeaResult) GetKeywordIdeaMetrics() *common.KeywordPlanHistoricalMetrics {
	if m != nil {
		return m.KeywordIdeaMetrics
	}
	return nil
}

func init() {
	proto.RegisterType((*GenerateKeywordIdeasRequest)(nil), "google.ads.googleads.v2.services.GenerateKeywordIdeasRequest")
	proto.RegisterType((*KeywordAndUrlSeed)(nil), "google.ads.googleads.v2.services.KeywordAndUrlSeed")
	proto.RegisterType((*KeywordSeed)(nil), "google.ads.googleads.v2.services.KeywordSeed")
	proto.RegisterType((*UrlSeed)(nil), "google.ads.googleads.v2.services.UrlSeed")
	proto.RegisterType((*GenerateKeywordIdeaResponse)(nil), "google.ads.googleads.v2.services.GenerateKeywordIdeaResponse")
	proto.RegisterType((*GenerateKeywordIdeaResult)(nil), "google.ads.googleads.v2.services.GenerateKeywordIdeaResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/keyword_plan_idea_service.proto", fileDescriptor_5cad9891eddcc8ab)
}

var fileDescriptor_5cad9891eddcc8ab = []byte{
	// 748 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x4f, 0xdb, 0x48,
	0x14, 0x5f, 0x3b, 0x2c, 0x81, 0xc9, 0x6a, 0x25, 0x46, 0x68, 0x37, 0x0b, 0x68, 0x37, 0x8a, 0x38,
	0xb0, 0x48, 0x6b, 0xaf, 0xcc, 0x61, 0x59, 0xd3, 0x48, 0x0d, 0x55, 0x9b, 0xa0, 0xaa, 0x08, 0x99,
	0x92, 0x43, 0x15, 0xc9, 0x1a, 0xec, 0x87, 0x65, 0xc5, 0x99, 0x49, 0x67, 0xc6, 0xd0, 0x3f, 0xe2,
	0xc2, 0xb5, 0xc7, 0x7e, 0x83, 0x1e, 0xfb, 0x1d, 0xfa, 0x05, 0x38, 0x55, 0xe2, 0x2b, 0xf4, 0xd4,
	0x4f, 0x51, 0xd9, 0x9e, 0x49, 0x02, 0x24, 0x0d, 0xe5, 0xf6, 0xfc, 0xde, 0xef, 0xfd, 0xde, 0x6f,
	0xde, 0xbc, 0x79, 0x46, 0x0f, 0x23, 0xc6, 0xa2, 0x04, 0x6c, 0x12, 0x0a, 0xbb, 0x30, 0x33, 0xeb,
	0xd4, 0xb1, 0x05, 0xf0, 0xd3, 0x38, 0x00, 0x61, 0xf7, 0xe0, 0xf5, 0x19, 0xe3, 0xa1, 0x3f, 0x48,
	0x08, 0xf5, 0xe3, 0x10, 0x88, 0xaf, 0x42, 0xd6, 0x80, 0x33, 0xc9, 0x70, 0xad, 0x48, 0xb3, 0x48,
	0x28, 0xac, 0x21, 0x83, 0x75, 0xea, 0x58, 0x9a, 0x61, 0x65, 0x7b, 0x5a, 0x8d, 0x80, 0xf5, 0xfb,
	0x8c, 0x5e, 0xaf, 0x50, 0xf8, 0x0a, 0xee, 0xe9, 0x99, 0x40, 0xd3, 0xfe, 0x0d, 0x69, 0x14, 0xe4,
	0x19, 0xe3, 0x3d, 0x95, 0xb9, 0xa6, 0x33, 0x07, 0xb1, 0x4d, 0x28, 0x65, 0x92, 0xc8, 0x98, 0x51,
	0xa1, 0xa2, 0x7f, 0xaa, 0x68, 0xfe, 0x75, 0x9c, 0x9e, 0xd8, 0x67, 0x9c, 0x0c, 0x06, 0xc0, 0x75,
	0xfc, 0xf7, 0xb1, 0xec, 0x20, 0x89, 0x81, 0xca, 0x22, 0x50, 0xff, 0x3c, 0x87, 0x56, 0x5b, 0x40,
	0x81, 0x13, 0x09, 0x4f, 0x8b, 0xea, 0x7b, 0x21, 0x10, 0xe1, 0xc1, 0xcb, 0x14, 0x84, 0xc4, 0x7f,
	0xa1, 0x4a, 0x90, 0x0a, 0xc9, 0xfa, 0xc0, 0xfd, 0x38, 0xac, 0x1a, 0x35, 0x63, 0x63, 0xd1, 0x43,
	0xda, 0xb5, 0x17, 0xe2, 0x6d, 0xb4, 0x90, 0x10, 0x1a, 0xa5, 0x24, 0x82, 0x6a, 0xb9, 0x66, 0x6c,
	0x54, 0x9c, 0x35, 0xd5, 0x35, 0x4b, 0x8b, 0xb1, 0x0e, 0x25, 0x8f, 0x69, 0xd4, 0x21, 0x49, 0x0a,
	0xde, 0x10, 0x8d, 0xf7, 0xd1, 0x72, 0x04, 0xcc, 0x97, 0x84, 0x47, 0x20, 0xfd, 0x80, 0x51, 0x21,
	0x09, 0x95, 0xa2, 0xba, 0x50, 0x2b, 0xcd, 0x64, 0xc1, 0x11, 0xb0, 0xe7, 0x79, 0xe2, 0x23, 0x9d,
	0x87, 0xdf, 0xa0, 0xe5, 0x49, 0xfd, 0xab, 0x2e, 0xd6, 0x8c, 0x8d, 0x5f, 0x9d, 0xb6, 0x35, 0xed,
	0x5a, 0xf3, 0xd6, 0x5b, 0xea, 0xf0, 0x07, 0x09, 0xa1, 0xfb, 0x45, 0xe2, 0x63, 0x9a, 0xf6, 0x27,
	0xb8, 0x3d, 0xdc, 0xbb, 0xe5, 0xc3, 0x27, 0xa3, 0xda, 0x84, 0x86, 0x7e, 0xca, 0x13, 0x5f, 0x00,
	0x84, 0x55, 0x33, 0xef, 0xc8, 0x96, 0x35, 0x6b, 0xa4, 0x74, 0x9d, 0x26, 0x0d, 0x8f, 0x78, 0x72,
	0x08, 0x10, 0xb6, 0x7f, 0xf2, 0x96, 0x7a, 0x37, 0x9d, 0xd8, 0x43, 0xbf, 0xe8, 0x3a, 0x39, 0x7f,
	0x29, 0xe7, 0xff, 0xe7, 0xce, 0xfc, 0x8a, 0xb9, 0xd2, 0x1b, 0x7d, 0xe2, 0x27, 0x68, 0x61, 0xa8,
	0xf7, 0xe7, 0x9c, 0xef, 0xef, 0xd9, 0x7c, 0x23, 0x95, 0xe5, 0xb4, 0x30, 0x77, 0xe7, 0xd1, 0x5c,
	0xc6, 0x51, 0x3f, 0x47, 0x4b, 0xb7, 0x4e, 0x83, 0x2d, 0x54, 0x4a, 0x79, 0x92, 0xcf, 0xcf, 0xac,
	0xbb, 0xcd, 0x80, 0xd9, 0x58, 0x29, 0x8d, 0xa2, 0x6a, 0xde, 0x61, 0x20, 0x86, 0xe8, 0x7a, 0x0b,
	0x55, 0xc6, 0x0e, 0x7b, 0x8d, 0xc8, 0xf8, 0x21, 0xa2, 0xff, 0x51, 0xf9, 0x9e, 0xea, 0xeb, 0x72,
	0xe2, 0xa3, 0xf2, 0x40, 0x0c, 0x18, 0x15, 0x80, 0x8f, 0x50, 0x99, 0x83, 0x48, 0x13, 0xa9, 0x25,
	0xed, 0xcc, 0x6e, 0xf8, 0x64, 0xbe, 0x34, 0x91, 0x9e, 0xe6, 0xaa, 0x7f, 0x32, 0xd0, 0x1f, 0x53,
	0x61, 0xf8, 0x5f, 0x34, 0x27, 0xe1, 0x95, 0x54, 0x23, 0xf9, 0xfd, 0x43, 0xe4, 0x48, 0x4c, 0x47,
	0x43, 0x9d, 0xaf, 0xc9, 0x3e, 0x48, 0x1e, 0x07, 0x42, 0x0d, 0xdd, 0x83, 0xa9, 0x9a, 0xd5, 0xc6,
	0x1b, 0x7b, 0x3a, 0xed, 0x58, 0x48, 0xc6, 0xe3, 0x80, 0x24, 0xcf, 0x0a, 0x8e, 0xe1, 0x23, 0xca,
	0x04, 0x2a, 0x9f, 0xf3, 0xce, 0x44, 0xbf, 0x8d, 0x25, 0x65, 0xa1, 0xc3, 0xe2, 0xf8, 0xf8, 0xca,
	0x40, 0xcb, 0x93, 0xd6, 0x14, 0x6e, 0xdc, 0xab, 0x73, 0x7a, 0xbd, 0xad, 0x34, 0xee, 0xdb, 0xf8,
	0xfc, 0x22, 0xeb, 0x8d, 0x8b, 0xab, 0x2f, 0xef, 0xcd, 0xff, 0xea, 0x4e, 0xbe, 0xfc, 0xd5, 0x52,
	0x14, 0xf6, 0xdb, 0xb1, 0x95, 0xd9, 0xd8, 0x3c, 0x77, 0xa3, 0x09, 0x0a, 0x5c, 0x63, 0x73, 0x65,
	0xf5, 0xb2, 0x59, 0x1d, 0x15, 0x55, 0xd6, 0x20, 0x16, 0x59, 0x07, 0x77, 0x2f, 0x4c, 0xb4, 0x1e,
	0xb0, 0xfe, 0x4c, 0x81, 0xbb, 0xab, 0x93, 0x7b, 0x76, 0x90, 0x5d, 0xec, 0x81, 0xf1, 0xa2, 0xad,
	0x08, 0x22, 0x96, 0x6d, 0x5e, 0x8b, 0xf1, 0xc8, 0x8e, 0x80, 0xe6, 0xd7, 0x6e, 0x8f, 0x4a, 0x4e,
	0xff, 0x5f, 0xee, 0x68, 0xe3, 0x83, 0x59, 0x6a, 0x35, 0x9b, 0x1f, 0xcd, 0x5a, 0xab, 0x20, 0x6c,
	0x86, 0xc2, 0x2a, 0xcc, 0xcc, 0xea, 0x38, 0x96, 0x2a, 0x2c, 0x2e, 0x35, 0xa4, 0xdb, 0x0c, 0x45,
	0x77, 0x08, 0xe9, 0x76, 0x9c, 0xae, 0x86, 0x7c, 0x35, 0xd7, 0x0b, 0xbf, 0xeb, 0x36, 0x43, 0xe1,
	0xba, 0x43, 0x90, 0xeb, 0x76, 0x1c, 0xd7, 0xd5, 0xb0, 0xe3, 0xf9, 0x5c, 0xe7, 0xd6, 0xb7, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1b, 0x81, 0x2f, 0xaa, 0xd6, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeywordPlanIdeaServiceClient is the client API for KeywordPlanIdeaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordPlanIdeaServiceClient interface {
	// Returns a list of keyword ideas.
	GenerateKeywordIdeas(ctx context.Context, in *GenerateKeywordIdeasRequest, opts ...grpc.CallOption) (*GenerateKeywordIdeaResponse, error)
}

type keywordPlanIdeaServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeywordPlanIdeaServiceClient(cc *grpc.ClientConn) KeywordPlanIdeaServiceClient {
	return &keywordPlanIdeaServiceClient{cc}
}

func (c *keywordPlanIdeaServiceClient) GenerateKeywordIdeas(ctx context.Context, in *GenerateKeywordIdeasRequest, opts ...grpc.CallOption) (*GenerateKeywordIdeaResponse, error) {
	out := new(GenerateKeywordIdeaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.KeywordPlanIdeaService/GenerateKeywordIdeas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanIdeaServiceServer is the server API for KeywordPlanIdeaService service.
type KeywordPlanIdeaServiceServer interface {
	// Returns a list of keyword ideas.
	GenerateKeywordIdeas(context.Context, *GenerateKeywordIdeasRequest) (*GenerateKeywordIdeaResponse, error)
}

// UnimplementedKeywordPlanIdeaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeywordPlanIdeaServiceServer struct {
}

func (*UnimplementedKeywordPlanIdeaServiceServer) GenerateKeywordIdeas(ctx context.Context, req *GenerateKeywordIdeasRequest) (*GenerateKeywordIdeaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKeywordIdeas not implemented")
}

func RegisterKeywordPlanIdeaServiceServer(s *grpc.Server, srv KeywordPlanIdeaServiceServer) {
	s.RegisterService(&_KeywordPlanIdeaService_serviceDesc, srv)
}

func _KeywordPlanIdeaService_GenerateKeywordIdeas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKeywordIdeasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordIdeas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.KeywordPlanIdeaService/GenerateKeywordIdeas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordIdeas(ctx, req.(*GenerateKeywordIdeasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordPlanIdeaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.KeywordPlanIdeaService",
	HandlerType: (*KeywordPlanIdeaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateKeywordIdeas",
			Handler:    _KeywordPlanIdeaService_GenerateKeywordIdeas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/keyword_plan_idea_service.proto",
}
