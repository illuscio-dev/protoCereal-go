package oneof

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"reflect"
)

type CodecBuilder struct {
	oneOfInterface reflect.Type
	// Map of the encoded bson type to the oneof field type it decodes to
	decodeMap map[bsonTypeKey]reflect.Type
	// For message types that can be encoded into an embedded document, we actually CAN
	// disambiguate between the target decode types, since we can embed the type name
	// in the written document.
	embeddedDocTypeMap map[string]reflect.Type
}

func (builder *CodecBuilder) ValidateType(oneOfType reflect.Type) error {
	// Our concrete type must implement our oneof type
	if !oneOfType.Implements(builder.oneOfInterface) {
		return fmt.Errorf(
			"'%v' does not implement oneof interface '%v'",
			oneOfType,
			builder.oneOfInterface,
		)
		// Our concrete type must be a pointer
	} else if oneOfType.Kind() != reflect.Ptr {
		return fmt.Errorf(
			"'%v' is not a pointer type",
			oneOfType,
		)
		// The number of fields must be 1
	} else if oneOfType.Elem().NumField() != 1 {
		return fmt.Errorf(
			"'%v' contains %v fields, expected 1",
			oneOfType,
			oneOfType.NumField(),
		)
	}
	return nil
}

func (builder *CodecBuilder) autoRegisterOneOfType(
	oneOfType reflect.Type,
) error {
	var err error
	if err = builder.ValidateType(oneOfType) ; err != nil {
		return err
	}

	innerType := oneOfType.Elem().Field(0).Type
	// First try to register it if it's a known type that serialized to something
	// other than an embedded doc
	bsonTypes, known := builder.autoRegisterKnownTypes(oneOfType, innerType)

	// If it's not known, we're next try to register it as an unknown type.
	if !known {
		bsonTypes, err = builder.autoRegisterUnknownTypes(oneOfType, innerType)
		if err != nil {
			return err
		}
	}

	for _, thisBsonKey := range bsonTypes {
		err := builder.RegisterOneOfType(
			oneOfType, thisBsonKey.BsonType(), thisBsonKey.BinaryType(),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (builder *CodecBuilder) autoRegisterUnknownTypes(
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
	case reflect.Ptr:
		// If this inner type is not a struct, it's out of spec for protobuf structures
		// and we cannot handle it.
		if innerType.Elem().Kind() != reflect.Struct {
			return nil, fmt.Errorf(
				"inner pointer type '%v' does not point to struct in" +
					" oneof wrapper '%v'",
				innerType,
				oneOfType,
			)
		}
		// If the pointer type does not implement the proto message interface, we can't
		// handle it.
		if !innerType.Implements(protoMessageInterface ) {
			return nil, fmt.Errorf(
				"inner value type '%v' is not proto message for " +
					" oneof wrapper '%v'",
				innerType,
				oneOfType,
			)
		}

		bsonTypes = []bsonTypeKey{newSimpleKey(bsontype.EmbeddedDocument)}
	default:
		return nil, fmt.Errorf(
			"could not determine bson type for inner type '%v' for oneof" +
				" wrapper '%v'",
			innerType,
			oneOfType,
		)
	}

	return bsonTypes, nil
}

// inspects the type to see if it is a known type (like decimal or UUID) that we
// can handle specially since we have for-knowledge of it's serialized type.
func (builder *CodecBuilder) autoRegisterKnownTypes(
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
			typeKeys, newBsonTypeKey(bsontype.Binary, bsontype.BinaryGeneric),
		)
	case wrapperBoolType:
		typeKeys =  append(typeKeys, newSimpleKey(bsontype.Boolean))
	case wrapperBytesType:
		typeKeys = append(
			typeKeys, newBsonTypeKey(bsontype.Binary, bsontype.BinaryGeneric),
		)
	case wrapperDoubleType:
		typeKeys =  append(typeKeys, newSimpleKey(bsontype.Double))
	case wrapperFloatType:
		typeKeys =  append(typeKeys, newSimpleKey(bsontype.Double))
	case wrapperInt32Type:
		typeKeys =  append(typeKeys, newSimpleKey(bsontype.Int32))
		typeKeys =  append(typeKeys, newSimpleKey(bsontype.Int64))
	case wrapperInt64Type:
		typeKeys =  append(typeKeys, newSimpleKey(bsontype.Int32))
		typeKeys =  append(typeKeys, newSimpleKey(bsontype.Int64))
	case wrapperStringType:
		typeKeys =  append(typeKeys, newSimpleKey(bsontype.String))
	case wrapperUInt32Type:
		typeKeys =  append(typeKeys, newSimpleKey(bsontype.Int32))
		typeKeys =  append(typeKeys, newSimpleKey(bsontype.Int64))
	case wrapperUInt64Type:
		typeKeys =  append(typeKeys, newSimpleKey(bsontype.Int32))
		typeKeys =  append(typeKeys, newSimpleKey(bsontype.Int64))
	default:
		known = false
	}

	return typeKeys, known
}

func (builder *CodecBuilder) AutoRegisterOneOfTypes(
	oneOfTypes ...reflect.Type,
) error {
	for _, thisType := range oneOfTypes {
		err := builder.autoRegisterOneOfType(thisType)
		if err != nil {
			return err
		}
	}

	return nil
}

func (builder *CodecBuilder) RegisterOneOfType(
	oneOfType reflect.Type, bsonType bsontype.Type, binaryType byte,
) error  {
	if err := builder.ValidateType(oneOfType) ; err != nil {
		return err
	}

	typeKey := newBsonTypeKey(bsonType, binaryType)

	if existingType, ok := builder.decodeMap[typeKey] ; ok {
		return fmt.Errorf(
			"bson type '%v' is already registered to oneof type '%v'",
			bsonType,
			existingType,
		)
	}

	builder.decodeMap[typeKey] = oneOfType.Elem()

	return nil
}

func (builder *CodecBuilder) Register(registryBuilder *bsoncodec.RegistryBuilder) {
	codec := &oneOfCodec{
		oneOfInterface: builder.oneOfInterface,
		decodeMap:      builder.decodeMap,
	}

	registryBuilder.RegisterCodec(codec.oneOfInterface, codec)
}

func NewCodecBuilder(oneOfInterface reflect.Type) (*CodecBuilder, error) {
	if oneOfInterface.Kind() != reflect.Interface {
		return nil, fmt.Errorf("'%v' is not an interface", oneOfInterface)
	}

	return &CodecBuilder{
		oneOfInterface: oneOfInterface,
		decodeMap:      make(map[bsonTypeKey]reflect.Type),
	}, nil
}
