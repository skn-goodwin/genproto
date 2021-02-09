// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.13.0
// source: product-manager/v1/product_manager.proto

package product_manager

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Поле resource_id используется в качестве идентификатора REST ресурса в связке с google/api/annotations.proto
	//  string resource_id = 1;
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_manager_v1_product_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_manager_v1_product_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_product_manager_v1_product_manager_proto_rawDescGZIP(), []int{0}
}

func (x *FetchRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Поле resource_id используется в качестве идентификатора REST ресурса в связке с google/api/annotations.proto
	//  string resource_id = 1;
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	PageSize      int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_manager_v1_product_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_manager_v1_product_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_product_manager_v1_product_manager_proto_rawDescGZIP(), []int{1}
}

func (x *ListRequest) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*ListResponse_Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_manager_v1_product_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_manager_v1_product_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_product_manager_v1_product_manager_proto_rawDescGZIP(), []int{2}
}

func (x *ListResponse) GetProducts() []*ListResponse_Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type ListResponse_Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price            string `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"` // Используем тип string для предотвращения проблем с мантиссой
	PriceChangeCount int32  `protobuf:"varint,3,opt,name=price_change_count,json=priceChangeCount,proto3" json:"price_change_count,omitempty"`
	UpdateUnix       int64  `protobuf:"varint,4,opt,name=update_unix,json=updateUnix,proto3" json:"update_unix,omitempty"`
}

func (x *ListResponse_Product) Reset() {
	*x = ListResponse_Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_manager_v1_product_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse_Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse_Product) ProtoMessage() {}

func (x *ListResponse_Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_manager_v1_product_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse_Product.ProtoReflect.Descriptor instead.
func (*ListResponse_Product) Descriptor() ([]byte, []int) {
	return file_product_manager_v1_product_manager_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListResponse_Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListResponse_Product) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *ListResponse_Product) GetPriceChangeCount() int32 {
	if x != nil {
		return x.PriceChangeCount
	}
	return 0
}

func (x *ListResponse_Product) GetUpdateUnix() int64 {
	if x != nil {
		return x.UpdateUnix
	}
	return 0
}

var File_product_manager_v1_product_manager_proto protoreflect.FileDescriptor

var file_product_manager_v1_product_manager_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x61, 0x74, 0x6c, 0x61,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x52, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xe0, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x74,
	0x6c, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x1a, 0x82, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x70, 0x72, 0x69, 0x63, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x32, 0xbe, 0x01, 0x0a, 0x15,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x27,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x59, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xdd, 0x01, 0x0a,
	0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x13,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0xa2, 0x02, 0x05, 0x41, 0x50, 0x4d, 0x56, 0x31, 0xaa, 0x02, 0x18, 0x41, 0x74, 0x6c, 0x61,
	0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x5c, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x1a, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_manager_v1_product_manager_proto_rawDescOnce sync.Once
	file_product_manager_v1_product_manager_proto_rawDescData = file_product_manager_v1_product_manager_proto_rawDesc
)

func file_product_manager_v1_product_manager_proto_rawDescGZIP() []byte {
	file_product_manager_v1_product_manager_proto_rawDescOnce.Do(func() {
		file_product_manager_v1_product_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_manager_v1_product_manager_proto_rawDescData)
	})
	return file_product_manager_v1_product_manager_proto_rawDescData
}

var file_product_manager_v1_product_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_product_manager_v1_product_manager_proto_goTypes = []interface{}{
	(*FetchRequest)(nil),         // 0: atlant.product_manager.v1.FetchRequest
	(*ListRequest)(nil),          // 1: atlant.product_manager.v1.ListRequest
	(*ListResponse)(nil),         // 2: atlant.product_manager.v1.ListResponse
	(*ListResponse_Product)(nil), // 3: atlant.product_manager.v1.ListResponse.Product
	(*emptypb.Empty)(nil),        // 4: google.protobuf.Empty
}
var file_product_manager_v1_product_manager_proto_depIdxs = []int32{
	3, // 0: atlant.product_manager.v1.ListResponse.products:type_name -> atlant.product_manager.v1.ListResponse.Product
	0, // 1: atlant.product_manager.v1.ProductManagerService.Fetch:input_type -> atlant.product_manager.v1.FetchRequest
	1, // 2: atlant.product_manager.v1.ProductManagerService.List:input_type -> atlant.product_manager.v1.ListRequest
	4, // 3: atlant.product_manager.v1.ProductManagerService.Fetch:output_type -> google.protobuf.Empty
	2, // 4: atlant.product_manager.v1.ProductManagerService.List:output_type -> atlant.product_manager.v1.ListResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_product_manager_v1_product_manager_proto_init() }
func file_product_manager_v1_product_manager_proto_init() {
	if File_product_manager_v1_product_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_manager_v1_product_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_product_manager_v1_product_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_product_manager_v1_product_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_product_manager_v1_product_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse_Product); i {
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
			RawDescriptor: file_product_manager_v1_product_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_manager_v1_product_manager_proto_goTypes,
		DependencyIndexes: file_product_manager_v1_product_manager_proto_depIdxs,
		MessageInfos:      file_product_manager_v1_product_manager_proto_msgTypes,
	}.Build()
	File_product_manager_v1_product_manager_proto = out.File
	file_product_manager_v1_product_manager_proto_rawDesc = nil
	file_product_manager_v1_product_manager_proto_goTypes = nil
	file_product_manager_v1_product_manager_proto_depIdxs = nil
}
