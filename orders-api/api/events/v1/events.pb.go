// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: v1/events.proto

package v1events

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderBooked struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *OrderBooked) Reset() {
	*x = OrderBooked{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderBooked) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBooked) ProtoMessage() {}

func (x *OrderBooked) ProtoReflect() protoreflect.Message {
	mi := &file_v1_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderBooked.ProtoReflect.Descriptor instead.
func (*OrderBooked) Descriptor() ([]byte, []int) {
	return file_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *OrderBooked) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderBooked) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderBooked) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_v1_events_proto protoreflect.FileDescriptor

var file_v1_events_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x65,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x52,
	0x42, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x73, 0x72,
	0x61, 0x64, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_events_proto_rawDescOnce sync.Once
	file_v1_events_proto_rawDescData = file_v1_events_proto_rawDesc
)

func file_v1_events_proto_rawDescGZIP() []byte {
	file_v1_events_proto_rawDescOnce.Do(func() {
		file_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_events_proto_rawDescData)
	})
	return file_v1_events_proto_rawDescData
}

var file_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v1_events_proto_goTypes = []interface{}{
	(*OrderBooked)(nil), // 0: OrderBooked
}
var file_v1_events_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_events_proto_init() }
func file_v1_events_proto_init() {
	if File_v1_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderBooked); i {
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
			RawDescriptor: file_v1_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_events_proto_goTypes,
		DependencyIndexes: file_v1_events_proto_depIdxs,
		MessageInfos:      file_v1_events_proto_msgTypes,
	}.Build()
	File_v1_events_proto = out.File
	file_v1_events_proto_rawDesc = nil
	file_v1_events_proto_goTypes = nil
	file_v1_events_proto_depIdxs = nil
}
