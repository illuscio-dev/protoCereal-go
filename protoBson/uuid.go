package protoBson

import (
	"github.com/illuscio-dev/protoCereal-go/cereal"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"reflect"
)

// Save the zero value for a nil timestamp.
var uuidZeroValue = reflect.Zero(reflect.TypeOf(new(cereal.UUID)))

// CODEC FOR MARSHALLING AND UNMARSHALLING UUIDs
type protoUUIDCodec struct{}

func (codec protoUUIDCodec) DecodeValue(
	context bsoncodec.DecodeContext, reader bsonrw.ValueReader, value reflect.Value,
) error {
	// Handle nil value
	if reader.ReadNull() == nil {
		value.Set(uuidZeroValue)
		return nil
	}

	bin, _, err := reader.ReadBinary()
	if err != nil {
		return err
	}

	protoVal := &cereal.UUID{
		Value: bin,
	}

	value.Set(reflect.ValueOf(protoVal))
	return nil
}

func (codec protoUUIDCodec) EncodeValue(
	context bsoncodec.EncodeContext, writer bsonrw.ValueWriter, value reflect.Value,
) error {
	// Handle nil value
	if value.IsNil() {
		return writer.WriteNull()
	}

	valueUUID := value.Interface().(*cereal.UUID)
	err := writer.WriteBinaryWithSubtype(valueUUID.Value, bsontype.BinaryUUID)
	if err != nil {
		return err
	}
	return nil
}
