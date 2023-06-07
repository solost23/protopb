// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protos/face_recognition/face_recognition.proto

package face_recognition

import (
	common "github.com/solost23/protopb/gen/go/protos/common"
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

type GenerateFaceEncodingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// 文件流列表
	Data [][]byte `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	// 文件类型
	UploadType string `protobuf:"bytes,3,opt,name=uploadType,proto3" json:"uploadType,omitempty"`
}

func (x *GenerateFaceEncodingRequest) Reset() {
	*x = GenerateFaceEncodingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_face_recognition_face_recognition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateFaceEncodingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateFaceEncodingRequest) ProtoMessage() {}

func (x *GenerateFaceEncodingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_face_recognition_face_recognition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateFaceEncodingRequest.ProtoReflect.Descriptor instead.
func (*GenerateFaceEncodingRequest) Descriptor() ([]byte, []int) {
	return file_protos_face_recognition_face_recognition_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateFaceEncodingRequest) GetHeader() *common.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GenerateFaceEncodingRequest) GetData() [][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GenerateFaceEncodingRequest) GetUploadType() string {
	if x != nil {
		return x.UploadType
	}
	return ""
}

type GenerateFaceEncodingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo     *common.ErrorInfo `protobuf:"bytes,1,opt,name=errorInfo,proto3" json:"errorInfo,omitempty"`
	FaceEncodings []string          `protobuf:"bytes,2,rep,name=faceEncodings,proto3" json:"faceEncodings,omitempty"`
}

func (x *GenerateFaceEncodingResponse) Reset() {
	*x = GenerateFaceEncodingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_face_recognition_face_recognition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateFaceEncodingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateFaceEncodingResponse) ProtoMessage() {}

func (x *GenerateFaceEncodingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_face_recognition_face_recognition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateFaceEncodingResponse.ProtoReflect.Descriptor instead.
func (*GenerateFaceEncodingResponse) Descriptor() ([]byte, []int) {
	return file_protos_face_recognition_face_recognition_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateFaceEncodingResponse) GetErrorInfo() *common.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

func (x *GenerateFaceEncodingResponse) GetFaceEncodings() []string {
	if x != nil {
		return x.FaceEncodings
	}
	return nil
}

type CompareFacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// 文件流
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// 文件类型
	UploadType string `protobuf:"bytes,3,opt,name=uploadType,proto3" json:"uploadType,omitempty"`
}

func (x *CompareFacesRequest) Reset() {
	*x = CompareFacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_face_recognition_face_recognition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareFacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareFacesRequest) ProtoMessage() {}

func (x *CompareFacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_face_recognition_face_recognition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompareFacesRequest.ProtoReflect.Descriptor instead.
func (*CompareFacesRequest) Descriptor() ([]byte, []int) {
	return file_protos_face_recognition_face_recognition_proto_rawDescGZIP(), []int{2}
}

func (x *CompareFacesRequest) GetHeader() *common.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CompareFacesRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CompareFacesRequest) GetUploadType() string {
	if x != nil {
		return x.UploadType
	}
	return ""
}

type CompareFacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo *common.ErrorInfo `protobuf:"bytes,1,opt,name=errorInfo,proto3" json:"errorInfo,omitempty"`
	// 第一个对比上的用户的userId
	UserId int32 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	// 是否找到
	IsFound bool `protobuf:"varint,3,opt,name=isFound,proto3" json:"isFound,omitempty"`
}

func (x *CompareFacesResponse) Reset() {
	*x = CompareFacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_face_recognition_face_recognition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareFacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareFacesResponse) ProtoMessage() {}

func (x *CompareFacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_face_recognition_face_recognition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompareFacesResponse.ProtoReflect.Descriptor instead.
func (*CompareFacesResponse) Descriptor() ([]byte, []int) {
	return file_protos_face_recognition_face_recognition_proto_rawDescGZIP(), []int{3}
}

func (x *CompareFacesResponse) GetErrorInfo() *common.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

func (x *CompareFacesResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CompareFacesResponse) GetIsFound() bool {
	if x != nil {
		return x.IsFound
	}
	return false
}

var File_protos_face_recognition_face_recognition_proto protoreflect.FileDescriptor

var file_protos_face_recognition_face_recognition_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1d, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x66,
	0x61, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x1b,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x1c,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x61,
	0x63, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0d, 0x66, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x22, 0x85, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x46, 0x61, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x72, 0x65, 0x46, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x46, 0x6f, 0x75, 0x6e,
	0x64, 0x32, 0x9c, 0x02, 0x0a, 0x0f, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x8f, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x46, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x3a,
	0x2e, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x66, 0x61,
	0x63, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x67, 0x6f, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x46, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x77, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x65, 0x46, 0x61, 0x63, 0x65, 0x73, 0x12, 0x32, 0x2e, 0x67, 0x6f, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x46,
	0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x67, 0x6f,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x65, 0x46, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x73, 0x74, 0x32, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x62, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x61,
	0x63, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_face_recognition_face_recognition_proto_rawDescOnce sync.Once
	file_protos_face_recognition_face_recognition_proto_rawDescData = file_protos_face_recognition_face_recognition_proto_rawDesc
)

func file_protos_face_recognition_face_recognition_proto_rawDescGZIP() []byte {
	file_protos_face_recognition_face_recognition_proto_rawDescOnce.Do(func() {
		file_protos_face_recognition_face_recognition_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_face_recognition_face_recognition_proto_rawDescData)
	})
	return file_protos_face_recognition_face_recognition_proto_rawDescData
}

var file_protos_face_recognition_face_recognition_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_face_recognition_face_recognition_proto_goTypes = []interface{}{
	(*GenerateFaceEncodingRequest)(nil),  // 0: go_interface.face_recognition.GenerateFaceEncodingRequest
	(*GenerateFaceEncodingResponse)(nil), // 1: go_interface.face_recognition.GenerateFaceEncodingResponse
	(*CompareFacesRequest)(nil),          // 2: go_interface.face_recognition.CompareFacesRequest
	(*CompareFacesResponse)(nil),         // 3: go_interface.face_recognition.CompareFacesResponse
	(*common.RequestHeader)(nil),         // 4: go_interface.common.requestHeader
	(*common.ErrorInfo)(nil),             // 5: go_interface.common.errorInfo
}
var file_protos_face_recognition_face_recognition_proto_depIdxs = []int32{
	4, // 0: go_interface.face_recognition.GenerateFaceEncodingRequest.header:type_name -> go_interface.common.requestHeader
	5, // 1: go_interface.face_recognition.GenerateFaceEncodingResponse.errorInfo:type_name -> go_interface.common.errorInfo
	4, // 2: go_interface.face_recognition.CompareFacesRequest.header:type_name -> go_interface.common.requestHeader
	5, // 3: go_interface.face_recognition.CompareFacesResponse.errorInfo:type_name -> go_interface.common.errorInfo
	0, // 4: go_interface.face_recognition.faceRecognition.GenerateFaceEncoding:input_type -> go_interface.face_recognition.GenerateFaceEncodingRequest
	2, // 5: go_interface.face_recognition.faceRecognition.CompareFaces:input_type -> go_interface.face_recognition.CompareFacesRequest
	1, // 6: go_interface.face_recognition.faceRecognition.GenerateFaceEncoding:output_type -> go_interface.face_recognition.GenerateFaceEncodingResponse
	3, // 7: go_interface.face_recognition.faceRecognition.CompareFaces:output_type -> go_interface.face_recognition.CompareFacesResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_face_recognition_face_recognition_proto_init() }
func file_protos_face_recognition_face_recognition_proto_init() {
	if File_protos_face_recognition_face_recognition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_face_recognition_face_recognition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateFaceEncodingRequest); i {
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
		file_protos_face_recognition_face_recognition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateFaceEncodingResponse); i {
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
		file_protos_face_recognition_face_recognition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompareFacesRequest); i {
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
		file_protos_face_recognition_face_recognition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompareFacesResponse); i {
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
			RawDescriptor: file_protos_face_recognition_face_recognition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_face_recognition_face_recognition_proto_goTypes,
		DependencyIndexes: file_protos_face_recognition_face_recognition_proto_depIdxs,
		MessageInfos:      file_protos_face_recognition_face_recognition_proto_msgTypes,
	}.Build()
	File_protos_face_recognition_face_recognition_proto = out.File
	file_protos_face_recognition_face_recognition_proto_rawDesc = nil
	file_protos_face_recognition_face_recognition_proto_goTypes = nil
	file_protos_face_recognition_face_recognition_proto_depIdxs = nil
}