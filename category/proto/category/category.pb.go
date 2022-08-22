// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/category/category.proto

/*
Package go_micro_service_category is a generated protocol buffer package.

包名，通过protoc生成go文件

It is generated from these files:
	proto/category/category.proto

It has these top-level messages:
	CategoryRequest
	CategoryResponse
	DeleteCategoryRequest
	DeleteCategoryResponse
	FindByNameRequest
	FindByIdRequest
	FindByLevelRequest
	FindByParentRequest
	FindAllRequest
	FindAllResponse
	CreateCategoryResponse
	UpdateCategoryResponse
*/
package go_micro_service_category

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 定义一个分类的请求消息，包含category_name、category_level等5个字段
type CategoryRequest struct {
	CategoryName        string `protobuf:"bytes,1,opt,name=category_name,json=categoryName" json:"category_name,omitempty"`
	CategoryLevel       uint32 `protobuf:"varint,2,opt,name=category_level,json=categoryLevel" json:"category_level,omitempty"`
	CategoryParent      int64  `protobuf:"varint,3,opt,name=category_parent,json=categoryParent" json:"category_parent,omitempty"`
	CategoryImage       string `protobuf:"bytes,4,opt,name=category_image,json=categoryImage" json:"category_image,omitempty"`
	CategoryDescription string `protobuf:"bytes,5,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
}

func (m *CategoryRequest) Reset()                    { *m = CategoryRequest{} }
func (m *CategoryRequest) String() string            { return proto.CompactTextString(m) }
func (*CategoryRequest) ProtoMessage()               {}
func (*CategoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CategoryRequest) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

func (m *CategoryRequest) GetCategoryLevel() uint32 {
	if m != nil {
		return m.CategoryLevel
	}
	return 0
}

func (m *CategoryRequest) GetCategoryParent() int64 {
	if m != nil {
		return m.CategoryParent
	}
	return 0
}

func (m *CategoryRequest) GetCategoryImage() string {
	if m != nil {
		return m.CategoryImage
	}
	return ""
}

func (m *CategoryRequest) GetCategoryDescription() string {
	if m != nil {
		return m.CategoryDescription
	}
	return ""
}

// 定义一个分类的响应消息，包含id、category_name、category_level等6个字段
type CategoryResponse struct {
	Id                  int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	CategoryName        string `protobuf:"bytes,2,opt,name=category_name,json=categoryName" json:"category_name,omitempty"`
	CategoryLevel       uint32 `protobuf:"varint,3,opt,name=category_level,json=categoryLevel" json:"category_level,omitempty"`
	CategoryParent      int64  `protobuf:"varint,4,opt,name=category_parent,json=categoryParent" json:"category_parent,omitempty"`
	CategoryImages      string `protobuf:"bytes,5,opt,name=category_images,json=categoryImages" json:"category_images,omitempty"`
	CategoryDescription string `protobuf:"bytes,6,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
}

func (m *CategoryResponse) Reset()                    { *m = CategoryResponse{} }
func (m *CategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*CategoryResponse) ProtoMessage()               {}
func (*CategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CategoryResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CategoryResponse) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

func (m *CategoryResponse) GetCategoryLevel() uint32 {
	if m != nil {
		return m.CategoryLevel
	}
	return 0
}

func (m *CategoryResponse) GetCategoryParent() int64 {
	if m != nil {
		return m.CategoryParent
	}
	return 0
}

func (m *CategoryResponse) GetCategoryImages() string {
	if m != nil {
		return m.CategoryImages
	}
	return ""
}

func (m *CategoryResponse) GetCategoryDescription() string {
	if m != nil {
		return m.CategoryDescription
	}
	return ""
}

