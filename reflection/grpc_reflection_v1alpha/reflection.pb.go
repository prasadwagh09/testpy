// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reflection.proto

/*
Package grpc_reflection_v1alpha is a generated protocol buffer package.

It is generated from these files:
	reflection.proto

It has these top-level messages:
	ServerReflectionRequest
	ExtensionRequest
	ServerReflectionResponse
	FileDescriptorResponse
	ExtensionNumberResponse
	ListServiceResponse
	ServiceResponse
	ErrorResponse
*/
package grpc_reflection_v1alpha

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// The message sent by the client when calling ServerReflectionInfo method.
type ServerReflectionRequest struct {
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	// To use reflection service, the client should set one of the following
	// fields in message_request. The server distinguishes requests by their
	// defined field and then handles them using corresponding methods.
	//
	// Types that are valid to be assigned to MessageRequest:
	//	*ServerReflectionRequest_FileByFilename
	//	*ServerReflectionRequest_FileContainingSymbol
	//	*ServerReflectionRequest_FileContainingExtension
	//	*ServerReflectionRequest_AllExtensionNumbersOfType
	//	*ServerReflectionRequest_ListServices
	MessageRequest isServerReflectionRequest_MessageRequest `protobuf_oneof:"message_request"`
}

func (m *ServerReflectionRequest) Reset()                    { *m = ServerReflectionRequest{} }
func (m *ServerReflectionRequest) String() string            { return proto.CompactTextString(m) }
func (*ServerReflectionRequest) ProtoMessage()               {}
func (*ServerReflectionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isServerReflectionRequest_MessageRequest interface {
	isServerReflectionRequest_MessageRequest()
}

type ServerReflectionRequest_FileByFilename struct {
	FileByFilename string `protobuf:"bytes,3,opt,name=file_by_filename,json=fileByFilename,oneof"`
}
type ServerReflectionRequest_FileContainingSymbol struct {
	FileContainingSymbol string `protobuf:"bytes,4,opt,name=file_containing_symbol,json=fileContainingSymbol,oneof"`
}
type ServerReflectionRequest_FileContainingExtension struct {
	FileContainingExtension *ExtensionRequest `protobuf:"bytes,5,opt,name=file_containing_extension,json=fileContainingExtension,oneof"`
}
type ServerReflectionRequest_AllExtensionNumbersOfType struct {
	AllExtensionNumbersOfType string `protobuf:"bytes,6,opt,name=all_extension_numbers_of_type,json=allExtensionNumbersOfType,oneof"`
}
type ServerReflectionRequest_ListServices struct {
	ListServices string `protobuf:"bytes,7,opt,name=list_services,json=listServices,oneof"`
}

func (*ServerReflectionRequest_FileByFilename) isServerReflectionRequest_MessageRequest()            {}
func (*ServerReflectionRequest_FileContainingSymbol) isServerReflectionRequest_MessageRequest()      {}
func (*ServerReflectionRequest_FileContainingExtension) isServerReflectionRequest_MessageRequest()   {}
func (*ServerReflectionRequest_AllExtensionNumbersOfType) isServerReflectionRequest_MessageRequest() {}
func (*ServerReflectionRequest_ListServices) isServerReflectionRequest_MessageRequest()              {}

func (m *ServerReflectionRequest) GetMessageRequest() isServerReflectionRequest_MessageRequest {
	if m != nil {
		return m.MessageRequest
	}
	return nil
}

func (m *ServerReflectionRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ServerReflectionRequest) GetFileByFilename() string {
	if x, ok := m.GetMessageRequest().(*ServerReflectionRequest_FileByFilename); ok {
		return x.FileByFilename
	}
	return ""
}

func (m *ServerReflectionRequest) GetFileContainingSymbol() string {
	if x, ok := m.GetMessageRequest().(*ServerReflectionRequest_FileContainingSymbol); ok {
		return x.FileContainingSymbol
	}
	return ""
}

func (m *ServerReflectionRequest) GetFileContainingExtension() *ExtensionRequest {
	if x, ok := m.GetMessageRequest().(*ServerReflectionRequest_FileContainingExtension); ok {
		return x.FileContainingExtension
	}
	return nil
}

func (m *ServerReflectionRequest) GetAllExtensionNumbersOfType() string {
	if x, ok := m.GetMessageRequest().(*ServerReflectionRequest_AllExtensionNumbersOfType); ok {
		return x.AllExtensionNumbersOfType
	}
	return ""
}

func (m *ServerReflectionRequest) GetListServices() string {
	if x, ok := m.GetMessageRequest().(*ServerReflectionRequest_ListServices); ok {
		return x.ListServices
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ServerReflectionRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ServerReflectionRequest_OneofMarshaler, _ServerReflectionRequest_OneofUnmarshaler, _ServerReflectionRequest_OneofSizer, []interface{}{
		(*ServerReflectionRequest_FileByFilename)(nil),
		(*ServerReflectionRequest_FileContainingSymbol)(nil),
		(*ServerReflectionRequest_FileContainingExtension)(nil),
		(*ServerReflectionRequest_AllExtensionNumbersOfType)(nil),
		(*ServerReflectionRequest_ListServices)(nil),
	}
}

func _ServerReflectionRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ServerReflectionRequest)
	// message_request
	switch x := m.MessageRequest.(type) {
	case *ServerReflectionRequest_FileByFilename:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.FileByFilename)
	case *ServerReflectionRequest_FileContainingSymbol:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.FileContainingSymbol)
	case *ServerReflectionRequest_FileContainingExtension:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FileContainingExtension); err != nil {
			return err
		}
	case *ServerReflectionRequest_AllExtensionNumbersOfType:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.AllExtensionNumbersOfType)
	case *ServerReflectionRequest_ListServices:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ListServices)
	case nil:
	default:
		return fmt.Errorf("ServerReflectionRequest.MessageRequest has unexpected type %T", x)
	}
	return nil
}

