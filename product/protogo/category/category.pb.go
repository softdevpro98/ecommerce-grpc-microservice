// Code generated by protoc-gen-go. DO NOT EDIT.
// source: category.proto

/*
Package category is a generated protocol buffer package.

It is generated from these files:
	category.proto

It has these top-level messages:
	CategoryQueryRequest
	CategoryResponse
	CategoryListReponse
*/
package category

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

type CategoryQueryRequest struct {
	ID int32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *CategoryQueryRequest) Reset()                    { *m = CategoryQueryRequest{} }
func (m *CategoryQueryRequest) String() string            { return proto.CompactTextString(m) }
func (*CategoryQueryRequest) ProtoMessage()               {}
func (*CategoryQueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CategoryQueryRequest) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type CategoryResponse struct {
	ID          int32  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description" json:"Description,omitempty"`
	Image       string `protobuf:"bytes,4,opt,name=Image" json:"Image,omitempty"`
}

func (m *CategoryResponse) Reset()                    { *m = CategoryResponse{} }
func (m *CategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*CategoryResponse) ProtoMessage()               {}
func (*CategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CategoryResponse) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *CategoryResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CategoryResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CategoryResponse) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type CategoryListReponse struct {
	CategoryList []*CategoryResponse `protobuf:"bytes,1,rep,name=CategoryList" json:"CategoryList,omitempty"`
}

func (m *CategoryListReponse) Reset()                    { *m = CategoryListReponse{} }
func (m *CategoryListReponse) String() string            { return proto.CompactTextString(m) }
func (*CategoryListReponse) ProtoMessage()               {}
func (*CategoryListReponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CategoryListReponse) GetCategoryList() []*CategoryResponse {
	if m != nil {
		return m.CategoryList
	}
	return nil
}

func init() {
	proto.RegisterType((*CategoryQueryRequest)(nil), "category.CategoryQueryRequest")
	proto.RegisterType((*CategoryResponse)(nil), "category.CategoryResponse")
	proto.RegisterType((*CategoryListReponse)(nil), "category.CategoryListReponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CategoryService service

type CategoryServiceClient interface {
	FindByID(ctx context.Context, in *CategoryQueryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	FindAll(ctx context.Context, in *CategoryQueryRequest, opts ...grpc.CallOption) (CategoryService_FindAllClient, error)
}

type categoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewCategoryServiceClient(cc *grpc.ClientConn) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) FindByID(ctx context.Context, in *CategoryQueryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := grpc.Invoke(ctx, "/category.CategoryService/FindByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) FindAll(ctx context.Context, in *CategoryQueryRequest, opts ...grpc.CallOption) (CategoryService_FindAllClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CategoryService_serviceDesc.Streams[0], c.cc, "/category.CategoryService/FindAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &categoryServiceFindAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CategoryService_FindAllClient interface {
	Recv() (*CategoryResponse, error)
	grpc.ClientStream
}

type categoryServiceFindAllClient struct {
	grpc.ClientStream
}

func (x *categoryServiceFindAllClient) Recv() (*CategoryResponse, error) {
	m := new(CategoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for CategoryService service

type CategoryServiceServer interface {
	FindByID(context.Context, *CategoryQueryRequest) (*CategoryResponse, error)
	FindAll(*CategoryQueryRequest, CategoryService_FindAllServer) error
}

func RegisterCategoryServiceServer(s *grpc.Server, srv CategoryServiceServer) {
	s.RegisterService(&_CategoryService_serviceDesc, srv)
}

func _CategoryService_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/category.CategoryService/FindByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).FindByID(ctx, req.(*CategoryQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_FindAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CategoryQueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CategoryServiceServer).FindAll(m, &categoryServiceFindAllServer{stream})
}

type CategoryService_FindAllServer interface {
	Send(*CategoryResponse) error
	grpc.ServerStream
}

type categoryServiceFindAllServer struct {
	grpc.ServerStream
}

func (x *categoryServiceFindAllServer) Send(m *CategoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CategoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "category.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindByID",
			Handler:    _CategoryService_FindByID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindAll",
			Handler:       _CategoryService_FindAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "category.proto",
}

func init() { proto.RegisterFile("category.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4e, 0x2c, 0x49,
	0x4d, 0xcf, 0x2f, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xd4,
	0xb8, 0x44, 0x9c, 0xa1, 0xec, 0xc0, 0xd2, 0xd4, 0xa2, 0xca, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x21, 0x3e, 0x2e, 0x26, 0x4f, 0x17, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x26, 0x4f,
	0x17, 0xa5, 0x3c, 0x2e, 0x01, 0x98, 0xba, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x74,
	0x35, 0x42, 0x42, 0x5c, 0x2c, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x60, 0xb6, 0x90, 0x02, 0x17, 0xb7, 0x4b, 0x6a, 0x71, 0x72, 0x51, 0x66, 0x41, 0x49, 0x66, 0x7e,
	0x9e, 0x04, 0x33, 0x58, 0x0a, 0x59, 0x48, 0x48, 0x84, 0x8b, 0xd5, 0x33, 0x37, 0x31, 0x3d, 0x55,
	0x82, 0x05, 0x2c, 0x07, 0xe1, 0x28, 0x85, 0x72, 0x09, 0xc3, 0xec, 0xf3, 0xc9, 0x2c, 0x2e, 0x09,
	0x4a, 0x85, 0x58, 0x69, 0xc7, 0xc5, 0x83, 0x2c, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0x24,
	0xa5, 0x07, 0xf7, 0x1f, 0xba, 0x23, 0x83, 0x50, 0xd4, 0x1b, 0x2d, 0x62, 0xe4, 0xe2, 0x87, 0x09,
	0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xb9, 0x71, 0x71, 0xb8, 0x65, 0xe6, 0xa5, 0x38,
	0x55, 0x7a, 0xba, 0x08, 0xc9, 0x61, 0x9a, 0x84, 0x1c, 0x2c, 0x52, 0x78, 0x6c, 0x12, 0x72, 0xe7,
	0x62, 0x07, 0x99, 0xe3, 0x98, 0x93, 0x43, 0x89, 0x31, 0x06, 0x8c, 0x49, 0x6c, 0xe0, 0x48, 0x32,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x36, 0x69, 0x9d, 0xe7, 0xb6, 0x01, 0x00, 0x00,
}