type DeleteCategoryRequest struct {
	CategoryId int64 `protobuf:"varint,1,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
}

func (m *DeleteCategoryRequest) Reset()                    { *m = DeleteCategoryRequest{} }
func (m *DeleteCategoryRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteCategoryRequest) ProtoMessage()               {}
func (*DeleteCategoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeleteCategoryRequest) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

type DeleteCategoryResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *DeleteCategoryResponse) Reset()                    { *m = DeleteCategoryResponse{} }
func (m *DeleteCategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteCategoryResponse) ProtoMessage()               {}
func (*DeleteCategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeleteCategoryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

//
type FindByNameRequest struct {
	CategoryName string `protobuf:"bytes,1,opt,name=category_name,json=categoryName" json:"category_name,omitempty"`
}

func (m *FindByNameRequest) Reset()                    { *m = FindByNameRequest{} }
func (m *FindByNameRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByNameRequest) ProtoMessage()               {}
func (*FindByNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FindByNameRequest) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

type FindByIdRequest struct {
	CategoryId int64 `protobuf:"varint,1,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
}

func (m *FindByIdRequest) Reset()                    { *m = FindByIdRequest{} }
func (m *FindByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByIdRequest) ProtoMessage()               {}
func (*FindByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FindByIdRequest) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

type FindByLevelRequest struct {
	Level uint32 `protobuf:"varint,1,opt,name=level" json:"level,omitempty"`
}

func (m *FindByLevelRequest) Reset()                    { *m = FindByLevelRequest{} }
func (m *FindByLevelRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByLevelRequest) ProtoMessage()               {}
func (*FindByLevelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FindByLevelRequest) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type FindByParentRequest struct {
	ParentId int64 `protobuf:"varint,1,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
}

func (m *FindByParentRequest) Reset()                    { *m = FindByParentRequest{} }
func (m *FindByParentRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByParentRequest) ProtoMessage()               {}
func (*FindByParentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FindByParentRequest) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

type FindAllRequest struct {
}

func (m *FindAllRequest) Reset()                    { *m = FindAllRequest{} }
func (m *FindAllRequest) String() string            { return proto.CompactTextString(m) }
func (*FindAllRequest) ProtoMessage()               {}
func (*FindAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type FindAllResponse struct {
	Category []*CategoryResponse `protobuf:"bytes,1,rep,name=category" json:"category,omitempty"`
}

func (m *FindAllResponse) Reset()                    { *m = FindAllResponse{} }
func (m *FindAllResponse) String() string            { return proto.CompactTextString(m) }
func (*FindAllResponse) ProtoMessage()               {}
func (*FindAllResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *FindAllResponse) GetCategory() []*CategoryResponse {
	if m != nil {
		return m.Category
	}
	return nil
}

type CreateCategoryResponse struct {
	Message    string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	CategoryId int64  `protobuf:"varint,2,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
}

func (m *CreateCategoryResponse) Reset()                    { *m = CreateCategoryResponse{} }
func (m *CreateCategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateCategoryResponse) ProtoMessage()               {}
func (*CreateCategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CreateCategoryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateCategoryResponse) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

type UpdateCategoryResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *UpdateCategoryResponse) Reset()                    { *m = UpdateCategoryResponse{} }
func (m *UpdateCategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateCategoryResponse) ProtoMessage()               {}
func (*UpdateCategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UpdateCategoryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CategoryRequest)(nil), "go.micro.service.category.CategoryRequest")
	proto.RegisterType((*CategoryResponse)(nil), "go.micro.service.category.CategoryResponse")
	proto.RegisterType((*DeleteCategoryRequest)(nil), "go.micro.service.category.DeleteCategoryRequest")
	proto.RegisterType((*DeleteCategoryResponse)(nil), "go.micro.service.category.DeleteCategoryResponse")
	proto.RegisterType((*FindByNameRequest)(nil), "go.micro.service.category.FindByNameRequest")
	proto.RegisterType((*FindByIdRequest)(nil), "go.micro.service.category.FindByIdRequest")
	proto.RegisterType((*FindByLevelRequest)(nil), "go.micro.service.category.FindByLevelRequest")
	proto.RegisterType((*FindByParentRequest)(nil), "go.micro.service.category.FindByParentRequest")
	proto.RegisterType((*FindAllRequest)(nil), "go.micro.service.category.FindAllRequest")
	proto.RegisterType((*FindAllResponse)(nil), "go.micro.service.category.FindAllResponse")
	proto.RegisterType((*CreateCategoryResponse)(nil), "go.micro.service.category.CreateCategoryResponse")
	proto.RegisterType((*UpdateCategoryResponse)(nil), "go.micro.service.category.UpdateCategoryResponse")
}

