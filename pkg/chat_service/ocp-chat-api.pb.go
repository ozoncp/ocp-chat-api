// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: ocp-chat-api.proto

package chat_service

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClassroomId uint64 `protobuf:"varint,2,opt,name=classroom_id,json=classroomId,proto3" json:"classroom_id,omitempty"`
	Link        string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_chat_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_chat_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_ocp_chat_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateRequest) GetClassroomId() uint64 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

func (x *CreateRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_chat_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_chat_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_ocp_chat_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_ocp_chat_api_proto protoreflect.FileDescriptor

var file_ocp_chat_api_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x63, 0x70, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x56,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x3e, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x4a, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x74, 0x41, 0x70,
	0x69, 0x12, 0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12,
	0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ocp_chat_api_proto_rawDescOnce sync.Once
	file_ocp_chat_api_proto_rawDescData = file_ocp_chat_api_proto_rawDesc
)

func file_ocp_chat_api_proto_rawDescGZIP() []byte {
	file_ocp_chat_api_proto_rawDescOnce.Do(func() {
		file_ocp_chat_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocp_chat_api_proto_rawDescData)
	})
	return file_ocp_chat_api_proto_rawDescData
}

var (
	file_ocp_chat_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
	file_ocp_chat_api_proto_goTypes  = []interface{}{
		(*CreateRequest)(nil),  // 0: chat_api.CreateRequest
		(*CreateResponse)(nil), // 1: chat_api.CreateResponse
	}
)

var file_ocp_chat_api_proto_depIdxs = []int32{
	0, // 0: chat_api.ChatApi.CreateChat:input_type -> chat_api.CreateRequest
	1, // 1: chat_api.ChatApi.CreateChat:output_type -> chat_api.CreateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ocp_chat_api_proto_init() }
func file_ocp_chat_api_proto_init() {
	if File_ocp_chat_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ocp_chat_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_ocp_chat_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
			RawDescriptor: file_ocp_chat_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ocp_chat_api_proto_goTypes,
		DependencyIndexes: file_ocp_chat_api_proto_depIdxs,
		MessageInfos:      file_ocp_chat_api_proto_msgTypes,
	}.Build()
	File_ocp_chat_api_proto = out.File
	file_ocp_chat_api_proto_rawDesc = nil
	file_ocp_chat_api_proto_goTypes = nil
	file_ocp_chat_api_proto_depIdxs = nil
}