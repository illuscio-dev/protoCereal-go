package protoBson

import (
	"fmt"
	"github.com/illuscio-dev/protoCereal-go/protoBson/common"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
	"reflect"
)

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

	err = common.EncodeStructWithTypeInfo(context, writer, messageValue, anyVal.TypeUrl)
	if err != nil {
		return fmt.Errorf("error encoding anypb.Any: %w", err)
	}
	return nil
}

func messageTypeFromURL(typeURL string) (reflect.Type, error) {
	messageType, err := protoregistry.GlobalTypes.FindMessageByURL(typeURL)
	if err != nil {
		return nil, fmt.Errorf(
			"could not find proto message for url '%v': %w", typeURL, err,
		)
	}

	return reflect.TypeOf(messageType.New().Interface()).Elem(), nil
}

func (codec protoAnyCodec) DecodeValue(
	context bsoncodec.DecodeContext, reader bsonrw.ValueReader, value reflect.Value,
) error {
	if reader.ReadNull() == nil {
		value.Set(reflect.Zero(value.Type()))
		return nil
	}

	anyPayload, err := common.DecodeStructWithTypeInfo(
		context, reader, messageTypeFromURL,
	)
	if err != nil {
		return fmt.Errorf("error extracting typed message: %w", err)
	}

	newMessage, ok := anyPayload.Interface().(proto.Message)
	if !ok {
		return fmt.Errorf(
			"unmarshalled message of type '%v' could not be converted"+
				" to proto.Message",
			anyPayload.Type(),
		)
	}

	anyValue, err := anypb.New(newMessage)
	if err != nil {
		return fmt.Errorf("error setting inner payload of Any message: %w", err)
	}

	value.Set(reflect.ValueOf(anyValue))

	return nil
}
