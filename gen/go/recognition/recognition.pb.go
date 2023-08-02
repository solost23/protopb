// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: recognition/recognition.proto

package recognition

import (
	common "github.com/solost23/protopb/common"
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
	Data   [][]byte              `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"` // 文件流列表
}

func (x *GenerateFaceEncodingRequest) Reset() {
	*x = GenerateFaceEncodingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recognition_recognition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateFaceEncodingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateFaceEncodingRequest) ProtoMessage() {}

func (x *GenerateFaceEncodingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recognition_recognition_proto_msgTypes[0]
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
	return file_recognition_recognition_proto_rawDescGZIP(), []int{0}
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
		mi := &file_recognition_recognition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateFaceEncodingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateFaceEncodingResponse) ProtoMessage() {}

func (x *GenerateFaceEncodingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recognition_recognition_proto_msgTypes[1]
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
	return file_recognition_recognition_proto_rawDescGZIP(), []int{1}
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
	Data   []byte                `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"` // 文件流
}

func (x *CompareFacesRequest) Reset() {
	*x = CompareFacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recognition_recognition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareFacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareFacesRequest) ProtoMessage() {}

func (x *CompareFacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recognition_recognition_proto_msgTypes[2]
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
	return file_recognition_recognition_proto_rawDescGZIP(), []int{2}
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

type CompareFacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo *common.ErrorInfo `protobuf:"bytes,1,opt,name=errorInfo,proto3" json:"errorInfo,omitempty"`
	UserId    string            `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`    // 第一个对比上的用户的userId
	IsFound   bool              `protobuf:"varint,3,opt,name=isFound,proto3" json:"isFound,omitempty"` // 是否找到
}

func (x *CompareFacesResponse) Reset() {
	*x = CompareFacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recognition_recognition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareFacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareFacesResponse) ProtoMessage() {}

func (x *CompareFacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recognition_recognition_proto_msgTypes[3]
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
	return file_recognition_recognition_proto_rawDescGZIP(), []int{3}
}

func (x *CompareFacesResponse) GetErrorInfo() *common.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

func (x *CompareFacesResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CompareFacesResponse) GetIsFound() bool {
	if x != nil {
		return x.IsFound
	}
	return false
}

var File_recognition_recognition_proto protoreflect.FileDescriptor

var file_recognition_recognition_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65,
	0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x60, 0x0a, 0x1b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x61, 0x63,
	0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x75, 0x0a, 0x1c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46,
	0x61, 0x63, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x61, 0x63,
	0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x58, 0x0a, 0x13, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x65, 0x46, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x79, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x46,
	0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x46, 0x6f, 0x75, 0x6e, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x32,
	0xda, 0x01, 0x0a, 0x16, 0x46, 0x61, 0x63, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x14, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x28, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x72,
	0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x46, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x65, 0x46, 0x61, 0x63, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x46, 0x61, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x46,
	0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a, 0x27,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x73,
	0x74, 0x32, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x63, 0x6f,
	0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recognition_recognition_proto_rawDescOnce sync.Once
	file_recognition_recognition_proto_rawDescData = file_recognition_recognition_proto_rawDesc
)

func file_recognition_recognition_proto_rawDescGZIP() []byte {
	file_recognition_recognition_proto_rawDescOnce.Do(func() {
		file_recognition_recognition_proto_rawDescData = protoimpl.X.CompressGZIP(file_recognition_recognition_proto_rawDescData)
	})
	return file_recognition_recognition_proto_rawDescData
}

var file_recognition_recognition_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_recognition_recognition_proto_goTypes = []interface{}{
	(*GenerateFaceEncodingRequest)(nil),  // 0: recognition.GenerateFaceEncodingRequest
	(*GenerateFaceEncodingResponse)(nil), // 1: recognition.GenerateFaceEncodingResponse
	(*CompareFacesRequest)(nil),          // 2: recognition.CompareFacesRequest
	(*CompareFacesResponse)(nil),         // 3: recognition.CompareFacesResponse
	(*common.RequestHeader)(nil),         // 4: common.RequestHeader
	(*common.ErrorInfo)(nil),             // 5: common.ErrorInfo
}
var file_recognition_recognition_proto_depIdxs = []int32{
	4, // 0: recognition.GenerateFaceEncodingRequest.header:type_name -> common.RequestHeader
	5, // 1: recognition.GenerateFaceEncodingResponse.errorInfo:type_name -> common.ErrorInfo
	4, // 2: recognition.CompareFacesRequest.header:type_name -> common.RequestHeader
	5, // 3: recognition.CompareFacesResponse.errorInfo:type_name -> common.ErrorInfo
	0, // 4: recognition.FaceRecognitionService.GenerateFaceEncoding:input_type -> recognition.GenerateFaceEncodingRequest
	2, // 5: recognition.FaceRecognitionService.CompareFaces:input_type -> recognition.CompareFacesRequest
	1, // 6: recognition.FaceRecognitionService.GenerateFaceEncoding:output_type -> recognition.GenerateFaceEncodingResponse
	3, // 7: recognition.FaceRecognitionService.CompareFaces:output_type -> recognition.CompareFacesResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_recognition_recognition_proto_init() }
func file_recognition_recognition_proto_init() {
	if File_recognition_recognition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recognition_recognition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_recognition_recognition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_recognition_recognition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_recognition_recognition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_recognition_recognition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recognition_recognition_proto_goTypes,
		DependencyIndexes: file_recognition_recognition_proto_depIdxs,
		MessageInfos:      file_recognition_recognition_proto_msgTypes,
	}.Build()
	File_recognition_recognition_proto = out.File
	file_recognition_recognition_proto_rawDesc = nil
	file_recognition_recognition_proto_goTypes = nil
	file_recognition_recognition_proto_depIdxs = nil
}
