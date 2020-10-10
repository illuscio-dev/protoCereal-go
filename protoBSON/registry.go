package protoBson

import (
	"errors"
	"fmt"
	"github.com/illuscio-dev/protoCereal-go/messagesCereal"
	"github.com/illuscio-dev/protoCereal-go/protoBSON/enum"
	"github.com/illuscio-dev/protoCereal-go/protoBSON/oneof"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"reflect"
	"strings"
)

func validateWrapperType(wrapperType reflect.Type) error {
	// Check that this is a pointer to a struct
	if wrapperType.Kind() != reflect.Ptr ||
		wrapperType.Elem().Kind() != reflect.Struct {

		return fmt.Errorf(
			"custom wrapper type '%v' is not pointer to a struct", wrapperType,
		)
	}

	// Dereference to the underlying struct
	structType := wrapperType.Elem()

	// Iterate through all the fields, counting the public ones and remembering the
	// last one's name.
	fieldCount := structType.NumField()
	publicCount := 0
	publicFieldName := ""
	for i := 0; i < fieldCount; i++ {
		fieldInfo := structType.Field(i)

		firstLetter := string([]rune(fieldInfo.Name)[0])
		if strings.ToUpper(firstLetter) == firstLetter {
			publicCount++
			publicFieldName = fieldInfo.Name
		}
	}

	// Check that there is only one public field.
	if publicCount != 1 {
		return fmt.Errorf(
			"custom wrapper type '%v' must have exactly 1 public field, but"+
				" contains %v",
			wrapperType,
			publicCount,
		)
	}

	// Check that it is called 'Value' (conforming to the google wrapper type
	// convention).
	if publicFieldName != "Value" {
		return fmt.Errorf(
			"custom wrapper message '%v' does not have 'Value' field",
			wrapperType,
		)
	}

	return nil
}

// Register the bson codecs that come with protoCereal.
func registerProtoCerealCodecs(builder *bsoncodec.RegistryBuilder, opts *Opts) {
	builder.RegisterCodec(reflect.TypeOf(new(anypb.Any)), protoAnyCodec{})
	builder.RegisterCodec(reflect.TypeOf(new(messagesCereal.UUID)), protoUUIDCodec{})
	builder.RegisterCodec(
		reflect.TypeOf(new(messagesCereal.Decimal)), protoDecimalCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(messagesCereal.RawData)), protoRawDataCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(timestamppb.Timestamp)), protoTimestampCodec{},
	)
	// Wrapper types
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.BoolValue)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.BytesValue)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.FloatValue)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.DoubleValue)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.Int32Value)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.Int64Value)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.StringValue)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.UInt32Value)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.UInt64Value)), protoWrapperCodec{},
	)

}

func buildAndRegisterOneOfCodecs(builder *bsoncodec.RegistryBuilder, opts *Opts) error {
	// Register our one-of codecs with the registry
	for _, oneOfMessage := range opts.oneOfMessages {
		oneOfBuilders, err := oneof.CodecBuildersForMessage(
			oneOfMessage, opts.oneOfElementInfo,
		)
		if err != nil {
			return fmt.Errorf("error creating oneof codec: %w", err)
		}
		for _, thisOneOfBuilder := range oneOfBuilders {
			thisOneOfBuilder.Register(builder)
		}

	}

	return nil
}

func registerEnumStringCodec(builder *bsoncodec.RegistryBuilder, opts *Opts) {
	if !opts.enumStrings {
		return
	}

	enumCodec := new(enum.CodecEnumStringer)
	enumInterfaceType := reflect.TypeOf((*enum.ProtoEnum)(nil)).Elem()
	builder.RegisterHookEncoder(enumInterfaceType, enumCodec)
	builder.RegisterHookDecoder(enumInterfaceType, enumCodec)
}

func registerCustomWrappers(builder *bsoncodec.RegistryBuilder, opts *Opts) error {
	// Add custom wrapper type
	for _, wrapper := range opts.customWrappers {
		wrapperType := reflect.TypeOf(wrapper)
		err := validateWrapperType(wrapperType)
		if err != nil {
			return err
		}
		builder.RegisterCodec(wrapperType, protoWrapperCodec{})
	}

	return nil
}

// Register a the cereal codecs onto a registry builder.
func RegisterCerealCodecs(builder *bsoncodec.RegistryBuilder, opts *Opts) error {
	if builder == nil {
		return errors.New("registry builder cannot be nil")
	}

	if opts == nil {
		opts = NewMongoOpts()
	}

	// Default types
	if opts.addDefaultCodecs {
		bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(builder)
		bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(builder)
	}

	registerProtoCerealCodecs(builder, opts)
	if err := buildAndRegisterOneOfCodecs(builder, opts); err != nil {
		return err
	}

	registerEnumStringCodec(builder, opts)

	err := registerCustomWrappers(builder, opts)
	return err
}

func NewCerealRegistryBuilder(opts *Opts) (*bsoncodec.RegistryBuilder, error) {
	builder := bsoncodec.NewRegistryBuilder()
	err := RegisterCerealCodecs(builder, opts)
	if err != nil {
		return nil, err
	}

	return builder, err
}
