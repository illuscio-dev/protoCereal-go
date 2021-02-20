package oneof

import (
	"fmt"
	"github.com/illuscio-dev/protoCereal-go/protobson/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/proto"
	"reflect"
)

// Codec for handling encoding and decoding of one-of fields. This default encoder is
// restrained by needing to map 1-to-1 bson types to decode types. Since the original
// type information is lost when the original value is encoded, we need to be able
// to infer from the bson value what the original value was.
//
// This means that while we can support one-of unions like [string, int], we cannot
// support unions like [float, double] as both are encoded to a bsontype.Double, so we
// cannot know on the decode side what sort of value to use.
type oneOfCodec struct {
	oneOfInterface reflect.Type
	// Map of the encoded bson type to the oneof field type it decodes to
	decodeMap map[bsonTypeKey]reflect.Type
	// For message types that can be encoded into an embedded document, we actually CAN
	// disambiguate between the target decode types, since we can embed the type name
	// in the written document. This map contains a map from the full proto type name
	// to it's **de-referenced one of wrapper type**.
	embeddedDocTypeMap map[string]reflect.Type
}

func (codec *oneOfCodec) checkEncodeWithType(
	wrapperValue reflect.Value,
	innerValue reflect.Value,
) bool {
	if len(codec.embeddedDocTypeMap) <= 1 {
		return false
	}

	if innerValue.Kind() != reflect.Ptr || innerValue.Elem().Kind() != reflect.Struct {
		return false
	}

	// Dereference through the interface and pointer to get the wrapper type
	wrapperType := wrapperValue.Elem().Elem().Type()
	for _, embeddedType := range codec.embeddedDocTypeMap {
		if embeddedType == wrapperType {
			return true
		}
	}
	return false
}

func (codec *oneOfCodec) encodeWhichEmbeddedDocument(
	context bsoncodec.EncodeContext,
	writer bsonrw.ValueWriter,
	innerValue reflect.Value,
) error {
	// Use the proto message name as the identifier
	typeIdentifier := embeddedDocKeyFromValue(innerValue)
	protoMessage := innerValue.Interface().(proto.Message)

	err := common.EncodeStructWithTypeInfo(
		context, writer, protoMessage, typeIdentifier,
	)
	if err != nil {
		return fmt.Errorf(
			"error encoding subdocument with type identifier '%v': %w",
			typeIdentifier,
			err,
		)
	}
	return nil
}

func (codec *oneOfCodec) encodeGenericInnerValue(
	context bsoncodec.EncodeContext,
	writer bsonrw.ValueWriter,
	innerValue reflect.Value,
) error {
	encoder, err := context.LookupEncoder(innerValue.Type())
	if err != nil {
		return fmt.Errorf(
			"encoder lookup error: %w",
			err,
		)
	}

	err = encoder.EncodeValue(context, writer, innerValue)
	if err != nil {
		return fmt.Errorf("encoder err: %w", err)
	}

	return nil
}

func (codec *oneOfCodec) EncodeValue(
	context bsoncodec.EncodeContext, writer bsonrw.ValueWriter, value reflect.Value,
) error {
	if value.IsNil() {
		return writer.WriteNull()
	}

	// The value we get is going to be an interface value to a pointer, so we need
	// to elem() it once to get past the interface, then again to dereference the
	// pointer
	innerValue := value.Elem().Elem().Field(0)
	// If this is a message type and the value is nil, write nil
	if innerValue.Kind() == reflect.Ptr && innerValue.IsNil() {
		return writer.WriteNull()
	}

	encodeWithType := codec.checkEncodeWithType(value, innerValue)

	var err error
	if encodeWithType {
		err = codec.encodeWhichEmbeddedDocument(context, writer, innerValue)
		if err != nil {
			return err
		}
	} else {
		err = codec.encodeGenericInnerValue(context, writer, innerValue)
	}

	if err != nil {
		return fmt.Errorf(
			"error marshalling inner '%v' value for oneof wrapper '%v': %w",
			innerValue.Type(),
			value.Type(),
			err,
		)
	}

	return nil
}

