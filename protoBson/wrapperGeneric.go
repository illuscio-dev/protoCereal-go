package protoBson

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"reflect"
)

// Codec for encoding gRPC wrappers.BoolValue, which exposes an API for having nullable
// boolean values.
type protoWrapperCodec struct {
}

func (codec protoWrapperCodec) EncodeValue(
	context bsoncodec.EncodeContext, writer bsonrw.ValueWriter, value reflect.Value,
) error {
	// If this is a nil value, write a nil value
	if value.IsNil() {
		return writer.WriteNull()
	}

	// get the inner value of the wrapper
	innerValue := value.Elem().FieldByName("Value")

	// get the encoder for this inner value so we can encode it directly
	innerEncoder, err := context.LookupEncoder(innerValue.Type())
	if err != nil {
		return fmt.Errorf(
			"could not get encoder for inner type '%v' of protobuf wrapper"+
				" type '%v': %w",
			innerValue.Type(),
			value.Type(),
			err,
		)
	}

	// encode the inner value directly
	err = innerEncoder.EncodeValue(context, writer, innerValue)
	if err != nil {
		return fmt.Errorf(
			"error encoding inner type '%v' of protobuf wrapper type '%v': %w",
			innerValue.Type(),
			value.Type(),
			err,
		)
	}

	return nil
}

func (codec protoWrapperCodec) DecodeValue(
	context bsoncodec.DecodeContext, reader bsonrw.ValueReader, value reflect.Value,
) error {
	valueType := value.Type()

	// If there is no error reading null then we have a null value
	if reader.ReadNull() == nil {
		// Set to a nil pointer by getting the zero value of our type
		value.Set(reflect.Zero(value.Type()))
		return nil
	}

	// Create an initialized pointer reference.
	newWrapper := reflect.New(valueType.Elem())

	// Get the reflect.Value for the value field of the new wrapper object.
	innerValue := newWrapper.Elem().FieldByName("Value")

	// Get the encoder for this inner value so we can encode it directly
	innerDecoder, err := context.LookupDecoder(innerValue.Type())
	if err != nil {
		return fmt.Errorf(
			"could not get encoder for inner type '%v' of protobuf wrapper"+
				" type '%v': %w",
			innerValue.Type(),
			value.Type(),
			err,
		)
	}

	// encode the inner value directly
	err = innerDecoder.DecodeValue(context, reader, innerValue)
	if err != nil {
		return fmt.Errorf(
			"error encoding inner type '%v' of protobuf wrapper type '%v': %w",
			innerValue.Type(),
			value.Type(),
			err,
		)
	}

	// Set the value to our new wra[[er value
	value.Set(newWrapper)
	return nil
}
