package oneof

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"reflect"
)

var ErrWrongOneOfInterface = errors.New("oneof interface not implemented")

type ElementInfo struct {
	InnerGoType   reflect.Type
	BsonType      bsontype.Type
	BinarySubType byte
}

type CodecBuilder struct {
	// The master interface for the possible field values.
	oneOfInterface reflect.Type
	// Map of the encoded bson type to the oneof field type it decodes to
	decodeMap map[bsonTypeKey]reflect.Type
	// For message types that can be encoded into an embedded document, we actually CAN
	// disambiguate between the target decode types, since we can embed the type name
	// in the written document.
	embeddedDocTypeMap map[string]reflect.Type
	// Manual element info passed in from higher-level builders. This may include
	// elements from other one-ofs.
	innerTypeMap map[reflect.Type][]bsonTypeKey
}

// Validate that a type conforms to the expected constraints to a concrete one-of
// wrapper type spat out by the protoc codegen.
func (builder *CodecBuilder) validateValueWrapperType(oneOfType reflect.Type) error {
	// Our concrete type must implement our oneof type
	if !oneOfType.Implements(builder.oneOfInterface) {
		return fmt.Errorf(
			"%w: '%v' does not implement '%v'",
			ErrWrongOneOfInterface,
			oneOfType,
			builder.oneOfInterface,
		)
		// Our concrete type must be a pointer
	}
	return nil
}

