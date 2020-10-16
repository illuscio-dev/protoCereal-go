package protoBson

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/proto"
	"reflect"
	"strings"
)

var ErrWrapperMessageType = errors.New("wrapper codec expected pointer to struct")
var ErrWrapperEmbeddedField = errors.New(
	"wrapper messages should not have embedded fields",
)
var ErrWrapperPublicFieldCount = errors.New(
	"wrapper expected to have exactly 1 public field",
)

// Codec for encoding gRPC wrappers.BoolValue, which exposes an API for having nullable
// boolean values.
type protoWrapperCodec struct {
	FieldIndex int
}

func (codec protoWrapperCodec) EncodeValue(
	context bsoncodec.EncodeContext, writer bsonrw.ValueWriter, value reflect.Value,
) error {
	// If this is a nil value, write a nil value
	if value.IsNil() {
		return writer.WriteNull()
	}

	// get the inner value of the wrapper
	innerValue := value.Elem().Field(codec.FieldIndex)

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

// Builds the wrapper codec for a given proto message.
//
// Wrapper messages must:
//
//	- Have exactly 1 non-embedded public field. This field's value will be extracted
//	  and marshalled as the bson value.
func newWrapperCodec(protoMessage proto.Message) (protoWrapperCodec, error) {
	pointerType := reflect.TypeOf(protoMessage)
	if pointerType.Kind() != reflect.Ptr {
		return protoWrapperCodec{}, fmt.Errorf(
			"%w, got '%v'", ErrWrapperMessageType, pointerType,
		)
	}

	structType := pointerType.Elem()
	if structType.Kind() != reflect.Struct {
		return protoWrapperCodec{}, fmt.Errorf(
			"%w, got '%v'", ErrWrapperMessageType, pointerType,
		)
	}

	pubCount := 0
	var fieldIndex int
	for i := 0; i < structType.NumField(); i++ {
		fieldInfo := structType.Field(i)
		fieldStartsWith := string([]rune(fieldInfo.Name)[0])
		if strings.ToUpper(fieldStartsWith) != fieldStartsWith {
			continue
		}
		pubCount++
		if len(fieldInfo.Index) != 1 {
			return protoWrapperCodec{}, fmt.Errorf(
				"%w, '%v' of type '%v' is embedded",
				ErrWrapperEmbeddedField,
				fieldInfo.Name,
				pointerType,
			)
		}
		fieldIndex = fieldInfo.Index[0]
	}

	if pubCount != 1 {
		return protoWrapperCodec{}, fmt.Errorf(
			"%w, found %v public fields for type '%v'",
			ErrWrapperPublicFieldCount,
			pubCount,
			pointerType,
		)
	}

	return protoWrapperCodec{FieldIndex: fieldIndex}, nil
}

func mustWrapperCodec(protoMessage proto.Message) protoWrapperCodec {
	newCodec, err := newWrapperCodec(protoMessage)
	if err != nil {
		panic(err)
	}
	return newCodec
}
