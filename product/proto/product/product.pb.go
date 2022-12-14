// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/product/product.proto

/*
Package go_micro_service_product is a generated protocol buffer package.

It is generated from these files:
	proto/product/product.proto

It has these top-level messages:
	ProductInfo
	ProductImage
	ProductSize
	ProductSeo
	ResponseProduct
	RequestID
	Response
	RequestAll
	AllProduct
*/
package go_micro_service_product

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

// 定义一个商品信息的消息
// message是固定的关键字
// ProductInfo是自定义类名，后面会使用到ProductInfo。
type ProductInfo struct {
	Id                 int64           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ProductName        string          `protobuf:"bytes,2,opt,name=product_name,json=productName" json:"product_name,omitempty"`
	ProductSku         string          `protobuf:"bytes,3,opt,name=product_sku,json=productSku" json:"product_sku,omitempty"`
	ProductPrice       float64         `protobuf:"fixed64,4,opt,name=product_price,json=productPrice" json:"product_price,omitempty"`
	ProductDescription string          `protobuf:"bytes,5,opt,name=product_description,json=productDescription" json:"product_description,omitempty"`
	ProductCategoryId  int64           `protobuf:"varint,6,opt,name=product_category_id,json=productCategoryId" json:"product_category_id,omitempty"`
	ProductImage       []*ProductImage `protobuf:"bytes,7,rep,name=product_image,json=productImage" json:"product_image,omitempty"`
	ProductSize        []*ProductSize  `protobuf:"bytes,8,rep,name=product_size,json=productSize" json:"product_size,omitempty"`
	ProductSeo         *ProductSeo     `protobuf:"bytes,9,opt,name=product_seo,json=productSeo" json:"product_seo,omitempty"`
}

func (m *ProductInfo) Reset()                    { *m = ProductInfo{} }
func (m *ProductInfo) String() string            { return proto.CompactTextString(m) }
func (*ProductInfo) ProtoMessage()               {}
func (*ProductInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProductInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductInfo) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *ProductInfo) GetProductSku() string {
	if m != nil {
		return m.ProductSku
	}
	return ""
}

func (m *ProductInfo) GetProductPrice() float64 {
	if m != nil {
		return m.ProductPrice
	}
	return 0
}

func (m *ProductInfo) GetProductDescription() string {
	if m != nil {
		return m.ProductDescription
	}
	return ""
}

func (m *ProductInfo) GetProductCategoryId() int64 {
	if m != nil {
		return m.ProductCategoryId
	}
	return 0
}

func (m *ProductInfo) GetProductImage() []*ProductImage {
	if m != nil {
		return m.ProductImage
	}
	return nil
}

func (m *ProductInfo) GetProductSize() []*ProductSize {
	if m != nil {
		return m.ProductSize
	}
	return nil
}

func (m *ProductInfo) GetProductSeo() *ProductSeo {
	if m != nil {
		return m.ProductSeo
	}
	return nil
}

// 定义一个商品图片的消息
type ProductImage struct {
	Id        int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ImageName string `protobuf:"bytes,2,opt,name=image_name,json=imageName" json:"image_name,omitempty"`
	ImageCode string `protobuf:"bytes,3,opt,name=image_code,json=imageCode" json:"image_code,omitempty"`
	ImageUrl  string `protobuf:"bytes,4,opt,name=image_url,json=imageUrl" json:"image_url,omitempty"`
}

func (m *ProductImage) Reset()                    { *m = ProductImage{} }
func (m *ProductImage) String() string            { return proto.CompactTextString(m) }
func (*ProductImage) ProtoMessage()               {}
func (*ProductImage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProductImage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductImage) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *ProductImage) GetImageCode() string {
	if m != nil {
		return m.ImageCode
	}
	return ""
}

func (m *ProductImage) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

// 定义一个商品尺寸的消息
type ProductSize struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SizeName string `protobuf:"bytes,2,opt,name=size_name,json=sizeName" json:"size_name,omitempty"`
	SizeCode string `protobuf:"bytes,3,opt,name=size_code,json=sizeCode" json:"size_code,omitempty"`
}