func _ServerReflectionRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ServerReflectionRequest)
	switch tag {
	case 3: // message_request.file_by_filename
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.MessageRequest = &ServerReflectionRequest_FileByFilename{x}
		return true, err
	case 4: // message_request.file_containing_symbol
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.MessageRequest = &ServerReflectionRequest_FileContainingSymbol{x}
		return true, err
	case 5: // message_request.file_containing_extension
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ExtensionRequest)
		err := b.DecodeMessage(msg)
		m.MessageRequest = &ServerReflectionRequest_FileContainingExtension{msg}
		return true, err
	case 6: // message_request.all_extension_numbers_of_type
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.MessageRequest = &ServerReflectionRequest_AllExtensionNumbersOfType{x}
		return true, err
	case 7: // message_request.list_services
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.MessageRequest = &ServerReflectionRequest_ListServices{x}
		return true, err
	default:
		return false, nil
	}
}

func _ServerReflectionRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ServerReflectionRequest)
	// message_request
	switch x := m.MessageRequest.(type) {
	case *ServerReflectionRequest_FileByFilename:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.FileByFilename)))
		n += len(x.FileByFilename)
	case *ServerReflectionRequest_FileContainingSymbol:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.FileContainingSymbol)))
		n += len(x.FileContainingSymbol)
	case *ServerReflectionRequest_FileContainingExtension:
		s := proto.Size(x.FileContainingExtension)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ServerReflectionRequest_AllExtensionNumbersOfType:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AllExtensionNumbersOfType)))
		n += len(x.AllExtensionNumbersOfType)
	case *ServerReflectionRequest_ListServices:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.ListServices)))
		n += len(x.ListServices)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The type name and extension number sent by the client when requesting
// file_containing_extension.
type ExtensionRequest struct {
	// Fully-qualified type name. The format should be <package>.<type>
	ContainingType  string `protobuf:"bytes,1,opt,name=containing_type,json=containingType" json:"containing_type,omitempty"`
	ExtensionNumber int32  `protobuf:"varint,2,opt,name=extension_number,json=extensionNumber" json:"extension_number,omitempty"`
}

