package protobson_test

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/illuscio-dev/protoCereal-go/cereal"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

// Checks that timestamp messages are encoded as a single field witha bson datetime
// value
func TestCodec_Timestamp_BsonVal(t *testing.T) {
	assert := assert.New(t)

	type hasUUID struct {
		Value *timestamp.Timestamp
	}

	original := &hasUUID{
		Value: timestamppb.New(time.Now().UTC()),
	}

	encoded, err := bson.MarshalWithRegistry(testRegistry, original)
	assert.NoError(err, "encoding err")

	decodedMap := make(map[string]interface{})
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, decodedMap)

	for key, value := range decodedMap {
		fmt.Printf("%v: %v\n", key, value)
	}
	assert.IsType(primitive.DateTime(0), decodedMap["value"])
}

// Tests that we can round-trip a struct with a timestamp message
func TestCodec_Timestamp_RoundTrip(t *testing.T) {
	assert := assert.New(t)

	type hasTimestamp struct {
		Value *timestamp.Timestamp
	}

	original := &hasTimestamp{
		Value: timestamppb.New(time.Now().UTC()),
	}
	cereal.ClipTimestamp(original.Value)

	encoded, err := bson.MarshalWithRegistry(testRegistry, original)
	assert.NoError(err, "encoding err")

	decoded := new(hasTimestamp)
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, decoded)
	assert.NoError(err, "decoding error")

	assert.True(
		original.Value.AsTime().Equal(decoded.Value.AsTime()),
		"val equality",
	)

}

// Tests that we can round-trip a struct with a timestamp message
func TestCodec_Timestamp_RoundTrip_Null(t *testing.T) {
	assert := assert.New(t)

	type hasUUID struct {
		Value *timestamp.Timestamp
	}

	original := &hasUUID{
		Value: nil,
	}

	encoded, err := bson.MarshalWithRegistry(testRegistry, original)
	assert.NoError(err, "encoding err")

	decoded := new(hasUUID)
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, decoded)
	assert.NoError(err, "decoding error")

	assert.Nil(
		original.Value, "nil value",
	)
}
