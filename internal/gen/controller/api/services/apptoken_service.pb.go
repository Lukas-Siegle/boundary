// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: controller/api/services/v1/apptoken_service.proto

package services

import (
	apptokens "github.com/hashicorp/boundary/sdk/pbs/controller/api/resources/apptokens"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateAppTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *apptokens.AppToken `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *CreateAppTokenRequest) Reset() {
	*x = CreateAppTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_apptoken_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppTokenRequest) ProtoMessage() {}

func (x *CreateAppTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_apptoken_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateAppTokenRequest) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_apptoken_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAppTokenRequest) GetItem() *apptokens.AppToken {
	if x != nil {
		return x.Item
	}
	return nil
}

type CreateAppTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *apptokens.AppToken `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *CreateAppTokenResponse) Reset() {
	*x = CreateAppTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_apptoken_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppTokenResponse) ProtoMessage() {}

func (x *CreateAppTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_apptoken_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppTokenResponse.ProtoReflect.Descriptor instead.
func (*CreateAppTokenResponse) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_apptoken_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAppTokenResponse) GetItem() *apptokens.AppToken {
	if x != nil {
		return x.Item
	}
	return nil
}

var File_controller_api_services_v1_apptoken_service_proto protoreflect.FileDescriptor

var file_controller_api_services_v1_apptoken_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x34, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x22, 0x5d, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x32, 0xa6, 0x01, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x92, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x70, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70,
	0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x70, 0x70, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_api_services_v1_apptoken_service_proto_rawDescOnce sync.Once
	file_controller_api_services_v1_apptoken_service_proto_rawDescData = file_controller_api_services_v1_apptoken_service_proto_rawDesc
)

func file_controller_api_services_v1_apptoken_service_proto_rawDescGZIP() []byte {
	file_controller_api_services_v1_apptoken_service_proto_rawDescOnce.Do(func() {
		file_controller_api_services_v1_apptoken_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_api_services_v1_apptoken_service_proto_rawDescData)
	})
	return file_controller_api_services_v1_apptoken_service_proto_rawDescData
}

var file_controller_api_services_v1_apptoken_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controller_api_services_v1_apptoken_service_proto_goTypes = []interface{}{
	(*CreateAppTokenRequest)(nil),  // 0: controller.api.services.v1.CreateAppTokenRequest
	(*CreateAppTokenResponse)(nil), // 1: controller.api.services.v1.CreateAppTokenResponse
	(*apptokens.AppToken)(nil),     // 2: controller.api.resources.apptokens.v1.AppToken
}
var file_controller_api_services_v1_apptoken_service_proto_depIdxs = []int32{
	2, // 0: controller.api.services.v1.CreateAppTokenRequest.item:type_name -> controller.api.resources.apptokens.v1.AppToken
	2, // 1: controller.api.services.v1.CreateAppTokenResponse.item:type_name -> controller.api.resources.apptokens.v1.AppToken
	0, // 2: controller.api.services.v1.AppTokenService.CreateAppToken:input_type -> controller.api.services.v1.CreateAppTokenRequest
	1, // 3: controller.api.services.v1.AppTokenService.CreateAppToken:output_type -> controller.api.services.v1.CreateAppTokenResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_controller_api_services_v1_apptoken_service_proto_init() }
func file_controller_api_services_v1_apptoken_service_proto_init() {
	if File_controller_api_services_v1_apptoken_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_api_services_v1_apptoken_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppTokenRequest); i {
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
		file_controller_api_services_v1_apptoken_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppTokenResponse); i {
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
			RawDescriptor: file_controller_api_services_v1_apptoken_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controller_api_services_v1_apptoken_service_proto_goTypes,
		DependencyIndexes: file_controller_api_services_v1_apptoken_service_proto_depIdxs,
		MessageInfos:      file_controller_api_services_v1_apptoken_service_proto_msgTypes,
	}.Build()
	File_controller_api_services_v1_apptoken_service_proto = out.File
	file_controller_api_services_v1_apptoken_service_proto_rawDesc = nil
	file_controller_api_services_v1_apptoken_service_proto_goTypes = nil
	file_controller_api_services_v1_apptoken_service_proto_depIdxs = nil
}