func (m *ExtensionRequest) Reset()                    { *m = ExtensionRequest{} }
func (m *ExtensionRequest) String() string            { return proto.CompactTextString(m) }
func (*ExtensionRequest) ProtoMessage()               {}
func (*ExtensionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ExtensionRequest) GetContainingType() string {
	if m != nil {
		return m.ContainingType
	}
	return ""
}

func (m *ExtensionRequest) GetExtensionNumber() int32 {
	if m != nil {
		return m.ExtensionNumber
	}
	return 0
}

// The message sent by the server to answer ServerReflectionInfo method.
type ServerReflectionResponse struct {
	ValidHost       string                   `protobuf:"bytes,1,opt,name=valid_host,json=validHost" json:"valid_host,omitempty"`
	OriginalRequest *ServerReflectionRequest `protobuf:"bytes,2,opt,name=original_request,json=originalRequest" json:"original_request,omitempty"`
	// The server set one of the following fields accroding to the message_request
	// in the request.
	//
	// Types that are valid to be assigned to MessageResponse:
	//	*ServerReflectionResponse_FileDescriptorResponse
	//	*ServerReflectionResponse_AllExtensionNumbersResponse
	//	*ServerReflectionResponse_ListServicesResponse
	//	*ServerReflectionResponse_ErrorResponse
	MessageResponse isServerReflectionResponse_MessageResponse `protobuf_oneof:"message_response"`
}

func (m *ServerReflectionResponse) Reset()                    { *m = ServerReflectionResponse{} }
func (m *ServerReflectionResponse) String() string            { return proto.CompactTextString(m) }
func (*ServerReflectionResponse) ProtoMessage()               {}
func (*ServerReflectionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isServerReflectionResponse_MessageResponse interface {
	isServerReflectionResponse_MessageResponse()
}

type ServerReflectionResponse_FileDescriptorResponse struct {
	FileDescriptorResponse *FileDescriptorResponse `protobuf:"bytes,4,opt,name=file_descriptor_response,json=fileDescriptorResponse,oneof"`
}
type ServerReflectionResponse_AllExtensionNumbersResponse struct {
	AllExtensionNumbersResponse *ExtensionNumberResponse `protobuf:"bytes,5,opt,name=all_extension_numbers_response,json=allExtensionNumbersResponse,oneof"`
}
type ServerReflectionResponse_ListServicesResponse struct {
	ListServicesResponse *ListServiceResponse `protobuf:"bytes,6,opt,name=list_services_response,json=listServicesResponse,oneof"`
}
type ServerReflectionResponse_ErrorResponse struct {
	ErrorResponse *ErrorResponse `protobuf:"bytes,7,opt,name=error_response,json=errorResponse,oneof"`
}

func (*ServerReflectionResponse_FileDescriptorResponse) isServerReflectionResponse_MessageResponse() {}
func (*ServerReflectionResponse_AllExtensionNumbersResponse) isServerReflectionResponse_MessageResponse() {
}
func (*ServerReflectionResponse_ListServicesResponse) isServerReflectionResponse_MessageResponse() {}
func (*ServerReflectionResponse_ErrorResponse) isServerReflectionResponse_MessageResponse()        {}

func (m *ServerReflectionResponse) GetMessageResponse() isServerReflectionResponse_MessageResponse {
	if m != nil {
		return m.MessageResponse
	}
	return nil
}

func (m *ServerReflectionResponse) GetValidHost() string {
	if m != nil {
		return m.ValidHost
	}
	return ""
}

func (m *ServerReflectionResponse) GetOriginalRequest() *ServerReflectionRequest {
	if m != nil {
		return m.OriginalRequest
	}
	return nil
}

func (m *ServerReflectionResponse) GetFileDescriptorResponse() *FileDescriptorResponse {
	if x, ok := m.GetMessageResponse().(*ServerReflectionResponse_FileDescriptorResponse); ok {
		return x.FileDescriptorResponse
	}
	return nil
}

func (m *ServerReflectionResponse) GetAllExtensionNumbersResponse() *ExtensionNumberResponse {
	if x, ok := m.GetMessageResponse().(*ServerReflectionResponse_AllExtensionNumbersResponse); ok {
		return x.AllExtensionNumbersResponse
	}
	return nil
}

func (m *ServerReflectionResponse) GetListServicesResponse() *ListServiceResponse {
	if x, ok := m.GetMessageResponse().(*ServerReflectionResponse_ListServicesResponse); ok {
		return x.ListServicesResponse
	}
	return nil
}

func (m *ServerReflectionResponse) GetErrorResponse() *ErrorResponse {
	if x, ok := m.GetMessageResponse().(*ServerReflectionResponse_ErrorResponse); ok {
		return x.ErrorResponse
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ServerReflectionResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ServerReflectionResponse_OneofMarshaler, _ServerReflectionResponse_OneofUnmarshaler, _ServerReflectionResponse_OneofSizer, []interface{}{
		(*ServerReflectionResponse_FileDescriptorResponse)(nil),
		(*ServerReflectionResponse_AllExtensionNumbersResponse)(nil),
		(*ServerReflectionResponse_ListServicesResponse)(nil),
		(*ServerReflectionResponse_ErrorResponse)(nil),
	}
}

func _ServerReflectionResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ServerReflectionResponse)
	// message_response
	switch x := m.MessageResponse.(type) {
	case *ServerReflectionResponse_FileDescriptorResponse:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FileDescriptorResponse); err != nil {
			return err
		}
	case *ServerReflectionResponse_AllExtensionNumbersResponse:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AllExtensionNumbersResponse); err != nil {
			return err
		}
	case *ServerReflectionResponse_ListServicesResponse:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListServicesResponse); err != nil {
			return err
		}
	case *ServerReflectionResponse_ErrorResponse:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ErrorResponse); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ServerReflectionResponse.MessageResponse has unexpected type %T", x)
	}
	return nil
}

