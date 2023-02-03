// Copyright 2016 The gRPC Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Service exported by server reflection.  A more complete description of how
// server reflection works can be found at
// https://github.com/grpc/grpc/blob/master/doc/server-reflection.md
//
// The canonical version of this proto can be found at
// https://github.com/grpc/grpc-proto/blob/master/grpc/reflection/v1/reflection.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: grpc/reflection/v1/reflection.proto

package grpc_reflection_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The message sent by the client when calling ServerReflectionInfo method.
type ServerReflectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// To use reflection service, the client should set one of the following
	// fields in message_request. The server distinguishes requests by their
	// defined field and then handles them using corresponding methods.
	//
	// Types that are assignable to MessageRequest:
	//
	//	*ServerReflectionRequest_FileByFilename
	//	*ServerReflectionRequest_FileContainingSymbol
	//	*ServerReflectionRequest_FileContainingExtension
	//	*ServerReflectionRequest_AllExtensionNumbersOfType
	//	*ServerReflectionRequest_ListServices
	MessageRequest isServerReflectionRequest_MessageRequest `protobuf_oneof:"message_request"`
}

func (x *ServerReflectionRequest) Reset() {
	*x = ServerReflectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerReflectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerReflectionRequest) ProtoMessage() {}

func (x *ServerReflectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerReflectionRequest.ProtoReflect.Descriptor instead.
func (*ServerReflectionRequest) Descriptor() ([]byte, []int) {
	return file_grpc_reflection_v1_reflection_proto_rawDescGZIP(), []int{0}
}

func (x *ServerReflectionRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (m *ServerReflectionRequest) GetMessageRequest() isServerReflectionRequest_MessageRequest {
	if m != nil {
		return m.MessageRequest
	}
	return nil
}

func (x *ServerReflectionRequest) GetFileByFilename() string {
	if x, ok := x.GetMessageRequest().(*ServerReflectionRequest_FileByFilename); ok {
		return x.FileByFilename
	}
	return ""
}

func (x *ServerReflectionRequest) GetFileContainingSymbol() string {
	if x, ok := x.GetMessageRequest().(*ServerReflectionRequest_FileContainingSymbol); ok {
		return x.FileContainingSymbol
	}
	return ""
}

func (x *ServerReflectionRequest) GetFileContainingExtension() *ExtensionRequest {
	if x, ok := x.GetMessageRequest().(*ServerReflectionRequest_FileContainingExtension); ok {
		return x.FileContainingExtension
	}
	return nil
}

func (x *ServerReflectionRequest) GetAllExtensionNumbersOfType() string {
	if x, ok := x.GetMessageRequest().(*ServerReflectionRequest_AllExtensionNumbersOfType); ok {
		return x.AllExtensionNumbersOfType
	}
	return ""
}

func (x *ServerReflectionRequest) GetListServices() string {
	if x, ok := x.GetMessageRequest().(*ServerReflectionRequest_ListServices); ok {
		return x.ListServices
	}
	return ""
}

type isServerReflectionRequest_MessageRequest interface {
	isServerReflectionRequest_MessageRequest()
}

type ServerReflectionRequest_FileByFilename struct {
	// Find a proto file by the file name.
	FileByFilename string `protobuf:"bytes,3,opt,name=file_by_filename,json=fileByFilename,proto3,oneof"`
}

type ServerReflectionRequest_FileContainingSymbol struct {
	// Find the proto file that declares the given fully-qualified symbol name.
	// This field should be a fully-qualified symbol name
	// (e.g. <package>.<service>[.<method>] or <package>.<type>).
	FileContainingSymbol string `protobuf:"bytes,4,opt,name=file_containing_symbol,json=fileContainingSymbol,proto3,oneof"`
}

type ServerReflectionRequest_FileContainingExtension struct {
	// Find the proto file which defines an extension extending the given
	// message type with the given field number.
	FileContainingExtension *ExtensionRequest `protobuf:"bytes,5,opt,name=file_containing_extension,json=fileContainingExtension,proto3,oneof"`
}

type ServerReflectionRequest_AllExtensionNumbersOfType struct {
	// Finds the tag numbers used by all known extensions of the given message
	// type, and appends them to ExtensionNumberResponse in an undefined order.
	// Its corresponding method is best-effort: it's not guaranteed that the
	// reflection service will implement this method, and it's not guaranteed
	// that this method will provide all extensions. Returns
	// StatusCode::UNIMPLEMENTED if it's not implemented.
	// This field should be a fully-qualified type name. The format is
	// <package>.<type>
	AllExtensionNumbersOfType string `protobuf:"bytes,6,opt,name=all_extension_numbers_of_type,json=allExtensionNumbersOfType,proto3,oneof"`
}

type ServerReflectionRequest_ListServices struct {
	// List the full names of registered services. The content will not be
	// checked.
	ListServices string `protobuf:"bytes,7,opt,name=list_services,json=listServices,proto3,oneof"`
}

func (*ServerReflectionRequest_FileByFilename) isServerReflectionRequest_MessageRequest() {}

func (*ServerReflectionRequest_FileContainingSymbol) isServerReflectionRequest_MessageRequest() {}

func (*ServerReflectionRequest_FileContainingExtension) isServerReflectionRequest_MessageRequest() {}

func (*ServerReflectionRequest_AllExtensionNumbersOfType) isServerReflectionRequest_MessageRequest() {
}

func (*ServerReflectionRequest_ListServices) isServerReflectionRequest_MessageRequest() {}

// The type name and extension number sent by the client when requesting
// file_containing_extension.
type ExtensionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Fully-qualified type name. The format should be <package>.<type>
	ContainingType  string `protobuf:"bytes,1,opt,name=containing_type,json=containingType,proto3" json:"containing_type,omitempty"`
	ExtensionNumber int32  `protobuf:"varint,2,opt,name=extension_number,json=extensionNumber,proto3" json:"extension_number,omitempty"`
}

func (x *ExtensionRequest) Reset() {
	*x = ExtensionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionRequest) ProtoMessage() {}

func (x *ExtensionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionRequest.ProtoReflect.Descriptor instead.
func (*ExtensionRequest) Descriptor() ([]byte, []int) {
	return file_grpc_reflection_v1_reflection_proto_rawDescGZIP(), []int{1}
}

func (x *ExtensionRequest) GetContainingType() string {
	if x != nil {
		return x.ContainingType
	}
	return ""
}

func (x *ExtensionRequest) GetExtensionNumber() int32 {
	if x != nil {
		return x.ExtensionNumber
	}
	return 0
}

// The message sent by the server to answer ServerReflectionInfo method.
type ServerReflectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidHost       string                   `protobuf:"bytes,1,opt,name=valid_host,json=validHost,proto3" json:"valid_host,omitempty"`
	OriginalRequest *ServerReflectionRequest `protobuf:"bytes,2,opt,name=original_request,json=originalRequest,proto3" json:"original_request,omitempty"`
	// The server sets one of the following fields according to the message_request
	// in the request.
	//
	// Types that are assignable to MessageResponse:
	//
	//	*ServerReflectionResponse_FileDescriptorResponse
	//	*ServerReflectionResponse_AllExtensionNumbersResponse
	//	*ServerReflectionResponse_ListServicesResponse
	//	*ServerReflectionResponse_ErrorResponse
	MessageResponse isServerReflectionResponse_MessageResponse `protobuf_oneof:"message_response"`
}

func (x *ServerReflectionResponse) Reset() {
	*x = ServerReflectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerReflectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerReflectionResponse) ProtoMessage() {}

func (x *ServerReflectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerReflectionResponse.ProtoReflect.Descriptor instead.
func (*ServerReflectionResponse) Descriptor() ([]byte, []int) {
	return file_grpc_reflection_v1_reflection_proto_rawDescGZIP(), []int{2}
}

func (x *ServerReflectionResponse) GetValidHost() string {
	if x != nil {
		return x.ValidHost
	}
	return ""
}

func (x *ServerReflectionResponse) GetOriginalRequest() *ServerReflectionRequest {
	if x != nil {
		return x.OriginalRequest
	}
	return nil
}

func (m *ServerReflectionResponse) GetMessageResponse() isServerReflectionResponse_MessageResponse {
	if m != nil {
		return m.MessageResponse
	}
	return nil
}

func (x *ServerReflectionResponse) GetFileDescriptorResponse() *FileDescriptorResponse {
	if x, ok := x.GetMessageResponse().(*ServerReflectionResponse_FileDescriptorResponse); ok {
		return x.FileDescriptorResponse
	}
	return nil
}

func (x *ServerReflectionResponse) GetAllExtensionNumbersResponse() *ExtensionNumberResponse {
	if x, ok := x.GetMessageResponse().(*ServerReflectionResponse_AllExtensionNumbersResponse); ok {
		return x.AllExtensionNumbersResponse
	}
	return nil
}

func (x *ServerReflectionResponse) GetListServicesResponse() *ListServiceResponse {
	if x, ok := x.GetMessageResponse().(*ServerReflectionResponse_ListServicesResponse); ok {
		return x.ListServicesResponse
	}
	return nil
}

func (x *ServerReflectionResponse) GetErrorResponse() *ErrorResponse {
	if x, ok := x.GetMessageResponse().(*ServerReflectionResponse_ErrorResponse); ok {
		return x.ErrorResponse
	}
	return nil
}

type isServerReflectionResponse_MessageResponse interface {
	isServerReflectionResponse_MessageResponse()
}

type ServerReflectionResponse_FileDescriptorResponse struct {
	// This message is used to answer file_by_filename, file_containing_symbol,
	// file_containing_extension requests with transitive dependencies.
	// As the repeated label is not allowed in oneof fields, we use a
	// FileDescriptorResponse message to encapsulate the repeated fields.
	// The reflection service is allowed to avoid sending FileDescriptorProtos
	// that were previously sent in response to earlier requests in the stream.
	FileDescriptorResponse *FileDescriptorResponse `protobuf:"bytes,4,opt,name=file_descriptor_response,json=fileDescriptorResponse,proto3,oneof"`
}

type ServerReflectionResponse_AllExtensionNumbersResponse struct {
	// This message is used to answer all_extension_numbers_of_type requests.
	AllExtensionNumbersResponse *ExtensionNumberResponse `protobuf:"bytes,5,opt,name=all_extension_numbers_response,json=allExtensionNumbersResponse,proto3,oneof"`
}

type ServerReflectionResponse_ListServicesResponse struct {
	// This message is used to answer list_services requests.
	ListServicesResponse *ListServiceResponse `protobuf:"bytes,6,opt,name=list_services_response,json=listServicesResponse,proto3,oneof"`
}

type ServerReflectionResponse_ErrorResponse struct {
	// This message is used when an error occurs.
	ErrorResponse *ErrorResponse `protobuf:"bytes,7,opt,name=error_response,json=errorResponse,proto3,oneof"`
}

func (*ServerReflectionResponse_FileDescriptorResponse) isServerReflectionResponse_MessageResponse() {
}

func (*ServerReflectionResponse_AllExtensionNumbersResponse) isServerReflectionResponse_MessageResponse() {
}

func (*ServerReflectionResponse_ListServicesResponse) isServerReflectionResponse_MessageResponse() {}

func (*ServerReflectionResponse_ErrorResponse) isServerReflectionResponse_MessageResponse() {}

// Serialized FileDescriptorProto messages sent by the server answering
// a file_by_filename, file_containing_symbol, or file_containing_extension
// request.
type FileDescriptorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Serialized FileDescriptorProto messages. We avoid taking a dependency on
	// descriptor.proto, which uses proto2 only features, by making them opaque
	// bytes instead.
	FileDescriptorProto [][]byte `protobuf:"bytes,1,rep,name=file_descriptor_proto,json=fileDescriptorProto,proto3" json:"file_descriptor_proto,omitempty"`
}

func (x *FileDescriptorResponse) Reset() {
	*x = FileDescriptorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDescriptorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDescriptorResponse) ProtoMessage() {}

func (x *FileDescriptorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDescriptorResponse.ProtoReflect.Descriptor instead.
func (*FileDescriptorResponse) Descriptor() ([]byte, []int) {
	return file_grpc_reflection_v1_reflection_proto_rawDescGZIP(), []int{3}
}

func (x *FileDescriptorResponse) GetFileDescriptorProto() [][]byte {
	if x != nil {
		return x.FileDescriptorProto
	}
	return nil
}

// A list of extension numbers sent by the server answering
// all_extension_numbers_of_type request.
type ExtensionNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full name of the base type, including the package name. The format
	// is <package>.<type>
	BaseTypeName    string  `protobuf:"bytes,1,opt,name=base_type_name,json=baseTypeName,proto3" json:"base_type_name,omitempty"`
	ExtensionNumber []int32 `protobuf:"varint,2,rep,packed,name=extension_number,json=extensionNumber,proto3" json:"extension_number,omitempty"`
}

func (x *ExtensionNumberResponse) Reset() {
	*x = ExtensionNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionNumberResponse) ProtoMessage() {}

func (x *ExtensionNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionNumberResponse.ProtoReflect.Descriptor instead.
func (*ExtensionNumberResponse) Descriptor() ([]byte, []int) {
	return file_grpc_reflection_v1_reflection_proto_rawDescGZIP(), []int{4}
}

func (x *ExtensionNumberResponse) GetBaseTypeName() string {
	if x != nil {
		return x.BaseTypeName
	}
	return ""
}

func (x *ExtensionNumberResponse) GetExtensionNumber() []int32 {
	if x != nil {
		return x.ExtensionNumber
	}
	return nil
}

// A list of ServiceResponse sent by the server answering list_services request.
type ListServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The information of each service may be expanded in the future, so we use
	// ServiceResponse message to encapsulate it.
	Service []*ServiceResponse `protobuf:"bytes,1,rep,name=service,proto3" json:"service,omitempty"`
}

func (x *ListServiceResponse) Reset() {
	*x = ListServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceResponse) ProtoMessage() {}

func (x *ListServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceResponse.ProtoReflect.Descriptor instead.
func (*ListServiceResponse) Descriptor() ([]byte, []int) {
	return file_grpc_reflection_v1_reflection_proto_rawDescGZIP(), []int{5}
}

func (x *ListServiceResponse) GetService() []*ServiceResponse {
	if x != nil {
		return x.Service
	}
	return nil
}

// The information of a single service used by ListServiceResponse to answer
// list_services request.
type ServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full name of a registered service, including its package name. The format
	// is <package>.<service>
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ServiceResponse) Reset() {
	*x = ServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceResponse) ProtoMessage() {}

func (x *ServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceResponse.ProtoReflect.Descriptor instead.
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return file_grpc_reflection_v1_reflection_proto_rawDescGZIP(), []int{6}
}

func (x *ServiceResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The error code and error message sent by the server when an error occurs.
type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field uses the error codes defined in grpc::StatusCode.
	ErrorCode    int32  `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_reflection_v1_reflection_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_grpc_reflection_v1_reflection_proto_rawDescGZIP(), []int{7}
}

func (x *ErrorResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *ErrorResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_grpc_reflection_v1_reflection_proto protoreflect.FileDescriptor

var file_grpc_reflection_v1_reflection_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x72, 0x65, 0x66, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xf3, 0x02, 0x0a, 0x17, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x62, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x46, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x14, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x62, 0x0a,
	0x19, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x17, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x42, 0x0a, 0x1d, 0x61, 0x6c, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x19, 0x61, 0x6c, 0x6c, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x4f,
	0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c,
	0x6c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x11, 0x0a, 0x0f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x66, 0x0a, 0x10, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x10,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xae, 0x04, 0x0a, 0x18, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x48,
	0x6f, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0f, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x66, 0x0a, 0x18, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x16, 0x66, 0x69, 0x6c,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x1e, 0x61, 0x6c, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x1b, 0x61, 0x6c, 0x6c, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x16, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x72,
	0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x00, 0x52, 0x14, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x0a, 0x16, 0x46, 0x69, 0x6c, 0x65,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x13, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x17, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x54, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x53, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x89, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x75, 0x0a, 0x14, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x66,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x66, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x66, 0x0a, 0x15, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x72, 0x65, 0x66, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65, 0x66, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_reflection_v1_reflection_proto_rawDescOnce sync.Once
	file_grpc_reflection_v1_reflection_proto_rawDescData = file_grpc_reflection_v1_reflection_proto_rawDesc
)

func file_grpc_reflection_v1_reflection_proto_rawDescGZIP() []byte {
	file_grpc_reflection_v1_reflection_proto_rawDescOnce.Do(func() {
		file_grpc_reflection_v1_reflection_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_reflection_v1_reflection_proto_rawDescData)
	})
	return file_grpc_reflection_v1_reflection_proto_rawDescData
}

var file_grpc_reflection_v1_reflection_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_reflection_v1_reflection_proto_goTypes = []interface{}{
	(*ServerReflectionRequest)(nil),  // 0: grpc.reflection.v1.ServerReflectionRequest
	(*ExtensionRequest)(nil),         // 1: grpc.reflection.v1.ExtensionRequest
	(*ServerReflectionResponse)(nil), // 2: grpc.reflection.v1.ServerReflectionResponse
	(*FileDescriptorResponse)(nil),   // 3: grpc.reflection.v1.FileDescriptorResponse
	(*ExtensionNumberResponse)(nil),  // 4: grpc.reflection.v1.ExtensionNumberResponse
	(*ListServiceResponse)(nil),      // 5: grpc.reflection.v1.ListServiceResponse
	(*ServiceResponse)(nil),          // 6: grpc.reflection.v1.ServiceResponse
	(*ErrorResponse)(nil),            // 7: grpc.reflection.v1.ErrorResponse
}
var file_grpc_reflection_v1_reflection_proto_depIdxs = []int32{
	1, // 0: grpc.reflection.v1.ServerReflectionRequest.file_containing_extension:type_name -> grpc.reflection.v1.ExtensionRequest
	0, // 1: grpc.reflection.v1.ServerReflectionResponse.original_request:type_name -> grpc.reflection.v1.ServerReflectionRequest
	3, // 2: grpc.reflection.v1.ServerReflectionResponse.file_descriptor_response:type_name -> grpc.reflection.v1.FileDescriptorResponse
	4, // 3: grpc.reflection.v1.ServerReflectionResponse.all_extension_numbers_response:type_name -> grpc.reflection.v1.ExtensionNumberResponse
	5, // 4: grpc.reflection.v1.ServerReflectionResponse.list_services_response:type_name -> grpc.reflection.v1.ListServiceResponse
	7, // 5: grpc.reflection.v1.ServerReflectionResponse.error_response:type_name -> grpc.reflection.v1.ErrorResponse
	6, // 6: grpc.reflection.v1.ListServiceResponse.service:type_name -> grpc.reflection.v1.ServiceResponse
	0, // 7: grpc.reflection.v1.ServerReflection.ServerReflectionInfo:input_type -> grpc.reflection.v1.ServerReflectionRequest
	2, // 8: grpc.reflection.v1.ServerReflection.ServerReflectionInfo:output_type -> grpc.reflection.v1.ServerReflectionResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_grpc_reflection_v1_reflection_proto_init() }
func file_grpc_reflection_v1_reflection_proto_init() {
	if File_grpc_reflection_v1_reflection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_reflection_v1_reflection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerReflectionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_reflection_v1_reflection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_reflection_v1_reflection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerReflectionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_reflection_v1_reflection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDescriptorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_reflection_v1_reflection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionNumberResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_reflection_v1_reflection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServiceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_reflection_v1_reflection_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_reflection_v1_reflection_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_grpc_reflection_v1_reflection_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ServerReflectionRequest_FileByFilename)(nil),
		(*ServerReflectionRequest_FileContainingSymbol)(nil),
		(*ServerReflectionRequest_FileContainingExtension)(nil),
		(*ServerReflectionRequest_AllExtensionNumbersOfType)(nil),
		(*ServerReflectionRequest_ListServices)(nil),
	}
	file_grpc_reflection_v1_reflection_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ServerReflectionResponse_FileDescriptorResponse)(nil),
		(*ServerReflectionResponse_AllExtensionNumbersResponse)(nil),
		(*ServerReflectionResponse_ListServicesResponse)(nil),
		(*ServerReflectionResponse_ErrorResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_reflection_v1_reflection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_reflection_v1_reflection_proto_goTypes,
		DependencyIndexes: file_grpc_reflection_v1_reflection_proto_depIdxs,
		MessageInfos:      file_grpc_reflection_v1_reflection_proto_msgTypes,
	}.Build()
	File_grpc_reflection_v1_reflection_proto = out.File
	file_grpc_reflection_v1_reflection_proto_rawDesc = nil
	file_grpc_reflection_v1_reflection_proto_goTypes = nil
	file_grpc_reflection_v1_reflection_proto_depIdxs = nil
}
