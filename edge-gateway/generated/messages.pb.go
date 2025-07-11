// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: messages.proto

package messages

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DeviceEvent is what's sent from the devices to the Edge Gateway
type DeviceEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeviceId      string                 `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Payload       string                 `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeviceEvent) Reset() {
	*x = DeviceEvent{}
	mi := &file_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceEvent) ProtoMessage() {}

func (x *DeviceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceEvent.ProtoReflect.Descriptor instead.
func (*DeviceEvent) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceEvent) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *DeviceEvent) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DeviceEvent) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

var File_messages_proto protoreflect.FileDescriptor

const file_messages_proto_rawDesc = "" +
	"\n" +
	"\x0emessages.proto\x12\bmessages\"X\n" +
	"\vDeviceEvent\x12\x1b\n" +
	"\tdevice_id\x18\x01 \x01(\tR\bdeviceId\x12\x12\n" +
	"\x04type\x18\x02 \x01(\tR\x04type\x12\x18\n" +
	"\apayload\x18\x03 \x01(\tR\apayloadBEZCgithub.com/KelsyF/secure-edge-cloud/edge-gateway/generated/messagesb\x06proto3"

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData []byte
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_messages_proto_rawDesc), len(file_messages_proto_rawDesc)))
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_messages_proto_goTypes = []any{
	(*DeviceEvent)(nil), // 0: messages.DeviceEvent
}
var file_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_messages_proto_rawDesc), len(file_messages_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
