package protoBson_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/illuscio-dev/protoCereal-go/messagesCereal"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestCodec_UUID_BsonVal(t *testing.T) {
	assert := assert.New(t)

	type hasUUID struct {
		Value *messagesCereal.UUID
	}

	original := &hasUUID{
		Value: messagesCereal.UUIDFromGoogle(uuid.Must(uuid.NewRandom())),
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
		Value *messagesCereal.UUID
	}

	original := &hasUUID{
		Value: messagesCereal.UUIDFromGoogle(uuid.Must(uuid.NewRandom())),
	}

	encoded, err := bson.MarshalWithRegistry(testRegistry, original)
	assert.NoError(err, "encoding err")

	decoded := new(hasUUID)
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, decoded)
	assert.NoError(err, "decoding error")

	assert.Equal(original.Value.Value, decoded.Value.Value, "uuid val")
}

func TestCodec_UUID_RoundTrip_Null(t *testing.T) {
	assert := assert.New(t)

	type hasUUID struct {
		Value *messagesCereal.UUID
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
