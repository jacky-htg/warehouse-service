// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: inventories/stock_service.proto

package inventories

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ClosingStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year  int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
}

func (x *ClosingStockRequest) Reset() {
	*x = ClosingStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventories_stock_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClosingStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClosingStockRequest) ProtoMessage() {}

func (x *ClosingStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventories_stock_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClosingStockRequest.ProtoReflect.Descriptor instead.
func (*ClosingStockRequest) Descriptor() ([]byte, []int) {
	return file_inventories_stock_service_proto_rawDescGZIP(), []int{0}
}

func (x *ClosingStockRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *ClosingStockRequest) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

type StockListInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId string `protobuf:"bytes,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
}

func (x *StockListInput) Reset() {
	*x = StockListInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventories_stock_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockListInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockListInput) ProtoMessage() {}

func (x *StockListInput) ProtoReflect() protoreflect.Message {
	mi := &file_inventories_stock_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockListInput.ProtoReflect.Descriptor instead.
func (*StockListInput) Descriptor() ([]byte, []int) {
	return file_inventories_stock_service_proto_rawDescGZIP(), []int{1}
}

func (x *StockListInput) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

type StockInfoInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId  string `protobuf:"bytes,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *StockInfoInput) Reset() {
	*x = StockInfoInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventories_stock_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockInfoInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockInfoInput) ProtoMessage() {}

