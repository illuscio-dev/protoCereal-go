package protoBson_test

import (
	"fmt"
	"github.com/illuscio-dev/protoCereal-go/messagesCereal"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestCodec_RawData_BsonVal(t *testing.T) {
	assert := assert.New(t)

	type hasVal struct {
		Value *messagesCereal.RawData
	}

	original := &hasVal{
		Value: &messagesCereal.RawData{
			Data: []byte("some data"),
		},
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
		bsontype.BinaryUserDefined,
		binValue.Subtype,
		"value is correct binary subtype",
	)
}

func TestCodec_RawData_RoundTrip(t *testing.T) {
	assert := assert.New(t)

	type hasVal struct {
		Value *messagesCereal.RawData
	}

	original := &hasVal{
		Value: &messagesCereal.RawData{
			Data: []byte("some data"),
		},
	}

	encoded, err := bson.MarshalWithRegistry(testRegistry, original)
	assert.NoError(err, "encoding err")

	decoded := new(hasVal)
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, decoded)
	assert.NoError(err, "decoding error")

	assert.Equal(original.Value.Data, decoded.Value.Data, "uuid val")
}

func TestCodec_RawData_RoundTrip_Null(t *testing.T) {
	assert := assert.New(t)

	type hasVal struct {
		Value *messagesCereal.RawData
	}

	original := &hasVal{
		Value: nil,
	}

	encoded, err := bson.MarshalWithRegistry(testRegistry, original)
	assert.NoError(err, "encoding err")

	decoded := new(hasVal)
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, decoded)
	assert.NoError(err, "decoding error")

	assert.Nil(decoded.Value, "nil val")
}
