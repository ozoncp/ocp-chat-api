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

type CreateChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClassroomId uint64 `protobuf:"varint,2,opt,name=classroom_id,json=classroomId,proto3" json:"classroom_id,omitempty"`
	Link        string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *CreateChatRequest) Reset() {
	*x = CreateChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_chat_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatRequest) ProtoMessage() {}

func (x *CreateChatRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateChatRequest.ProtoReflect.Descriptor instead.
func (*CreateChatRequest) Descriptor() ([]byte, []int) {
	return file_ocp_chat_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateChatRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateChatRequest) GetClassroomId() uint64 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

func (x *CreateChatRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type CreateChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateChatResponse) Reset() {
	*x = CreateChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_chat_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatResponse) ProtoMessage() {}

func (x *CreateChatResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateChatResponse.ProtoReflect.Descriptor instead.
func (*CreateChatResponse) Descriptor() ([]byte, []int) {
	return file_ocp_chat_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateChatResponse) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateChatResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DescribeChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DescribeChatRequest) Reset() {
	*x = DescribeChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_chat_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeChatRequest) ProtoMessage() {}

func (x *DescribeChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_chat_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeChatRequest.ProtoReflect.Descriptor instead.
func (*DescribeChatRequest) Descriptor() ([]byte, []int) {
	return file_ocp_chat_api_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeChatRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DescribeChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClassroomId uint64 `protobuf:"varint,2,opt,name=classroom_id,json=classroomId,proto3" json:"classroom_id,omitempty"`
	Link        string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *DescribeChatResponse) Reset() {
	*x = DescribeChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_chat_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeChatResponse) ProtoMessage() {}

func (x *DescribeChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_chat_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeChatResponse.ProtoReflect.Descriptor instead.
func (*DescribeChatResponse) Descriptor() ([]byte, []int) {
	return file_ocp_chat_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeChatResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DescribeChatResponse) GetClassroomId() uint64 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

func (x *DescribeChatResponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type ListChatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ListChatsRequest) Reset() {
	*x = ListChatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_chat_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChatsRequest) ProtoMessage() {}

func (x *ListChatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_chat_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChatsRequest.ProtoReflect.Descriptor instead.
func (*ListChatsRequest) Descriptor() ([]byte, []int) {
	return file_ocp_chat_api_proto_rawDescGZIP(), []int{4}
}

func (x *ListChatsRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListChatOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClassroomId uint64 `protobuf:"varint,2,opt,name=classroom_id,json=classroomId,proto3" json:"classroom_id,omitempty"`
	Link        string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *ListChatOne) Reset() {
	*x = ListChatOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_chat_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChatOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChatOne) ProtoMessage() {}

func (x *ListChatOne) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_chat_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChatOne.ProtoReflect.Descriptor instead.
func (*ListChatOne) Descriptor() ([]byte, []int) {
	return file_ocp_chat_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListChatOne) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListChatOne) GetClassroomId() uint64 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

func (x *ListChatOne) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type ListChatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Packet []*ListChatOne `protobuf:"bytes,1,rep,name=packet,proto3" json:"packet,omitempty"`
}

func (x *ListChatsResponse) Reset() {
	*x = ListChatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_chat_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChatsResponse) ProtoMessage() {}

func (x *ListChatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_chat_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChatsResponse.ProtoReflect.Descriptor instead.
func (*ListChatsResponse) Descriptor() ([]byte, []int) {
	return file_ocp_chat_api_proto_rawDescGZIP(), []int{6}
}

func (x *ListChatsResponse) GetPacket() []*ListChatOne {
	if x != nil {
		return x.Packet
	}
	return nil
}

type RemoveChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveChatRequest) Reset() {
	*x = RemoveChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_chat_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveChatRequest) ProtoMessage() {}

func (x *RemoveChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_chat_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveChatRequest.ProtoReflect.Descriptor instead.
func (*RemoveChatRequest) Descriptor() ([]byte, []int) {
	return file_ocp_chat_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveChatRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RemoveChatResponse) Reset() {
	*x = RemoveChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_chat_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveChatResponse) ProtoMessage() {}

func (x *RemoveChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_chat_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveChatResponse.ProtoReflect.Descriptor instead.
func (*RemoveChatResponse) Descriptor() ([]byte, []int) {
	return file_ocp_chat_api_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveChatResponse) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RemoveChatResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_ocp_chat_api_proto protoreflect.FileDescriptor

var file_ocp_chat_api_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x63, 0x70, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x5a,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x42, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x25,
	0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x22, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x54, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x68, 0x61, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x42,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x06, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb0, 0x02, 0x0a, 0x07,
	0x43, 0x68, 0x61, 0x74, 0x41, 0x70, 0x69, 0x12, 0x47, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4d, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43,
	0x68, 0x61, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11,
	0x5a, 0x0f, 0x2e, 0x2f, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	file_ocp_chat_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
	file_ocp_chat_api_proto_goTypes  = []interface{}{
		(*CreateChatRequest)(nil),    // 0: chat_api.CreateChatRequest
		(*CreateChatResponse)(nil),   // 1: chat_api.CreateChatResponse
		(*DescribeChatRequest)(nil),  // 2: chat_api.DescribeChatRequest
		(*DescribeChatResponse)(nil), // 3: chat_api.DescribeChatResponse
		(*ListChatsRequest)(nil),     // 4: chat_api.ListChatsRequest
		(*ListChatOne)(nil),          // 5: chat_api.ListChatOne
		(*ListChatsResponse)(nil),    // 6: chat_api.ListChatsResponse
		(*RemoveChatRequest)(nil),    // 7: chat_api.RemoveChatRequest
		(*RemoveChatResponse)(nil),   // 8: chat_api.RemoveChatResponse
	}
)

var file_ocp_chat_api_proto_depIdxs = []int32{
	5, // 0: chat_api.ListChatsResponse.packet:type_name -> chat_api.ListChatOne
	0, // 1: chat_api.ChatApi.CreateChat:input_type -> chat_api.CreateChatRequest
	2, // 2: chat_api.ChatApi.DescribeChat:input_type -> chat_api.DescribeChatRequest
	4, // 3: chat_api.ChatApi.ListChats:input_type -> chat_api.ListChatsRequest
	7, // 4: chat_api.ChatApi.RemoveChat:input_type -> chat_api.RemoveChatRequest
	1, // 5: chat_api.ChatApi.CreateChat:output_type -> chat_api.CreateChatResponse
	3, // 6: chat_api.ChatApi.DescribeChat:output_type -> chat_api.DescribeChatResponse
	6, // 7: chat_api.ChatApi.ListChats:output_type -> chat_api.ListChatsResponse
	8, // 8: chat_api.ChatApi.RemoveChat:output_type -> chat_api.RemoveChatResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ocp_chat_api_proto_init() }
func file_ocp_chat_api_proto_init() {
	if File_ocp_chat_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ocp_chat_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatRequest); i {
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
			switch v := v.(*CreateChatResponse); i {
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
		file_ocp_chat_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeChatRequest); i {
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
		file_ocp_chat_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeChatResponse); i {
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
		file_ocp_chat_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChatsRequest); i {
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
		file_ocp_chat_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChatOne); i {
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
		file_ocp_chat_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChatsResponse); i {
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
		file_ocp_chat_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveChatRequest); i {
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
		file_ocp_chat_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveChatResponse); i {
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
			NumMessages:   9,
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
