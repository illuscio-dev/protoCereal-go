package protoBson_test

import (
	"fmt"
	"github.com/illuscio-dev/protoCereal-go/messagesCereal_test"
	"github.com/illuscio-dev/protoCereal-go/protoBSON"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/anypb"
	"testing"
)

func TestCodec_Any_MapValues(t *testing.T) {
	assert := assert.New(t)

	type hasVal struct {
		Value *anypb.Any
	}

	anyVal, err := anypb.New(
		&messagesCereal_test.TestProto{
			FieldString: "some string",
			FieldInt32:  42,
		},
	)

	original := &hasVal{
		Value: anyVal,
	}

	encoded, err := bson.MarshalWithRegistry(testRegistry, original)
	assert.NoError(err, "encoding err")
	if err != nil {
		t.FailNow()
	}

	fmt.Println("encoded!")

	testMap := make(map[string]interface{})
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, testMap)
	assert.NoError(err, "decoding to map")
	fmt.Println("test map of data:", testMap)

	testMap, ok := testMap["value"].(map[string]interface{})
	if !ok {
		t.Error("could not get nested doc")
	}

	stringValueInterface, ok := testMap["field_string"]
	if !ok {
		t.Error("field_string not in encoded doc")
	}

	int32Value, ok := stringValueInterface.(string)
	if !ok {
		t.Error("string_value is not string")
	}

	assert.Equal("some string", int32Value)

	typeUrlVal, ok := testMap[protoBson.AnyTypeUrlField]
	assert.True(ok)
	assert.NotZero(typeUrlVal)
}

func TestCodec_Any_RoundTrip(t *testing.T) {
	assert := assert.New(t)

	type hasVal struct {
		Value *anypb.Any
	}

	anyVal, err := anypb.New(
		&messagesCereal_test.TestProto{
			FieldString: "some string",
			FieldInt32:  42,
		},
	)

	original := &hasVal{
		Value: anyVal,
	}

	encoded, err := bson.MarshalWithRegistry(
		testRegistry, original,
	)
	assert.NoError(err, "encoding err")
	if err != nil {
		t.FailNow()
	}

	testMap := make(map[string]interface{})
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, testMap)
	assert.NoError(err, "decoding to map")
	fmt.Println("test map of data:", testMap)

	decoded := new(hasVal)
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, decoded)
	assert.NoError(err, "decoding error")
	if err != nil {
		t.FailNow()
	}

	fmt.Println(decoded)
	messageInterface, err := decoded.Value.UnmarshalNew()
	assert.NoError(err, "unmarshal any proto")
	if err != nil {
		t.FailNow()
	}

	message, ok := messageInterface.(*messagesCereal_test.TestProto)
	if !ok {
		t.Errorf("could not convert any payload to dicom header")
	}
	assert.Equal("some string", message.FieldString)
	assert.Equal(int32(42), message.FieldInt32)
}

func TestAnyCodec_Nil(t *testing.T) {
	assert := assert.New(t)

	type hasVal struct {
		Value *anypb.Any
	}

	original := &hasVal{
		Value: nil,
	}

	encoded, err := bson.MarshalWithRegistry(
		testRegistry, original,
	)
	assert.NoError(err, "encoding err")
	if err != nil {
		t.FailNow()
	}

	decoded := new(hasVal)
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, decoded)
	assert.NoError(err, "decoding error")
	if err != nil {
		t.FailNow()
	}

	assert.Nil(decoded.Value)
}
