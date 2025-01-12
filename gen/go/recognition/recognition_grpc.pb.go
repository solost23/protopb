// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: recognition/recognition.proto

package recognition

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FaceRecognitionService_GenerateFaceEncoding_FullMethodName = "/recognition.FaceRecognitionService/GenerateFaceEncoding"
	FaceRecognitionService_CompareFaces_FullMethodName         = "/recognition.FaceRecognitionService/CompareFaces"
)

// FaceRecognitionServiceClient is the client API for FaceRecognitionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FaceRecognitionServiceClient interface {
	// 批量生成人像编码
	GenerateFaceEncoding(ctx context.Context, in *GenerateFaceEncodingRequest, opts ...grpc.CallOption) (*GenerateFaceEncodingResponse, error)
	// 对比人像库，返回符合条件人列表的第一个符合人UserID & true, 若没有符合人，那么直接返回false
	CompareFaces(ctx context.Context, in *CompareFacesRequest, opts ...grpc.CallOption) (*CompareFacesResponse, error)
}

type faceRecognitionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFaceRecognitionServiceClient(cc grpc.ClientConnInterface) FaceRecognitionServiceClient {
	return &faceRecognitionServiceClient{cc}
}

func (c *faceRecognitionServiceClient) GenerateFaceEncoding(ctx context.Context, in *GenerateFaceEncodingRequest, opts ...grpc.CallOption) (*GenerateFaceEncodingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateFaceEncodingResponse)
	err := c.cc.Invoke(ctx, FaceRecognitionService_GenerateFaceEncoding_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faceRecognitionServiceClient) CompareFaces(ctx context.Context, in *CompareFacesRequest, opts ...grpc.CallOption) (*CompareFacesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CompareFacesResponse)
	err := c.cc.Invoke(ctx, FaceRecognitionService_CompareFaces_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FaceRecognitionServiceServer is the server API for FaceRecognitionService service.
// All implementations must embed UnimplementedFaceRecognitionServiceServer
// for forward compatibility.
type FaceRecognitionServiceServer interface {
	// 批量生成人像编码
	GenerateFaceEncoding(context.Context, *GenerateFaceEncodingRequest) (*GenerateFaceEncodingResponse, error)
	// 对比人像库，返回符合条件人列表的第一个符合人UserID & true, 若没有符合人，那么直接返回false
	CompareFaces(context.Context, *CompareFacesRequest) (*CompareFacesResponse, error)
	mustEmbedUnimplementedFaceRecognitionServiceServer()
}

// UnimplementedFaceRecognitionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFaceRecognitionServiceServer struct{}

func (UnimplementedFaceRecognitionServiceServer) GenerateFaceEncoding(context.Context, *GenerateFaceEncodingRequest) (*GenerateFaceEncodingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateFaceEncoding not implemented")
}
func (UnimplementedFaceRecognitionServiceServer) CompareFaces(context.Context, *CompareFacesRequest) (*CompareFacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompareFaces not implemented")
}
func (UnimplementedFaceRecognitionServiceServer) mustEmbedUnimplementedFaceRecognitionServiceServer() {
}
func (UnimplementedFaceRecognitionServiceServer) testEmbeddedByValue() {}

// UnsafeFaceRecognitionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FaceRecognitionServiceServer will
// result in compilation errors.
type UnsafeFaceRecognitionServiceServer interface {
	mustEmbedUnimplementedFaceRecognitionServiceServer()
}

func RegisterFaceRecognitionServiceServer(s grpc.ServiceRegistrar, srv FaceRecognitionServiceServer) {
	// If the following call pancis, it indicates UnimplementedFaceRecognitionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FaceRecognitionService_ServiceDesc, srv)
}

func _FaceRecognitionService_GenerateFaceEncoding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateFaceEncodingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaceRecognitionServiceServer).GenerateFaceEncoding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FaceRecognitionService_GenerateFaceEncoding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaceRecognitionServiceServer).GenerateFaceEncoding(ctx, req.(*GenerateFaceEncodingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaceRecognitionService_CompareFaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompareFacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaceRecognitionServiceServer).CompareFaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FaceRecognitionService_CompareFaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaceRecognitionServiceServer).CompareFaces(ctx, req.(*CompareFacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FaceRecognitionService_ServiceDesc is the grpc.ServiceDesc for FaceRecognitionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FaceRecognitionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "recognition.FaceRecognitionService",
	HandlerType: (*FaceRecognitionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateFaceEncoding",
			Handler:    _FaceRecognitionService_GenerateFaceEncoding_Handler,
		},
		{
			MethodName: "CompareFaces",
			Handler:    _FaceRecognitionService_CompareFaces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recognition/recognition.proto",
}
