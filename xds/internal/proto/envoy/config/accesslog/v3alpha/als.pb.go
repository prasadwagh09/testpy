// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/accesslog/v3alpha/als.proto

package envoy_config_accesslog_v3alpha

import (
	fmt "fmt"
	core "google.golang.org/grpc/xds/internal/proto/envoy/api/v3alpha/core"
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

type HttpGrpcAccessLogConfig struct {
	CommonConfig                    *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	AdditionalRequestHeadersToLog   []string                   `protobuf:"bytes,2,rep,name=additional_request_headers_to_log,json=additionalRequestHeadersToLog,proto3" json:"additional_request_headers_to_log,omitempty"`
	AdditionalResponseHeadersToLog  []string                   `protobuf:"bytes,3,rep,name=additional_response_headers_to_log,json=additionalResponseHeadersToLog,proto3" json:"additional_response_headers_to_log,omitempty"`
	AdditionalResponseTrailersToLog []string                   `protobuf:"bytes,4,rep,name=additional_response_trailers_to_log,json=additionalResponseTrailersToLog,proto3" json:"additional_response_trailers_to_log,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                   `json:"-"`
	XXX_unrecognized                []byte                     `json:"-"`
	XXX_sizecache                   int32                      `json:"-"`
}

func (m *HttpGrpcAccessLogConfig) Reset()         { *m = HttpGrpcAccessLogConfig{} }
func (m *HttpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*HttpGrpcAccessLogConfig) ProtoMessage()    {}
func (*HttpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_938e3858beb7bdc4, []int{0}
}

func (m *HttpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *HttpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *HttpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpGrpcAccessLogConfig.Merge(m, src)
}
func (m *HttpGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Size(m)
}
func (m *HttpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HttpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *HttpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalRequestHeadersToLog() []string {
	if m != nil {
		return m.AdditionalRequestHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseHeadersToLog() []string {
	if m != nil {
		return m.AdditionalResponseHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseTrailersToLog() []string {
	if m != nil {
		return m.AdditionalResponseTrailersToLog
	}
	return nil
}

type TcpGrpcAccessLogConfig struct {
	CommonConfig         *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TcpGrpcAccessLogConfig) Reset()         { *m = TcpGrpcAccessLogConfig{} }
func (m *TcpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*TcpGrpcAccessLogConfig) ProtoMessage()    {}
func (*TcpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_938e3858beb7bdc4, []int{1}
}

func (m *TcpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *TcpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *TcpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpGrpcAccessLogConfig.Merge(m, src)
}
func (m *TcpGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Size(m)
}
func (m *TcpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TcpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *TcpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

type CommonGrpcAccessLogConfig struct {
	LogName              string                `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	GrpcService          *core.GrpcService     `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	BufferFlushInterval  *duration.Duration    `protobuf:"bytes,3,opt,name=buffer_flush_interval,json=bufferFlushInterval,proto3" json:"buffer_flush_interval,omitempty"`
	BufferSizeBytes      *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=buffer_size_bytes,json=bufferSizeBytes,proto3" json:"buffer_size_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CommonGrpcAccessLogConfig) Reset()         { *m = CommonGrpcAccessLogConfig{} }
func (m *CommonGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*CommonGrpcAccessLogConfig) ProtoMessage()    {}
func (*CommonGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_938e3858beb7bdc4, []int{2}
}

func (m *CommonGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *CommonGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (m *CommonGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonGrpcAccessLogConfig.Merge(m, src)
}
func (m *CommonGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Size(m)
}
func (m *CommonGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonGrpcAccessLogConfig proto.InternalMessageInfo

func (m *CommonGrpcAccessLogConfig) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *CommonGrpcAccessLogConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetBufferFlushInterval() *duration.Duration {
	if m != nil {
		return m.BufferFlushInterval
	}
	return nil
}

func (m *CommonGrpcAccessLogConfig) GetBufferSizeBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.BufferSizeBytes
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v3alpha.HttpGrpcAccessLogConfig")
	proto.RegisterType((*TcpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v3alpha.TcpGrpcAccessLogConfig")
	proto.RegisterType((*CommonGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v3alpha.CommonGrpcAccessLogConfig")
}

func init() {
	proto.RegisterFile("envoy/config/accesslog/v3alpha/als.proto", fileDescriptor_938e3858beb7bdc4)
}

var fileDescriptor_938e3858beb7bdc4 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0x36, 0x58, 0xe7, 0x0d, 0x01, 0x41, 0xb0, 0xae, 0x82, 0x52, 0xba, 0x4b, 0x41,
	0x28, 0x91, 0xd6, 0x13, 0xdc, 0x96, 0x21, 0xe8, 0x50, 0x85, 0xaa, 0xac, 0xc0, 0x31, 0x72, 0xd3,
	0x57, 0xd7, 0x92, 0x9b, 0x67, 0x6c, 0xa7, 0xd0, 0x1d, 0x39, 0xf2, 0xa7, 0x70, 0xe6, 0x7f, 0xe2,
	0x7f, 0xd8, 0x09, 0xc5, 0xce, 0xba, 0x8e, 0x6d, 0x70, 0xe4, 0x96, 0xc4, 0xdf, 0xf7, 0x7b, 0xdf,
	0x8b, 0x3f, 0xd2, 0x81, 0x7c, 0x8e, 0x8b, 0x28, 0xc3, 0x7c, 0xc2, 0x59, 0x44, 0xb3, 0x0c, 0xb4,
	0x16, 0xc8, 0xa2, 0x79, 0x97, 0x0a, 0x39, 0xa5, 0x11, 0x15, 0x3a, 0x94, 0x0a, 0x0d, 0x06, 0x4d,
	0xab, 0x0c, 0x9d, 0x32, 0x5c, 0x2a, 0xc3, 0x4a, 0xd9, 0x78, 0xe6, 0x48, 0x54, 0xf2, 0xa5, 0x39,
	0x43, 0x05, 0x11, 0x53, 0x32, 0x4b, 0x35, 0xa8, 0x39, 0xcf, 0xc0, 0xa1, 0x1a, 0x4d, 0x86, 0xc8,
	0x04, 0x44, 0xf6, 0x6d, 0x54, 0x4c, 0xa2, 0x71, 0xa1, 0xa8, 0xe1, 0x98, 0x5f, 0x77, 0xfe, 0x45,
	0x51, 0x29, 0x41, 0x55, 0x51, 0x1a, 0x3b, 0x73, 0x2a, 0xf8, 0x98, 0x1a, 0x88, 0xce, 0x1e, 0xdc,
	0x41, 0xfb, 0x97, 0x4f, 0x76, 0x7a, 0xc6, 0xc8, 0xb7, 0x4a, 0x66, 0x07, 0x36, 0x61, 0x1f, 0xd9,
	0xa1, 0x4d, 0x1c, 0x4c, 0xc9, 0xed, 0x0c, 0x67, 0x33, 0xcc, 0x53, 0xb7, 0x42, 0xdd, 0x6b, 0x79,
	0x9d, 0xad, 0xfd, 0x97, 0xe1, 0xdf, 0xf7, 0x0a, 0x0f, 0xad, 0xe9, 0x0a, 0x62, 0x5c, 0x3b, 0x8d,
	0x6f, 0x7e, 0xf7, 0xfc, 0xbb, 0x5e, 0xb2, 0xed, 0xc8, 0xd5, 0xa4, 0x1e, 0x79, 0x4a, 0xc7, 0x63,
	0x5e, 0x2e, 0x44, 0x45, 0xaa, 0xe0, 0x73, 0x01, 0xda, 0xa4, 0x53, 0xa0, 0x63, 0x50, 0x3a, 0x35,
	0x98, 0x0a, 0x64, 0x75, 0xbf, 0xb5, 0xd6, 0xd9, 0x4c, 0x1e, 0x9f, 0x0b, 0x13, 0xa7, 0xeb, 0x39,
	0xd9, 0x10, 0xfb, 0xc8, 0x82, 0x77, 0xa4, 0x7d, 0x81, 0xa4, 0x25, 0xe6, 0x1a, 0xfe, 0x44, 0xad,
	0x59, 0x54, 0x73, 0x15, 0xe5, 0x84, 0x17, 0x58, 0x7d, 0xb2, 0x77, 0x15, 0xcb, 0x28, 0xca, 0xc5,
	0x0a, 0x6c, 0xdd, 0xc2, 0x9e, 0x5c, 0x86, 0x0d, 0x2b, 0xa1, 0xa5, 0xb5, 0xbf, 0x79, 0xe4, 0xe1,
	0x30, 0xfb, 0xbf, 0x3f, 0xba, 0xfd, 0xd3, 0x27, 0xbb, 0xd7, 0xba, 0x82, 0x36, 0xa9, 0x09, 0x64,
	0x69, 0x4e, 0x67, 0x60, 0x23, 0x6c, 0xc6, 0x1b, 0xa7, 0xf1, 0xba, 0xf2, 0x5b, 0x5e, 0xb2, 0x21,
	0x90, 0xbd, 0xa7, 0x33, 0x08, 0x06, 0x64, 0x7b, 0xb5, 0x9f, 0x75, 0xdf, 0x46, 0xdd, 0xab, 0xa2,
	0x52, 0xc9, 0x97, 0xe9, 0xca, 0x2e, 0x87, 0xe5, 0x98, 0x63, 0x27, 0x5d, 0x09, 0xb5, 0xc5, 0xce,
	0x3f, 0x07, 0x9f, 0xc8, 0x83, 0x51, 0x31, 0x99, 0x80, 0x4a, 0x27, 0xa2, 0xd0, 0xd3, 0x94, 0xe7,
	0x06, 0xd4, 0x9c, 0x8a, 0xfa, 0x9a, 0x45, 0xef, 0x86, 0xae, 0xdb, 0xe1, 0x59, 0xb7, 0xc3, 0xd7,
	0x55, 0xf7, 0x2d, 0xf0, 0x87, 0xe7, 0x3f, 0xbf, 0x91, 0xdc, 0x77, 0x84, 0x37, 0x25, 0xe0, 0xa8,
	0xf2, 0x07, 0x3d, 0x72, 0xaf, 0x02, 0x6b, 0x7e, 0x02, 0xe9, 0x68, 0x61, 0x40, 0xd7, 0xd7, 0x2d,
	0xf4, 0xd1, 0x25, 0xe8, 0x87, 0xa3, 0xdc, 0x74, 0xf7, 0x3f, 0x52, 0x51, 0x40, 0x72, 0xc7, 0xd9,
	0x8e, 0xf9, 0x09, 0xc4, 0xa5, 0x29, 0x7e, 0x45, 0x5e, 0x70, 0x74, 0x2b, 0x4a, 0x85, 0x5f, 0x17,
	0xff, 0xb8, 0x98, 0xb8, 0x76, 0x20, 0xf4, 0xa0, 0x24, 0x0f, 0xbc, 0xd1, 0x2d, 0x3b, 0xa2, 0xfb,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x39, 0x22, 0xfe, 0xce, 0x38, 0x04, 0x00, 0x00,
}
