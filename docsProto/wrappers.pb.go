// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: cereal_proto/docs/wrappers.proto

package docsProto

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

// Wrapper for fixed int64.
type Fixed64Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: bson:"value"
	Value uint64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty" bson:"value"`
}

func (x *Fixed64Value) Reset() {
	*x = Fixed64Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cereal_proto_docs_wrappers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fixed64Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fixed64Value) ProtoMessage() {}

func (x *Fixed64Value) ProtoReflect() protoreflect.Message {
	mi := &file_cereal_proto_docs_wrappers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fixed64Value.ProtoReflect.Descriptor instead.
func (*Fixed64Value) Descriptor() ([]byte, []int) {
	return file_cereal_proto_docs_wrappers_proto_rawDescGZIP(), []int{0}
}

func (x *Fixed64Value) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_cereal_proto_docs_wrappers_proto protoreflect.FileDescriptor

var file_cereal_proto_docs_wrappers_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64,
	0x6f, 0x63, 0x73, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x22, 0x24,
	0x0a, 0x0c, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6c, 0x6c, 0x75, 0x73, 0x63, 0x69, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x2d, 0x67, 0x6f, 0x2f, 0x64,
	0x6f, 0x63, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cereal_proto_docs_wrappers_proto_rawDescOnce sync.Once
	file_cereal_proto_docs_wrappers_proto_rawDescData = file_cereal_proto_docs_wrappers_proto_rawDesc
)

func file_cereal_proto_docs_wrappers_proto_rawDescGZIP() []byte {
	file_cereal_proto_docs_wrappers_proto_rawDescOnce.Do(func() {
		file_cereal_proto_docs_wrappers_proto_rawDescData = protoimpl.X.CompressGZIP(file_cereal_proto_docs_wrappers_proto_rawDescData)
	})
	return file_cereal_proto_docs_wrappers_proto_rawDescData
}

var file_cereal_proto_docs_wrappers_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cereal_proto_docs_wrappers_proto_goTypes = []interface{}{
	(*Fixed64Value)(nil), // 0: cereal_doc.Fixed64Value
}
var file_cereal_proto_docs_wrappers_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cereal_proto_docs_wrappers_proto_init() }
func file_cereal_proto_docs_wrappers_proto_init() {
	if File_cereal_proto_docs_wrappers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cereal_proto_docs_wrappers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fixed64Value); i {
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
			RawDescriptor: file_cereal_proto_docs_wrappers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cereal_proto_docs_wrappers_proto_goTypes,
		DependencyIndexes: file_cereal_proto_docs_wrappers_proto_depIdxs,
		MessageInfos:      file_cereal_proto_docs_wrappers_proto_msgTypes,
	}.Build()
	File_cereal_proto_docs_wrappers_proto = out.File
	file_cereal_proto_docs_wrappers_proto_rawDesc = nil
	file_cereal_proto_docs_wrappers_proto_goTypes = nil
	file_cereal_proto_docs_wrappers_proto_depIdxs = nil
}
