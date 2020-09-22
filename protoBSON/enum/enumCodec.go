package enum

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/reflect/protoreflect"
	"reflect"
)

// Interface that generated enum types conform to.
type ProtoEnum interface {
	String() string
	Descriptor() protoreflect.EnumDescriptor
	Type() protoreflect.EnumType
	Number() protoreflect.EnumNumber
}

type CodecEnumStringer struct {
}

func (codec *CodecEnumStringer) EncodeValue(
	context bsoncodec.EncodeContext, writer bsonrw.ValueWriter, value reflect.Value,
) error {
	enumVal, ok := value.Interface().(ProtoEnum)
	if !ok {
		return fmt.Errorf(
			"non-enum type '%v' passed to enum codec",
			value.Type(),
		)
	}
	err := writer.WriteString(enumVal.String())
	if err != nil {
		return fmt.Errorf("error writing enum of type '%v'", value.Type())
	}
	return nil
}

func (codec *CodecEnumStringer) DecodeValue(
	context bsoncodec.DecodeContext, reader bsonrw.ValueReader, value reflect.Value,
) error {
	// Convert this value to it's enum interface.
	enumVal, ok := value.Interface().(ProtoEnum)
	if !ok {
		return fmt.Errorf(
			"non-enum type '%v' passed to enum decoder",
			value.Type(),
		)
	}

	// Read the serialized enum name and get it's proto descriptor
	enumName, err := reader.ReadString()
	if err != nil {
		return fmt.Errorf("error reading enum name: %w", err)
	}
	valueDescriptor := enumVal.Descriptor().Values().ByName(protoreflect.Name(enumName))

	// Set the int value for the enum
	value.SetInt(int64(valueDescriptor.Number()))

	return nil
}