// Automatically register a new concrete one-of wrapper type for the interface
// we are building this codec for.
func (builder *CodecBuilder) deduceConcreteTypeEncoding(
	oneOfType reflect.Type,
) error {
	var err error
	if err = builder.validateValueWrapperType(oneOfType); err != nil {
		return err
	}

	innerType := oneOfType.Elem().Field(0).Type

	// first see if this inner type has been manually added
	bsonTypes, known := builder.innerTypeMap[innerType]

	// First try to register it if it's a known type that serialized to something
	// other than an embedded doc
	if !known {
		bsonTypes, known = builder.deduceKnownTypeEncoding(oneOfType, innerType)
	}

	// If it's not known, we're next try to register it as an unknown type.
	if !known {
		bsonTypes, err = builder.deduceUnknownTypeEncoding(oneOfType, innerType)
		if err != nil {
			return err
		}
	}

	for _, thisBsonKey := range bsonTypes {
		err := builder.AddConcrete(
			oneOfType, thisBsonKey.BsonType(), thisBsonKey.BinaryType(),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

//revive:disable:cyclomatic

// automatically register a concrete one-of type for an inner type for which we don't
// have privileged information about what the inner type encodes / decodes from. We
// make a best-guess.
func (builder *CodecBuilder) deduceUnknownTypeEncoding(
	oneOfType reflect.Type, innerType reflect.Type,
) (bsonTypes []bsonTypeKey, err error) {
	switch innerType.Kind() {
	case reflect.Bool:
		bsonTypes = []bsonTypeKey{newSimpleKey(bsontype.Boolean)}
	case reflect.Int32:
		bsonTypes = []bsonTypeKey{newSimpleKey(bsontype.Int32)}
	// Because ints & units are compressed where possible by the bson encoder, we need
	// to take up both int32 and int64 for UInt32, UInt64 and int64
	case reflect.Int64:
		bsonTypes = []bsonTypeKey{
			newSimpleKey(bsontype.Int32), newSimpleKey(bsontype.Int64),
		}
	case reflect.Uint32:
		bsonTypes = []bsonTypeKey{
			newSimpleKey(bsontype.Int32), newSimpleKey(bsontype.Int64),
		}
	case reflect.Uint64:
		bsonTypes = []bsonTypeKey{
			newSimpleKey(bsontype.Int32), newSimpleKey(bsontype.Int64),
		}
	case reflect.String:
		bsonTypes = []bsonTypeKey{newSimpleKey(bsontype.String)}
	case reflect.Float64:
		bsonTypes = []bsonTypeKey{newSimpleKey(bsontype.Double)}
	case reflect.Float32:
		bsonTypes = []bsonTypeKey{newSimpleKey(bsontype.Double)}
	// NOTE: bytes are always represented as a []uint8. Because repeated fields are
	// disallowed in one-ofs, we can assume an array represents a bytes field.
	case reflect.Slice:
		if innerType != bytesFieldType {
			return nil, fmt.Errorf(
				"inner slice type '%v' is not a bytes field for"+
					" oneof wrapper '%v'",
				innerType,
				oneOfType,
			)
		}
		bsonTypes = []bsonTypeKey{newBsonTypeKey(bsontype.Binary, 0x0)}
	case reflect.Ptr:
		// If this inner type is not a struct, it's out of spec for protobuf structures
		// and we cannot handle it.
		if innerType.Elem().Kind() != reflect.Struct {
			return nil, fmt.Errorf(
				"inner pointer type '%v' does not point to struct in"+
					" oneof wrapper '%v'",
				innerType,
				oneOfType,
			)
		}
		// If the pointer type does not implement the proto message interface, we can't
		// handle it.
		if !innerType.Implements(protoMessageInterface) {
			return nil, fmt.Errorf(
				"inner value type '%v' is not proto message for "+
					" oneof wrapper '%v'",
				innerType,
				oneOfType,
			)
		}
		bsonTypes = []bsonTypeKey{newSimpleKey(bsontype.EmbeddedDocument)}

	default:
		return nil, fmt.Errorf(
			"could not determine bson type for inner type '%v' for oneof"+
				" wrapper '%v'",
			innerType,
			oneOfType,
		)
	}

	return bsonTypes, nil
}

// inspects the type to see if it is a known type (like decimal or UUID) that we
// can handle specially since we have foreknowledge of it's serialized type.
func (builder *CodecBuilder) deduceKnownTypeEncoding(
	oneOfType reflect.Type, innerType reflect.Type,
) (typeKeys []bsonTypeKey, known bool) {
	known = true

	switch innerType {
	case decimalType:
		typeKeys = append(typeKeys, newSimpleKey(bsontype.Decimal128))
	case timestampType:
		typeKeys = append(typeKeys, newSimpleKey(bsontype.DateTime))
	case uuidType:
		typeKeys = append(
			typeKeys, newBsonTypeKey(bsontype.Binary, bsontype.BinaryUUID),
		)
	case rawDataType:
		typeKeys = append(
			typeKeys, newBsonTypeKey(bsontype.Binary, bsontype.BinaryUserDefined),
		)
	case wrapperBoolType:
		typeKeys = append(typeKeys, newSimpleKey(bsontype.Boolean))
	case wrapperBytesType:
		typeKeys = append(
			typeKeys, newBsonTypeKey(bsontype.Binary, bsontype.BinaryGeneric),
		)
	case wrapperDoubleType:
		typeKeys = append(typeKeys, newSimpleKey(bsontype.Double))
	case wrapperFloatType:
		typeKeys = append(typeKeys, newSimpleKey(bsontype.Double))
	case wrapperInt32Type:
		typeKeys = append(typeKeys, newSimpleKey(bsontype.Int32))
		typeKeys = append(typeKeys, newSimpleKey(bsontype.Int64))
	case wrapperInt64Type:
		typeKeys = append(typeKeys, newSimpleKey(bsontype.Int32))
		typeKeys = append(typeKeys, newSimpleKey(bsontype.Int64))
	case wrapperStringType:
		typeKeys = append(typeKeys, newSimpleKey(bsontype.String))
	case wrapperUInt32Type:
		typeKeys = append(typeKeys, newSimpleKey(bsontype.Int32))
		typeKeys = append(typeKeys, newSimpleKey(bsontype.Int64))
	case wrapperUInt64Type:
		typeKeys = append(typeKeys, newSimpleKey(bsontype.Int32))
		typeKeys = append(typeKeys, newSimpleKey(bsontype.Int64))
	default:
		known = false
	}

	return typeKeys, known
}

//revive:enable:cyclomatic

// Try to automatically deduce the encoding / decoding mapping of the wrapper types.
func (builder *CodecBuilder) AutoAddConcrete(
	oneOfTypes ...reflect.Type,
) error {
	for _, thisType := range oneOfTypes {
		err := builder.deduceConcreteTypeEncoding(thisType)
		if err != nil {
			return err
		}
	}

	return nil
}

// Register a concrete one of type and the bson type it is encoded to. We don't know
// what other codecs are in the registry, so to properly round-trip a one-of field
// we need to create a 1-1 mapping of the wrapper type and it's bson encoded type,
// except for the special case of message types that encoded to embedded docs, since
// we can add a field with the proto message type.
func (builder *CodecBuilder) AddConcrete(
	oneOfType reflect.Type, bsonType bsontype.Type, binaryType byte,
) error {
	if err := builder.validateValueWrapperType(oneOfType); err != nil {
		return err
	}

	typeKey := newBsonTypeKey(bsonType, binaryType)

	existingType, ok := builder.decodeMap[typeKey]
	if bsonType == bsontype.EmbeddedDocument {
		// Get the de-referenced struct type of the wrapper
		innerType := structTypeFromOneOfWrapperType(oneOfType)
		// Make a key from the fully qualified type path and add it to our embedded
		// struct types
		docTypeKey := embeddedDocKeyFromType(innerType)
		// Store the oneOf type wrapper for this type
		builder.embeddedDocTypeMap[docTypeKey] = oneOfType.Elem()
	} else if ok {
		return fmt.Errorf(
			"bson type '%v' is already registered to oneof type '%v'",
			bsonType,
			existingType,
		)
	}

	// Store the struct type for the one-of wrapper
	builder.decodeMap[typeKey] = oneOfType.Elem()

	return nil
}

// auto-register the interface and concrete type for a messages one-of field.
func (builder *CodecBuilder) fromMessageOneOfField(
	message proto.Message, oneOfField protoreflect.OneofDescriptor,
) error {
	// Get the proto field name for our master one-of
	protoFieldName := string(oneOfField.Name())
	// We're going to store the corresponding go struct field name here.
	goFieldName := ""

	// Find the struct field that has this proto field name in it's tag and remember the
	// name so we can inspect it later..
	messageType := reflect.TypeOf(message).Elem()
	for i := 0; i < messageType.NumField(); i++ {
		thisField := messageType.Field(i)
		tagVal, ok := thisField.Tag.Lookup("protobuf_oneof")
		if ok && tagVal == protoFieldName {
			goFieldName = thisField.Name
			break
		}
	}

	if goFieldName == "" {
		return fmt.Errorf(
			"could not find go stuct field for proto field %v", protoFieldName,
		)
	}

	// Store the master interface for this one-of (it's a private interface, but that's
	// okay for reflection-based operations.)
	wrapperInterfaceField, _ := messageType.FieldByName(goFieldName)
	builder.oneOfInterface = wrapperInterfaceField.Type

	// We'll store a list of our concrete wrapper types here.
	concreteWrapperTypes := make([]reflect.Type, 0)

	// Proto message one-ofs are actually a list of sub-fields under the hood with some
	// syntactic and language-interface specific sugar. We are going to go through and
	// set each of these subfields to a new value, then inspect the go struct to get
	// the one-of wrapper type for that specific sub-field so we can register it as
	// part of the codec.
	subFields := oneOfField.Fields()
	for i := 0; i < subFields.Len(); i++ {
		// Get the possible sub-field descriptor
		thisPossible := subFields.Get(i)

		// make a new message
		newMessage := message.ProtoReflect().New()

		// Now set this field to its initial value for this one-of type
		newMessage.Set(thisPossible, newMessage.NewField(thisPossible))

		// Now we want to extract the actual go wrapper value for this one-of type
		goWrapperType := reflect.
			// Get the go reflect value of the new protoreflect message
			ValueOf(newMessage.Interface()).
			// Dereference the pointer
			Elem().
			// Get the one-of struct field
			FieldByName(goFieldName).
			// Push through the interface to the concrete value
			Elem().
			// Get the type (this will be a pointer to it)
			Type()

		concreteWrapperTypes = append(concreteWrapperTypes, goWrapperType)
	}

	err := builder.AutoAddConcrete(concreteWrapperTypes...)
	if err != nil {
		return err
	}

	return nil
}

func (builder *CodecBuilder) Register(registryBuilder *bsoncodec.RegistryBuilder) {
	codec := &oneOfCodec{
		oneOfInterface:     builder.oneOfInterface,
		decodeMap:          builder.decodeMap,
		embeddedDocTypeMap: builder.embeddedDocTypeMap,
	}

	registryBuilder.RegisterCodec(codec.oneOfInterface, codec)
}

func NewCodecBuilder() *CodecBuilder {
	return &CodecBuilder{
		decodeMap:          make(map[bsonTypeKey]reflect.Type),
		embeddedDocTypeMap: make(map[string]reflect.Type),
		innerTypeMap:       make(map[reflect.Type][]bsonTypeKey),
	}
}

func CodecBuildersForMessage(
	message proto.Message, elementInfo []*ElementInfo,
) ([]*CodecBuilder, error) {
	innerTypeMap := make(map[reflect.Type][]bsonTypeKey)

	// Create the inner type map for manually added types
	for _, elementInfo := range elementInfo {
		// Some go types can be encoded to multiple bson types depending on the
		// situation. For instance, int64s can be encoded as int32s when small enough,
		// so we need to support registering multiple bson types. We need to fetch
		// the current list if any and add the new bson type to it.
		bsonKey := newBsonTypeKey(elementInfo.BsonType, elementInfo.BinarySubType)
		bsonTypes, _ := innerTypeMap[elementInfo.InnerGoType]
		bsonTypes = append(bsonTypes, bsonKey)
		innerTypeMap[elementInfo.InnerGoType] = bsonTypes
	}

	// Register the one-of fields for a message
	oneOfs := message.ProtoReflect().Descriptor().Oneofs()
	builders := make([]*CodecBuilder, 0, oneOfs.Len())

	for i := 0; i < oneOfs.Len(); i++ {
		oneOfField := oneOfs.Get(i)

		builder := NewCodecBuilder()
		builder.innerTypeMap = innerTypeMap

		err := builder.fromMessageOneOfField(message, oneOfField)
		if err != nil {
			return nil, fmt.Errorf(
				"error registering one-of field '%v' of type '%v': %w",
				oneOfField.Name(),
				reflect.TypeOf(message),
				err,
			)
		}

		builders = append(builders, builder)
	}

	return builders, nil
}
