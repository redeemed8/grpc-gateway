// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.4
// source: myservice.proto

package pbfile

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Num      int64  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_myservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_myservice_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *Message) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type SimpleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Num      int64     `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Language string    `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`
	Msg      *InnerMsg `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SimpleMessage) Reset() {
	*x = SimpleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleMessage) ProtoMessage() {}

func (x *SimpleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_myservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleMessage.ProtoReflect.Descriptor instead.
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return file_myservice_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SimpleMessage) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *SimpleMessage) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *SimpleMessage) GetMsg() *InnerMsg {
	if x != nil {
		return x.Msg
	}
	return nil
}

type InnerMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	F1 string `protobuf:"bytes,1,opt,name=f1,proto3" json:"f1,omitempty"`
	F2 string `protobuf:"bytes,2,opt,name=f2,proto3" json:"f2,omitempty"`
}

func (x *InnerMsg) Reset() {
	*x = InnerMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerMsg) ProtoMessage() {}

func (x *InnerMsg) ProtoReflect() protoreflect.Message {
	mi := &file_myservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerMsg.ProtoReflect.Descriptor instead.
func (*InnerMsg) Descriptor() ([]byte, []int) {
	return file_myservice_proto_rawDescGZIP(), []int{2}
}

func (x *InnerMsg) GetF1() string {
	if x != nil {
		return x.F1
	}
	return ""
}

func (x *InnerMsg) GetF2() string {
	if x != nil {
		return x.F2
	}
	return ""
}

var File_myservice_proto protoreflect.FileDescriptor

var file_myservice_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x11, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x47, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x74, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2a,
	0x0a, 0x08, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x66, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x66, 0x32, 0x32, 0x7f, 0x0a, 0x09, 0x4d, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12,
	0x12, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x45, 0x63, 0x68, 0x6f, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x18, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x70,
	0x62, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_myservice_proto_rawDescOnce sync.Once
	file_myservice_proto_rawDescData = file_myservice_proto_rawDesc
)

func file_myservice_proto_rawDescGZIP() []byte {
	file_myservice_proto_rawDescOnce.Do(func() {
		file_myservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_myservice_proto_rawDescData)
	})
	return file_myservice_proto_rawDescData
}

var file_myservice_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_myservice_proto_goTypes = []interface{}{
	(*Message)(nil),       // 0: myservice.Message
	(*SimpleMessage)(nil), // 1: myservice.SimpleMessage
	(*InnerMsg)(nil),      // 2: myservice.InnerMsg
}
var file_myservice_proto_depIdxs = []int32{
	2, // 0: myservice.SimpleMessage.msg:type_name -> myservice.InnerMsg
	0, // 1: myservice.MyService.Echo:input_type -> myservice.Message
	1, // 2: myservice.MyService.EchoSimple:input_type -> myservice.SimpleMessage
	0, // 3: myservice.MyService.Echo:output_type -> myservice.Message
	1, // 4: myservice.MyService.EchoSimple:output_type -> myservice.SimpleMessage
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_myservice_proto_init() }
func file_myservice_proto_init() {
	if File_myservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_myservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_myservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleMessage); i {
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
		file_myservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerMsg); i {
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
			RawDescriptor: file_myservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_myservice_proto_goTypes,
		DependencyIndexes: file_myservice_proto_depIdxs,
		MessageInfos:      file_myservice_proto_msgTypes,
	}.Build()
	File_myservice_proto = out.File
	file_myservice_proto_rawDesc = nil
	file_myservice_proto_goTypes = nil
	file_myservice_proto_depIdxs = nil
}
