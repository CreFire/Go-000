// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: Api/msg.proto

package Api

import (
	proto "github.com/golang/protobuf/proto"
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

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Number int32  `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	AcTime int32  `protobuf:"varint,3,opt,name=AcTime,proto3" json:"AcTime,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Api_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Api_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_Api_msg_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *SearchRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *SearchRequest) GetAcTime() int32 {
	if x != nil {
		return x.AcTime
	}
	return 0
}

type SearchReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SearchReply) Reset() {
	*x = SearchReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Api_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReply) ProtoMessage() {}

func (x *SearchReply) ProtoReflect() protoreflect.Message {
	mi := &file_Api_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReply.ProtoReflect.Descriptor instead.
func (*SearchReply) Descriptor() ([]byte, []int) {
	return file_Api_msg_proto_rawDescGZIP(), []int{1}
}

func (x *SearchReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SearchReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SearchReply) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_Api_msg_proto protoreflect.FileDescriptor

var file_Api_msg_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x41, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x57, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x41, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0x33, 0x0a, 0x0b, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x24, 0x0a, 0x04, 0x6d, 0x73, 0x67, 0x46, 0x12, 0x0e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x05, 0x5a, 0x03, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Api_msg_proto_rawDescOnce sync.Once
	file_Api_msg_proto_rawDescData = file_Api_msg_proto_rawDesc
)

func file_Api_msg_proto_rawDescGZIP() []byte {
	file_Api_msg_proto_rawDescOnce.Do(func() {
		file_Api_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_Api_msg_proto_rawDescData)
	})
	return file_Api_msg_proto_rawDescData
}

var file_Api_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Api_msg_proto_goTypes = []interface{}{
	(*SearchRequest)(nil), // 0: SearchRequest
	(*SearchReply)(nil),   // 1: SearchReply
}
var file_Api_msg_proto_depIdxs = []int32{
	0, // 0: GrpcService.msgF:input_type -> SearchRequest
	1, // 1: GrpcService.msgF:output_type -> SearchReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Api_msg_proto_init() }
func file_Api_msg_proto_init() {
	if File_Api_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Api_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_Api_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReply); i {
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
			RawDescriptor: file_Api_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Api_msg_proto_goTypes,
		DependencyIndexes: file_Api_msg_proto_depIdxs,
		MessageInfos:      file_Api_msg_proto_msgTypes,
	}.Build()
	File_Api_msg_proto = out.File
	file_Api_msg_proto_rawDesc = nil
	file_Api_msg_proto_goTypes = nil
	file_Api_msg_proto_depIdxs = nil
}
