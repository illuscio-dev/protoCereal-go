package protoBson

import (
	"errors"
	"fmt"
	"github.com/illuscio-dev/protoCereal-go/messagesCereal"
	"github.com/illuscio-dev/protoCereal-go/protoBSON/enum"
	"github.com/illuscio-dev/protoCereal-go/protoBSON/oneof"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"reflect"
	"strings"
)

// Holds options for registering codecs.
type Opts struct {
	addDefaultCodecs bool
	enumStrings      bool
	oneOfBuilders    []*oneof.CodecBuilder
	customWrappers   []proto.Message
}

// Whether to add the default mongo codecs to the registry. Default: true.
func (opts *Opts) WithAddDefaultCodecs(add bool) *Opts {
	opts.addDefaultCodecs = add
	return opts
}

// Whether to add the default mongo codecs to the registry. Default: true.
func (opts *Opts) WithEnumStrings(enable bool) *Opts {
	opts.enumStrings = enable
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

// Add a new wrapper type (as wrappers.StringValue) that is not one of protoCereal's
// default wrapper codecs.
func (opts *Opts) WithCustomWrappers(wrapperMessages ...proto.Message) *Opts {
	opts.customWrappers = append(opts.customWrappers, wrapperMessages...)
	return opts
}

// Create a new mongo opts object with default values.
func NewMongoOpts() *Opts {
	return new(Opts).WithAddDefaultCodecs(true)
}

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

	// Register our one-of codecs with the registry
	for _, oneOfBuilder := range opts.oneOfBuilders {
		oneOfBuilder.Register(builder)
	}

	if opts.enumStrings {
		enumCodec := new(enum.CodecEnumStringer)
		enumInterfaceType := reflect.TypeOf((*enum.ProtoEnum)(nil)).Elem()
		builder.RegisterHookEncoder(enumInterfaceType, enumCodec)
		builder.RegisterHookDecoder(enumInterfaceType, enumCodec)
	}

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
