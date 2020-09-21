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
type Opts struct {
	addDefaultCodecs bool
	oneOfBuilders    []*oneof.CodecBuilder
}

// Whether to add the default mongo codecs to the registry. Default: true.
func (opts *Opts) WithAddDefaultCodecs(add bool) *Opts {
	opts.addDefaultCodecs = add
	return opts
}

// Extract one-of fields from these message types and register their encoders / decoders
// with the registry.
func (opts *Opts) WithOneOfFields(messages ...proto.Message) *Opts {
	for _, thisMessage := range messages {
		oneOfBuilders, err := oneof.CodecBuildersForMessage(thisMessage)
		if err != nil {
			panic(fmt.Errorf("error creating oneof codec: %w", err))
		}
		opts.oneOfBuilders = append(opts.oneOfBuilders, oneOfBuilders...)
	}

	return opts
}

// Create a new mongo opts object with default values.
func NewMongoOpts() *Opts {
	return new(Opts).WithAddDefaultCodecs(true)
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
