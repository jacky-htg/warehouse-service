// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: inventories/delivery_return_message.proto

package inventories

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type DeliveryReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyId  string                  `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	BranchId   string                  `protobuf:"bytes,3,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	BranchName string                  `protobuf:"bytes,4,opt,name=branch_name,json=branchName,proto3" json:"branch_name,omitempty"`
	Delivery   *Delivery               `protobuf:"bytes,5,opt,name=delivery,proto3" json:"delivery,omitempty"`
	Code       string                  `protobuf:"bytes,6,opt,name=code,proto3" json:"code,omitempty"`
	ReturnDate *timestamp.Timestamp    `protobuf:"bytes,7,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
	Remark     string                  `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	CreatedAt  *timestamp.Timestamp    `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy  string                  `protobuf:"bytes,10,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt  *timestamp.Timestamp    `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy  string                  `protobuf:"bytes,12,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	Details    []*DeliveryReturnDetail `protobuf:"bytes,13,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *DeliveryReturn) Reset() {
	*x = DeliveryReturn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventories_delivery_return_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryReturn) ProtoMessage() {}

func (x *DeliveryReturn) ProtoReflect() protoreflect.Message {
	mi := &file_inventories_delivery_return_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryReturn.ProtoReflect.Descriptor instead.
func (*DeliveryReturn) Descriptor() ([]byte, []int) {
	return file_inventories_delivery_return_message_proto_rawDescGZIP(), []int{0}
}

func (x *DeliveryReturn) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeliveryReturn) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *DeliveryReturn) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *DeliveryReturn) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

func (x *DeliveryReturn) GetDelivery() *Delivery {
	if x != nil {
		return x.Delivery
	}
	return nil
}

func (x *DeliveryReturn) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DeliveryReturn) GetReturnDate() *timestamp.Timestamp {
	if x != nil {
		return x.ReturnDate
	}
	return nil
}

func (x *DeliveryReturn) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *DeliveryReturn) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DeliveryReturn) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *DeliveryReturn) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *DeliveryReturn) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *DeliveryReturn) GetDetails() []*DeliveryReturnDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_inventories_delivery_return_message_proto protoreflect.FileDescriptor

var file_inventories_delivery_return_message_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x77, 0x69, 0x72,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x30, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x04, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x77, 0x69, 0x72,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x44, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x47, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x77,
	0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x65, 0x72, 0x70, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x50, 0x01, 0x5a, 0x1a, 0x70, 0x62, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x3b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventories_delivery_return_message_proto_rawDescOnce sync.Once
	file_inventories_delivery_return_message_proto_rawDescData = file_inventories_delivery_return_message_proto_rawDesc
)

func file_inventories_delivery_return_message_proto_rawDescGZIP() []byte {
	file_inventories_delivery_return_message_proto_rawDescOnce.Do(func() {
		file_inventories_delivery_return_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventories_delivery_return_message_proto_rawDescData)
	})
	return file_inventories_delivery_return_message_proto_rawDescData
}

var file_inventories_delivery_return_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_inventories_delivery_return_message_proto_goTypes = []interface{}{
	(*DeliveryReturn)(nil),       // 0: wiradata.inventories.DeliveryReturn
	(*Delivery)(nil),             // 1: wiradata.inventories.Delivery
	(*timestamp.Timestamp)(nil),  // 2: google.protobuf.Timestamp
	(*DeliveryReturnDetail)(nil), // 3: wiradata.inventories.DeliveryReturnDetail
}
var file_inventories_delivery_return_message_proto_depIdxs = []int32{
	1, // 0: wiradata.inventories.DeliveryReturn.delivery:type_name -> wiradata.inventories.Delivery
	2, // 1: wiradata.inventories.DeliveryReturn.return_date:type_name -> google.protobuf.Timestamp
	2, // 2: wiradata.inventories.DeliveryReturn.created_at:type_name -> google.protobuf.Timestamp
	2, // 3: wiradata.inventories.DeliveryReturn.updated_at:type_name -> google.protobuf.Timestamp
	3, // 4: wiradata.inventories.DeliveryReturn.details:type_name -> wiradata.inventories.DeliveryReturnDetail
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_inventories_delivery_return_message_proto_init() }
func file_inventories_delivery_return_message_proto_init() {
	if File_inventories_delivery_return_message_proto != nil {
		return
	}
	file_inventories_delivery_return_detail_message_proto_init()
	file_inventories_delivery_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_inventories_delivery_return_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryReturn); i {
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
			RawDescriptor: file_inventories_delivery_return_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_inventories_delivery_return_message_proto_goTypes,
		DependencyIndexes: file_inventories_delivery_return_message_proto_depIdxs,
		MessageInfos:      file_inventories_delivery_return_message_proto_msgTypes,
	}.Build()
	File_inventories_delivery_return_message_proto = out.File
	file_inventories_delivery_return_message_proto_rawDesc = nil
	file_inventories_delivery_return_message_proto_goTypes = nil
	file_inventories_delivery_return_message_proto_depIdxs = nil
}
