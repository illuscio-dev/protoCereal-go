package protobson

import (
	"errors"
	"fmt"
	"github.com/illuscio-dev/protoCereal-go/cereal"
	"github.com/illuscio-dev/protoCereal-go/protobson/enum"
	"github.com/illuscio-dev/protoCereal-go/protobson/oneof"
	"github.com/illuscio-dev/protoCereal-go/protobson/wrapper"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"reflect"
)

// Register the bson codecs that come with protoCereal.
func registerProtoCerealCodecs(builder *bsoncodec.RegistryBuilder, opts *Opts) {
	builder.RegisterCodec(reflect.TypeOf(new(anypb.Any)), protoAnyCodec{})
	builder.RegisterCodec(reflect.TypeOf(new(cereal.UUID)), protoUUIDCodec{})
	builder.RegisterCodec(
		reflect.TypeOf(new(cereal.Decimal)), protoDecimalCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(cereal.RawData)), protoRawDataCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(timestamppb.Timestamp)), protoTimestampCodec{},
	)
	// Wrapper types
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.BoolValue)),
		wrapper.MustWrapperCodec(new(wrapperspb.BoolValue)),
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.BytesValue)),
		wrapper.MustWrapperCodec(new(wrapperspb.BytesValue)),
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.FloatValue)),
		wrapper.MustWrapperCodec(new(wrapperspb.FloatValue)),
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.DoubleValue)),
		wrapper.MustWrapperCodec(new(wrapperspb.DoubleValue)),
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.Int32Value)),
		wrapper.MustWrapperCodec(new(wrapperspb.Int32Value)),
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.Int64Value)),
		wrapper.MustWrapperCodec(new(wrapperspb.Int64Value)),
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.StringValue)),
		wrapper.MustWrapperCodec(new(wrapperspb.StringValue)),
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.UInt32Value)),
		wrapper.MustWrapperCodec(new(wrapperspb.UInt32Value)),
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrapperspb.UInt64Value)),
		wrapper.MustWrapperCodec(new(wrapperspb.UInt64Value)),
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
	for _, wrapperType := range opts.customWrappers {
		wrapperCodec, err := wrapper.NewWrapperCodec(wrapperType)
		if err != nil {
			return fmt.Errorf("error creating custom wrapper codec: %w", err)
		}
		builder.RegisterCodec(reflect.TypeOf(wrapperType), wrapperCodec)
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
