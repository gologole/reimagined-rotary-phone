// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: api/proto/dborder.proto

package protogen

import (
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

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Desc      string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Price     int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	OverPrice int32  `protobuf:"varint,4,opt,name=overPrice,proto3" json:"overPrice,omitempty"`
	Date      string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	CourierId int32  `protobuf:"varint,6,opt,name=courierId,proto3" json:"courierId,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_dborder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_dborder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_dborder_proto_rawDescGZIP(), []int{0}
}

func (x *OrderResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderResponse) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *OrderResponse) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderResponse) GetOverPrice() int32 {
	if x != nil {
		return x.OverPrice
	}
	return 0
}

func (x *OrderResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *OrderResponse) GetCourierId() int32 {
	if x != nil {
		return x.CourierId
	}
	return 0
}

type GetAllOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllOrdersRequest) Reset() {
	*x = GetAllOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_dborder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOrdersRequest) ProtoMessage() {}

func (x *GetAllOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_dborder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllOrdersRequest.ProtoReflect.Descriptor instead.
func (*GetAllOrdersRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_dborder_proto_rawDescGZIP(), []int{1}
}

type GetAllOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*OrderResponse `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetAllOrdersResponse) Reset() {
	*x = GetAllOrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_dborder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOrdersResponse) ProtoMessage() {}

func (x *GetAllOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_dborder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllOrdersResponse.ProtoReflect.Descriptor instead.
func (*GetAllOrdersResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_dborder_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllOrdersResponse) GetOrders() []*OrderResponse {
	if x != nil {
		return x.Orders
	}
	return nil
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Desc      string `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	Price     int32  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	OverPrice int32  `protobuf:"varint,3,opt,name=overPrice,proto3" json:"overPrice,omitempty"`
	CourierId int32  `protobuf:"varint,4,opt,name=courierId,proto3" json:"courierId,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_dborder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_dborder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_dborder_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrderRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CreateOrderRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateOrderRequest) GetOverPrice() int32 {
	if x != nil {
		return x.OverPrice
	}
	return 0
}

func (x *CreateOrderRequest) GetCourierId() int32 {
	if x != nil {
		return x.CourierId
	}
	return 0
}

type DeleteOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteOrderRequest) Reset() {
	*x = DeleteOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_dborder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderRequest) ProtoMessage() {}

func (x *DeleteOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_dborder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrderRequest.ProtoReflect.Descriptor instead.
func (*DeleteOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_dborder_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteOrderRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_proto_dborder_proto protoreflect.FileDescriptor

var file_api_proto_dborder_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x62, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x22, 0x99, 0x01, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6f, 0x76, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x7a, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x76, 0x65,
	0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6f, 0x76,
	0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x49, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0xd0, 0x01, 0x0a, 0x05,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_dborder_proto_rawDescOnce sync.Once
	file_api_proto_dborder_proto_rawDescData = file_api_proto_dborder_proto_rawDesc
)

func file_api_proto_dborder_proto_rawDescGZIP() []byte {
	file_api_proto_dborder_proto_rawDescOnce.Do(func() {
		file_api_proto_dborder_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_dborder_proto_rawDescData)
	})
	return file_api_proto_dborder_proto_rawDescData
}

var file_api_proto_dborder_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_dborder_proto_goTypes = []any{
	(*OrderResponse)(nil),        // 0: order.OrderResponse
	(*GetAllOrdersRequest)(nil),  // 1: order.GetAllOrdersRequest
	(*GetAllOrdersResponse)(nil), // 2: order.GetAllOrdersResponse
	(*CreateOrderRequest)(nil),   // 3: order.CreateOrderRequest
	(*DeleteOrderRequest)(nil),   // 4: order.DeleteOrderRequest
}
var file_api_proto_dborder_proto_depIdxs = []int32{
	0, // 0: order.GetAllOrdersResponse.orders:type_name -> order.OrderResponse
	1, // 1: order.Order.GetAllOrders:input_type -> order.GetAllOrdersRequest
	3, // 2: order.Order.CreateOrder:input_type -> order.CreateOrderRequest
	4, // 3: order.Order.DeleteOrder:input_type -> order.DeleteOrderRequest
	2, // 4: order.Order.GetAllOrders:output_type -> order.GetAllOrdersResponse
	0, // 5: order.Order.CreateOrder:output_type -> order.OrderResponse
	0, // 6: order.Order.DeleteOrder:output_type -> order.OrderResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_dborder_proto_init() }
func file_api_proto_dborder_proto_init() {
	if File_api_proto_dborder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_dborder_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*OrderResponse); i {
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
		file_api_proto_dborder_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllOrdersRequest); i {
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
		file_api_proto_dborder_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllOrdersResponse); i {
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
		file_api_proto_dborder_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOrderRequest); i {
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
		file_api_proto_dborder_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteOrderRequest); i {
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
			RawDescriptor: file_api_proto_dborder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_dborder_proto_goTypes,
		DependencyIndexes: file_api_proto_dborder_proto_depIdxs,
		MessageInfos:      file_api_proto_dborder_proto_msgTypes,
	}.Build()
	File_api_proto_dborder_proto = out.File
	file_api_proto_dborder_proto_rawDesc = nil
	file_api_proto_dborder_proto_goTypes = nil
	file_api_proto_dborder_proto_depIdxs = nil
}
