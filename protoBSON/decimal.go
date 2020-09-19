package protoBson

import (
	"fmt"
	"github.com/illuscio-dev/protoCereal-go/messagesCereal"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

var decimalZeroVal = reflect.Zero(reflect.TypeOf(new(messagesCereal.Decimal)))

type protoDecimalCodec struct{}

func (codec protoDecimalCodec) EncodeValue(
	context bsoncodec.EncodeContext, writer bsonrw.ValueWriter, value reflect.Value,
) error {
	if value.IsNil() {
		return writer.WriteNull()
	}

	decimalProto, ok := value.Interface().(*messagesCereal.Decimal)
	if !ok {
		return fmt.Errorf(
			"type '%v' passed to decimal codec was not protobuf decimal pointer",
			value.Type(),
		)
	}

	decimalBson := primitive.NewDecimal128(decimalProto.High, decimalProto.Low)

	err := writer.WriteDecimal128(decimalBson)
	if err != nil {
		return fmt.Errorf("error writing decimal value: %w", err)
	}

	return nil
}

func (codec protoDecimalCodec) DecodeValue(
	context bsoncodec.DecodeContext, reader bsonrw.ValueReader, value reflect.Value,
) error {
	if err := reader.ReadNull(); err == nil {
		value.Set(decimalZeroVal)
		return nil
	}

	decimalBson, err := reader.ReadDecimal128()
	if err != nil {
		return fmt.Errorf(
			"could not read decimal 128 for proto decimal: %w", err,
		)
	}

	high, low := decimalBson.GetBytes()

	decimalProto := &messagesCereal.Decimal{
		High: high,
		Low:  low,
	}

	value.Set(reflect.ValueOf(decimalProto))

	return nil
}
