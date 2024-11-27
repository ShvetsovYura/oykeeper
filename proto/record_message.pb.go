// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: proto/record_message.proto

package proto

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

type RecordCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item string `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *RecordCreateRequest) Reset() {
	*x = RecordCreateRequest{}
	mi := &file_proto_record_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecordCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordCreateRequest) ProtoMessage() {}

func (x *RecordCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_record_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordCreateRequest.ProtoReflect.Descriptor instead.
func (*RecordCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_record_message_proto_rawDescGZIP(), []int{0}
}

func (x *RecordCreateRequest) GetItem() string {
	if x != nil {
		return x.Item
	}
	return ""
}

type RecordCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RecordCreateResponse) Reset() {
	*x = RecordCreateResponse{}
	mi := &file_proto_record_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecordCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordCreateResponse) ProtoMessage() {}

func (x *RecordCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_record_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordCreateResponse.ProtoReflect.Descriptor instead.
func (*RecordCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_record_message_proto_rawDescGZIP(), []int{1}
}

type AttrbuteCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AttrbuteCreateRequest) Reset() {
	*x = AttrbuteCreateRequest{}
	mi := &file_proto_record_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AttrbuteCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttrbuteCreateRequest) ProtoMessage() {}

func (x *AttrbuteCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_record_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttrbuteCreateRequest.ProtoReflect.Descriptor instead.
func (*AttrbuteCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_record_message_proto_rawDescGZIP(), []int{2}
}

func (x *AttrbuteCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttrbuteCreateRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type AttributeCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AttributeCreateResponse) Reset() {
	*x = AttributeCreateResponse{}
	mi := &file_proto_record_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AttributeCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeCreateResponse) ProtoMessage() {}

func (x *AttributeCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_record_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeCreateResponse.ProtoReflect.Descriptor instead.
func (*AttributeCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_record_message_proto_rawDescGZIP(), []int{3}
}

func (x *AttributeCreateResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttributeCreateResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_proto_record_message_proto protoreflect.FileDescriptor

var file_proto_record_message_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x16,
	0x0a, 0x14, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0x0a, 0x15, 0x41, 0x74, 0x74, 0x72, 0x62, 0x75,
	0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x43, 0x0a, 0x17, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x11,
	0x5a, 0x0f, 0x6f, 0x79, 0x2d, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_record_message_proto_rawDescOnce sync.Once
	file_proto_record_message_proto_rawDescData = file_proto_record_message_proto_rawDesc
)

func file_proto_record_message_proto_rawDescGZIP() []byte {
	file_proto_record_message_proto_rawDescOnce.Do(func() {
		file_proto_record_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_record_message_proto_rawDescData)
	})
	return file_proto_record_message_proto_rawDescData
}

var file_proto_record_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_record_message_proto_goTypes = []any{
	(*RecordCreateRequest)(nil),     // 0: proto.RecordCreateRequest
	(*RecordCreateResponse)(nil),    // 1: proto.RecordCreateResponse
	(*AttrbuteCreateRequest)(nil),   // 2: proto.AttrbuteCreateRequest
	(*AttributeCreateResponse)(nil), // 3: proto.AttributeCreateResponse
}
var file_proto_record_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_record_message_proto_init() }
func file_proto_record_message_proto_init() {
	if File_proto_record_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_record_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_record_message_proto_goTypes,
		DependencyIndexes: file_proto_record_message_proto_depIdxs,
		MessageInfos:      file_proto_record_message_proto_msgTypes,
	}.Build()
	File_proto_record_message_proto = out.File
	file_proto_record_message_proto_rawDesc = nil
	file_proto_record_message_proto_goTypes = nil
	file_proto_record_message_proto_depIdxs = nil
}
