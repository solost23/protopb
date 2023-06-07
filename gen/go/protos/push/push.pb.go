// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protos/push/push.proto

package push

import (
	common "github.com/solost23/protopb/gen/go/protos/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic       string `protobuf:"bytes,6,opt,name=topic,proto3" json:"topic,omitempty"`
	Name        string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Addr        string `protobuf:"bytes,8,opt,name=addr,proto3" json:"addr,omitempty"`
	ContentType string `protobuf:"bytes,9,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Content     string `protobuf:"bytes,10,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Email) Reset() {
	*x = Email{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_push_push_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_protos_push_push_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_protos_push_push_proto_rawDescGZIP(), []int{0}
}

func (x *Email) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Email) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Email) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Email) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Email) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SendEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Email  *Email                `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SendEmailRequest) Reset() {
	*x = SendEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_push_push_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailRequest) ProtoMessage() {}

func (x *SendEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_push_push_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailRequest.ProtoReflect.Descriptor instead.
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return file_protos_push_push_proto_rawDescGZIP(), []int{1}
}

func (x *SendEmailRequest) GetHeader() *common.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *SendEmailRequest) GetEmail() *Email {
	if x != nil {
		return x.Email
	}
	return nil
}

type SendEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo *common.ErrorInfo `protobuf:"bytes,1,opt,name=errorInfo,proto3" json:"errorInfo,omitempty"`
}

func (x *SendEmailResponse) Reset() {
	*x = SendEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_push_push_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailResponse) ProtoMessage() {}

func (x *SendEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_push_push_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailResponse.ProtoReflect.Descriptor instead.
func (*SendEmailResponse) Descriptor() ([]byte, []int) {
	return file_protos_push_push_proto_rawDescGZIP(), []int{2}
}

func (x *SendEmailResponse) GetErrorInfo() *common.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

type SendLarkTextByUnionIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header   *common.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	UnionIds []string              `protobuf:"bytes,2,rep,name=unionIds,proto3" json:"unionIds,omitempty"`
	Content  string                `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendLarkTextByUnionIdsRequest) Reset() {
	*x = SendLarkTextByUnionIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_push_push_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendLarkTextByUnionIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendLarkTextByUnionIdsRequest) ProtoMessage() {}

func (x *SendLarkTextByUnionIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_push_push_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendLarkTextByUnionIdsRequest.ProtoReflect.Descriptor instead.
func (*SendLarkTextByUnionIdsRequest) Descriptor() ([]byte, []int) {
	return file_protos_push_push_proto_rawDescGZIP(), []int{3}
}

func (x *SendLarkTextByUnionIdsRequest) GetHeader() *common.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *SendLarkTextByUnionIdsRequest) GetUnionIds() []string {
	if x != nil {
		return x.UnionIds
	}
	return nil
}

func (x *SendLarkTextByUnionIdsRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SendLarkTextByUnionIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo *common.ErrorInfo `protobuf:"bytes,1,opt,name=errorInfo,proto3" json:"errorInfo,omitempty"`
}

func (x *SendLarkTextByUnionIdsResponse) Reset() {
	*x = SendLarkTextByUnionIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_push_push_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendLarkTextByUnionIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendLarkTextByUnionIdsResponse) ProtoMessage() {}