func (m *ProductSize) Reset()                    { *m = ProductSize{} }
func (m *ProductSize) String() string            { return proto.CompactTextString(m) }
func (*ProductSize) ProtoMessage()               {}
func (*ProductSize) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProductSize) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductSize) GetSizeName() string {
	if m != nil {
		return m.SizeName
	}
	return ""
}

func (m *ProductSize) GetSizeCode() string {
	if m != nil {
		return m.SizeCode
	}
	return ""
}

// 定义一个商品seo的消息
type ProductSeo struct {
	Id             int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SeoTitle       string `protobuf:"bytes,2,opt,name=seo_title,json=seoTitle" json:"seo_title,omitempty"`
	SeoKeywords    string `protobuf:"bytes,3,opt,name=seo_keywords,json=seoKeywords" json:"seo_keywords,omitempty"`
	SeoDescription string `protobuf:"bytes,4,opt,name=seo_description,json=seoDescription" json:"seo_description,omitempty"`
	SeoCode        string `protobuf:"bytes,5,opt,name=seo_code,json=seoCode" json:"seo_code,omitempty"`
}

func (m *ProductSeo) Reset()                    { *m = ProductSeo{} }
func (m *ProductSeo) String() string            { return proto.CompactTextString(m) }
func (*ProductSeo) ProtoMessage()               {}
func (*ProductSeo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ProductSeo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductSeo) GetSeoTitle() string {
	if m != nil {
		return m.SeoTitle
	}
	return ""
}

func (m *ProductSeo) GetSeoKeywords() string {
	if m != nil {
		return m.SeoKeywords
	}
	return ""
}

func (m *ProductSeo) GetSeoDescription() string {
	if m != nil {
		return m.SeoDescription
	}
	return ""
}

func (m *ProductSeo) GetSeoCode() string {
	if m != nil {
		return m.SeoCode
	}
	return ""
}

// 定义一个响应商品的消息
type ResponseProduct struct {
	ProductId int64 `protobuf:"varint,1,opt,name=product_id,json=productId" json:"product_id,omitempty"`
}

func (m *ResponseProduct) Reset()                    { *m = ResponseProduct{} }
func (m *ResponseProduct) String() string            { return proto.CompactTextString(m) }
func (*ResponseProduct) ProtoMessage()               {}
func (*ResponseProduct) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ResponseProduct) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

// 定义一个请求ID的消息
type RequestID struct {
	ProductId int64 `protobuf:"varint,1,opt,name=product_id,json=productId" json:"product_id,omitempty"`
}

func (m *RequestID) Reset()                    { *m = RequestID{} }
func (m *RequestID) String() string            { return proto.CompactTextString(m) }
func (*RequestID) ProtoMessage()               {}
func (*RequestID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RequestID) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

// 定义一个响应的消息
type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// 定义一个返回所有的消息
type RequestAll struct {
}

func (m *RequestAll) Reset()                    { *m = RequestAll{} }
func (m *RequestAll) String() string            { return proto.CompactTextString(m) }
func (*RequestAll) ProtoMessage()               {}
func (*RequestAll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// 定义一个所有商品的消息
type AllProduct struct {
	ProductInfo []*ProductInfo `protobuf:"bytes,1,rep,name=product_info,json=productInfo" json:"product_info,omitempty"`
}

func (m *AllProduct) Reset()                    { *m = AllProduct{} }
func (m *AllProduct) String() string            { return proto.CompactTextString(m) }
func (*AllProduct) ProtoMessage()               {}
func (*AllProduct) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AllProduct) GetProductInfo() []*ProductInfo {
	if m != nil {
		return m.ProductInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*ProductInfo)(nil), "go.micro.service.product.ProductInfo")
	proto.RegisterType((*ProductImage)(nil), "go.micro.service.product.ProductImage")
	proto.RegisterType((*ProductSize)(nil), "go.micro.service.product.ProductSize")
	proto.RegisterType((*ProductSeo)(nil), "go.micro.service.product.ProductSeo")
	proto.RegisterType((*ResponseProduct)(nil), "go.micro.service.product.ResponseProduct")
	proto.RegisterType((*RequestID)(nil), "go.micro.service.product.RequestID")
	proto.RegisterType((*Response)(nil), "go.micro.service.product.Response")
	proto.RegisterType((*RequestAll)(nil), "go.micro.service.product.RequestAll")
	proto.RegisterType((*AllProduct)(nil), "go.micro.service.product.AllProduct")
}

