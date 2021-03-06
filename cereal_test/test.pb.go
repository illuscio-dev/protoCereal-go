// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.13.0
// source: cereal_proto/test/test.proto

package cereal_test

import (
	proto "github.com/golang/protobuf/proto"
	cereal "github.com/illuscio-dev/protoCereal-go/cereal"
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

type Wizard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Wizard) Reset() {
	*x = Wizard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cereal_proto_test_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wizard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wizard) ProtoMessage() {}

func (x *Wizard) ProtoReflect() protoreflect.Message {
	mi := &file_cereal_proto_test_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wizard.ProtoReflect.Descriptor instead.
func (*Wizard) Descriptor() ([]byte, []int) {
	return file_cereal_proto_test_test_proto_rawDescGZIP(), []int{0}
}

func (x *Wizard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Witch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Witch) Reset() {
	*x = Witch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cereal_proto_test_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Witch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Witch) ProtoMessage() {}

func (x *Witch) ProtoReflect() protoreflect.Message {
	mi := &file_cereal_proto_test_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Witch.ProtoReflect.Descriptor instead.
func (*Witch) Descriptor() ([]byte, []int) {
	return file_cereal_proto_test_test_proto_rawDescGZIP(), []int{1}
}

func (x *Witch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: bson:"field_string"
	FieldString string `protobuf:"bytes,1,opt,name=field_string,json=fieldString,proto3" json:"field_string,omitempty" bson:"field_string"`
	// @inject_tag: bson:"field_int32"
	FieldInt32 int32 `protobuf:"varint,2,opt,name=field_int32,json=fieldInt32,proto3" json:"field_int32,omitempty" bson:"field_int32"`
}

func (x *TestProto) Reset() {
	*x = TestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cereal_proto_test_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestProto) ProtoMessage() {}

func (x *TestProto) ProtoReflect() protoreflect.Message {
	mi := &file_cereal_proto_test_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestProto.ProtoReflect.Descriptor instead.
func (*TestProto) Descriptor() ([]byte, []int) {
	return file_cereal_proto_test_test_proto_rawDescGZIP(), []int{2}
}

func (x *TestProto) GetFieldString() string {
	if x != nil {
		return x.FieldString
	}
	return ""
}

func (x *TestProto) GetFieldInt32() int32 {
	if x != nil {
		return x.FieldInt32
	}
	return 0
}

type TestOneOfFirst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: bson:"some_value"
	//
	// Types that are assignable to SomeValue:
	//	*TestOneOfFirst_FieldString
	//	*TestOneOfFirst_FieldInt32
	//	*TestOneOfFirst_FieldBool
	//	*TestOneOfFirst_FieldDouble
	//	*TestOneOfFirst_FieldDecimal
	//	*TestOneOfFirst_FieldUuid
	//	*TestOneOfFirst_FieldRaw
	//	*TestOneOfFirst_FieldWizard
	SomeValue isTestOneOfFirst_SomeValue `protobuf_oneof:"some_value" bson:"some_value"`
}

func (x *TestOneOfFirst) Reset() {
	*x = TestOneOfFirst{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cereal_proto_test_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestOneOfFirst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestOneOfFirst) ProtoMessage() {}

func (x *TestOneOfFirst) ProtoReflect() protoreflect.Message {
	mi := &file_cereal_proto_test_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestOneOfFirst.ProtoReflect.Descriptor instead.
func (*TestOneOfFirst) Descriptor() ([]byte, []int) {
	return file_cereal_proto_test_test_proto_rawDescGZIP(), []int{3}
}

func (m *TestOneOfFirst) GetSomeValue() isTestOneOfFirst_SomeValue {
	if m != nil {
		return m.SomeValue
	}
	return nil
}

func (x *TestOneOfFirst) GetFieldString() string {
	if x, ok := x.GetSomeValue().(*TestOneOfFirst_FieldString); ok {
		return x.FieldString
	}
	return ""
}

func (x *TestOneOfFirst) GetFieldInt32() int32 {
	if x, ok := x.GetSomeValue().(*TestOneOfFirst_FieldInt32); ok {
		return x.FieldInt32
	}
	return 0
}

func (x *TestOneOfFirst) GetFieldBool() bool {
	if x, ok := x.GetSomeValue().(*TestOneOfFirst_FieldBool); ok {
		return x.FieldBool
	}
	return false
}

func (x *TestOneOfFirst) GetFieldDouble() float64 {
	if x, ok := x.GetSomeValue().(*TestOneOfFirst_FieldDouble); ok {
		return x.FieldDouble
	}
	return 0
}

func (x *TestOneOfFirst) GetFieldDecimal() *cereal.Decimal {
	if x, ok := x.GetSomeValue().(*TestOneOfFirst_FieldDecimal); ok {
		return x.FieldDecimal
	}
	return nil
}

func (x *TestOneOfFirst) GetFieldUuid() *cereal.UUID {
	if x, ok := x.GetSomeValue().(*TestOneOfFirst_FieldUuid); ok {
		return x.FieldUuid
	}
	return nil
}

func (x *TestOneOfFirst) GetFieldRaw() *cereal.RawData {
	if x, ok := x.GetSomeValue().(*TestOneOfFirst_FieldRaw); ok {
		return x.FieldRaw
	}
	return nil
}

func (x *TestOneOfFirst) GetFieldWizard() *Wizard {
	if x, ok := x.GetSomeValue().(*TestOneOfFirst_FieldWizard); ok {
		return x.FieldWizard
	}
	return nil
}

type isTestOneOfFirst_SomeValue interface {
	isTestOneOfFirst_SomeValue()
}

type TestOneOfFirst_FieldString struct {
	FieldString string `protobuf:"bytes,1,opt,name=field_string,json=fieldString,proto3,oneof"`
}

type TestOneOfFirst_FieldInt32 struct {
	FieldInt32 int32 `protobuf:"varint,2,opt,name=field_int32,json=fieldInt32,proto3,oneof"`
}

type TestOneOfFirst_FieldBool struct {
	FieldBool bool `protobuf:"varint,3,opt,name=field_bool,json=fieldBool,proto3,oneof"`
}

type TestOneOfFirst_FieldDouble struct {
	FieldDouble float64 `protobuf:"fixed64,5,opt,name=field_double,json=fieldDouble,proto3,oneof"`
}

type TestOneOfFirst_FieldDecimal struct {
	FieldDecimal *cereal.Decimal `protobuf:"bytes,6,opt,name=field_decimal,json=fieldDecimal,proto3,oneof"`
}

type TestOneOfFirst_FieldUuid struct {
	FieldUuid *cereal.UUID `protobuf:"bytes,7,opt,name=field_uuid,json=fieldUuid,proto3,oneof"`
}

type TestOneOfFirst_FieldRaw struct {
	FieldRaw *cereal.RawData `protobuf:"bytes,8,opt,name=field_raw,json=fieldRaw,proto3,oneof"`
}

type TestOneOfFirst_FieldWizard struct {
	FieldWizard *Wizard `protobuf:"bytes,9,opt,name=field_wizard,json=fieldWizard,proto3,oneof"`
}

func (*TestOneOfFirst_FieldString) isTestOneOfFirst_SomeValue() {}

func (*TestOneOfFirst_FieldInt32) isTestOneOfFirst_SomeValue() {}

func (*TestOneOfFirst_FieldBool) isTestOneOfFirst_SomeValue() {}

func (*TestOneOfFirst_FieldDouble) isTestOneOfFirst_SomeValue() {}

func (*TestOneOfFirst_FieldDecimal) isTestOneOfFirst_SomeValue() {}

func (*TestOneOfFirst_FieldUuid) isTestOneOfFirst_SomeValue() {}

func (*TestOneOfFirst_FieldRaw) isTestOneOfFirst_SomeValue() {}

func (*TestOneOfFirst_FieldWizard) isTestOneOfFirst_SomeValue() {}

// Used to test the ability to encode and decode multiple sub-messages as embedded docs.
type TestOneOfMultiMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Mage:
	//	*TestOneOfMultiMessage_Witch
	//	*TestOneOfMultiMessage_Wizard
	Mage isTestOneOfMultiMessage_Mage `protobuf_oneof:"mage"`
}

func (x *TestOneOfMultiMessage) Reset() {
	*x = TestOneOfMultiMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cereal_proto_test_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestOneOfMultiMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestOneOfMultiMessage) ProtoMessage() {}

func (x *TestOneOfMultiMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cereal_proto_test_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestOneOfMultiMessage.ProtoReflect.Descriptor instead.
func (*TestOneOfMultiMessage) Descriptor() ([]byte, []int) {
	return file_cereal_proto_test_test_proto_rawDescGZIP(), []int{4}
}

func (m *TestOneOfMultiMessage) GetMage() isTestOneOfMultiMessage_Mage {
	if m != nil {
		return m.Mage
	}
	return nil
}

func (x *TestOneOfMultiMessage) GetWitch() *Witch {
	if x, ok := x.GetMage().(*TestOneOfMultiMessage_Witch); ok {
		return x.Witch
	}
	return nil
}

func (x *TestOneOfMultiMessage) GetWizard() *Wizard {
	if x, ok := x.GetMage().(*TestOneOfMultiMessage_Wizard); ok {
		return x.Wizard
	}
	return nil
}

type isTestOneOfMultiMessage_Mage interface {
	isTestOneOfMultiMessage_Mage()
}

type TestOneOfMultiMessage_Witch struct {
	Witch *Witch `protobuf:"bytes,1,opt,name=witch,proto3,oneof"`
}

type TestOneOfMultiMessage_Wizard struct {
	Wizard *Wizard `protobuf:"bytes,2,opt,name=wizard,proto3,oneof"`
}

func (*TestOneOfMultiMessage_Witch) isTestOneOfMultiMessage_Mage() {}

func (*TestOneOfMultiMessage_Wizard) isTestOneOfMultiMessage_Mage() {}

// Used to test custom wrapper types.
type ListValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: bson:"value"
	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" bson:"value"`
}

func (x *ListValue) Reset() {
	*x = ListValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cereal_proto_test_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListValue) ProtoMessage() {}

func (x *ListValue) ProtoReflect() protoreflect.Message {
	mi := &file_cereal_proto_test_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListValue.ProtoReflect.Descriptor instead.
func (*ListValue) Descriptor() ([]byte, []int) {
	return file_cereal_proto_test_test_proto_rawDescGZIP(), []int{5}
}

func (x *ListValue) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_cereal_proto_test_test_proto protoreflect.FileDescriptor

var file_cereal_proto_test_test_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x63, 0x65, 0x72,
	0x65, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72,
	0x61, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a,
	0x06, 0x57, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1b, 0x0a, 0x05, 0x57,
	0x69, 0x74, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x22, 0xfd, 0x02, 0x0a, 0x0e, 0x54, 0x65,
	0x73, 0x74, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0c,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x21, 0x0a, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x12, 0x1f, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x6f,
	0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x23, 0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d,
	0x61, 0x6c, 0x48, 0x00, 0x52, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x63, 0x69, 0x6d,
	0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x2e,
	0x55, 0x55, 0x49, 0x44, 0x48, 0x00, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x75, 0x69,
	0x64, 0x12, 0x2e, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x72, 0x61, 0x77, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x2e, 0x52, 0x61,
	0x77, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x61,
	0x77, 0x12, 0x38, 0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x77, 0x69, 0x7a, 0x61, 0x72,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x48, 0x00, 0x52, 0x0b,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x57, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x73,
	0x6f, 0x6d, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7a, 0x0a, 0x15, 0x54, 0x65, 0x73,
	0x74, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x77, 0x69, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x57, 0x69, 0x74, 0x63, 0x68, 0x48, 0x00, 0x52, 0x05, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x2d,
	0x0a, 0x06, 0x77, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x7a,
	0x61, 0x72, 0x64, 0x48, 0x00, 0x52, 0x06, 0x77, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x42, 0x06, 0x0a,
	0x04, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6c, 0x6c, 0x75, 0x73, 0x63, 0x69, 0x6f, 0x2d,
	0x64, 0x65, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x2d,
	0x67, 0x6f, 0x2f, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cereal_proto_test_test_proto_rawDescOnce sync.Once
	file_cereal_proto_test_test_proto_rawDescData = file_cereal_proto_test_test_proto_rawDesc
)

func file_cereal_proto_test_test_proto_rawDescGZIP() []byte {
	file_cereal_proto_test_test_proto_rawDescOnce.Do(func() {
		file_cereal_proto_test_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_cereal_proto_test_test_proto_rawDescData)
	})
	return file_cereal_proto_test_test_proto_rawDescData
}

var file_cereal_proto_test_test_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cereal_proto_test_test_proto_goTypes = []interface{}{
	(*Wizard)(nil),                // 0: cereal_test.Wizard
	(*Witch)(nil),                 // 1: cereal_test.Witch
	(*TestProto)(nil),             // 2: cereal_test.TestProto
	(*TestOneOfFirst)(nil),        // 3: cereal_test.TestOneOfFirst
	(*TestOneOfMultiMessage)(nil), // 4: cereal_test.TestOneOfMultiMessage
	(*ListValue)(nil),             // 5: cereal_test.ListValue
	(*cereal.Decimal)(nil),        // 6: cereal.Decimal
	(*cereal.UUID)(nil),           // 7: cereal.UUID
	(*cereal.RawData)(nil),        // 8: cereal.RawData
}
var file_cereal_proto_test_test_proto_depIdxs = []int32{
	6, // 0: cereal_test.TestOneOfFirst.field_decimal:type_name -> cereal.Decimal
	7, // 1: cereal_test.TestOneOfFirst.field_uuid:type_name -> cereal.UUID
	8, // 2: cereal_test.TestOneOfFirst.field_raw:type_name -> cereal.RawData
	0, // 3: cereal_test.TestOneOfFirst.field_wizard:type_name -> cereal_test.Wizard
	1, // 4: cereal_test.TestOneOfMultiMessage.witch:type_name -> cereal_test.Witch
	0, // 5: cereal_test.TestOneOfMultiMessage.wizard:type_name -> cereal_test.Wizard
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cereal_proto_test_test_proto_init() }
func file_cereal_proto_test_test_proto_init() {
	if File_cereal_proto_test_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cereal_proto_test_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wizard); i {
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
		file_cereal_proto_test_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Witch); i {
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
		file_cereal_proto_test_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestProto); i {
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
		file_cereal_proto_test_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestOneOfFirst); i {
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
		file_cereal_proto_test_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestOneOfMultiMessage); i {
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
		file_cereal_proto_test_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListValue); i {
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
	file_cereal_proto_test_test_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*TestOneOfFirst_FieldString)(nil),
		(*TestOneOfFirst_FieldInt32)(nil),
		(*TestOneOfFirst_FieldBool)(nil),
		(*TestOneOfFirst_FieldDouble)(nil),
		(*TestOneOfFirst_FieldDecimal)(nil),
		(*TestOneOfFirst_FieldUuid)(nil),
		(*TestOneOfFirst_FieldRaw)(nil),
		(*TestOneOfFirst_FieldWizard)(nil),
	}
	file_cereal_proto_test_test_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*TestOneOfMultiMessage_Witch)(nil),
		(*TestOneOfMultiMessage_Wizard)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cereal_proto_test_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cereal_proto_test_test_proto_goTypes,
		DependencyIndexes: file_cereal_proto_test_test_proto_depIdxs,
		MessageInfos:      file_cereal_proto_test_test_proto_msgTypes,
	}.Build()
	File_cereal_proto_test_test_proto = out.File
	file_cereal_proto_test_test_proto_rawDesc = nil
	file_cereal_proto_test_test_proto_goTypes = nil
	file_cereal_proto_test_test_proto_depIdxs = nil
}
