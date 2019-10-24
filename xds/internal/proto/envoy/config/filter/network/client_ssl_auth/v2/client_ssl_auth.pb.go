// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/client_ssl_auth/v2/client_ssl_auth.proto

package envoy_config_filter_network_client_ssl_auth_v2

import (
	fmt "fmt"
	core "google.golang.org/grpc/xds/internal/proto/envoy/api/v2/core"
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

type ClientSSLAuth struct {
	AuthApiCluster       string             `protobuf:"bytes,1,opt,name=auth_api_cluster,json=authApiCluster,proto3" json:"auth_api_cluster,omitempty"`
	StatPrefix           string             `protobuf:"bytes,2,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	RefreshDelay         *duration.Duration `protobuf:"bytes,3,opt,name=refresh_delay,json=refreshDelay,proto3" json:"refresh_delay,omitempty"`
	IpWhiteList          []*core.CidrRange  `protobuf:"bytes,4,rep,name=ip_white_list,json=ipWhiteList,proto3" json:"ip_white_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ClientSSLAuth) Reset()         { *m = ClientSSLAuth{} }
func (m *ClientSSLAuth) String() string { return proto.CompactTextString(m) }
func (*ClientSSLAuth) ProtoMessage()    {}
func (*ClientSSLAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c05e9c9b57da130, []int{0}
}

func (m *ClientSSLAuth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientSSLAuth.Unmarshal(m, b)
}
func (m *ClientSSLAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientSSLAuth.Marshal(b, m, deterministic)
}
func (m *ClientSSLAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientSSLAuth.Merge(m, src)
}
func (m *ClientSSLAuth) XXX_Size() int {
	return xxx_messageInfo_ClientSSLAuth.Size(m)
}
func (m *ClientSSLAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientSSLAuth.DiscardUnknown(m)
}

var xxx_messageInfo_ClientSSLAuth proto.InternalMessageInfo

func (m *ClientSSLAuth) GetAuthApiCluster() string {
	if m != nil {
		return m.AuthApiCluster
	}
	return ""
}

func (m *ClientSSLAuth) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ClientSSLAuth) GetRefreshDelay() *duration.Duration {
	if m != nil {
		return m.RefreshDelay
	}
	return nil
}

func (m *ClientSSLAuth) GetIpWhiteList() []*core.CidrRange {
	if m != nil {
		return m.IpWhiteList
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientSSLAuth)(nil), "envoy.config.filter.network.client_ssl_auth.v2.ClientSSLAuth")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/client_ssl_auth/v2/client_ssl_auth.proto", fileDescriptor_2c05e9c9b57da130)
}

var fileDescriptor_2c05e9c9b57da130 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x6b, 0xe3, 0x30,
	0x18, 0xc5, 0x71, 0x12, 0xee, 0x38, 0xf9, 0x72, 0x1c, 0x5e, 0xce, 0x17, 0x8e, 0xab, 0xe9, 0xe4,
	0x49, 0xa2, 0xee, 0x5a, 0x4a, 0xe3, 0x64, 0xcc, 0x10, 0x9c, 0x42, 0x47, 0xa1, 0xc4, 0xb2, 0xfd,
	0x51, 0x61, 0x09, 0x49, 0x76, 0x92, 0x7f, 0xba, 0x7f, 0x40, 0xa7, 0x22, 0xcb, 0x19, 0x9a, 0xad,
	0x9b, 0xa4, 0xf7, 0xbe, 0xdf, 0xf7, 0x78, 0x42, 0x6b, 0xde, 0xf6, 0xf2, 0x4c, 0x0e, 0xb2, 0xad,
	0xa0, 0x26, 0x15, 0x08, 0xcb, 0x35, 0x69, 0xb9, 0x3d, 0x4a, 0xfd, 0x4a, 0x0e, 0x02, 0x78, 0x6b,
	0xa9, 0x31, 0x82, 0xb2, 0xce, 0x36, 0xa4, 0xcf, 0xae, 0x9f, 0xb0, 0xd2, 0xd2, 0xca, 0x08, 0x0f,
	0x14, 0xec, 0x29, 0xd8, 0x53, 0xf0, 0x48, 0xc1, 0xd7, 0x23, 0x7d, 0xb6, 0xb8, 0xf1, 0x5b, 0x99,
	0x82, 0x81, 0x29, 0x35, 0x27, 0xac, 0x2c, 0x35, 0x37, 0xc6, 0x03, 0x17, 0xff, 0x6b, 0x29, 0x6b,
	0xc1, 0xc9, 0x70, 0xdb, 0x77, 0x15, 0x29, 0x3b, 0xcd, 0x2c, 0xc8, 0x76, 0xd4, 0xff, 0xf4, 0x4c,
	0x40, 0xc9, 0x2c, 0x27, 0x97, 0x83, 0x17, 0x6e, 0xdf, 0x02, 0x34, 0x5f, 0x0d, 0x0b, 0x77, 0xbb,
	0xcd, 0xb2, 0xb3, 0x4d, 0x74, 0x87, 0x7e, 0xbb, 0xb5, 0x94, 0x29, 0xa0, 0x07, 0xd1, 0x19, 0xcb,
	0x75, 0x1c, 0x24, 0x41, 0xfa, 0x23, 0xff, 0xfe, 0x9e, 0xcf, 0xf4, 0x24, 0x09, 0x8a, 0x5f, 0xce,
	0xb0, 0x54, 0xb0, 0xf2, 0x72, 0x94, 0xa2, 0xd0, 0x58, 0x66, 0xa9, 0xd2, 0xbc, 0x82, 0x53, 0x3c,
	0xf9, 0xec, 0x46, 0x4e, 0xdb, 0x0e, 0x52, 0xf4, 0x88, 0xe6, 0x9a, 0x57, 0x9a, 0x9b, 0x86, 0x96,
	0x5c, 0xb0, 0x73, 0x3c, 0x4d, 0x82, 0x34, 0xcc, 0xfe, 0x62, 0x9f, 0x1f, 0x5f, 0xf2, 0xe3, 0xf5,
	0x98, 0xbf, 0xf8, 0x39, 0xfa, 0xd7, 0xce, 0x1e, 0x3d, 0xa1, 0x39, 0x28, 0x7a, 0x6c, 0xc0, 0x72,
	0x2a, 0xc0, 0xd8, 0x78, 0x96, 0x4c, 0xd3, 0x30, 0xfb, 0x37, 0x16, 0xca, 0x14, 0xe0, 0x3e, 0xc3,
	0xae, 0x20, 0xbc, 0x82, 0x52, 0x17, 0xac, 0xad, 0x79, 0x11, 0x82, 0x7a, 0x71, 0x13, 0x1b, 0x30,
	0x36, 0x7f, 0x46, 0x0f, 0x20, 0xbd, 0x5d, 0x69, 0x79, 0x3a, 0x7f, 0xf1, 0x2b, 0xf2, 0x68, 0x6c,
	0xcb, 0x08, 0xd7, 0xd6, 0xd6, 0xe5, 0xdd, 0x06, 0xfb, 0x6f, 0x43, 0xf0, 0xfb, 0x8f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xaf, 0x45, 0x93, 0xbe, 0x1f, 0x02, 0x00, 0x00,
}
