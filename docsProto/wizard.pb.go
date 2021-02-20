// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.13.0
// source: cereal_proto/docs/wizard.proto

package docsProto

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// The hogwarts houses.
type Houses int32

const (
	Houses_GRYFFINDOR Houses = 0
	Houses_RAVENCLAW  Houses = 1
	Houses_HUFFLEPUFF Houses = 2
	Houses_SLYTHERIN  Houses = 3
)

// Enum value maps for Houses.
var (
	Houses_name = map[int32]string{
		0: "GRYFFINDOR",
		1: "RAVENCLAW",
		2: "HUFFLEPUFF",
		3: "SLYTHERIN",
	}
	Houses_value = map[string]int32{
		"GRYFFINDOR": 0,
		"RAVENCLAW":  1,
		"HUFFLEPUFF": 2,
		"SLYTHERIN":  3,
	}
)

func (x Houses) Enum() *Houses {
	p := new(Houses)
	*p = x
	return p
}

func (x Houses) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Houses) Descriptor() protoreflect.EnumDescriptor {
	return file_cereal_proto_docs_wizard_proto_enumTypes[0].Descriptor()
}

func (Houses) Type() protoreflect.EnumType {
	return &file_cereal_proto_docs_wizard_proto_enumTypes[0]
}

func (x Houses) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Houses.Descriptor instead.
func (Houses) EnumDescriptor() ([]byte, []int) {
	return file_cereal_proto_docs_wizard_proto_rawDescGZIP(), []int{0}
}

// Information about a Wizard's wand.
type Wand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Core string `protobuf:"bytes,1,opt,name=core,proto3" json:"core,omitempty"`
}

func (x *Wand) Reset() {
	*x = Wand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cereal_proto_docs_wizard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wand) ProtoMessage() {}

func (x *Wand) ProtoReflect() protoreflect.Message {
	mi := &file_cereal_proto_docs_wizard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wand.ProtoReflect.Descriptor instead.
func (*Wand) Descriptor() ([]byte, []int) {
	return file_cereal_proto_docs_wizard_proto_rawDescGZIP(), []int{0}
}

func (x *Wand) GetCore() string {
	if x != nil {
		return x.Core
	}
	return ""
}

// Information about a Wizard's sword.
type Sword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metal string `protobuf:"bytes,2,opt,name=metal,proto3" json:"metal,omitempty"`
}

func (x *Sword) Reset() {
	*x = Sword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cereal_proto_docs_wizard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sword) ProtoMessage() {}

func (x *Sword) ProtoReflect() protoreflect.Message {
	mi := &file_cereal_proto_docs_wizard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sword.ProtoReflect.Descriptor instead.
func (*Sword) Descriptor() ([]byte, []int) {
	return file_cereal_proto_docs_wizard_proto_rawDescGZIP(), []int{1}
}

func (x *Sword) GetMetal() string {
	if x != nil {
		return x.Metal
	}
	return ""
}

// Information about a Hogwarts wizard.
type Wizard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the Wizard.
	// @inject_tag: bson:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	// A unique Identifier for the Wizard.
	// @inject_tag: bson:"id"
	Id *cereal.UUID `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" bson:"id"`
	// The exact moment the Wizard was sorted.
	// @inject_tag: bson:"sorted_at"
	SortedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=sorted_at,json=sortedAt,proto3" json:"sorted_at,omitempty" bson:"sorted_at"`
	// The house this wizard belongs to.
	// @inject_tag: bson:"hogwarts_house"
	HogwartsHouse Houses `protobuf:"varint,4,opt,name=hogwarts_house,json=hogwartsHouse,proto3,enum=cereal_doc.Houses" json:"hogwarts_house,omitempty" bson:"hogwarts_house"`
	// The current balance of the wizard's Gringott's account in Galleons.
	// @inject_tag: bson:"gingotts_balance"
	GingottsBalance *cereal.Decimal `protobuf:"bytes,5,opt,name=gingotts_balance,json=gingottsBalance,proto3" json:"gingotts_balance,omitempty" bson:"gingotts_balance"`
	// Name of the Wizard's familiar. Nil if Wizard does not have a familiar.
	// @inject_tag: bson:"familiar_name"
	FamiliarName *wrappers.StringValue `protobuf:"bytes,6,opt,name=familiar_name,json=familiarName,proto3" json:"familiar_name,omitempty" bson:"familiar_name"`
	// Image of the Wizard.
	// @inject_tag: bson:"portrait"
	Portrait *cereal.RawData `protobuf:"bytes,7,opt,name=portrait,proto3" json:"portrait,omitempty" bson:"portrait"`
	// The preferred weapon of this wizard.
	// @inject_tag: bson:"weapon"
	//
	// Types that are assignable to Weapon:
	//	*Wizard_Wand
	//	*Wizard_Sword
	Weapon isWizard_Weapon `protobuf_oneof:"weapon" bson:"weapon"`
	// An object the wizard is destined to obtain.
	// @inject_tag: bson:"destined_object"
	DestinedObject *any.Any `protobuf:"bytes,10,opt,name=destined_object,json=destinedObject,proto3" json:"destined_object,omitempty" bson:"destined_object"`
}

