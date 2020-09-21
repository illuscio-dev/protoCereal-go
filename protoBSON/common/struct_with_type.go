package common

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/proto"
	"reflect"
)

const TypeUrlField = "pb_type"

// We're going to use this struct to marshall and unmarshall the Any field
type typedData struct {
	// Holds the type url for the original message type of this message.
	TypeIdentifier string `bson:"pb_type"`
	// In order to inline the fields, we have to use a map or pointer type here, sadly.
	// We cannot use a pointer since we won't know the struct type ahead of time.
	// We COULD use BsonRaw type but we would not be able to inline the contained
	// fields, adding an additional layer of nesting.
	Data bson.M `bson:",inline"`
}

var typedDataType = reflect.TypeOf(typedData{})

type typeIdentifierField struct {
	// Holds the type url for the original message type of this message.
	TypeIdentifier string `bson:"pb_type"`
}

var typeIdentifierType = reflect.TypeOf(typeIdentifierField{})

func EncodeStructWithTypeInfo(
	context bsoncodec.EncodeContext,
	writer bsonrw.ValueWriter,
	value proto.Message,
	typeIdentifier string,
) error {
	// Marshall the inner message type
	messageBytes, err := bson.MarshalWithContext(context, value)
	if err != nil {
		return fmt.Errorf(
			"could not marshal payload, type '%v': %w",
			reflect.TypeOf(value),
			err,
		)
	}

	// Now Unmarshall this inner type into a generic bson document
	// TODO: write some custom marshalling logic here to extract the fields ourselves
	//   and make this round-trip unnecessary.
	messageM := make(bson.M, 0)
	err = bson.UnmarshalWithRegistry(context.Registry, messageBytes, messageM)
	if err != nil {
		return fmt.Errorf(
			"error unmarshalling payload, type '%v' to map: %w",
			reflect.TypeOf(value),
			err,
		)
	}
	dbValue := typedData{
		TypeIdentifier: typeIdentifier,
		Data:           messageM,
	}

	dbEncoder, err := context.LookupEncoder(typedDataType)
	if err != nil {
		return fmt.Errorf(
			"no encoder for intermediate type 'typedData': %w", err,
		)
	}

	err = dbEncoder.EncodeValue(context, writer, reflect.ValueOf(dbValue))
	if err != nil {
		return fmt.Errorf(
			"error encoding for intermediate type 'typedData': %w", err,
		)
	}

	return nil
}

// Decode a document with the type information appended as a field. typeResolver must
// hand back the (non-pointer) message type
func DecodeStructWithTypeInfo(
	context bsoncodec.DecodeContext,
	reader bsonrw.ValueReader,
	typeResolver func(string) (reflect.Type, error),
) (reflect.Value, error) {
	// First we are going to extract the url field
	urlData := new(typeIdentifierField)
	rawDataValue := reflect.ValueOf(urlData)

	// Get the raw bytes from the reader. We'll need to convert our reader into a
	// BytesReader
	bytesReader := ConvertToBytesReader(reader)
	bsonType, valueBytes, err := bytesReader.ReadValueBytes(make([]byte, 0))
	if err != nil {
		return reflect.Value{}, fmt.Errorf("error reading raw bytes: %w", err)
	}

	// Get the decoder for the intermediate struct
	decoder, err := context.LookupDecoder(typeIdentifierType)
	if err != nil {
		return reflect.Value{}, fmt.Errorf(
			"could not fetch decoder for struct type: %w", err,
		)
	}

	// Create a reader to do this first extraction
	reader = bsonrw.NewBSONValueReader(bsonType, valueBytes)
	err = decoder.DecodeValue(context, reader, rawDataValue.Elem())
	if err != nil {
		return reflect.Value{}, fmt.Errorf(
			"error decoding proto Any from bson: %w", err,
		)
	}

	// Now get the message type
	messageType, err := typeResolver(urlData.TypeIdentifier)
	if err != nil {
		// If the message type is not in our registry we can't decode it.
		return reflect.Value{}, fmt.Errorf(
			"error resolving ptotobuf message type '%v': %w",
			urlData.TypeIdentifier,
			err,
		)
	}

	// Next, load the bson document into the protobuf message type.
	// Let's get a new message of the correct type
	newMessageValue := reflect.New(messageType)
	decoder, err = context.LookupDecoder(messageType)
	if err != nil {
		return reflect.Value{}, fmt.Errorf(
			"could not fetch decoder for struct type '%v': %w",
			reflect.TypeOf(newMessageValue),
			err,
		)
	}

	// Make a second reader to extract the actual message.
	reader = bsonrw.NewBSONValueReader(bsonType, valueBytes)
	err = decoder.DecodeValue(context, reader, newMessageValue.Elem())
	if err != nil {
		return reflect.Value{}, fmt.Errorf(
			"error unmarshalling to Any proto payload: %w", err,
		)
	}

	return newMessageValue, nil
}
