package protobson_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/illuscio-dev/protoCereal-go/cereal"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoUUID "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"testing"
)

func TestCodec_UUID_BsonVal(t *testing.T) {
	assert := assert.New(t)

	type hasUUID struct {
		Value *cereal.UUID
	}

	original := &hasUUID{
		Value: cereal.UUIDFromGoogle(uuid.Must(uuid.NewRandom())),
	}

	encoded, err := bson.MarshalWithRegistry(testRegistry, original)
	assert.NoError(err, "encoding err")

	decoded := make(map[string]interface{})
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, decoded)
	assert.NoError(err, "decoding error")

	for key, value := range decoded {
		fmt.Printf("%v: %v\n", key, value)
	}

	assert.IsType(primitive.Binary{}, decoded["value"], "uuid val")
	binValue, ok := decoded["value"].(primitive.Binary)
	if !ok {
		t.FailNow()
	}

	assert.Equal(
		bsontype.BinaryUUID,
		binValue.Subtype,
		"value is correct binary subtype",
	)
}

func TestCodec_UUID_RoundTrip(t *testing.T) {
	assert := assert.New(t)

	type hasUUID struct {
		Value *cereal.UUID
	}

	original := &hasUUID{
		Value: cereal.UUIDFromGoogle(uuid.Must(uuid.NewRandom())),
	}

	encoded, err := bson.MarshalWithRegistry(testRegistry, original)
	assert.NoError(err, "encoding err")

	decoded := new(hasUUID)
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, decoded)
	assert.NoError(err, "decoding error")

	assert.Equal(original.Value.Bin, decoded.Value.Bin, "uuid val")
}

func TestCodec_UUID_RoundTrip_Null(t *testing.T) {
	assert := assert.New(t)

	type hasUUID struct {
		Value *cereal.UUID
	}

	original := &hasUUID{
		Value: nil,
	}

	encoded := make([]byte, 0)
	encoded, err := bson.MarshalAppendWithRegistry(
		testRegistry, encoded, original,
	)
	assert.NoError(err, "encoding err")

	decoded := new(hasUUID)
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, decoded)
	assert.NoError(err, "decoding error")

	assert.Nil(original.Value, "uuid nil")
}

func TestUUID_Google_RoundTrip(t *testing.T) {
	assert := assert.New(t)

	uuidVal := cereal.MustUUIDRandom()
	assert.NotZero(uuidVal, "UUID is not zero value")

	googleVal := uuidVal.MustGoogle()
	if !assert.IsType(uuid.UUID{}, googleVal, "value is google") {
		t.FailNow()
	}

	messageLoaded := cereal.UUIDFromGoogle(googleVal)
	assert.Equal(uuidVal.Bin, messageLoaded.Bin, "loaded and original")
}

func TestUUID_Google_Zero(t *testing.T) {
	zeroGoogle := uuid.UUID{}

	nilMessage := (*cereal.UUID)(nil)
	googleVal, err := nilMessage.ToGoogle()
	assert.NoError(t, err, "error converting to google value")
	assert.Equal(t, zeroGoogle, googleVal)
}

func TestUUID_Google_BadLength(t *testing.T) {
	assert := assert.New(t)

	badMessage := &cereal.UUID{
		Bin: []byte{0x0, 0x1},
	}

	zeroVal, err := badMessage.ToGoogle()
	assert.EqualError(
		err, "proto uuid message must be 16 bytes: 2 bytes found",
	)
	assert.Zero(zeroVal)
}

func TestUUID_Google_BadLengthMust(t *testing.T) {
	assert := assert.New(t)

	testFunc := func() {
		badMessage := &cereal.UUID{
			Bin: []byte{0x0, 0x1},
		}

		_ = badMessage.MustGoogle()
	}

	assert.Panics(testFunc, "panic on MustGoogle wrong length")

}

func TestUUID_Mongo_RoundTrip(t *testing.T) {
	assert := assert.New(t)

	uuidVal := cereal.MustUUIDRandom()
	assert.NotZero(uuidVal, "UUID is not zero value")

	mongoVal := uuidVal.MustMongo()
	if !assert.IsType(mongoUUID.UUID{}, mongoVal, "value is google") {
		t.FailNow()
	}

	messageLoaded := cereal.UUIDFromMongo(mongoVal)
	assert.Equal(uuidVal.Bin, messageLoaded.Bin, "loaded and original")
}

func TestUUID_Mongo_Zero(t *testing.T) {
	zeroMongo := mongoUUID.UUID{}

	nilMessage := (*cereal.UUID)(nil)
	googleVal, err := nilMessage.ToMongo()
	assert.NoError(t, err, "error converting to google value")
	assert.Equal(t, zeroMongo, googleVal)
}

func TestUUID_Mongo_BadLength(t *testing.T) {
	assert := assert.New(t)

	badMessage := &cereal.UUID{
		Bin: []byte{0x0, 0x1},
	}

	zeroVal, err := badMessage.ToMongo()
	assert.EqualError(
		err, "proto uuid message must be 16 bytes: 2 bytes found",
	)
	assert.Zero(zeroVal)
}

func TestUUID_Mongo_BadLengthMust(t *testing.T) {
	assert := assert.New(t)

	testFunc := func() {
		badMessage := &cereal.UUID{
			Bin: []byte{0x0, 0x1},
		}

		_ = badMessage.MustMongo()
	}

	assert.Panics(testFunc, "panic on MustGoogle wrong length")
}
