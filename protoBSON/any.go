package protoBson

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
	"reflect"
)

const AnyTypeUrlField = "pb_type_url"

// We're going to use this struct to marshall and unmarshall the Any field
type anyData struct {
	// Holds the type url for the original message type of this message.
	TypeUrl string `bson:"pb_type_url"`
	// In order to inline the fields, we have to use a map or pointer type here, sadly.
	// We cannot use a pointer since we won't know the struct type ahead of time.
	// We COULD use BsonRaw type but we would not be able to inline the contained
	// fields, adding an additional layer of nesting.
	Data bson.M `bson:",inline"`
}

var anyDataType = reflect.TypeOf(anyData{})

// Save the zero value for a nil any.
var anyZeroValue = reflect.Zero(reflect.TypeOf(new(anypb.Any)))

// Codec for marshalling and unmarshalling anypb.Any.
type protoAnyCodec struct {
}

func (codec protoAnyCodec) EncodeValue(
	context bsoncodec.EncodeContext, writer bsonrw.ValueWriter, value reflect.Value,
) error {
	if value.IsNil() {
		return writer.WriteNull()
	}

	// Type assert tp the any message type
	anyVal := value.Interface().(*anypb.Any)

	// Next we need to write the rest of the fields
	messageValue, err := anyVal.UnmarshalNew()
	if err != nil {
		return fmt.Errorf("could not unmarshall anypb.Any: %w", err)
	}

	// Marshall the inner message type
	messageBytes, err := bson.MarshalWithContext(context, messageValue)
	if err != nil {
		return fmt.Errorf(
			"could not marshal anypb.Any payload type '%v': %w",
			reflect.TypeOf(messageValue),
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
			"error unmarshalling anypb.Any payload type '%v' to map: %w",
			reflect.TypeOf(messageValue),
			err,
		)
	}

	dbValue := anyData{
		TypeUrl: anyVal.TypeUrl,
		Data:    messageM,
	}

	dbEncoder, err := context.LookupEncoder(anyDataType)
	if err != nil {
		return fmt.Errorf(
			"no encoder for anypb.Any intermediate type anyData: %w", err,
		)
	}

	err = dbEncoder.EncodeValue(context, writer, reflect.ValueOf(dbValue))
	if err != nil {
		return fmt.Errorf(
			"error encoding anypb.Any intermediate type anyData: %w", err,
		)
	}

	return nil
}

func (codec protoAnyCodec) DecodeValue(
	context bsoncodec.DecodeContext, reader bsonrw.ValueReader, value reflect.Value,
) error {
	// Check for nil value first.
	if reader.ReadNull() == nil {
		value.Set(anyZeroValue)
		return nil
	}

	// TODO: write some custom unmarshalling logic here to extract the fields ourselves
	//   and make this round-trip unnecessary.
	rawData := new(anyData)
	rawDataValue := reflect.ValueOf(rawData)

	// We can use the cached anyDataType var to ge the type here.
	decoder, err := context.LookupDecoder(anyDataType)
	if err != nil {
		return fmt.Errorf("could not fetch decoder for struct type: %w", err)
	}

	err = decoder.DecodeValue(context, reader, rawDataValue.Elem())
	if err != nil {
		return fmt.Errorf(
			"error decoding proto Any from bson: %w", err,
		)
	}

	// Now get the message type
	messageType, err := protoregistry.GlobalTypes.FindMessageByURL(rawData.TypeUrl)
	if err != nil {
		// If the message type is not in our registry we can't decode it.
		return fmt.Errorf(
			"error fetching ptotobuf message type '%v"+
				" descibed for anypb.Any value: %w",
			rawData.TypeUrl,
			err,
		)
	}

	// Next, load the bson document into the protobuf message type.
	// Let's get a new message of the correct type
	newMessage := messageType.New().Interface()

	// First we have to marshall back out the generic document
	messageBytes, err := bson.MarshalWithRegistry(context.Registry, rawData.Data)
	if err != nil {
		return fmt.Errorf("error marhsalling inner data to bson: %w", err)
	}

	err = bson.UnmarshalWithRegistry(context.Registry, messageBytes, newMessage)
	if err != nil {
		return fmt.Errorf("error unmarshalling to Any proto payload: %w", err)
	}

	anyValue, err := anypb.New(newMessage)
	if err != nil {
		return fmt.Errorf("error setting inner payload of Any message: %w", err)
	}

	value.Set(reflect.ValueOf(anyValue))

	return nil
}
