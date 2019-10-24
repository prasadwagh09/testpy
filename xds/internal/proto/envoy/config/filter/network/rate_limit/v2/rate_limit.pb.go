// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/rate_limit/v2/rate_limit.proto

package envoy_config_filter_network_rate_limit_v2

import (
	fmt "fmt"
	ratelimit "google.golang.org/grpc/xds/internal/proto/envoy/api/v2/ratelimit"
	v2 "google.golang.org/grpc/xds/internal/proto/envoy/config/ratelimit/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type RateLimit struct {
	StatPrefix           string                           `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	Domain               string                           `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Descriptors          []*ratelimit.RateLimitDescriptor `protobuf:"bytes,3,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	Timeout              *duration.Duration               `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	FailureModeDeny      bool                             `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	RateLimitService     *v2.RateLimitServiceConfig       `protobuf:"bytes,6,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_34e9a222968daa71, []int{0}
}

func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetDescriptors() []*ratelimit.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimit) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitService() *v2.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.network.rate_limit.v2.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/rate_limit/v2/rate_limit.proto", fileDescriptor_34e9a222968daa71)
}

var fileDescriptor_34e9a222968daa71 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x8f, 0x94, 0x30,
	0x18, 0xc7, 0x53, 0x76, 0x77, 0x76, 0xb6, 0x24, 0xba, 0xf6, 0x22, 0xee, 0x41, 0x89, 0x26, 0x06,
	0x35, 0x69, 0x23, 0x7b, 0x30, 0xf1, 0x88, 0x73, 0x53, 0x93, 0x09, 0x1e, 0x3c, 0x92, 0xee, 0xf0,
	0x30, 0x79, 0x22, 0x50, 0x52, 0x0a, 0x0e, 0x5f, 0xc1, 0xa3, 0x1f, 0x77, 0xbc, 0x18, 0x28, 0x2f,
	0x33, 0x9e, 0xbc, 0xb5, 0xfc, 0x5f, 0x9e, 0xfe, 0x68, 0xe9, 0x47, 0x28, 0x5b, 0xd5, 0x89, 0x9d,
	0x2a, 0x33, 0xdc, 0x8b, 0x0c, 0x73, 0x03, 0x5a, 0x94, 0x60, 0x7e, 0x2a, 0xfd, 0x43, 0x68, 0x69,
	0x20, 0xc9, 0xb1, 0x40, 0x23, 0xda, 0xf0, 0x64, 0xc7, 0x2b, 0xad, 0x8c, 0x62, 0x6f, 0x86, 0x2c,
	0xb7, 0x59, 0x6e, 0xb3, 0x7c, 0xcc, 0xf2, 0x13, 0x77, 0x1b, 0xde, 0xbd, 0xb6, 0x63, 0x64, 0x85,
	0x53, 0x93, 0xad, 0x9d, 0x57, 0xb6, 0xf2, 0xee, 0xd5, 0xd9, 0x71, 0x16, 0x5f, 0x1f, 0xca, 0xeb,
	0xd1, 0xf4, 0x7c, 0xaf, 0xd4, 0x3e, 0x07, 0x31, 0xec, 0x1e, 0x9a, 0x4c, 0xa4, 0x8d, 0x96, 0x06,
	0x55, 0x39, 0xea, 0x4f, 0x5b, 0x99, 0x63, 0x2a, 0x0d, 0x88, 0x69, 0x61, 0x85, 0x97, 0x7f, 0x1c,
	0x7a, 0x13, 0x4b, 0x03, 0x5f, 0xfa, 0x4e, 0x16, 0x50, 0xb7, 0x36, 0xd2, 0x24, 0x95, 0x86, 0x0c,
	0x0f, 0x1e, 0xf1, 0x49, 0x70, 0x13, 0x5d, 0x1f, 0xa3, 0x4b, 0xed, 0xf8, 0x24, 0xa6, 0xbd, 0xb6,
	0x1d, 0x24, 0xf6, 0x82, 0xae, 0x52, 0x55, 0x48, 0x2c, 0x3d, 0xe7, 0xdc, 0x34, 0x7e, 0x66, 0xdf,
	0xa9, 0x9b, 0x42, 0xbd, 0xd3, 0x58, 0x19, 0xa5, 0x6b, 0xef, 0xc2, 0xbf, 0x08, 0xdc, 0xf0, 0x1d,
	0xb7, 0xff, 0x47, 0x56, 0xc8, 0xdb, 0x90, 0x2f, 0xa8, 0xf3, 0x11, 0x36, 0x73, 0x26, 0x5a, 0x1f,
	0xa3, 0xab, 0xdf, 0xc4, 0x59, 0x93, 0xf8, 0xb4, 0x89, 0xdd, 0xd3, 0x6b, 0x83, 0x05, 0xa8, 0xc6,
	0x78, 0x97, 0x3e, 0x09, 0xdc, 0xf0, 0x19, 0xb7, 0xf0, 0x7c, 0x82, 0xe7, 0x9b, 0x11, 0x3e, 0x9e,
	0x9c, 0xec, 0x2d, 0x7d, 0x92, 0x49, 0xcc, 0x1b, 0x0d, 0x49, 0xa1, 0x52, 0x48, 0x52, 0x28, 0x3b,
	0xef, 0xca, 0x27, 0xc1, 0x3a, 0x7e, 0x3c, 0x0a, 0x5f, 0x55, 0x0a, 0x1b, 0x28, 0x3b, 0x86, 0x94,
	0x2d, 0x37, 0x95, 0xd4, 0xa0, 0x5b, 0xdc, 0x81, 0xb7, 0x1a, 0x66, 0xbd, 0xe7, 0x67, 0x17, 0xbc,
	0x00, 0xb4, 0xe1, 0xc2, 0xf0, 0xcd, 0x46, 0x3e, 0x0d, 0x9e, 0x01, 0xe3, 0x17, 0x71, 0x6e, 0x49,
	0x7c, 0xab, 0xff, 0x71, 0x44, 0x9f, 0xe9, 0x07, 0x54, 0xb6, 0xb2, 0xd2, 0xea, 0xd0, 0xf1, 0xff,
	0x7e, 0x3e, 0xd1, 0xa3, 0x79, 0xdc, 0xb6, 0xc7, 0xde, 0x92, 0x87, 0xd5, 0xc0, 0x7f, 0xff, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x97, 0x7d, 0x01, 0x48, 0xc0, 0x02, 0x00, 0x00,
}