func (x *Wizard) Reset() {
	*x = Wizard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cereal_proto_docs_wizard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wizard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wizard) ProtoMessage() {}

func (x *Wizard) ProtoReflect() protoreflect.Message {
	mi := &file_cereal_proto_docs_wizard_proto_msgTypes[2]
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
	return file_cereal_proto_docs_wizard_proto_rawDescGZIP(), []int{2}
}

func (x *Wizard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Wizard) GetId() *cereal.UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Wizard) GetSortedAt() *timestamp.Timestamp {
	if x != nil {
		return x.SortedAt
	}
	return nil
}

func (x *Wizard) GetHogwartsHouse() Houses {
	if x != nil {
		return x.HogwartsHouse
	}
	return Houses_GRYFFINDOR
}

func (x *Wizard) GetGingottsBalance() *cereal.Decimal {
	if x != nil {
		return x.GingottsBalance
	}
	return nil
}

func (x *Wizard) GetFamiliarName() *wrappers.StringValue {
	if x != nil {
		return x.FamiliarName
	}
	return nil
}

func (x *Wizard) GetPortrait() *cereal.RawData {
	if x != nil {
		return x.Portrait
	}
	return nil
}

func (m *Wizard) GetWeapon() isWizard_Weapon {
	if m != nil {
		return m.Weapon
	}
	return nil
}

func (x *Wizard) GetWand() *Wand {
	if x, ok := x.GetWeapon().(*Wizard_Wand); ok {
		return x.Wand
	}
	return nil
}

func (x *Wizard) GetSword() *Sword {
	if x, ok := x.GetWeapon().(*Wizard_Sword); ok {
		return x.Sword
	}
	return nil
}

func (x *Wizard) GetDestinedObject() *any.Any {
	if x != nil {
		return x.DestinedObject
	}
	return nil
}

type isWizard_Weapon interface {
	isWizard_Weapon()
}

type Wizard_Wand struct {
	Wand *Wand `protobuf:"bytes,8,opt,name=wand,proto3,oneof"`
}

type Wizard_Sword struct {
	Sword *Sword `protobuf:"bytes,9,opt,name=sword,proto3,oneof"`
}

func (*Wizard_Wand) isWizard_Weapon() {}

func (*Wizard_Sword) isWizard_Weapon() {}

var File_cereal_proto_docs_wizard_proto protoreflect.FileDescriptor

var file_cereal_proto_docs_wizard_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64,
	0x6f, 0x63, 0x73, 0x2f, 0x77, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x1a, 0x1a, 0x63, 0x65,
	0x72, 0x65, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x6d,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x72, 0x61, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x04, 0x57, 0x61,
	0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x22, 0xf6, 0x03, 0x0a, 0x06, 0x57, 0x69, 0x7a, 0x61, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0e, 0x68,
	0x6f, 0x67, 0x77, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x63,
	0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x52, 0x0d, 0x68, 0x6f, 0x67, 0x77, 0x61, 0x72, 0x74,
	0x73, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x10, 0x67, 0x69, 0x6e, 0x67, 0x6f, 0x74,
	0x74, 0x73, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61,
	0x6c, 0x52, 0x0f, 0x67, 0x69, 0x6e, 0x67, 0x6f, 0x74, 0x74, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x69, 0x61, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x69, 0x61,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c,
	0x2e, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x77, 0x61, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x2e, 0x57, 0x61,
	0x6e, 0x64, 0x48, 0x00, 0x52, 0x04, 0x77, 0x61, 0x6e, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x65, 0x72, 0x65,
	0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x2e, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x00, 0x52, 0x05,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x3d, 0x0a, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x65,
	0x64, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x65, 0x64, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x2a, 0x46,
	0x0a, 0x06, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x52, 0x59, 0x46,
	0x46, 0x49, 0x4e, 0x44, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x41, 0x56, 0x45,
	0x4e, 0x43, 0x4c, 0x41, 0x57, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x55, 0x46, 0x46, 0x4c,
	0x45, 0x50, 0x55, 0x46, 0x46, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x4c, 0x59, 0x54, 0x48,
	0x45, 0x52, 0x49, 0x4e, 0x10, 0x03, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6c, 0x6c, 0x75, 0x73, 0x63, 0x69, 0x6f, 0x2d, 0x64, 0x65,
	0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x65, 0x72, 0x65, 0x61, 0x6c, 0x2d, 0x67, 0x6f,
	0x2f, 0x64, 0x6f, 0x63, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cereal_proto_docs_wizard_proto_rawDescOnce sync.Once
	file_cereal_proto_docs_wizard_proto_rawDescData = file_cereal_proto_docs_wizard_proto_rawDesc
)