func _ServerReflectionResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ServerReflectionResponse)
	switch tag {
	case 4: // message_response.file_descriptor_response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FileDescriptorResponse)
		err := b.DecodeMessage(msg)
		m.MessageResponse = &ServerReflectionResponse_FileDescriptorResponse{msg}
		return true, err
	case 5: // message_response.all_extension_numbers_response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ExtensionNumberResponse)
		err := b.DecodeMessage(msg)
		m.MessageResponse = &ServerReflectionResponse_AllExtensionNumbersResponse{msg}
		return true, err
	case 6: // message_response.list_services_response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ListServiceResponse)
		err := b.DecodeMessage(msg)
		m.MessageResponse = &ServerReflectionResponse_ListServicesResponse{msg}
		return true, err
	case 7: // message_response.error_response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ErrorResponse)
		err := b.DecodeMessage(msg)
		m.MessageResponse = &ServerReflectionResponse_ErrorResponse{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ServerReflectionResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ServerReflectionResponse)
	// message_response
	switch x := m.MessageResponse.(type) {
	case *ServerReflectionResponse_FileDescriptorResponse:
		s := proto.Size(x.FileDescriptorResponse)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ServerReflectionResponse_AllExtensionNumbersResponse:
		s := proto.Size(x.AllExtensionNumbersResponse)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ServerReflectionResponse_ListServicesResponse:
		s := proto.Size(x.ListServicesResponse)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ServerReflectionResponse_ErrorResponse:
		s := proto.Size(x.ErrorResponse)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Serialized FileDescriptorProto messages sent by the server answering
// a file_by_filename, file_containing_symbol, or file_containing_extension
// request.
type FileDescriptorResponse struct {
	// Serialized FileDescriptorProto messages. We avoid taking a dependency on
	// descriptor.proto, which uses proto2 only features, by making them opaque
	// bytes instead.
	FileDescriptorProto [][]byte `protobuf:"bytes,1,rep,name=file_descriptor_proto,json=fileDescriptorProto,proto3" json:"file_descriptor_proto,omitempty"`
}

func (m *FileDescriptorResponse) Reset()                    { *m = FileDescriptorResponse{} }
func (m *FileDescriptorResponse) String() string            { return proto.CompactTextString(m) }
func (*FileDescriptorResponse) ProtoMessage()               {}
func (*FileDescriptorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FileDescriptorResponse) GetFileDescriptorProto() [][]byte {
	if m != nil {
		return m.FileDescriptorProto
	}
	return nil
}

// A list of extension numbers sent by the server answering
// all_extension_numbers_of_type request.
type ExtensionNumberResponse struct {
	// Full name of the base type, including the package name. The format
	// is <package>.<type>
	BaseTypeName    string  `protobuf:"bytes,1,opt,name=base_type_name,json=baseTypeName" json:"base_type_name,omitempty"`
	ExtensionNumber []int32 `protobuf:"varint,2,rep,packed,name=extension_number,json=extensionNumber" json:"extension_number,omitempty"`
}

func (m *ExtensionNumberResponse) Reset()                    { *m = ExtensionNumberResponse{} }
func (m *ExtensionNumberResponse) String() string            { return proto.CompactTextString(m) }
func (*ExtensionNumberResponse) ProtoMessage()               {}
func (*ExtensionNumberResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ExtensionNumberResponse) GetBaseTypeName() string {
	if m != nil {
		return m.BaseTypeName
	}
	return ""
}

func (m *ExtensionNumberResponse) GetExtensionNumber() []int32 {
	if m != nil {
		return m.ExtensionNumber
	}
	return nil
}

// A list of ServiceResponse sent by the server answering list_services request.
type ListServiceResponse struct {
	// The information of each service may be expanded in the future, so we use
	// ServiceResponse message to encapsulate it.
	Service []*ServiceResponse `protobuf:"bytes,1,rep,name=service" json:"service,omitempty"`
}

func (m *ListServiceResponse) Reset()                    { *m = ListServiceResponse{} }
func (m *ListServiceResponse) String() string            { return proto.CompactTextString(m) }
func (*ListServiceResponse) ProtoMessage()               {}
func (*ListServiceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListServiceResponse) GetService() []*ServiceResponse {
	if m != nil {
		return m.Service
	}
	return nil
}

// The information of a single service used by ListServiceResponse to answer
// list_services request.
type ServiceResponse struct {
	// Full name of a registered service, including its package name. The format
	// is <package>.<service>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ServiceResponse) Reset()                    { *m = ServiceResponse{} }
func (m *ServiceResponse) String() string            { return proto.CompactTextString(m) }
func (*ServiceResponse) ProtoMessage()               {}
func (*ServiceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ServiceResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The error code and error message sent by the server when an error occurs.
type ErrorResponse struct {
	// This field uses the error codes defined in grpc::StatusCode.
	ErrorCode    int32  `protobuf:"varint,1,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage" json:"error_message,omitempty"`
}

func (m *ErrorResponse) Reset()                    { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string            { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()               {}
func (*ErrorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ErrorResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *ErrorResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*ServerReflectionRequest)(nil), "grpc.reflection.v1alpha.ServerReflectionRequest")
	proto.RegisterType((*ExtensionRequest)(nil), "grpc.reflection.v1alpha.ExtensionRequest")
	proto.RegisterType((*ServerReflectionResponse)(nil), "grpc.reflection.v1alpha.ServerReflectionResponse")
	proto.RegisterType((*FileDescriptorResponse)(nil), "grpc.reflection.v1alpha.FileDescriptorResponse")
	proto.RegisterType((*ExtensionNumberResponse)(nil), "grpc.reflection.v1alpha.ExtensionNumberResponse")
	proto.RegisterType((*ListServiceResponse)(nil), "grpc.reflection.v1alpha.ListServiceResponse")
	proto.RegisterType((*ServiceResponse)(nil), "grpc.reflection.v1alpha.ServiceResponse")
	proto.RegisterType((*ErrorResponse)(nil), "grpc.reflection.v1alpha.ErrorResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ServerReflection service

type ServerReflectionClient interface {
	// The reflection service is structured as a bidirectional stream, ensuring
	// all related requests go to a single server.
	ServerReflectionInfo(ctx context.Context, opts ...grpc.CallOption) (ServerReflection_ServerReflectionInfoClient, error)
}

type serverReflectionClient struct {
	cc *grpc.ClientConn
}

func NewServerReflectionClient(cc *grpc.ClientConn) ServerReflectionClient {
	return &serverReflectionClient{cc}
}

func (c *serverReflectionClient) ServerReflectionInfo(ctx context.Context, opts ...grpc.CallOption) (ServerReflection_ServerReflectionInfoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ServerReflection_serviceDesc.Streams[0], c.cc, "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverReflectionServerReflectionInfoClient{stream}
	return x, nil
}

type ServerReflection_ServerReflectionInfoClient interface {
	Send(*ServerReflectionRequest) error
	Recv() (*ServerReflectionResponse, error)
	grpc.ClientStream
}

type serverReflectionServerReflectionInfoClient struct {
	grpc.ClientStream
}

func (x *serverReflectionServerReflectionInfoClient) Send(m *ServerReflectionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serverReflectionServerReflectionInfoClient) Recv() (*ServerReflectionResponse, error) {
	m := new(ServerReflectionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ServerReflection service

type ServerReflectionServer interface {
	// The reflection service is structured as a bidirectional stream, ensuring
	// all related requests go to a single server.
	ServerReflectionInfo(ServerReflection_ServerReflectionInfoServer) error
}

func RegisterServerReflectionServer(s *grpc.Server, srv ServerReflectionServer) {
	s.RegisterService(&_ServerReflection_serviceDesc, srv)
}

func _ServerReflection_ServerReflectionInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServerReflectionServer).ServerReflectionInfo(&serverReflectionServerReflectionInfoServer{stream})
}

type ServerReflection_ServerReflectionInfoServer interface {
	Send(*ServerReflectionResponse) error
	Recv() (*ServerReflectionRequest, error)
	grpc.ServerStream
}

type serverReflectionServerReflectionInfoServer struct {
	grpc.ServerStream
}

func (x *serverReflectionServerReflectionInfoServer) Send(m *ServerReflectionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serverReflectionServerReflectionInfoServer) Recv() (*ServerReflectionRequest, error) {
	m := new(ServerReflectionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ServerReflection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.reflection.v1alpha.ServerReflection",
	HandlerType: (*ServerReflectionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerReflectionInfo",
			Handler:       _ServerReflection_ServerReflectionInfo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "reflection.proto",
}

func init() { proto.RegisterFile("reflection.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xd1, 0x6e, 0xd3, 0x4a,
	0x10, 0xad, 0xdb, 0xa4, 0x55, 0x26, 0x69, 0xe2, 0xbb, 0xed, 0x6d, 0x5c, 0x50, 0x51, 0x64, 0x28,
	0xa4, 0x08, 0x85, 0x36, 0x48, 0x7c, 0x40, 0x0a, 0x28, 0x48, 0xa5, 0x45, 0x1b, 0x5e, 0x10, 0x0f,
	0x2b, 0x27, 0x99, 0xa4, 0x06, 0xc7, 0x6b, 0x76, 0xdd, 0x40, 0x9e, 0xf8, 0x08, 0x3e, 0x8a, 0x5f,
	0xe2, 0x11, 0xed, 0xda, 0x71, 0x1c, 0xd7, 0x06, 0xf5, 0x29, 0xd1, 0x99, 0x99, 0x3d, 0x33, 0x73,
	0xce, 0x18, 0x4c, 0x81, 0x13, 0x0f, 0x47, 0xa1, 0xcb, 0xfd, 0x4e, 0x20, 0x78, 0xc8, 0x49, 0x73,
	0x2a, 0x82, 0x51, 0x27, 0x05, 0xcf, 0xcf, 0x1c, 0x2f, 0xb8, 0x76, 0xec, 0xdf, 0x9b, 0xd0, 0x1c,
	0xa0, 0x98, 0xa3, 0xa0, 0x49, 0x90, 0xe2, 0xd7, 0x1b, 0x94, 0x21, 0x21, 0x50, 0xba, 0xe6, 0x32,
	0xb4, 0x8c, 0x96, 0xd1, 0xae, 0x50, 0xfd, 0x9f, 0x3c, 0x05, 0x73, 0xe2, 0x7a, 0xc8, 0x86, 0x0b,
	0xa6, 0x7e, 0x7d, 0x67, 0x86, 0xd6, 0x96, 0x8a, 0xf7, 0x37, 0x68, 0x5d, 0x21, 0xbd, 0xc5, 0x9b,
	0x18, 0x27, 0x2f, 0xe1, 0x40, 0xe7, 0x8e, 0xb8, 0x1f, 0x3a, 0xae, 0xef, 0xfa, 0x53, 0x26, 0x17,
	0xb3, 0x21, 0xf7, 0xac, 0x52, 0x5c, 0xb1, 0xaf, 0xe2, 0xe7, 0x49, 0x78, 0xa0, 0xa3, 0x64, 0x0a,
	0x87, 0xd9, 0x3a, 0xfc, 0x1e, 0xa2, 0x2f, 0x5d, 0xee, 0x5b, 0xe5, 0x96, 0xd1, 0xae, 0x76, 0x4f,
	0x3a, 0x05, 0x03, 0x75, 0x5e, 0x2f, 0x33, 0xe3, 0x29, 0xfa, 0x1b, 0xb4, 0xb9, 0xce, 0x92, 0x64,
	0x90, 0x1e, 0x1c, 0x39, 0x9e, 0xb7, 0x7a, 0x9c, 0xf9, 0x37, 0xb3, 0x21, 0x0a, 0xc9, 0xf8, 0x84,
	0x85, 0x8b, 0x00, 0xad, 0xed, 0xb8, 0xcf, 0x43, 0xc7, 0xf3, 0x92, 0xb2, 0xcb, 0x28, 0xe9, 0x6a,
	0xf2, 0x61, 0x11, 0x20, 0x39, 0x86, 0x5d, 0xcf, 0x95, 0x21, 0x93, 0x28, 0xe6, 0xee, 0x08, 0xa5,
	0xb5, 0x13, 0xd7, 0xd4, 0x14, 0x3c, 0x88, 0xd1, 0xde, 0x7f, 0xd0, 0x98, 0xa1, 0x94, 0xce, 0x14,
	0x99, 0x88, 0x1a, 0xb3, 0x27, 0x60, 0x66, 0x9b, 0x25, 0x4f, 0xa0, 0x91, 0x9a, 0x5a, 0xf7, 0x10,
	0x6d, 0xbf, 0xbe, 0x82, 0x35, 0xed, 0x09, 0x98, 0xd9, 0xb6, 0xad, 0xcd, 0x96, 0xd1, 0x2e, 0xd3,
	0x06, 0xae, 0x37, 0x6a, 0xff, 0x2a, 0x81, 0x75, 0x5b, 0x62, 0x19, 0x70, 0x5f, 0x22, 0x39, 0x02,
	0x98, 0x3b, 0x9e, 0x3b, 0x66, 0x29, 0xa5, 0x2b, 0x1a, 0xe9, 0x2b, 0xb9, 0x3f, 0x81, 0xc9, 0x85,
	0x3b, 0x75, 0x7d, 0xc7, 0x5b, 0xf6, 0xad, 0x69, 0xaa, 0xdd, 0xd3, 0x42, 0x05, 0x0a, 0xec, 0x44,
	0x1b, 0xcb, 0x97, 0x96, 0xc3, 0x7e, 0x01, 0x4b, 0xeb, 0x3c, 0x46, 0x39, 0x12, 0x6e, 0x10, 0x72,
	0xc1, 0x44, 0xdc, 0x97, 0x76, 0x48, 0xb5, 0xfb, 0xbc, 0x90, 0x44, 0x99, 0xec, 0x55, 0x52, 0xb7,
	0x1c, 0xa7, 0xbf, 0x41, 0xb5, 0xe5, 0x6e, 0x47, 0xc8, 0x37, 0x78, 0x90, 0xaf, 0x75, 0x42, 0x59,
	0xfe, 0xc7, 0x5c, 0x19, 0x03, 0xa4, 0x38, 0xef, 0xe7, 0xd8, 0x23, 0x21, 0x1e, 0xc3, 0xc1, 0x9a,
	0x41, 0x56, 0x84, 0xdb, 0x9a, 0xf0, 0x59, 0x21, 0xe1, 0xc5, 0xca, 0x40, 0x29, 0xb2, 0xfd, 0xb4,
	0xaf, 0x12, 0x96, 0x2b, 0xa8, 0xa3, 0x10, 0xe9, 0x0d, 0xee, 0xe8, 0xd7, 0x1f, 0x17, 0x8f, 0xa3,
	0xd2, 0x53, 0xef, 0xee, 0x62, 0x1a, 0xe8, 0x11, 0x30, 0x57, 0x86, 0x8d, 0x30, 0xfb, 0x02, 0x0e,
	0xf2, 0xf7, 0x4e, 0xba, 0xf0, 0x7f, 0x56, 0x4a, 0xfd, 0xe1, 0xb1, 0x8c, 0xd6, 0x56, 0xbb, 0x46,
	0xf7, 0xd6, 0x45, 0x79, 0xaf, 0x42, 0xf6, 0x67, 0x68, 0x16, 0xac, 0x94, 0x3c, 0x82, 0xfa, 0xd0,
	0x91, 0xa8, 0x0f, 0x80, 0xe9, 0x6f, 0x4c, 0xe4, 0xcc, 0x9a, 0x42, 0x95, 0xff, 0x2f, 0xd5, 0xf7,
	0x25, 0xff, 0x06, 0xb6, 0xf2, 0x6e, 0xe0, 0x23, 0xec, 0xe5, 0x6c, 0x93, 0xf4, 0x60, 0x27, 0x96,
	0x45, 0x37, 0x5a, 0xed, 0xb6, 0xff, 0xea, 0xea, 0x54, 0x29, 0x5d, 0x16, 0xda, 0xc7, 0xd0, 0xc8,
	0x3e, 0x4b, 0xa0, 0x94, 0x6a, 0x5a, 0xff, 0xb7, 0x07, 0xb0, 0xbb, 0xb6, 0x71, 0x75, 0x79, 0x91,
	0x62, 0x23, 0x3e, 0x8e, 0x52, 0xcb, 0xb4, 0xa2, 0x91, 0x73, 0x3e, 0x46, 0xf2, 0x10, 0x22, 0x41,
	0x58, 0xac, 0x82, 0x3e, 0xbb, 0x0a, 0xad, 0x69, 0xf0, 0x5d, 0x84, 0x75, 0x7f, 0x1a, 0x60, 0x66,
	0xcf, 0x8d, 0xfc, 0x80, 0xfd, 0x2c, 0xf6, 0xd6, 0x9f, 0x70, 0x72, 0xe7, 0x8b, 0xbd, 0x77, 0x76,
	0x87, 0x8a, 0x68, 0xaa, 0xb6, 0x71, 0x6a, 0x0c, 0xb7, 0xb5, 0xf4, 0x2f, 0xfe, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xe9, 0x3f, 0x7b, 0x08, 0x87, 0x06, 0x00, 0x00,
}