func init() { proto.RegisterFile("proto/category/category.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0xed, 0xfe, 0xa4, 0x53, 0xea, 0x84, 0x6d, 0xa8, 0x4c, 0x10, 0x22, 0x5a, 0x84, 0x08,
	0x05, 0x0c, 0x35, 0x97, 0x5e, 0xa1, 0x11, 0x28, 0x12, 0x42, 0xc8, 0x88, 0x0b, 0x17, 0x64, 0xe2,
	0x51, 0x64, 0x14, 0xff, 0xd4, 0x36, 0x45, 0x3c, 0x0e, 0x6f, 0xc6, 0x3b, 0xf0, 0x02, 0xc8, 0xbb,
	0xde, 0x75, 0xec, 0x3a, 0x8e, 0x8d, 0x7a, 0xf3, 0xce, 0xce, 0xcc, 0x37, 0xdf, 0x37, 0x33, 0x6b,
	0xb8, 0x1f, 0xc5, 0x61, 0x1a, 0xbe, 0x58, 0x38, 0x29, 0x2e, 0xc3, 0xf8, 0x97, 0xfc, 0x30, 0x99,
	0x9d, 0xdc, 0x5d, 0x86, 0xa6, 0xef, 0x2d, 0xe2, 0xd0, 0x4c, 0x30, 0xbe, 0xf2, 0x16, 0x68, 0x0a,
	0x07, 0xfa, 0x47, 0x81, 0xc1, 0x45, 0x7e, 0xb0, 0xf1, 0xf2, 0x07, 0x26, 0x29, 0x79, 0x08, 0x47,
	0xe2, 0xfe, 0x6b, 0xe0, 0xf8, 0x68, 0x28, 0x13, 0x65, 0x7a, 0x60, 0xdf, 0x12, 0xc6, 0x0f, 0x8e,
	0x8f, 0xe4, 0x11, 0xe8, 0xd2, 0x69, 0x85, 0x57, 0xb8, 0x32, 0xd4, 0x89, 0x32, 0x3d, 0xb2, 0x65,
	0xe8, 0xfb, 0xcc, 0x48, 0x1e, 0xc3, 0x40, 0xba, 0x45, 0x4e, 0x8c, 0x41, 0x6a, 0x68, 0x13, 0x65,
	0xaa, 0xd9, 0x32, 0xfa, 0x23, 0xb3, 0x96, 0xf2, 0x79, 0xbe, 0xb3, 0x44, 0x63, 0x87, 0xa1, 0xca,
	0x7c, 0xf3, 0xcc, 0x48, 0xce, 0x60, 0x24, 0xdd, 0x5c, 0x4c, 0x16, 0xb1, 0x17, 0xa5, 0x5e, 0x18,
	0x18, 0xbb, 0xcc, 0xf9, 0x58, 0xdc, 0xcd, 0x8a, 0x2b, 0xfa, 0x57, 0x81, 0x61, 0x41, 0x31, 0x89,
	0xc2, 0x20, 0x41, 0xa2, 0x83, 0xea, 0xb9, 0x8c, 0x98, 0x66, 0xab, 0x9e, 0x7b, 0x9d, 0xb3, 0xda,
	0x8a, 0xb3, 0xd6, 0x92, 0xf3, 0x4e, 0x2d, 0xe7, 0x75, 0x47, 0xc6, 0x39, 0xc9, 0x79, 0xe8, 0x25,
	0xd2, 0xc9, 0x46, 0xd6, 0x7b, 0x9b, 0x59, 0x9f, 0xc3, 0x9d, 0x19, 0xae, 0x30, 0xc5, 0x6a, 0x77,
	0x1f, 0xc0, 0x61, 0x01, 0x2a, 0x24, 0x00, 0x09, 0xe8, 0x52, 0x0b, 0x4e, 0xaa, 0x91, 0xb9, 0x68,
	0x06, 0xec, 0xfb, 0x98, 0x24, 0x59, 0x73, 0xf8, 0x48, 0x88, 0x23, 0x3d, 0x87, 0xdb, 0x6f, 0xbd,
	0xc0, 0x7d, 0xc3, 0x74, 0xea, 0x32, 0x47, 0xd4, 0x82, 0x01, 0x8f, 0x9c, 0xbb, 0xad, 0x2b, 0x3c,
	0x05, 0xc2, 0x63, 0x98, 0xde, 0x22, 0x6c, 0x04, 0xbb, 0xbc, 0x29, 0x0a, 0x6b, 0x0a, 0x3f, 0x50,
	0x0b, 0x8e, 0xb9, 0x2f, 0xd7, 0x5c, 0x38, 0xdf, 0x83, 0x03, 0xde, 0x9a, 0x02, 0xa1, 0xcf, 0x0d,
	0x73, 0x97, 0x0e, 0x41, 0xcf, 0x62, 0x5e, 0xaf, 0x44, 0x6e, 0xfa, 0x85, 0x57, 0xc9, 0x2c, 0xb9,
	0x18, 0xef, 0xa0, 0x2f, 0x4a, 0x32, 0x94, 0x89, 0x36, 0x3d, 0xb4, 0x9e, 0x9a, 0x1b, 0xf7, 0xcc,
	0xac, 0x6a, 0x69, 0xcb, 0x60, 0xfa, 0x09, 0x4e, 0x2e, 0x62, 0x74, 0xba, 0xe8, 0x5d, 0x95, 0x48,
	0xad, 0x6b, 0xe2, 0xe7, 0xc8, 0xed, 0x94, 0xd4, 0xfa, 0xbd, 0x0f, 0x7d, 0xe1, 0x4e, 0x2e, 0x41,
	0x2f, 0x57, 0x45, 0x4e, 0x5b, 0xd1, 0x63, 0x7a, 0x8d, 0xcf, 0x9a, 0x7c, 0x6b, 0xc9, 0xd2, 0x5e,
	0x06, 0x59, 0xae, 0xf9, 0xc6, 0x20, 0xeb, 0xa5, 0xa0, 0x3d, 0xf2, 0x13, 0xf4, 0xf2, 0xac, 0x93,
	0x97, 0x0d, 0x69, 0x6a, 0x17, 0xaa, 0x11, 0xb8, 0x7e, 0x91, 0x18, 0x57, 0x36, 0xc2, 0xe2, 0x86,
	0x2f, 0x0e, 0x79, 0xd6, 0x90, 0xea, 0xda, 0x7e, 0x8d, 0xbb, 0xcc, 0x1b, 0xed, 0x11, 0x1f, 0x86,
	0x65, 0xc8, 0xf9, 0xac, 0x51, 0xe0, 0xca, 0x5a, 0x76, 0x85, 0x8b, 0xf9, 0xe2, 0x15, 0x70, 0xfc,
	0x71, 0x7c, 0xbe, 0x15, 0x71, 0x7d, 0xa9, 0xc7, 0xdb, 0x0a, 0x5c, 0xdb, 0x48, 0xda, 0x23, 0x29,
	0x8c, 0xca, 0x98, 0xf9, 0x43, 0x6b, 0x6e, 0x05, 0x2d, 0xbd, 0x0e, 0x1d, 0x51, 0xbf, 0xcb, 0xc7,
	0x41, 0x4e, 0xd1, 0x93, 0x36, 0x09, 0xfe, 0x03, 0xeb, 0xdb, 0x1e, 0xfb, 0xa3, 0xbf, 0xfa, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x39, 0x69, 0x5d, 0xa0, 0xf2, 0x07, 0x00, 0x00,
}
