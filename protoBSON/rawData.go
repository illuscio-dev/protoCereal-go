package protoBson

import (
	"github.com/illuscio-dev/protoCereal-go/messagesCereal"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"reflect"
)

// Save the zero value for a nil bool wrapper.
var dataZeroValue = reflect.Zero(reflect.TypeOf(new(messagesCereal.RawData)))

// CODEC FOR MARSHALLING AND UNMARSHALLING RAW DATA
type protoRawDataCodec struct{}

func (codec protoRawDataCodec) EncodeValue(
	context bsoncodec.EncodeContext, writer bsonrw.ValueWriter, value reflect.Value,
) error {
	// Handle null value case
	if value.IsNil() {
		return writer.WriteNull()
	}

	valueRaw := value.Interface().(*messagesCereal.RawData)
	err := writer.WriteBinaryWithSubtype(valueRaw.Data, bsontype.BinaryGeneric)
	if err != nil {
		return err
	}
	return nil
}

func (codec protoRawDataCodec) DecodeValue(
	context bsoncodec.DecodeContext, reader bsonrw.ValueReader, value reflect.Value,
) error {
	// Handle null value case
	if reader.ReadNull() == nil {
		value.Set(dataZeroValue)
		return nil
	}

	bin, _, err := reader.ReadBinary()
	if err != nil {
		return err
	}

	protoVal := &messagesCereal.RawData{
		Data: bin,
	}

	value.Set(reflect.ValueOf(protoVal))
	return nil
}
