package protosql

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"reflect"
)

// protoEnum is an interface tat describes the proto-reflect methods on a protobuf Enum
// value.
type protoEnum interface {
	Descriptor() protoreflect.EnumDescriptor
	Type() protoreflect.EnumType
	Number() protoreflect.EnumNumber
}

// EnumStringMarshaller marshals and unmarshalls a string for sql db values.
type EnumStringMarshaller struct {
	// Enum is the value to encode as a string when marshalling.
	//
	// When unmarshalling, this field must contain a value of the desired Enum type
	// to receive.
	Enum protoEnum
}

// Value encodes the protobuf Enum to a string.
func (enum EnumStringMarshaller) Value() (driver.Value, error) {
	// If this is nil, return nil.
	if enum.Enum == nil {
		return nil, nil
	}

	// Get the proto-budd defined name for the Enum value.
	name := enum.Enum.Descriptor().Values().Get(int(enum.Enum.Number())).Name()
	return string(name), nil
}

// Scan received a string value from the database and transforms it into a new value.
func (enum *EnumStringMarshaller) Scan(src interface{}) error {
	valueTarget := reflect.ValueOf(&enum.Enum).Elem()
	if src == nil {
		// Set the enum string marshaller to a fresh underlying value, thereby zeroing
		// out the inner interface to account for a nil.
		*enum = EnumStringMarshaller{}
		return nil
	}

	srcVal, ok := src.(string)
	if !ok {
		return newScanTypeErr("", src, enum.Enum)
	}

	if enum.Enum == nil {
		return errors.New(
			"cannot unmarshal proto Enum, Enum field is not set with concrete " +
				"value for type inspection",
		)
	}

	// Get the descriptor of the Enum value based on our string.
	enumDescriptor := enum.Enum.
		// Get the Enum descriptor of the type we are trying to
		Descriptor().
		Values().
		ByName(protoreflect.Name(srcVal))

	// Get the type of the enum.
	enumType := reflect.TypeOf(enum.Enum)

	// If the descriptor is nil, that means that value does not match any of the
	// possible Enum values.
	if enumDescriptor == nil {
		return fmt.Errorf(
			"received string does not match any known Enum names for type '%v'",
			enumType,
		)
	}

	// Convert the int32 representation of the enumNumber to the enum type.
	enumVal := reflect.ValueOf(int32(enumDescriptor.Number())).Convert(enumType)

	// Grab a pointer to the enum value field and set it's underlying value of our
	// converted enum value.
	valueTarget.Set(enumVal)

	return nil
}

// Enum creates a new EnumStringMarshaller for marshalling or unmarshalling protobuf
// enums to and from strings.
//
// To unmarshal an enum, a concrete value of the TYPE of enum you are expecting must be
// provided. This inner value will be nilled if the db passed back a null value.
//
// When marshalling, if enumVal is nil, a NULL value will be added to the database.
func Enum(enumVal protoEnum) *EnumStringMarshaller {
	return &EnumStringMarshaller{Enum: enumVal}
}
