package protoBson_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/illuscio-dev/protoCereal-go/messagesCereal"
	"github.com/illuscio-dev/protoCereal-go/messagesCereal_test"
	protoBson "github.com/illuscio-dev/protoCereal-go/protoBSON"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestOneOfFirst(t *testing.T) {
	cerealOpts := protoBson.NewMongoOpts()

	cerealOpts.WithOneOfFields(new(messagesCereal_test.TestOneOfFirst))

	uuidVal, err := uuid.NewRandom()
	if !assert.NoError(t, err, "uuid generation") {
		t.FailNow()
	}

	builder := bsoncodec.NewRegistryBuilder()
	err = protoBson.RegisterCerealCodecs(builder, cerealOpts)
	if !assert.NoError(t, err, "error registering codecs") {
		t.FailNow()
	}

	registry := builder.Build()

	type testCase struct {
		name            string
		value           messagesCereal_test.IsTestOneOfFirstSomeValue
		serializedValue interface{}
	}

	testCases := []testCase{
		{
			name: "int32",
			value: &messagesCereal_test.TestOneOfFirst_FieldInt32{
				FieldInt32: 42,
			},
			serializedValue: int32(42),
		},
		{
			name: "string",
			value: &messagesCereal_test.TestOneOfFirst_FieldString{
				FieldString: "Some Data",
			},
			serializedValue: "Some Data",
		},
		{
			name: "decimal",
			value: &messagesCereal_test.TestOneOfFirst_FieldDecimal{
				FieldDecimal: &messagesCereal.Decimal{
					High: 42,
					Low:  77,
				},
			},
			serializedValue: primitive.NewDecimal128(42, 77),
		},
		{
			name: "uuid",
			value: &messagesCereal_test.TestOneOfFirst_FieldUuid{
				FieldUuid: messagesCereal.UUIDFromGoogle(uuidVal),
			},
			serializedValue: primitive.Binary{
				Subtype: bsontype.BinaryUUID,
				Data:    uuidVal[:],
			},
		},
		{
			name: "raw",
			value: &messagesCereal_test.TestOneOfFirst_FieldRaw{
				FieldRaw: &messagesCereal.RawData{Data: []byte("some bin data")},
			},
			serializedValue: primitive.Binary{
				Subtype: bsontype.BinaryGeneric,
				Data:    []byte("some bin data"),
			},
		},
		{
			name: "wizard",
			value: &messagesCereal_test.TestOneOfFirst_FieldWizard{
				FieldWizard: &messagesCereal_test.Wizard{
					Name: "Harry Potter",
				},
			},
			serializedValue: primitive.M{
				"name": "Harry Potter",
			},
		},
		{
			name: "interior nil",
			value: &messagesCereal_test.TestOneOfFirst_FieldDecimal{
				FieldDecimal: nil,
			},
			serializedValue: nil,
		},
		{
			name:            "field nil",
			value:           nil,
			serializedValue: nil,
		},
	}

	var thisCase testCase

	runTest := func(t *testing.T) {
		assert := assert.New(t)

		original := &messagesCereal_test.TestOneOfFirst{
			SomeValue: thisCase.value,
		}

		encoded, err := bson.MarshalWithRegistry(registry, original)
		assert.NoError(err, "error marshalling data")
		if err != nil {
			t.FailNow()
		}

		mapData := bson.M{}
		err = bson.UnmarshalWithRegistry(registry, encoded, mapData)
		assert.NoError(err, "error unmarshalling to map")
		if err != nil {
			t.FailNow()
		}

		fmt.Println("de-serialized map:", mapData)

		mapValue, ok := mapData["some_value"]
		assert.True(ok, "map contains root field")
		if !ok {
			t.FailNow()
		}
		assert.Equal(thisCase.serializedValue, mapValue)

		unmarshalled := new(messagesCereal_test.TestOneOfFirst)
		err = bson.UnmarshalWithRegistry(registry, encoded, unmarshalled)
		assert.NoError(err, "error unmarshalling proto")
		if err != nil {
			t.FailNow()
		}

		// If the serialized value is expected to be nil, nil the whole wrapper.
		if thisCase.serializedValue == nil {
			original.SomeValue = nil
		}

		assert.Equal(original, unmarshalled, "unmarshalled equals original")
	}

	for _, thisCase = range testCases {
		t.Run(thisCase.name, runTest)
	}
}

func TestOneOfMultiMessageTargets(t *testing.T) {
	builder := bson.NewRegistryBuilder()
	cerealOpts := protoBson.NewMongoOpts().WithOneOfFields(
		new(messagesCereal_test.TestOneOfMultiMessage),
	)

	err := protoBson.RegisterCerealCodecs(builder, cerealOpts)
	if !assert.NoError(t, err, "cereal registration") {
		t.FailNow()
	}

	registry := builder.Build()

	type hasMessage struct {
		Message *messagesCereal_test.TestOneOfMultiMessage
	}

	assert := assert.New(t)

	original := &hasMessage{
		Message: &messagesCereal_test.TestOneOfMultiMessage{
			Mage: &messagesCereal_test.TestOneOfMultiMessage_Wizard{
				Wizard: &messagesCereal_test.Wizard{Name: "Harry Potter"},
			},
		},
	}

	encoded, err := bson.MarshalWithRegistry(registry, original)
	if !assert.NoError(err, "marshal message") {
		t.FailNow()
	}

	mapDecoded := bson.M{}
	err = bson.UnmarshalWithRegistry(registry, encoded, mapDecoded)
	if !assert.NoError(err, "unmarshal message.") {
		t.FailNow()
	}

	t.Log("MAP:", mapDecoded)
	if !assert.Contains(mapDecoded, "message") {
		t.FailNow()
	}

	messageMap := mapDecoded["message"].(bson.M)
	if !assert.Contains(messageMap, "mage") {
		t.FailNow()
	}

	mageMap := messageMap["mage"].(bson.M)
	if !assert.Contains(mageMap, "name") {
		t.FailNow()
	}

	if !assert.Equal("Harry Potter", mageMap["name"]) {
		t.FailNow()
	}

	decoded := new(hasMessage)
	err = bson.UnmarshalWithRegistry(registry, encoded, decoded)
	if !assert.NoError(err, "unmarshall to struct") {
		t.FailNow()
	}

	fmt.Println("DECODED:", decoded)
}

func TestAutoRegisterOneOfs(t *testing.T) {
	builder := bson.NewRegistryBuilder()
	cerealOpts := protoBson.
		NewMongoOpts().
		WithOneOfFields(new(messagesCereal_test.TestOneOfFirst))

	_ = protoBson.RegisterCerealCodecs(builder, cerealOpts)
}
