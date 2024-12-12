// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: anonymouschat.proto

// protobuf config

package chat

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConnectEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsConnected bool                   `protobuf:"varint,2,opt,name=is_connected,json=isConnected,proto3" json:"is_connected,omitempty"`
	Timestamp   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ConnectEvent) Reset() {
	*x = ConnectEvent{}
	mi := &file_anonymouschat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectEvent) ProtoMessage() {}

func (x *ConnectEvent) ProtoReflect() protoreflect.Message {
	mi := &file_anonymouschat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectEvent.ProtoReflect.Descriptor instead.
func (*ConnectEvent) Descriptor() ([]byte, []int) {
	return file_anonymouschat_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ConnectEvent) GetIsConnected() bool {
	if x != nil {
		return x.IsConnected
	}
	return false
}

func (x *ConnectEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_anonymouschat_proto protoreflect.FileDescriptor

var file_anonymouschat_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73,
	0x63, 0x68, 0x61, 0x74, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x52, 0x0a, 0x0b,
	0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_anonymouschat_proto_rawDescOnce sync.Once
	file_anonymouschat_proto_rawDescData = file_anonymouschat_proto_rawDesc
)

func file_anonymouschat_proto_rawDescGZIP() []byte {
	file_anonymouschat_proto_rawDescOnce.Do(func() {
		file_anonymouschat_proto_rawDescData = protoimpl.X.CompressGZIP(file_anonymouschat_proto_rawDescData)
	})
	return file_anonymouschat_proto_rawDescData
}

var file_anonymouschat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_anonymouschat_proto_goTypes = []any{
	(*ConnectEvent)(nil),          // 0: anonymouschat.ConnectEvent
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_anonymouschat_proto_depIdxs = []int32{
	1, // 0: anonymouschat.ConnectEvent.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: anonymouschat.ChatService.Connect:input_type -> anonymouschat.ConnectEvent
	0, // 2: anonymouschat.ChatService.Connect:output_type -> anonymouschat.ConnectEvent
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_anonymouschat_proto_init() }
func file_anonymouschat_proto_init() {
	if File_anonymouschat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anonymouschat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_anonymouschat_proto_goTypes,
		DependencyIndexes: file_anonymouschat_proto_depIdxs,
		MessageInfos:      file_anonymouschat_proto_msgTypes,
	}.Build()
	File_anonymouschat_proto = out.File
	file_anonymouschat_proto_rawDesc = nil
	file_anonymouschat_proto_goTypes = nil
	file_anonymouschat_proto_depIdxs = nil
}