// This function handled decoding binary values
func (codec *oneOfCodec) decodeBinaryInnerValue(
	context bsoncodec.DecodeContext,
	binData []byte,
	subType byte,
	innerValue reflect.Value,
	innerDecoder bsoncodec.ValueDecoder,
) error {
	// Because we've already read the value to get the subtype, we can't read it again.
	// We need to re-marshall the value and make a new reader to pass to the decoder.
	binValue := primitive.Binary{
		Subtype: subType,
		Data:    binData,
	}

	bsonType, bsonBytes, err := bson.MarshalValueWithRegistry(
		context.Registry, binValue,
	)
	if err != nil {
		return fmt.Errorf("error re-marshalling bin value: %w", err)
	}

	newReader := bsonrw.NewBSONValueReader(bsonType, bsonBytes)
	err = innerDecoder.DecodeValue(context, newReader, innerValue)
	if err != nil {
		return fmt.Errorf("error decoding binary val: %w", err)
	}

	return nil
}

func (codec *oneOfCodec) resolveEmbeddedType(typeKey string) (reflect.Type, error) {
	wrapperType, ok := codec.embeddedDocTypeMap[typeKey]
	if !ok {
		return nil, fmt.Errorf("embedded type not found: %v", typeKey)
	}

	embeddedType := wrapperType.Field(0).Type.Elem()

	return embeddedType, nil
}

func (codec *oneOfCodec) decodeWhichEmbeddedDocument(
	context bsoncodec.DecodeContext,
	reader bsonrw.ValueReader,
) (reflect.Value, error) {
	embeddedValue, err := common.DecodeStructWithTypeInfo(
		context, reader, codec.resolveEmbeddedType,
	)
	if err != nil {
		return reflect.Value{}, fmt.Errorf(
			"error decoding oneof embedded doc: %w", err,
		)
	}

	wrapperKey := embeddedDocKeyFromValue(embeddedValue)
	wrapperType, ok := codec.embeddedDocTypeMap[wrapperKey]
	if !ok {
		return reflect.Value{}, fmt.Errorf(
			"could not find one-of wrapper type for proto message type '%v'",
			wrapperKey,
		)
	}

	wrapperValue := reflect.New(wrapperType)
	wrapperValue.Elem().Field(0).Set(embeddedValue)

	return wrapperValue, nil
}

func (codec *oneOfCodec) decodeGeneralOneOf(
	context bsoncodec.DecodeContext, reader bsonrw.ValueReader,
) (reflect.Value, error) {
	bsonType := reader.Type()
	binaryType := byte(0x0)
	var binaryData []byte

	// If the bson type is binary, we have to read the value to discover the sub-type.
	if bsonType == bsontype.Binary {
		var err error
		binaryData, binaryType, err = reader.ReadBinary()
		if err != nil {
			return reflect.Value{}, fmt.Errorf(
				"error getting binary subtype: %w", err,
			)
		}
	}

	decodeType, ok := codec.decodeMap[newBsonTypeKey(bsonType, binaryType)]

	if !ok {
		return reflect.Value{}, fmt.Errorf(
			"oneof codec for oneof interface '%v' cannot decode bson type: '%v'",
			codec.oneOfInterface,
			bsonType,
		)
	}

	// Create a new value, this will create a pointer to a pointer
	newWrapper := reflect.New(decodeType)
	innerValue := newWrapper.Elem().Field(0)

	innerDecoder, err := context.LookupDecoder(innerValue.Type())
	if err != nil {
		return reflect.Value{}, fmt.Errorf(
			"could not find decoder for inner type '%v' of oneof type '%v': %w",
			innerValue.Type(),
			newWrapper.Type(),
			err,
		)
	}

	switch bsonType {
	case bsontype.Binary:
		err = codec.decodeBinaryInnerValue(
			context, binaryData, binaryType, innerValue, innerDecoder,
		)
	default:
		err = innerDecoder.DecodeValue(context, reader, innerValue)
	}

	return newWrapper, nil
}

func (codec *oneOfCodec) DecodeValue(
	context bsoncodec.DecodeContext, reader bsonrw.ValueReader, value reflect.Value,
) error {
	// If the value is null, set the value to it's zero value
	if reader.Type() == bsontype.Null {
		_ = reader.ReadNull()
		value.Set(reflect.Zero(value.Type()))
		return nil
	}

	var newWrapperValue reflect.Value
	var err error
	if reader.Type() == bsontype.EmbeddedDocument && len(codec.embeddedDocTypeMap) > 1 {
		newWrapperValue, err = codec.decodeWhichEmbeddedDocument(context, reader)
	} else {
		newWrapperValue, err = codec.decodeGeneralOneOf(context, reader)
	}

	if err != nil {
		return fmt.Errorf(
			"could not decode inner type of oneof type '%v': %w",
			value.Type(),
			err,
		)
	}

	// Set the wrapper value
	value.Set(newWrapperValue)
	return nil
}
