// Copyright 2017 gRPC authors.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: reflection/grpc_testing/proto2_ext.proto

package grpc_testing

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

type Extension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Whatzit *int32 `protobuf:"varint,1,opt,name=whatzit" json:"whatzit,omitempty"`
}

func (x *Extension) Reset() {
	*x = Extension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reflection_grpc_testing_proto2_ext_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extension) ProtoMessage() {}

func (x *Extension) ProtoReflect() protoreflect.Message {
	mi := &file_reflection_grpc_testing_proto2_ext_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extension.ProtoReflect.Descriptor instead.
func (*Extension) Descriptor() ([]byte, []int) {
	return file_reflection_grpc_testing_proto2_ext_proto_rawDescGZIP(), []int{0}
}

func (x *Extension) GetWhatzit() int32 {
	if x != nil && x.Whatzit != nil {
		return *x.Whatzit
	}
	return 0
}

var file_reflection_grpc_testing_proto2_ext_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*ToBeExtended)(nil),
		ExtensionType: (*int32)(nil),
		Field:         13,
		Name:          "grpc.testing.foo",
		Tag:           "varint,13,opt,name=foo",
		Filename:      "reflection/grpc_testing/proto2_ext.proto",
	},
	{
		ExtendedType:  (*ToBeExtended)(nil),
		ExtensionType: (*Extension)(nil),
		Field:         17,
		Name:          "grpc.testing.bar",
		Tag:           "bytes,17,opt,name=bar",
		Filename:      "reflection/grpc_testing/proto2_ext.proto",
	},
	{
		ExtendedType:  (*ToBeExtended)(nil),
		ExtensionType: (*SearchRequest)(nil),
		Field:         19,
		Name:          "grpc.testing.baz",
		Tag:           "bytes,19,opt,name=baz",
		Filename:      "reflection/grpc_testing/proto2_ext.proto",
	},
}

// Extension fields to ToBeExtended.
var (
	// optional int32 foo = 13;
	E_Foo = &file_reflection_grpc_testing_proto2_ext_proto_extTypes[0]
	// optional grpc.testing.Extension bar = 17;
	E_Bar = &file_reflection_grpc_testing_proto2_ext_proto_extTypes[1]
	// optional grpc.testing.SearchRequest baz = 19;
	E_Baz = &file_reflection_grpc_testing_proto2_ext_proto_extTypes[2]
)

var File_reflection_grpc_testing_proto2_ext_proto protoreflect.FileDescriptor

var file_reflection_grpc_testing_proto2_ext_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x5f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x24, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22,
	0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x25, 0x0a, 0x09, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x77, 0x68, 0x61, 0x74, 0x7a, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x77, 0x68, 0x61, 0x74, 0x7a, 0x69, 0x74, 0x3a, 0x2c, 0x0a, 0x03, 0x66, 0x6f, 0x6f,
	0x12, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x54, 0x6f, 0x42, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x3a, 0x45, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x12, 0x1a,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x6f,
	0x42, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x62, 0x61, 0x72, 0x3a, 0x49,
	0x0a, 0x03, 0x62, 0x61, 0x7a, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x6f, 0x42, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x03, 0x62, 0x61, 0x7a, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
}

var (
	file_reflection_grpc_testing_proto2_ext_proto_rawDescOnce sync.Once
	file_reflection_grpc_testing_proto2_ext_proto_rawDescData = file_reflection_grpc_testing_proto2_ext_proto_rawDesc
)

func file_reflection_grpc_testing_proto2_ext_proto_rawDescGZIP() []byte {
	file_reflection_grpc_testing_proto2_ext_proto_rawDescOnce.Do(func() {
		file_reflection_grpc_testing_proto2_ext_proto_rawDescData = protoimpl.X.CompressGZIP(file_reflection_grpc_testing_proto2_ext_proto_rawDescData)
	})
	return file_reflection_grpc_testing_proto2_ext_proto_rawDescData
}

var file_reflection_grpc_testing_proto2_ext_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_reflection_grpc_testing_proto2_ext_proto_goTypes = []interface{}{
	(*Extension)(nil),     // 0: grpc.testing.Extension
	(*ToBeExtended)(nil),  // 1: grpc.testing.ToBeExtended
	(*SearchRequest)(nil), // 2: grpc.testing.SearchRequest
}
var file_reflection_grpc_testing_proto2_ext_proto_depIdxs = []int32{
	1, // 0: grpc.testing.foo:extendee -> grpc.testing.ToBeExtended
	1, // 1: grpc.testing.bar:extendee -> grpc.testing.ToBeExtended
	1, // 2: grpc.testing.baz:extendee -> grpc.testing.ToBeExtended
	0, // 3: grpc.testing.bar:type_name -> grpc.testing.Extension
	2, // 4: grpc.testing.baz:type_name -> grpc.testing.SearchRequest
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	3, // [3:5] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reflection_grpc_testing_proto2_ext_proto_init() }
func file_reflection_grpc_testing_proto2_ext_proto_init() {
	if File_reflection_grpc_testing_proto2_ext_proto != nil {
		return
	}
	file_reflection_grpc_testing_proto2_proto_init()
	file_reflection_grpc_testing_test_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_reflection_grpc_testing_proto2_ext_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extension); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reflection_grpc_testing_proto2_ext_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_reflection_grpc_testing_proto2_ext_proto_goTypes,
		DependencyIndexes: file_reflection_grpc_testing_proto2_ext_proto_depIdxs,
		MessageInfos:      file_reflection_grpc_testing_proto2_ext_proto_msgTypes,
		ExtensionInfos:    file_reflection_grpc_testing_proto2_ext_proto_extTypes,
	}.Build()
	File_reflection_grpc_testing_proto2_ext_proto = out.File
	file_reflection_grpc_testing_proto2_ext_proto_rawDesc = nil
	file_reflection_grpc_testing_proto2_ext_proto_goTypes = nil
	file_reflection_grpc_testing_proto2_ext_proto_depIdxs = nil
}
