// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: v1/shipping.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ShipOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ShipOrderReq) Reset() {
	*x = ShipOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_shipping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipOrderReq) ProtoMessage() {}

func (x *ShipOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_shipping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipOrderReq.ProtoReflect.Descriptor instead.
func (*ShipOrderReq) Descriptor() ([]byte, []int) {
	return file_v1_shipping_proto_rawDescGZIP(), []int{0}
}

func (x *ShipOrderReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ShipOrderReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ShipOrderReply) Reset() {
	*x = ShipOrderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_shipping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipOrderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipOrderReply) ProtoMessage() {}

func (x *ShipOrderReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_shipping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipOrderReply.ProtoReflect.Descriptor instead.
func (*ShipOrderReply) Descriptor() ([]byte, []int) {
	return file_v1_shipping_proto_rawDescGZIP(), []int{1}
}

func (x *ShipOrderReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_v1_shipping_proto protoreflect.FileDescriptor

var file_v1_shipping_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x0c, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x61, 0x0a, 0x08, 0x53, 0x68, 0x69, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x12, 0x55, 0x0a, 0x09, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x21, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v1_shipping_proto_rawDescOnce sync.Once
	file_v1_shipping_proto_rawDescData = file_v1_shipping_proto_rawDesc
)

func file_v1_shipping_proto_rawDescGZIP() []byte {
	file_v1_shipping_proto_rawDescOnce.Do(func() {
		file_v1_shipping_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_shipping_proto_rawDescData)
	})
	return file_v1_shipping_proto_rawDescData
}

var file_v1_shipping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_shipping_proto_goTypes = []interface{}{
	(*ShipOrderReq)(nil),   // 0: shipping.service.v1.ShipOrderReq
	(*ShipOrderReply)(nil), // 1: shipping.service.v1.ShipOrderReply
}
var file_v1_shipping_proto_depIdxs = []int32{
	0, // 0: shipping.service.v1.Shipping.ShipOrder:input_type -> shipping.service.v1.ShipOrderReq
	1, // 1: shipping.service.v1.Shipping.ShipOrder:output_type -> shipping.service.v1.ShipOrderReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_shipping_proto_init() }
func file_v1_shipping_proto_init() {
	if File_v1_shipping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_shipping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipOrderReq); i {
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
		file_v1_shipping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipOrderReply); i {
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
			RawDescriptor: file_v1_shipping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_shipping_proto_goTypes,
		DependencyIndexes: file_v1_shipping_proto_depIdxs,
		MessageInfos:      file_v1_shipping_proto_msgTypes,
	}.Build()
	File_v1_shipping_proto = out.File
	file_v1_shipping_proto_rawDesc = nil
	file_v1_shipping_proto_goTypes = nil
	file_v1_shipping_proto_depIdxs = nil
}