func (x *SendLarkTextByUnionIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_push_push_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendLarkTextByUnionIdsResponse.ProtoReflect.Descriptor instead.
func (*SendLarkTextByUnionIdsResponse) Descriptor() ([]byte, []int) {
	return file_protos_push_push_proto_rawDescGZIP(), []int{4}
}

func (x *SendLarkTextByUnionIdsResponse) GetErrorInfo() *common.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

type SendLarkPostByUnionIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header   *common.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	UnionIds []string              `protobuf:"bytes,2,rep,name=unionIds,proto3" json:"unionIds,omitempty"`
	Content  *anypb.Any            `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendLarkPostByUnionIdsRequest) Reset() {
	*x = SendLarkPostByUnionIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_push_push_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendLarkPostByUnionIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendLarkPostByUnionIdsRequest) ProtoMessage() {}

func (x *SendLarkPostByUnionIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_push_push_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendLarkPostByUnionIdsRequest.ProtoReflect.Descriptor instead.
func (*SendLarkPostByUnionIdsRequest) Descriptor() ([]byte, []int) {
	return file_protos_push_push_proto_rawDescGZIP(), []int{5}
}

func (x *SendLarkPostByUnionIdsRequest) GetHeader() *common.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *SendLarkPostByUnionIdsRequest) GetUnionIds() []string {
	if x != nil {
		return x.UnionIds
	}
	return nil
}

func (x *SendLarkPostByUnionIdsRequest) GetContent() *anypb.Any {
	if x != nil {
		return x.Content
	}
	return nil
}

type SendLarkPostByUnionIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo *common.ErrorInfo `protobuf:"bytes,1,opt,name=errorInfo,proto3" json:"errorInfo,omitempty"`
}

func (x *SendLarkPostByUnionIdsResponse) Reset() {
	*x = SendLarkPostByUnionIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_push_push_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendLarkPostByUnionIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendLarkPostByUnionIdsResponse) ProtoMessage() {}

func (x *SendLarkPostByUnionIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_push_push_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendLarkPostByUnionIdsResponse.ProtoReflect.Descriptor instead.
func (*SendLarkPostByUnionIdsResponse) Descriptor() ([]byte, []int) {
	return file_protos_push_push_proto_rawDescGZIP(), []int{6}
}

func (x *SendLarkPostByUnionIdsResponse) GetErrorInfo() *common.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

var File_protos_push_push_proto protoreflect.FileDescriptor

var file_protos_push_push_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x2f, 0x70, 0x75,
	0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x7e, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x51, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x1d, 0x53, 0x65,
	0x6e, 0x64, 0x4c, 0x61, 0x72, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x42, 0x79, 0x55, 0x6e, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x5e, 0x0a,
	0x1e, 0x53, 0x65, 0x6e, 0x64, 0x4c, 0x61, 0x72, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x42, 0x79, 0x55,
	0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xa7, 0x01,
	0x0a, 0x1d, 0x53, 0x65, 0x6e, 0x64, 0x4c, 0x61, 0x72, 0x6b, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79,
	0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x5e, 0x0a, 0x1e, 0x53, 0x65, 0x6e, 0x64, 0x4c,
	0x61, 0x72, 0x6b, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67,
	0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xdc, 0x02, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68,
	0x12, 0x56, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x2e,
	0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x75, 0x73,
	0x68, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7d, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64,
	0x4c, 0x61, 0x72, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x42, 0x79, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x73, 0x12, 0x30, 0x2e, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4c, 0x61, 0x72, 0x6b, 0x54,
	0x65, 0x78, 0x74, 0x42, 0x79, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4c, 0x61, 0x72,
	0x6b, 0x54, 0x65, 0x78, 0x74, 0x42, 0x79, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7d, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x4c,
	0x61, 0x72, 0x6b, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x73, 0x12, 0x30, 0x2e, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4c, 0x61, 0x72, 0x6b, 0x50, 0x6f,
	0x73, 0x74, 0x42, 0x79, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4c, 0x61, 0x72, 0x6b,
	0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x73, 0x74, 0x32, 0x33, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_push_push_proto_rawDescOnce sync.Once
	file_protos_push_push_proto_rawDescData = file_protos_push_push_proto_rawDesc
)

func file_protos_push_push_proto_rawDescGZIP() []byte {
	file_protos_push_push_proto_rawDescOnce.Do(func() {
		file_protos_push_push_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_push_push_proto_rawDescData)
	})
	return file_protos_push_push_proto_rawDescData
}

var file_protos_push_push_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_push_push_proto_goTypes = []interface{}{
	(*Email)(nil),                          // 0: go_interface.push.Email
	(*SendEmailRequest)(nil),               // 1: go_interface.push.SendEmailRequest
	(*SendEmailResponse)(nil),              // 2: go_interface.push.SendEmailResponse
	(*SendLarkTextByUnionIdsRequest)(nil),  // 3: go_interface.push.SendLarkTextByUnionIdsRequest
	(*SendLarkTextByUnionIdsResponse)(nil), // 4: go_interface.push.SendLarkTextByUnionIdsResponse
	(*SendLarkPostByUnionIdsRequest)(nil),  // 5: go_interface.push.SendLarkPostByUnionIdsRequest
	(*SendLarkPostByUnionIdsResponse)(nil), // 6: go_interface.push.SendLarkPostByUnionIdsResponse
	(*common.RequestHeader)(nil),           // 7: go_interface.common.requestHeader
	(*common.ErrorInfo)(nil),               // 8: go_interface.common.errorInfo
	(*anypb.Any)(nil),                      // 9: google.protobuf.Any
}
var file_protos_push_push_proto_depIdxs = []int32{
	7,  // 0: go_interface.push.SendEmailRequest.header:type_name -> go_interface.common.requestHeader
	0,  // 1: go_interface.push.SendEmailRequest.email:type_name -> go_interface.push.Email
	8,  // 2: go_interface.push.SendEmailResponse.errorInfo:type_name -> go_interface.common.errorInfo
	7,  // 3: go_interface.push.SendLarkTextByUnionIdsRequest.header:type_name -> go_interface.common.requestHeader
	8,  // 4: go_interface.push.SendLarkTextByUnionIdsResponse.errorInfo:type_name -> go_interface.common.errorInfo
	7,  // 5: go_interface.push.SendLarkPostByUnionIdsRequest.header:type_name -> go_interface.common.requestHeader
	9,  // 6: go_interface.push.SendLarkPostByUnionIdsRequest.content:type_name -> google.protobuf.Any
	8,  // 7: go_interface.push.SendLarkPostByUnionIdsResponse.errorInfo:type_name -> go_interface.common.errorInfo
	1,  // 8: go_interface.push.Push.SendEmail:input_type -> go_interface.push.SendEmailRequest
	3,  // 9: go_interface.push.Push.SendLarkTextByUnionIds:input_type -> go_interface.push.SendLarkTextByUnionIdsRequest
	5,  // 10: go_interface.push.Push.SendLarkPostByUnionIds:input_type -> go_interface.push.SendLarkPostByUnionIdsRequest
	2,  // 11: go_interface.push.Push.SendEmail:output_type -> go_interface.push.SendEmailResponse
	4,  // 12: go_interface.push.Push.SendLarkTextByUnionIds:output_type -> go_interface.push.SendLarkTextByUnionIdsResponse
	6,  // 13: go_interface.push.Push.SendLarkPostByUnionIds:output_type -> go_interface.push.SendLarkPostByUnionIdsResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_protos_push_push_proto_init() }
func file_protos_push_push_proto_init() {
	if File_protos_push_push_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_push_push_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Email); i {
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
		file_protos_push_push_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailRequest); i {
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
		file_protos_push_push_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailResponse); i {
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
		file_protos_push_push_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendLarkTextByUnionIdsRequest); i {
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
		file_protos_push_push_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendLarkTextByUnionIdsResponse); i {
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
		file_protos_push_push_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendLarkPostByUnionIdsRequest); i {
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
		file_protos_push_push_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendLarkPostByUnionIdsResponse); i {
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
			RawDescriptor: file_protos_push_push_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_push_push_proto_goTypes,
		DependencyIndexes: file_protos_push_push_proto_depIdxs,
		MessageInfos:      file_protos_push_push_proto_msgTypes,
	}.Build()
	File_protos_push_push_proto = out.File
	file_protos_push_push_proto_rawDesc = nil
	file_protos_push_push_proto_goTypes = nil
	file_protos_push_push_proto_depIdxs = nil
}
