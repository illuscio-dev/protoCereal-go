package protoBson

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/illuscio-dev/protoCereal-go/messagesCereal"
	"github.com/illuscio-dev/protoCereal-go/protoBSON/oneof"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
)

// Holds options for registering codecs.
type MongoOpts struct {
	addDefaultCodecs bool
	oneOfBuilders    []*oneof.CodecBuilder
}

// Whether to add the default mongo codecs to the registry. Default: true.
func (opts *MongoOpts) WithAddDefaultCodecs(add bool) *MongoOpts {
	opts.addDefaultCodecs = add
	return opts
}

// Whether to add the default mongo codecs to the registry. Default: true.
func (opts *MongoOpts) WithOneOfType(
	oneOfInterface reflect.Type,
	oneOfTypes ...reflect.Type,
) *MongoOpts {
	builder, err := oneof.NewCodecBuilder(oneOfInterface)
	if err != nil {
		panic(fmt.Errorf("error creating oneOf builder: %w", err))
	}
	err = builder.AutoRegisterOneOfTypes(oneOfTypes...)
	if err != nil {
		panic(
			fmt.Errorf(
				"error registering oneOfTypes with codec builder: %w", err,
			),
		)
	}

	opts.oneOfBuilders = append(opts.oneOfBuilders, builder)
	return opts
}

// Extract one-of fields from this message type and register their encoders / decoders
func (opts *MongoOpts) WithOneOfFields(messages ...proto.Message) *MongoOpts {
	for _, thisMessage := range messages {
		oneOfBuilder := new(oneof.CodecBuilder)
		err := oneOfBuilder.RegisterOneOfFields(thisMessage)
		if err != nil {
			panic(fmt.Errorf("error creating oneof codec: %w", err))
		}
		opts.oneOfBuilders = append(opts.oneOfBuilders, oneOfBuilder)
	}

	return opts
}

// Create a new mongo opts object with default values.
func NewMongoOpts() *MongoOpts {
	return new(MongoOpts).WithAddDefaultCodecs(true)
}

// Register a the cereal codecs onto a registry builder.
func RegisterCerealCodecs(builder *bsoncodec.RegistryBuilder, opts *MongoOpts) error {
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

	// Custom types
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
		reflect.TypeOf(new(wrappers.BoolValue)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrappers.BytesValue)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrappers.FloatValue)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrappers.DoubleValue)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrappers.Int32Value)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrappers.Int64Value)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrappers.StringValue)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrappers.UInt32Value)), protoWrapperCodec{},
	)
	builder.RegisterCodec(
		reflect.TypeOf(new(wrappers.UInt64Value)), protoWrapperCodec{},
	)

	// Register our one-of codecs with the registry
	for _, oneOfBuilder := range opts.oneOfBuilders {
		oneOfBuilder.Register(builder)
	}

	return nil
}
