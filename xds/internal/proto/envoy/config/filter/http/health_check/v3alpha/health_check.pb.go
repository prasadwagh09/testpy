// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/health_check/v3alpha/health_check.proto

package envoy_config_filter_http_health_check_v3alpha

import (
	fmt "fmt"
	route "google.golang.org/grpc/xds/internal/proto/envoy/api/v3alpha/route"
	v3alpha "google.golang.org/grpc/xds/internal/proto/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type HealthCheck struct {
	PassThroughMode              *wrappers.BoolValue         `protobuf:"bytes,1,opt,name=pass_through_mode,json=passThroughMode,proto3" json:"pass_through_mode,omitempty"`
	CacheTime                    *duration.Duration          `protobuf:"bytes,3,opt,name=cache_time,json=cacheTime,proto3" json:"cache_time,omitempty"`
	ClusterMinHealthyPercentages map[string]*v3alpha.Percent `protobuf:"bytes,4,rep,name=cluster_min_healthy_percentages,json=clusterMinHealthyPercentages,proto3" json:"cluster_min_healthy_percentages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Headers                      []*route.HeaderMatcher      `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                    `json:"-"`
	XXX_unrecognized             []byte                      `json:"-"`
	XXX_sizecache                int32                       `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3e9e4bfd87bbf04, []int{0}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetPassThroughMode() *wrappers.BoolValue {
	if m != nil {
		return m.PassThroughMode
	}
	return nil
}

func (m *HealthCheck) GetCacheTime() *duration.Duration {
	if m != nil {
		return m.CacheTime
	}
	return nil
}

func (m *HealthCheck) GetClusterMinHealthyPercentages() map[string]*v3alpha.Percent {
	if m != nil {
		return m.ClusterMinHealthyPercentages
	}
	return nil
}

func (m *HealthCheck) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "envoy.config.filter.http.health_check.v3alpha.HealthCheck")
	proto.RegisterMapType((map[string]*v3alpha.Percent)(nil), "envoy.config.filter.http.health_check.v3alpha.HealthCheck.ClusterMinHealthyPercentagesEntry")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/health_check/v3alpha/health_check.proto", fileDescriptor_b3e9e4bfd87bbf04)
}

var fileDescriptor_b3e9e4bfd87bbf04 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0x66, 0x85, 0xcd, 0x3d, 0x50, 0x72, 0x21, 0x14, 0x04, 0x05, 0x24, 0xb4, 0x0b,
	0xb6, 0xd8, 0x2e, 0x13, 0x5c, 0xa6, 0x0c, 0xa4, 0x09, 0xa9, 0x52, 0x89, 0x26, 0x4e, 0x48, 0x91,
	0xe7, 0xbc, 0xc6, 0xd6, 0xdc, 0xd8, 0x72, 0x9c, 0x40, 0xbe, 0x02, 0x5f, 0x83, 0xcf, 0xc8, 0x85,
	0x13, 0x8a, 0xed, 0x86, 0x4d, 0x08, 0xd0, 0x2e, 0x95, 0xeb, 0xf7, 0x7f, 0xbf, 0xff, 0xf3, 0xff,
	0x05, 0x9d, 0x42, 0xdd, 0xa9, 0x9e, 0x30, 0x55, 0x6f, 0x44, 0x45, 0x36, 0x42, 0x5a, 0x30, 0x84,
	0x5b, 0xab, 0x09, 0x07, 0x2a, 0x2d, 0x2f, 0x18, 0x07, 0x76, 0x45, 0xba, 0x63, 0x2a, 0x35, 0xa7,
	0x37, 0x2e, 0xb1, 0x36, 0xca, 0xaa, 0xe4, 0x95, 0x23, 0x60, 0x4f, 0xc0, 0x9e, 0x80, 0x07, 0x02,
	0xbe, 0x21, 0x0e, 0x84, 0xc5, 0x0b, 0x6f, 0x48, 0xb5, 0x18, 0xa1, 0x46, 0xb5, 0x16, 0xfc, 0xaf,
	0x67, 0x2e, 0x96, 0x5e, 0x64, 0x7b, 0x0d, 0xa3, 0x4a, 0x83, 0x61, 0x50, 0xdb, 0xa0, 0x78, 0x52,
	0x29, 0x55, 0x49, 0x20, 0xee, 0xdf, 0x65, 0xbb, 0x21, 0x65, 0x6b, 0xa8, 0x15, 0xaa, 0xfe, 0x5b,
	0xfd, 0x8b, 0xa1, 0x5a, 0x83, 0x69, 0x42, 0xfd, 0x41, 0x47, 0xa5, 0x28, 0xa9, 0x05, 0xb2, 0x3b,
	0xf8, 0xc2, 0xf3, 0x1f, 0x31, 0x9a, 0x9d, 0xbb, 0xc1, 0xcf, 0x86, 0xb9, 0x93, 0x35, 0xba, 0xaf,
	0x69, 0xd3, 0x14, 0x96, 0x1b, 0xd5, 0x56, 0xbc, 0xd8, 0xaa, 0x12, 0xd2, 0x68, 0x19, 0x1d, 0xce,
	0x8e, 0x16, 0xd8, 0x9b, 0xe0, 0x9d, 0x09, 0xce, 0x94, 0x92, 0x9f, 0xa8, 0x6c, 0x21, 0xdb, 0xff,
	0x99, 0x4d, 0xbf, 0x45, 0x93, 0x79, 0x94, 0xdf, 0x1b, 0xda, 0x2f, 0x7c, 0xf7, 0x4a, 0x95, 0x90,
	0x9c, 0x20, 0xc4, 0x28, 0xe3, 0x50, 0x58, 0xb1, 0x85, 0x34, 0x76, 0xa8, 0x87, 0x7f, 0xa0, 0xde,
	0x85, 0xf7, 0xe4, 0x07, 0x4e, 0x7c, 0x21, 0xb6, 0x90, 0x7c, 0x8f, 0xd0, 0x53, 0x26, 0xdb, 0xc6,
	0x82, 0x29, 0xb6, 0xa2, 0x2e, 0x7c, 0xc0, 0x7d, 0x11, 0xa2, 0xa1, 0x15, 0x34, 0xe9, 0xde, 0x32,
	0x3e, 0x9c, 0x1d, 0x7d, 0xc6, 0xb7, 0xda, 0x0a, 0xbe, 0xf6, 0x62, 0x7c, 0xe6, 0x1d, 0x56, 0xa2,
	0xf6, 0xb7, 0xfd, 0xfa, 0x37, 0xfe, 0x7d, 0x6d, 0x4d, 0x9f, 0x3f, 0x66, 0xff, 0x90, 0x24, 0xa7,
	0xe8, 0x2e, 0x07, 0x5a, 0x82, 0x69, 0xd2, 0xa9, 0x1b, 0xe6, 0x65, 0x18, 0x86, 0x6a, 0x31, 0x1a,
	0xfa, 0x6d, 0x9f, 0x3b, 0xdd, 0x8a, 0x5a, 0xc6, 0xc1, 0xe4, 0xbb, 0xb6, 0x85, 0x44, 0xcf, 0xfe,
	0x3b, 0x44, 0x32, 0x47, 0xf1, 0x15, 0xf4, 0x6e, 0x15, 0x07, 0xf9, 0x70, 0x4c, 0x5e, 0xa3, 0x69,
	0x37, 0x84, 0x9f, 0x4e, 0x5c, 0xa6, 0x8f, 0x82, 0xed, 0xf0, 0x15, 0x8d, 0xbe, 0x01, 0x93, 0x7b,
	0xe5, 0x9b, 0xc9, 0x49, 0xf4, 0x61, 0x6f, 0x7f, 0x32, 0x8f, 0xb3, 0x8f, 0xe8, 0xad, 0x50, 0xbe,
	0x43, 0x1b, 0xf5, 0xb5, 0xbf, 0x5d, 0x80, 0xd9, 0xfc, 0x5a, 0x82, 0xeb, 0x61, 0x87, 0xeb, 0xe8,
	0xf2, 0x8e, 0x5b, 0xe6, 0xf1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xc1, 0xfe, 0x1b, 0x64,
	0x03, 0x00, 0x00,
}
