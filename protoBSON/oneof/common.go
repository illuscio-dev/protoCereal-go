package oneof

import (
	"google.golang.org/protobuf/proto"
	"reflect"
)

// Creates a key for an inner message type
func embeddedDocKeyFromType(innerMessageType reflect.Type) string {
	newValue := reflect.New(innerMessageType)
	return embeddedDocKeyFromValue(newValue)
}

func embeddedDocKeyFromValue(innerMessageType reflect.Value) string {
	fullName := innerMessageType.Interface().(proto.Message).
		ProtoReflect().
		Descriptor().
		FullName()
	return string(fullName)
}

// Extract struct from wrapper type. Must pass in pointer type to wrapper.
func structTypeFromOneOfWrapperType(wrapper reflect.Type) reflect.Type {
	return wrapper.
		// Dereference the wrapper
		Elem().
		// Get the single field it contains
		Field(0).
		// Get that fields type
		Type.
		// Dereference the pointer of the type to get at the underlying struct type.
		Elem()
}