func file_cereal_proto_docs_wizard_proto_rawDescGZIP() []byte {
	file_cereal_proto_docs_wizard_proto_rawDescOnce.Do(func() {
		file_cereal_proto_docs_wizard_proto_rawDescData = protoimpl.X.CompressGZIP(file_cereal_proto_docs_wizard_proto_rawDescData)
	})
	return file_cereal_proto_docs_wizard_proto_rawDescData
}

var file_cereal_proto_docs_wizard_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cereal_proto_docs_wizard_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cereal_proto_docs_wizard_proto_goTypes = []interface{}{
	(Houses)(0),                  // 0: cereal_doc.Houses
	(*Wand)(nil),                 // 1: cereal_doc.Wand
	(*Sword)(nil),                // 2: cereal_doc.Sword
	(*Wizard)(nil),               // 3: cereal_doc.Wizard
	(*cereal.UUID)(nil),          // 4: cereal.UUID
	(*timestamp.Timestamp)(nil),  // 5: google.protobuf.Timestamp
	(*cereal.Decimal)(nil),       // 6: cereal.Decimal
	(*wrappers.StringValue)(nil), // 7: google.protobuf.StringValue
	(*cereal.RawData)(nil),       // 8: cereal.RawData
	(*any.Any)(nil),              // 9: google.protobuf.Any
}
var file_cereal_proto_docs_wizard_proto_depIdxs = []int32{
	4, // 0: cereal_doc.Wizard.id:type_name -> cereal.UUID
	5, // 1: cereal_doc.Wizard.sorted_at:type_name -> google.protobuf.Timestamp
	0, // 2: cereal_doc.Wizard.hogwarts_house:type_name -> cereal_doc.Houses
	6, // 3: cereal_doc.Wizard.gingotts_balance:type_name -> cereal.Decimal
	7, // 4: cereal_doc.Wizard.familiar_name:type_name -> google.protobuf.StringValue
	8, // 5: cereal_doc.Wizard.portrait:type_name -> cereal.RawData
	1, // 6: cereal_doc.Wizard.wand:type_name -> cereal_doc.Wand
	2, // 7: cereal_doc.Wizard.sword:type_name -> cereal_doc.Sword
	9, // 8: cereal_doc.Wizard.destined_object:type_name -> google.protobuf.Any
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_cereal_proto_docs_wizard_proto_init() }
func file_cereal_proto_docs_wizard_proto_init() {
	if File_cereal_proto_docs_wizard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cereal_proto_docs_wizard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wand); i {
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
		file_cereal_proto_docs_wizard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sword); i {
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
		file_cereal_proto_docs_wizard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_cereal_proto_docs_wizard_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Wizard_Wand)(nil),
		(*Wizard_Sword)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cereal_proto_docs_wizard_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cereal_proto_docs_wizard_proto_goTypes,
		DependencyIndexes: file_cereal_proto_docs_wizard_proto_depIdxs,
		EnumInfos:         file_cereal_proto_docs_wizard_proto_enumTypes,
		MessageInfos:      file_cereal_proto_docs_wizard_proto_msgTypes,
	}.Build()
	File_cereal_proto_docs_wizard_proto = out.File
	file_cereal_proto_docs_wizard_proto_rawDesc = nil
	file_cereal_proto_docs_wizard_proto_goTypes = nil
	file_cereal_proto_docs_wizard_proto_depIdxs = nil
}
