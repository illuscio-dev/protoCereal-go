package protobson

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"time"
)

// Save the zero value for a nil timestamp.
var timestampZeroValue = reflect.Zero(reflect.TypeOf(new(timestamp.Timestamp)))

// CODEC FORM MARSHALLING AN UNMARSHALLING timestamp.Timestamp.
type protoTimestampCodec struct {
}

func (codec protoTimestampCodec) DecodeValue(
	context bsoncodec.DecodeContext, reader bsonrw.ValueReader, value reflect.Value,
) error {
	// Check for nil value first.
	if reader.ReadNull() == nil {
		value.Set(timestampZeroValue)
		return nil
	}

	var timeVal time.Time

	timeDecoder, err := context.LookupDecoder(reflect.TypeOf(timeVal))
	if err != nil {
		return nil
	}

	timeValReflect := reflect.ValueOf(&timeVal)
	err = timeDecoder.DecodeValue(context, reader, timeValReflect.Elem())
	if err != nil {
		return err
	}

	timestampVal := timestamppb.New(timeVal)

	value.Set(reflect.ValueOf(timestampVal))
	return nil
}

func (codec protoTimestampCodec) EncodeValue(
	context bsoncodec.EncodeContext, writer bsonrw.ValueWriter, value reflect.Value,
) error {
	// Check for nil value first.
	if value.IsNil() {
		return writer.WriteNull()
	}

	incoming := value.Interface().(*timestamp.Timestamp)
	timeVal := incoming.AsTime()

	timeEncoder, err := context.LookupEncoder(reflect.TypeOf(timeVal))
	if err != nil {
		return nil
	}

	return timeEncoder.EncodeValue(context, writer, reflect.ValueOf(timeVal))
}
