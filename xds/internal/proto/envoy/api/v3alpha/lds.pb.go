// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/lds.proto

package envoy_api_v3alpha

import (
	context "context"
	fmt "fmt"
	core "google.golang.org/grpc/xds/internal/proto/envoy/api/v3alpha/core"
	listener "google.golang.org/grpc/xds/internal/proto/envoy/api/v3alpha/listener"
	v2 "google.golang.org/grpc/xds/internal/proto/envoy/config/listener/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Listener_DrainType int32

const (
	Listener_DEFAULT     Listener_DrainType = 0
	Listener_MODIFY_ONLY Listener_DrainType = 1
)

var Listener_DrainType_name = map[int32]string{
	0: "DEFAULT",
	1: "MODIFY_ONLY",
}

var Listener_DrainType_value = map[string]int32{
	"DEFAULT":     0,
	"MODIFY_ONLY": 1,
}

func (x Listener_DrainType) String() string {
	return proto.EnumName(Listener_DrainType_name, int32(x))
}

func (Listener_DrainType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ba6811f9d8733fb9, []int{0, 0}
}

type Listener struct {
	Name                             string                            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address                          *core.Address                     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	FilterChains                     []*listener.FilterChain           `protobuf:"bytes,3,rep,name=filter_chains,json=filterChains,proto3" json:"filter_chains,omitempty"`
	PerConnectionBufferLimitBytes    *wrappers.UInt32Value             `protobuf:"bytes,5,opt,name=per_connection_buffer_limit_bytes,json=perConnectionBufferLimitBytes,proto3" json:"per_connection_buffer_limit_bytes,omitempty"`
	Metadata                         *core.Metadata                    `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	DeprecatedV1                     *Listener_DeprecatedV1            `protobuf:"bytes,7,opt,name=deprecated_v1,json=deprecatedV1,proto3" json:"deprecated_v1,omitempty"`
	DrainType                        Listener_DrainType                `protobuf:"varint,8,opt,name=drain_type,json=drainType,proto3,enum=envoy.api.v3alpha.Listener_DrainType" json:"drain_type,omitempty"`
	ListenerFilters                  []*listener.ListenerFilter        `protobuf:"bytes,9,rep,name=listener_filters,json=listenerFilters,proto3" json:"listener_filters,omitempty"`
	ListenerFiltersTimeout           *duration.Duration                `protobuf:"bytes,15,opt,name=listener_filters_timeout,json=listenerFiltersTimeout,proto3" json:"listener_filters_timeout,omitempty"`
	ContinueOnListenerFiltersTimeout bool                              `protobuf:"varint,17,opt,name=continue_on_listener_filters_timeout,json=continueOnListenerFiltersTimeout,proto3" json:"continue_on_listener_filters_timeout,omitempty"`
	Transparent                      *wrappers.BoolValue               `protobuf:"bytes,10,opt,name=transparent,proto3" json:"transparent,omitempty"`
	Freebind                         *wrappers.BoolValue               `protobuf:"bytes,11,opt,name=freebind,proto3" json:"freebind,omitempty"`
	SocketOptions                    []*core.SocketOption              `protobuf:"bytes,13,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
	TcpFastOpenQueueLength           *wrappers.UInt32Value             `protobuf:"bytes,12,opt,name=tcp_fast_open_queue_length,json=tcpFastOpenQueueLength,proto3" json:"tcp_fast_open_queue_length,omitempty"`
	TrafficDirection                 core.TrafficDirection             `protobuf:"varint,16,opt,name=traffic_direction,json=trafficDirection,proto3,enum=envoy.api.v3alpha.core.TrafficDirection" json:"traffic_direction,omitempty"`
	UdpListenerConfig                *listener.UdpListenerConfig       `protobuf:"bytes,18,opt,name=udp_listener_config,json=udpListenerConfig,proto3" json:"udp_listener_config,omitempty"`
	ApiListener                      *v2.ApiListener                   `protobuf:"bytes,19,opt,name=api_listener,json=apiListener,proto3" json:"api_listener,omitempty"`
	ConnectionBalanceConfig          *Listener_ConnectionBalanceConfig `protobuf:"bytes,20,opt,name=connection_balance_config,json=connectionBalanceConfig,proto3" json:"connection_balance_config,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}                          `json:"-"`
	XXX_unrecognized                 []byte                            `json:"-"`
	XXX_sizecache                    int32                             `json:"-"`
}

func (m *Listener) Reset()         { *m = Listener{} }
func (m *Listener) String() string { return proto.CompactTextString(m) }
func (*Listener) ProtoMessage()    {}
func (*Listener) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba6811f9d8733fb9, []int{0}
}

func (m *Listener) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener.Unmarshal(m, b)
}
func (m *Listener) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener.Marshal(b, m, deterministic)
}
func (m *Listener) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener.Merge(m, src)
}
func (m *Listener) XXX_Size() int {
	return xxx_messageInfo_Listener.Size(m)
}
func (m *Listener) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener.DiscardUnknown(m)
}

var xxx_messageInfo_Listener proto.InternalMessageInfo

func (m *Listener) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Listener) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Listener) GetFilterChains() []*listener.FilterChain {
	if m != nil {
		return m.FilterChains
	}
	return nil
}

func (m *Listener) GetPerConnectionBufferLimitBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.PerConnectionBufferLimitBytes
	}
	return nil
}

func (m *Listener) GetMetadata() *core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Listener) GetDeprecatedV1() *Listener_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

func (m *Listener) GetDrainType() Listener_DrainType {
	if m != nil {
		return m.DrainType
	}
	return Listener_DEFAULT
}

func (m *Listener) GetListenerFilters() []*listener.ListenerFilter {
	if m != nil {
		return m.ListenerFilters
	}
	return nil
}

func (m *Listener) GetListenerFiltersTimeout() *duration.Duration {
	if m != nil {
		return m.ListenerFiltersTimeout
	}
	return nil
}

func (m *Listener) GetContinueOnListenerFiltersTimeout() bool {
	if m != nil {
		return m.ContinueOnListenerFiltersTimeout
	}
	return false
}

func (m *Listener) GetTransparent() *wrappers.BoolValue {
	if m != nil {
		return m.Transparent
	}
	return nil
}

func (m *Listener) GetFreebind() *wrappers.BoolValue {
	if m != nil {
		return m.Freebind
	}
	return nil
}

func (m *Listener) GetSocketOptions() []*core.SocketOption {
	if m != nil {
		return m.SocketOptions
	}
	return nil
}

func (m *Listener) GetTcpFastOpenQueueLength() *wrappers.UInt32Value {
	if m != nil {
		return m.TcpFastOpenQueueLength
	}
	return nil
}

func (m *Listener) GetTrafficDirection() core.TrafficDirection {
	if m != nil {
		return m.TrafficDirection
	}
	return core.TrafficDirection_UNSPECIFIED
}

func (m *Listener) GetUdpListenerConfig() *listener.UdpListenerConfig {
	if m != nil {
		return m.UdpListenerConfig
	}
	return nil
}

func (m *Listener) GetApiListener() *v2.ApiListener {
	if m != nil {
		return m.ApiListener
	}
	return nil
}

func (m *Listener) GetConnectionBalanceConfig() *Listener_ConnectionBalanceConfig {
	if m != nil {
		return m.ConnectionBalanceConfig
	}
	return nil
}

type Listener_DeprecatedV1 struct {
	BindToPort           *wrappers.BoolValue `protobuf:"bytes,1,opt,name=bind_to_port,json=bindToPort,proto3" json:"bind_to_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Listener_DeprecatedV1) Reset()         { *m = Listener_DeprecatedV1{} }
func (m *Listener_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*Listener_DeprecatedV1) ProtoMessage()    {}
func (*Listener_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba6811f9d8733fb9, []int{0, 0}
}

func (m *Listener_DeprecatedV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener_DeprecatedV1.Unmarshal(m, b)
}
func (m *Listener_DeprecatedV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener_DeprecatedV1.Marshal(b, m, deterministic)
}
func (m *Listener_DeprecatedV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener_DeprecatedV1.Merge(m, src)
}
func (m *Listener_DeprecatedV1) XXX_Size() int {
	return xxx_messageInfo_Listener_DeprecatedV1.Size(m)
}
func (m *Listener_DeprecatedV1) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener_DeprecatedV1.DiscardUnknown(m)
}

var xxx_messageInfo_Listener_DeprecatedV1 proto.InternalMessageInfo

func (m *Listener_DeprecatedV1) GetBindToPort() *wrappers.BoolValue {
	if m != nil {
		return m.BindToPort
	}
	return nil
}

type Listener_ConnectionBalanceConfig struct {
	// Types that are valid to be assigned to BalanceType:
	//	*Listener_ConnectionBalanceConfig_ExactBalance_
	BalanceType          isListener_ConnectionBalanceConfig_BalanceType `protobuf_oneof:"balance_type"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *Listener_ConnectionBalanceConfig) Reset()         { *m = Listener_ConnectionBalanceConfig{} }
func (m *Listener_ConnectionBalanceConfig) String() string { return proto.CompactTextString(m) }
func (*Listener_ConnectionBalanceConfig) ProtoMessage()    {}
func (*Listener_ConnectionBalanceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba6811f9d8733fb9, []int{0, 1}
}

func (m *Listener_ConnectionBalanceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig.Unmarshal(m, b)
}
func (m *Listener_ConnectionBalanceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig.Marshal(b, m, deterministic)
}
func (m *Listener_ConnectionBalanceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener_ConnectionBalanceConfig.Merge(m, src)
}
func (m *Listener_ConnectionBalanceConfig) XXX_Size() int {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig.Size(m)
}
func (m *Listener_ConnectionBalanceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener_ConnectionBalanceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Listener_ConnectionBalanceConfig proto.InternalMessageInfo

type isListener_ConnectionBalanceConfig_BalanceType interface {
	isListener_ConnectionBalanceConfig_BalanceType()
}

type Listener_ConnectionBalanceConfig_ExactBalance_ struct {
	ExactBalance *Listener_ConnectionBalanceConfig_ExactBalance `protobuf:"bytes,1,opt,name=exact_balance,json=exactBalance,proto3,oneof"`
}

func (*Listener_ConnectionBalanceConfig_ExactBalance_) isListener_ConnectionBalanceConfig_BalanceType() {
}

func (m *Listener_ConnectionBalanceConfig) GetBalanceType() isListener_ConnectionBalanceConfig_BalanceType {
	if m != nil {
		return m.BalanceType
	}
	return nil
}

func (m *Listener_ConnectionBalanceConfig) GetExactBalance() *Listener_ConnectionBalanceConfig_ExactBalance {
	if x, ok := m.GetBalanceType().(*Listener_ConnectionBalanceConfig_ExactBalance_); ok {
		return x.ExactBalance
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Listener_ConnectionBalanceConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Listener_ConnectionBalanceConfig_ExactBalance_)(nil),
	}
}

type Listener_ConnectionBalanceConfig_ExactBalance struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Listener_ConnectionBalanceConfig_ExactBalance) Reset() {
	*m = Listener_ConnectionBalanceConfig_ExactBalance{}
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) String() string {
	return proto.CompactTextString(m)
}
func (*Listener_ConnectionBalanceConfig_ExactBalance) ProtoMessage() {}
func (*Listener_ConnectionBalanceConfig_ExactBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba6811f9d8733fb9, []int{0, 1, 0}
}

func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Unmarshal(m, b)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Marshal(b, m, deterministic)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Merge(m, src)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Size() int {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Size(m)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.DiscardUnknown(m)
}

var xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.api.v3alpha.Listener_DrainType", Listener_DrainType_name, Listener_DrainType_value)
	proto.RegisterType((*Listener)(nil), "envoy.api.v3alpha.Listener")
	proto.RegisterType((*Listener_DeprecatedV1)(nil), "envoy.api.v3alpha.Listener.DeprecatedV1")
	proto.RegisterType((*Listener_ConnectionBalanceConfig)(nil), "envoy.api.v3alpha.Listener.ConnectionBalanceConfig")
	proto.RegisterType((*Listener_ConnectionBalanceConfig_ExactBalance)(nil), "envoy.api.v3alpha.Listener.ConnectionBalanceConfig.ExactBalance")
}

func init() { proto.RegisterFile("envoy/api/v3alpha/lds.proto", fileDescriptor_ba6811f9d8733fb9) }

var fileDescriptor_ba6811f9d8733fb9 = []byte{
	// 1038 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcd, 0x72, 0x1b, 0x45,
	0x10, 0xce, 0xfa, 0x27, 0x96, 0x47, 0x3f, 0x96, 0xc7, 0x54, 0xbc, 0x11, 0x86, 0xc8, 0xc6, 0x54,
	0xe4, 0x50, 0xac, 0x12, 0x99, 0xe2, 0x90, 0xf2, 0x01, 0xcb, 0x8a, 0x2b, 0x09, 0x72, 0x6c, 0xd6,
	0x76, 0x8a, 0x1c, 0xa8, 0xad, 0xd1, 0x6e, 0x4b, 0x9e, 0xca, 0x7a, 0x66, 0x32, 0x33, 0x2b, 0xa2,
	0x2b, 0xc5, 0x81, 0xe2, 0xca, 0x43, 0xf0, 0x08, 0x3c, 0x08, 0xaf, 0xc0, 0x13, 0x70, 0xa2, 0x38,
	0x51, 0x3b, 0xbb, 0x2b, 0x2b, 0xfa, 0x23, 0x45, 0x15, 0xb7, 0xed, 0xe9, 0xef, 0xfb, 0xba, 0xa7,
	0xbb, 0xa7, 0x25, 0xf4, 0x21, 0xb0, 0x3e, 0x1f, 0xd4, 0x89, 0xa0, 0xf5, 0xfe, 0x3e, 0x09, 0xc5,
	0x15, 0xa9, 0x87, 0x81, 0x72, 0x84, 0xe4, 0x9a, 0xe3, 0x75, 0xe3, 0x74, 0x88, 0xa0, 0x4e, 0xea,
	0xac, 0xec, 0x4e, 0xe2, 0x7d, 0x2e, 0xa1, 0x4e, 0x82, 0x40, 0x82, 0x4a, 0x89, 0x95, 0xed, 0x19,
	0xa8, 0x0e, 0x51, 0x30, 0x1b, 0x12, 0x50, 0xe5, 0xf3, 0x3e, 0xc8, 0x41, 0x0a, 0xd9, 0x9b, 0x92,
	0x1b, 0x55, 0x1a, 0x18, 0xc8, 0xe1, 0x47, 0x0a, 0xfd, 0x62, 0x0e, 0x34, 0x0a, 0x84, 0x97, 0x19,
	0x9e, 0xcf, 0x59, 0x97, 0xf6, 0x52, 0xd6, 0x67, 0x09, 0x2b, 0x39, 0xbb, 0x21, 0xf4, 0x1b, 0xb1,
	0x90, 0x37, 0x16, 0x62, 0xab, 0xc7, 0x79, 0x2f, 0x04, 0x13, 0x83, 0x30, 0xc6, 0x35, 0xd1, 0x94,
	0xb3, 0xec, 0xc6, 0x1f, 0xa7, 0x5e, 0x63, 0x75, 0xa2, 0x6e, 0x3d, 0x88, 0xa4, 0x01, 0xcc, 0xf2,
	0x7f, 0x2f, 0x89, 0x10, 0x20, 0x33, 0xfe, 0x66, 0x9f, 0x84, 0x34, 0x20, 0x1a, 0xea, 0xd9, 0x47,
	0xe2, 0xd8, 0xf9, 0xad, 0x88, 0x72, 0xed, 0x34, 0x13, 0x8c, 0xd1, 0x12, 0x23, 0xd7, 0x60, 0x5b,
	0x55, 0xab, 0xb6, 0xea, 0x9a, 0x6f, 0x7c, 0x84, 0x56, 0xd2, 0xe2, 0xdb, 0x0b, 0x55, 0xab, 0x96,
	0x6f, 0xdc, 0x73, 0x26, 0xda, 0xe6, 0xc4, 0xd5, 0x77, 0x0e, 0x13, 0x58, 0x33, 0xf7, 0x77, 0x73,
	0xf9, 0x67, 0x6b, 0xa1, 0x6c, 0xb9, 0x19, 0x13, 0xb7, 0x51, 0xb1, 0x4b, 0x43, 0x1d, 0x17, 0xe8,
	0x8a, 0x50, 0xa6, 0xec, 0xc5, 0xea, 0x62, 0x2d, 0xdf, 0xb8, 0x3f, 0x45, 0x6a, 0x58, 0x96, 0x63,
	0x43, 0x38, 0x8a, 0xf1, 0x6e, 0xa1, 0x7b, 0x63, 0x28, 0xdc, 0x45, 0xdb, 0x22, 0xa9, 0x35, 0x03,
	0x3f, 0x2e, 0x82, 0xd7, 0x89, 0xba, 0x5d, 0x90, 0x5e, 0x48, 0xaf, 0xa9, 0xf6, 0x3a, 0x03, 0x0d,
	0xca, 0x5e, 0x36, 0xc9, 0x6e, 0x39, 0x49, 0x61, 0x9c, 0xac, 0x30, 0xce, 0xe5, 0x33, 0xa6, 0xf7,
	0x1b, 0x2f, 0x49, 0x18, 0x81, 0xfb, 0x91, 0x00, 0x79, 0x34, 0x54, 0x69, 0x1a, 0x91, 0x76, 0xac,
	0xd1, 0x8c, 0x25, 0xf0, 0x01, 0xca, 0x5d, 0x83, 0x26, 0x01, 0xd1, 0xc4, 0xbe, 0x6d, 0xe4, 0xaa,
	0xb3, 0xee, 0x7e, 0x92, 0xe2, 0xdc, 0x21, 0x03, 0x9f, 0xa0, 0x62, 0x00, 0x42, 0x82, 0x4f, 0x34,
	0x04, 0x5e, 0xff, 0x91, 0xbd, 0x62, 0x24, 0x6a, 0x53, 0x24, 0xb2, 0x06, 0x38, 0xad, 0x21, 0xe1,
	0xe5, 0x23, 0xb7, 0x10, 0x8c, 0x58, 0xb8, 0x85, 0x50, 0x20, 0x09, 0x65, 0x9e, 0x1e, 0x08, 0xb0,
	0x73, 0x55, 0xab, 0x56, 0x6a, 0x7c, 0x3a, 0x57, 0x2b, 0x46, 0x5f, 0x0c, 0x04, 0xb8, 0xab, 0x41,
	0xf6, 0x89, 0x2f, 0x51, 0x79, 0x38, 0xab, 0x49, 0x4d, 0x95, 0xbd, 0x6a, 0x7a, 0xf1, 0x60, 0x5e,
	0x2f, 0x32, 0xd1, 0xa4, 0x27, 0xee, 0x5a, 0xf8, 0x8e, 0xad, 0xf0, 0x39, 0xb2, 0xc7, 0x65, 0x3d,
	0x4d, 0xaf, 0x81, 0x47, 0xda, 0x5e, 0x33, 0xd7, 0xbe, 0x3b, 0xd1, 0x88, 0x56, 0x3a, 0xc1, 0xee,
	0x9d, 0x31, 0xb5, 0x8b, 0x84, 0x88, 0x5f, 0xa0, 0x5d, 0x9f, 0x33, 0x4d, 0x59, 0x04, 0x1e, 0x67,
	0xde, 0xcc, 0x00, 0xeb, 0x55, 0xab, 0x96, 0x73, 0xab, 0x19, 0xf6, 0x94, 0xb5, 0xa7, 0xeb, 0x1d,
	0xa0, 0xbc, 0x96, 0x84, 0x29, 0x41, 0x24, 0x30, 0x6d, 0x23, 0x93, 0x57, 0x65, 0x22, 0xaf, 0x26,
	0xe7, 0x61, 0x32, 0x1e, 0xa3, 0x70, 0xfc, 0x25, 0xca, 0x75, 0x25, 0x40, 0x87, 0xb2, 0xc0, 0xce,
	0xff, 0x2b, 0x75, 0x88, 0xc5, 0x5f, 0xa3, 0x92, 0xe2, 0xfe, 0x6b, 0xd0, 0x1e, 0x17, 0xe6, 0x45,
	0xdb, 0x45, 0x53, 0xef, 0xdd, 0x59, 0xa3, 0x74, 0x6e, 0xd0, 0xa7, 0x06, 0xec, 0x16, 0xd5, 0x88,
	0xa5, 0xf0, 0xb7, 0xa8, 0xa2, 0x7d, 0xe1, 0x75, 0x89, 0x8a, 0xe5, 0x80, 0x79, 0x6f, 0x22, 0x88,
	0xc0, 0x0b, 0x81, 0xf5, 0xf4, 0x95, 0x5d, 0x78, 0x8f, 0x91, 0xbf, 0xa3, 0x7d, 0x71, 0x4c, 0x94,
	0x3e, 0x15, 0xc0, 0xbe, 0x89, 0xc9, 0x6d, 0xc3, 0xc5, 0x97, 0x68, 0x5d, 0x4b, 0xd2, 0xed, 0x52,
	0xdf, 0x0b, 0xa8, 0x4c, 0x1e, 0x84, 0x5d, 0x36, 0x53, 0x56, 0x9b, 0x95, 0xe9, 0x45, 0x42, 0x68,
	0x65, 0x78, 0xb7, 0xac, 0xc7, 0x4e, 0xf0, 0x77, 0x68, 0x63, 0xca, 0x7e, 0xb4, 0xb1, 0xc9, 0xf4,
	0xf3, 0x79, 0x23, 0x77, 0x19, 0x88, 0xac, 0x8f, 0x47, 0x86, 0xe4, 0xae, 0x47, 0xe3, 0x47, 0xf8,
	0x29, 0x2a, 0x8c, 0xae, 0x52, 0x7b, 0xc3, 0xe8, 0x66, 0xcf, 0x22, 0x5d, 0xc6, 0x43, 0xc9, 0x7e,
	0xc3, 0x39, 0x14, 0x34, 0x93, 0x70, 0xf3, 0xe4, 0xc6, 0xc0, 0x1c, 0xdd, 0x1d, 0xdd, 0x27, 0x24,
	0x24, 0xcc, 0x87, 0x2c, 0xdd, 0x0f, 0x8c, 0xec, 0xfe, 0xbc, 0xd7, 0x36, 0xb2, 0x46, 0x12, 0x6e,
	0x9a, 0xf4, 0xa6, 0x3f, 0xdd, 0x51, 0x69, 0xa3, 0xc2, 0xe8, 0x6b, 0xc7, 0x07, 0xa8, 0x10, 0xcf,
	0x8b, 0xa7, 0xb9, 0x27, 0xb8, 0xd4, 0x66, 0x07, 0xcf, 0x9f, 0x31, 0x14, 0xe3, 0x2f, 0xf8, 0x19,
	0x97, 0xba, 0xf2, 0xab, 0x85, 0x36, 0x67, 0xa4, 0x80, 0x7b, 0xa8, 0x08, 0x6f, 0x89, 0xaf, 0xb3,
	0x5b, 0xa5, 0xd2, 0x5f, 0xfd, 0x87, 0xeb, 0x38, 0x4f, 0x62, 0xa1, 0xf4, 0xe8, 0xe9, 0x2d, 0xb7,
	0x00, 0x23, 0x76, 0xa5, 0x84, 0x0a, 0xa3, 0xfe, 0xe6, 0x06, 0x2a, 0x64, 0x85, 0x8c, 0x97, 0x16,
	0x5e, 0xfc, 0xab, 0x69, 0xed, 0xec, 0xa1, 0xd5, 0xe1, 0x66, 0xc2, 0x79, 0xb4, 0xd2, 0x7a, 0x72,
	0x7c, 0x78, 0xd9, 0xbe, 0x28, 0xdf, 0xc2, 0x6b, 0x28, 0x7f, 0x72, 0xda, 0x7a, 0x76, 0xfc, 0xca,
	0x3b, 0x7d, 0xd1, 0x7e, 0x55, 0xb6, 0x9e, 0x2f, 0xe5, 0x4a, 0xe5, 0xb5, 0xe7, 0x4b, 0xb9, 0xa5,
	0xf2, 0xb2, 0x5b, 0x8e, 0x14, 0x78, 0x5c, 0xd2, 0x1e, 0x65, 0x24, 0xf4, 0x02, 0xa5, 0x1b, 0x7f,
	0x2e, 0x20, 0x3b, 0xcb, 0xb7, 0x95, 0xfd, 0xb4, 0x9f, 0x83, 0xec, 0x53, 0x1f, 0xf0, 0x6b, 0x54,
	0x6a, 0x41, 0xa8, 0x49, 0x06, 0x50, 0x78, 0xda, 0x14, 0x1b, 0xc8, 0x90, 0xeb, 0xc2, 0x9b, 0x08,
	0x94, 0xae, 0xec, 0xbd, 0x07, 0x52, 0x09, 0xce, 0x14, 0xec, 0xdc, 0xaa, 0x59, 0x0f, 0x2d, 0xdc,
	0x41, 0x6b, 0xe7, 0x5a, 0x02, 0xb9, 0xbe, 0x89, 0xf6, 0xc9, 0x34, 0x8d, 0xf1, 0x40, 0xbb, 0xf3,
	0x41, 0xef, 0xc4, 0xf8, 0xd1, 0x42, 0xa5, 0x63, 0xd0, 0xfe, 0xd5, 0xff, 0x12, 0xe3, 0xfe, 0x0f,
	0xbf, 0xff, 0xf1, 0xcb, 0xc2, 0xf6, 0xce, 0xd6, 0xe4, 0x9f, 0xa5, 0xc7, 0xd9, 0xf3, 0x51, 0x8f,
	0xad, 0x07, 0xcd, 0x87, 0xe8, 0x1e, 0xe5, 0x89, 0xa4, 0x90, 0xfc, 0xed, 0x60, 0x52, 0xbd, 0x99,
	0x6b, 0x07, 0xea, 0x2c, 0x9e, 0xd6, 0x33, 0xeb, 0x27, 0xcb, 0xea, 0xdc, 0x36, 0x93, 0xbb, 0xff,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x9e, 0x3f, 0xc3, 0x0a, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ListenerDiscoveryServiceClient is the client API for ListenerDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ListenerDiscoveryServiceClient interface {
	DeltaListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_DeltaListenersClient, error)
	StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error)
	FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type listenerDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewListenerDiscoveryServiceClient(cc *grpc.ClientConn) ListenerDiscoveryServiceClient {
	return &listenerDiscoveryServiceClient{cc}
}

func (c *listenerDiscoveryServiceClient) DeltaListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_DeltaListenersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ListenerDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v3alpha.ListenerDiscoveryService/DeltaListeners", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceDeltaListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_DeltaListenersClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceDeltaListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceDeltaListenersClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceDeltaListenersClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ListenerDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v3alpha.ListenerDiscoveryService/StreamListeners", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceStreamListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_StreamListenersClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceStreamListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceStreamListenersClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v3alpha.ListenerDiscoveryService/FetchListeners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListenerDiscoveryServiceServer is the server API for ListenerDiscoveryService service.
type ListenerDiscoveryServiceServer interface {
	DeltaListeners(ListenerDiscoveryService_DeltaListenersServer) error
	StreamListeners(ListenerDiscoveryService_StreamListenersServer) error
	FetchListeners(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

// UnimplementedListenerDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedListenerDiscoveryServiceServer struct {
}

func (*UnimplementedListenerDiscoveryServiceServer) DeltaListeners(srv ListenerDiscoveryService_DeltaListenersServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaListeners not implemented")
}
func (*UnimplementedListenerDiscoveryServiceServer) StreamListeners(srv ListenerDiscoveryService_StreamListenersServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamListeners not implemented")
}
func (*UnimplementedListenerDiscoveryServiceServer) FetchListeners(ctx context.Context, req *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchListeners not implemented")
}

func RegisterListenerDiscoveryServiceServer(s *grpc.Server, srv ListenerDiscoveryServiceServer) {
	s.RegisterService(&_ListenerDiscoveryService_serviceDesc, srv)
}

func _ListenerDiscoveryService_DeltaListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).DeltaListeners(&listenerDiscoveryServiceDeltaListenersServer{stream})
}

type ListenerDiscoveryService_DeltaListenersServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceDeltaListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceDeltaListenersServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceDeltaListenersServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_StreamListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).StreamListeners(&listenerDiscoveryServiceStreamListenersServer{stream})
}

type ListenerDiscoveryService_StreamListenersServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceStreamListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceStreamListenersServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_FetchListeners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v3alpha.ListenerDiscoveryService/FetchListeners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ListenerDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v3alpha.ListenerDiscoveryService",
	HandlerType: (*ListenerDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchListeners",
			Handler:    _ListenerDiscoveryService_FetchListeners_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaListeners",
			Handler:       _ListenerDiscoveryService_DeltaListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamListeners",
			Handler:       _ListenerDiscoveryService_StreamListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v3alpha/lds.proto",
}