func (x *StockInfoInput) ProtoReflect() protoreflect.Message {
	mi := &file_inventories_stock_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockInfoInput.ProtoReflect.Descriptor instead.
func (*StockInfoInput) Descriptor() ([]byte, []int) {
	return file_inventories_stock_service_proto_rawDescGZIP(), []int{2}
}

func (x *StockInfoInput) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *StockInfoInput) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type StockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Qty     int32    `protobuf:"varint,2,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *StockInfo) Reset() {
	*x = StockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventories_stock_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockInfo) ProtoMessage() {}

func (x *StockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_inventories_stock_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockInfo.ProtoReflect.Descriptor instead.
func (*StockInfo) Descriptor() ([]byte, []int) {
	return file_inventories_stock_service_proto_rawDescGZIP(), []int{3}
}

func (x *StockInfo) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *StockInfo) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

type StockList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockInfo []*StockInfo `protobuf:"bytes,1,rep,name=stock_info,json=stockInfo,proto3" json:"stock_info,omitempty"`
}

func (x *StockList) Reset() {
	*x = StockList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventories_stock_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockList) ProtoMessage() {}

func (x *StockList) ProtoReflect() protoreflect.Message {
	mi := &file_inventories_stock_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockList.ProtoReflect.Descriptor instead.
func (*StockList) Descriptor() ([]byte, []int) {
	return file_inventories_stock_service_proto_rawDescGZIP(), []int{4}
}

func (x *StockList) GetStockInfo() []*StockInfo {
	if x != nil {
		return x.StockInfo
	}
	return nil
}

var File_inventories_stock_service_proto protoreflect.FileDescriptor

var file_inventories_stock_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x21, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a,
	0x13, 0x43, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x22, 0x2d,
	0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x22, 0x4c, 0x0a,
	0x0e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x09, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x37, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x69, 0x72, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x71, 0x74, 0x79, 0x22, 0x4b, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x32, 0x89, 0x02, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x57, 0x0a, 0x07, 0x43, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x2e, 0x77,
	0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x4d,
	0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x24, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1f, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x24, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1f, 0x2e, 0x77, 0x69, 0x72, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x47, 0x0a, 0x27,
	0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x65, 0x72, 0x70,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x1a, 0x70, 0x62, 0x2f, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x3b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventories_stock_service_proto_rawDescOnce sync.Once
	file_inventories_stock_service_proto_rawDescData = file_inventories_stock_service_proto_rawDesc
)

func file_inventories_stock_service_proto_rawDescGZIP() []byte {
	file_inventories_stock_service_proto_rawDescOnce.Do(func() {
		file_inventories_stock_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventories_stock_service_proto_rawDescData)
	})
	return file_inventories_stock_service_proto_rawDescData
}

var file_inventories_stock_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_inventories_stock_service_proto_goTypes = []interface{}{
	(*ClosingStockRequest)(nil), // 0: wiradata.inventories.ClosingStockRequest
	(*StockListInput)(nil),      // 1: wiradata.inventories.StockListInput
	(*StockInfoInput)(nil),      // 2: wiradata.inventories.StockInfoInput
	(*StockInfo)(nil),           // 3: wiradata.inventories.StockInfo
	(*StockList)(nil),           // 4: wiradata.inventories.StockList
	(*Product)(nil),             // 5: wiradata.inventories.Product
	(*MyBoolean)(nil),           // 6: wiradata.inventories.MyBoolean
}
var file_inventories_stock_service_proto_depIdxs = []int32{
	5, // 0: wiradata.inventories.StockInfo.product:type_name -> wiradata.inventories.Product
	3, // 1: wiradata.inventories.StockList.stock_info:type_name -> wiradata.inventories.StockInfo
	0, // 2: wiradata.inventories.StockService.Closing:input_type -> wiradata.inventories.ClosingStockRequest
	1, // 3: wiradata.inventories.StockService.List:input_type -> wiradata.inventories.StockListInput
	2, // 4: wiradata.inventories.StockService.Info:input_type -> wiradata.inventories.StockInfoInput
	6, // 5: wiradata.inventories.StockService.Closing:output_type -> wiradata.inventories.MyBoolean
	4, // 6: wiradata.inventories.StockService.List:output_type -> wiradata.inventories.StockList
	3, // 7: wiradata.inventories.StockService.Info:output_type -> wiradata.inventories.StockInfo
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_inventories_stock_service_proto_init() }
func file_inventories_stock_service_proto_init() {
	if File_inventories_stock_service_proto != nil {
		return
	}
	file_inventories_generic_message_proto_init()
	file_inventories_product_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_inventories_stock_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClosingStockRequest); i {
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
		file_inventories_stock_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockListInput); i {
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
		file_inventories_stock_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockInfoInput); i {
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
		file_inventories_stock_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockInfo); i {
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
		file_inventories_stock_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockList); i {
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
			RawDescriptor: file_inventories_stock_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventories_stock_service_proto_goTypes,
		DependencyIndexes: file_inventories_stock_service_proto_depIdxs,
		MessageInfos:      file_inventories_stock_service_proto_msgTypes,
	}.Build()
	File_inventories_stock_service_proto = out.File
	file_inventories_stock_service_proto_rawDesc = nil
	file_inventories_stock_service_proto_goTypes = nil
	file_inventories_stock_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StockServiceClient is the client API for StockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StockServiceClient interface {
	Closing(ctx context.Context, in *ClosingStockRequest, opts ...grpc.CallOption) (*MyBoolean, error)
	List(ctx context.Context, in *StockListInput, opts ...grpc.CallOption) (*StockList, error)
	Info(ctx context.Context, in *StockInfoInput, opts ...grpc.CallOption) (*StockInfo, error)
}

type stockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStockServiceClient(cc grpc.ClientConnInterface) StockServiceClient {
	return &stockServiceClient{cc}
}

func (c *stockServiceClient) Closing(ctx context.Context, in *ClosingStockRequest, opts ...grpc.CallOption) (*MyBoolean, error) {
	out := new(MyBoolean)
	err := c.cc.Invoke(ctx, "/wiradata.inventories.StockService/Closing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) List(ctx context.Context, in *StockListInput, opts ...grpc.CallOption) (*StockList, error) {
	out := new(StockList)
	err := c.cc.Invoke(ctx, "/wiradata.inventories.StockService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) Info(ctx context.Context, in *StockInfoInput, opts ...grpc.CallOption) (*StockInfo, error) {
	out := new(StockInfo)
	err := c.cc.Invoke(ctx, "/wiradata.inventories.StockService/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServiceServer is the server API for StockService service.
type StockServiceServer interface {
	Closing(context.Context, *ClosingStockRequest) (*MyBoolean, error)
	List(context.Context, *StockListInput) (*StockList, error)
	Info(context.Context, *StockInfoInput) (*StockInfo, error)
}

// UnimplementedStockServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStockServiceServer struct {
}

func (*UnimplementedStockServiceServer) Closing(context.Context, *ClosingStockRequest) (*MyBoolean, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Closing not implemented")
}
func (*UnimplementedStockServiceServer) List(context.Context, *StockListInput) (*StockList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedStockServiceServer) Info(context.Context, *StockInfoInput) (*StockInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}

func RegisterStockServiceServer(s *grpc.Server, srv StockServiceServer) {
	s.RegisterService(&_StockService_serviceDesc, srv)
}

func _StockService_Closing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClosingStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).Closing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.inventories.StockService/Closing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).Closing(ctx, req.(*ClosingStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockListInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.inventories.StockService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).List(ctx, req.(*StockListInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockInfoInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.inventories.StockService/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).Info(ctx, req.(*StockInfoInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _StockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wiradata.inventories.StockService",
	HandlerType: (*StockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Closing",
			Handler:    _StockService_Closing_Handler,
		},
		{
			MethodName: "List",
			Handler:    _StockService_List_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _StockService_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventories/stock_service.proto",
}
