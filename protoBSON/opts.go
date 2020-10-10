package protoBson

import (
	"github.com/illuscio-dev/protoCereal-go/protoBSON/oneof"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"google.golang.org/protobuf/proto"
	"reflect"
)

// Holds options for registering codecs.
type Opts struct {
	addDefaultCodecs bool
	enumStrings      bool
	oneOfElementInfo []*oneof.ElementInfo
	oneOfMessages    []proto.Message
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

// Add a proto oneof element type and it's bson mapping for generating. This is only
// needed if a oneof element's type cannot be automatically inferred (you have
// registered a custom encoder for it.)
func (opts *Opts) WithOneOfElementMapping(
	innerValue interface{}, bsonType bsontype.Type, binarySubType byte,
) *Opts {
	innerType := reflect.TypeOf(innerValue)

	elementInfo := &oneof.ElementInfo{
		InnerGoType:   innerType,
		BsonType:      bsonType,
		BinarySubType: binarySubType,
	}
	opts.oneOfElementInfo = append(opts.oneOfElementInfo, elementInfo)
	return opts
}

// Extract one-of fields from these message types and register their encoders / decoders
// with the registry.
func (opts *Opts) WithOneOfFields(messages ...proto.Message) *Opts {
	opts.oneOfMessages = append(opts.oneOfMessages, messages...)
	return opts
}

// Add a new wrapper type (as wrapperspb.StringValue) that is not one of protoCereal's
// default wrapper codecs.
func (opts *Opts) WithCustomWrappers(wrapperMessages ...proto.Message) *Opts {
	opts.customWrappers = append(opts.customWrappers, wrapperMessages...)
	return opts
}

// Create a new mongo opts object with default values.
func NewMongoOpts() *Opts {
	opts := &Opts{
		addDefaultCodecs: true,
		enumStrings:      true,
		oneOfMessages:    make([]proto.Message, 0),
		customWrappers:   make([]proto.Message, 0),
		oneOfElementInfo: make([]*oneof.ElementInfo, 0),
	}
	return opts
}
