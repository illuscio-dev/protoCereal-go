package oneof

import "go.mongodb.org/mongo-driver/bson/bsontype"

// We are going to store bson type information in a 2-byte array for mapping bson
// types to decode types. We need to bytes for:
//
//	BYTE 1: bsontype.Type
//  BYTE 2: binary type
//
// This allows us to register a codec for each binary type rather than only having one
// binary type available. This allows us to handle, for instance, a one of UUID /
// RawData field.
type bsonTypeKey [2]byte

// The bson type for this key
func (key bsonTypeKey) BsonType() bsontype.Type {
	return bsontype.Type(key[0])
}

// The binary type for this key, only relevant if BsonType is bsontype.Binary
func (key bsonTypeKey) BinaryType() byte {
	return key[1]
}

// Create a new bson type key
func newBsonTypeKey(bsonType bsontype.Type, binaryType byte) bsonTypeKey {
	return bsonTypeKey{byte(bsonType), binaryType}
}

// Create a new bson-type key with a Binary type of "Generic" useful for creating keys
// where the binary type is not relevant (the bson type is not Binary).
func newSimpleKey(bsonType bsontype.Type) bsonTypeKey {
	return newBsonTypeKey(bsonType, 0)
}