func init() { proto.RegisterFile("proto/product/product.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x6b, 0x68, 0xe3, 0x49, 0x2f, 0x74, 0x79, 0x31, 0x2d, 0x88, 0xb0, 0x2d, 0x10, 0x78,
	0x70, 0x51, 0xf8, 0x82, 0xd0, 0x80, 0x88, 0x2a, 0xa1, 0xca, 0xa5, 0xf0, 0x82, 0x30, 0xc1, 0x3b,
	0x8d, 0x56, 0x71, 0xbc, 0xc6, 0xeb, 0x80, 0xd2, 0xaf, 0xe1, 0x67, 0xf8, 0x19, 0xbe, 0x02, 0xed,
	0x7a, 0xd7, 0xb6, 0x0a, 0x6d, 0xc2, 0x53, 0x3c, 0x17, 0x9f, 0x39, 0x73, 0xe6, 0x28, 0x86, 0xfd,
	0x2c, 0x17, 0x85, 0x38, 0xca, 0x72, 0xc1, 0xe6, 0x71, 0x61, 0x7f, 0x03, 0x9d, 0x25, 0xfe, 0x44,
	0x04, 0x33, 0x1e, 0xe7, 0x22, 0x90, 0x98, 0x7f, 0xe7, 0x31, 0x06, 0xa6, 0x4e, 0x7f, 0xb9, 0xd0,
	0x39, 0x2d, 0x9f, 0x47, 0xe9, 0x85, 0x20, 0xdb, 0xb0, 0xc6, 0x99, 0xef, 0x74, 0x9d, 0x9e, 0x1b,
	0xae, 0x71, 0x46, 0x1e, 0xc1, 0xa6, 0x69, 0x8d, 0xd2, 0xf1, 0x0c, 0xfd, 0xb5, 0xae, 0xd3, 0xf3,
	0xc2, 0x8e, 0xc9, 0xbd, 0x1b, 0xcf, 0x90, 0x3c, 0x04, 0x1b, 0x46, 0x72, 0x3a, 0xf7, 0x5d, 0xdd,
	0x01, 0x26, 0x75, 0x36, 0x9d, 0x93, 0x03, 0xd8, 0xb2, 0x0d, 0x59, 0xce, 0x63, 0xf4, 0x6f, 0x75,
	0x9d, 0x9e, 0x13, 0x5a, 0xe0, 0x53, 0x95, 0x23, 0x47, 0x70, 0xd7, 0x36, 0x31, 0x94, 0x71, 0xce,
	0xb3, 0x82, 0x8b, 0xd4, 0xbf, 0xad, 0xd1, 0x88, 0x29, 0x0d, 0xeb, 0x0a, 0x09, 0xea, 0x17, 0xe2,
	0x71, 0x81, 0x13, 0x91, 0x2f, 0x22, 0xce, 0xfc, 0x75, 0x4d, 0x7d, 0xd7, 0x94, 0x8e, 0x4d, 0x65,
	0xc4, 0xc8, 0x49, 0xcd, 0x82, 0xcf, 0xc6, 0x13, 0xf4, 0x37, 0xba, 0x6e, 0xaf, 0xd3, 0x7f, 0x12,
	0x5c, 0xa7, 0x4d, 0x60, 0x75, 0x51, 0xdd, 0x15, 0x5b, 0x1d, 0x91, 0xb7, 0xb5, 0x2c, 0x92, 0x5f,
	0xa2, 0xdf, 0xd6, 0x58, 0x8f, 0x97, 0x62, 0x9d, 0xf1, 0x4b, 0xac, 0xd4, 0x53, 0x01, 0x79, 0xdd,
	0x50, 0x0f, 0x85, 0xef, 0x75, 0x9d, 0x5e, 0xa7, 0x7f, 0xb8, 0x1c, 0x08, 0x45, 0xad, 0x31, 0x0a,
	0xba, 0x80, 0xcd, 0x26, 0xdd, 0xbf, 0xee, 0xf8, 0x00, 0x40, 0x6f, 0xdd, 0xbc, 0xa2, 0xa7, 0x33,
	0xfa, 0x86, 0x55, 0x39, 0x16, 0x0c, 0xcd, 0x09, 0xcb, 0xf2, 0xb1, 0x60, 0x48, 0xf6, 0xa1, 0x0c,
	0xa2, 0x79, 0x9e, 0xe8, 0xeb, 0x79, 0x61, 0x5b, 0x27, 0xce, 0xf3, 0x84, 0x7e, 0xac, 0x1c, 0xa4,
	0x17, 0xba, 0x3a, 0x79, 0x1f, 0x3c, 0x25, 0x51, 0x73, 0x70, 0x5b, 0x25, 0xf4, 0x5c, 0x5b, 0x6c,
	0x8c, 0xd5, 0x45, 0x35, 0x95, 0xfe, 0x74, 0x00, 0xea, 0x75, 0xff, 0x09, 0x8c, 0x22, 0x2a, 0x78,
	0x91, 0xd4, 0xc0, 0x28, 0xde, 0xab, 0x58, 0xf9, 0x56, 0x15, 0xa7, 0xb8, 0xf8, 0x21, 0x72, 0x26,
	0x0d, 0x76, 0x47, 0xa2, 0x38, 0x31, 0x29, 0xf2, 0x14, 0x76, 0x54, 0x4b, 0xd3, 0x6d, 0xe5, 0x6a,
	0xdb, 0x12, 0x45, 0xd3, 0x69, 0xf7, 0x40, 0xe1, 0x96, 0x1c, 0x4b, 0x3f, 0x6e, 0x48, 0x14, 0x9a,
	0xe2, 0x0b, 0xd8, 0x09, 0x51, 0x66, 0x22, 0x95, 0x68, 0x98, 0x2a, 0x29, 0x2b, 0x9f, 0x59, 0xba,
	0x9e, 0x35, 0x0f, 0xa3, 0xcf, 0xc1, 0x0b, 0xf1, 0xdb, 0x1c, 0x65, 0x31, 0x1a, 0x2e, 0xeb, 0xbd,
	0x0f, 0x6d, 0x8b, 0x4e, 0xee, 0x80, 0x3b, 0x93, 0x13, 0xdd, 0xe3, 0x85, 0xea, 0x91, 0x6e, 0x02,
	0x18, 0xa4, 0x41, 0x92, 0xd0, 0x0f, 0x00, 0x83, 0x24, 0xb1, 0x24, 0x1a, 0xfe, 0xe4, 0xe9, 0x85,
	0xf0, 0x9d, 0x15, 0xfd, 0xa9, 0xfe, 0x03, 0x2a, 0x7f, 0xaa, 0xa0, 0xff, 0xdb, 0x85, 0x0d, 0x8b,
	0xfa, 0x05, 0x60, 0xc0, 0x98, 0x8d, 0x56, 0x43, 0xdb, 0x7b, 0x76, 0x7d, 0xdb, 0x15, 0xe9, 0x68,
	0x8b, 0x44, 0xb0, 0xf3, 0x86, 0xa7, 0x76, 0xc4, 0xab, 0xc5, 0x68, 0x48, 0x0e, 0x6e, 0x7a, 0xdf,
	0x08, 0xb9, 0xb7, 0x1a, 0x17, 0xda, 0x22, 0x9f, 0x60, 0xeb, 0x3c, 0x63, 0xe3, 0x02, 0xff, 0x73,
	0x0b, 0xba, 0x7c, 0x0b, 0xda, 0x22, 0x9f, 0x61, 0x77, 0x88, 0x09, 0x56, 0xe8, 0xab, 0x2f, 0xb0,
	0x2a, 0xfe, 0xb6, 0x92, 0xa7, 0x71, 0xe8, 0xc3, 0xa5, 0xe0, 0x83, 0x24, 0xd9, 0xbb, 0xa1, 0xab,
	0xc6, 0xa2, 0xad, 0xaf, 0xeb, 0xfa, 0x73, 0xf1, 0xf2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9,
	0x1f, 0x5c, 0x35, 0x4d, 0x06, 0x00, 0x00,
}
