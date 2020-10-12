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
	"log"
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
				Subtype: bsontype.BinaryUserDefined,
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

func TestOneOf_CustomMapping(t *testing.T) {
	cerealOpts := protoBson.
		NewMongoOpts().
		WithCustomWrappers(
			new(messagesCereal_test.DecimalList),
		).
		WithOneOfElementMapping(
			new(messagesCereal_test.DecimalList),
			bsontype.Array,
			0x0,
		).
		WithOneOfFields(new(messagesCereal_test.HasCustomOneOf))

	builder, err := protoBson.NewCerealRegistryBuilder(cerealOpts)
	if !assert.NoError(t, err, "create registry") {
		t.FailNow()
	}

	registry := builder.Build()

	type testCase struct {
		Name            string
		ElementValue    messagesCereal_test.IsHasCustomOneofList
		SerializedValue interface{}
	}

	var thisCase *testCase

	testCases := []*testCase{
		{
			Name: "String",
			ElementValue: &messagesCereal_test.HasCustomOneOf_StringValue{
				StringValue: "some value",
			},
			SerializedValue: "some value",
		},
		{
			Name: "Decimal",
			ElementValue: &messagesCereal_test.HasCustomOneOf_DecimalValue{
				DecimalValue: &messagesCereal.Decimal{
					High: 47,
					Low:  101,
				},
			},
			SerializedValue: primitive.NewDecimal128(47, 101),
		},
		{
			Name: "DecimalList",
			ElementValue: &messagesCereal_test.HasCustomOneOf_DecimalList{
				DecimalList: &messagesCereal_test.DecimalList{
					Value: []*messagesCereal.Decimal{
						{
							High: 100,
							Low:  101,
						},
						{
							High: 102,
							Low:  103,
						},
					},
				},
			},
			SerializedValue: bson.A{
				primitive.NewDecimal128(100, 101),
				primitive.NewDecimal128(102, 103),
			},
		},
	}

	testFunc := func(t *testing.T) {
		assert := assert.New(t)

		message := &messagesCereal_test.HasCustomOneOf{
			Many: thisCase.ElementValue,
		}

		serialized, err := bson.MarshalWithRegistry(registry, message)
		if !assert.NoError(err, "marshal to bson") {
			t.FailNow()
		}

		document := bson.M{}
		err = bson.UnmarshalWithRegistry(registry, serialized, document)

		if !assert.Contains(document, "many", "has key") {
			t.FailNow()
		}

		t.Log("DOCUMENT:", document)

		assert.Equal(
			thisCase.SerializedValue,
			document["many"],
			"correct serialized value",
		)

		deserialized := new(messagesCereal_test.HasCustomOneOf)
		err = bson.UnmarshalWithRegistry(registry, serialized, deserialized)
		if !assert.NoError(err, "unmarshall to proto") {
			t.FailNow()
		}

		assert.Equal(
			message, deserialized, "deserialized equals original",
		)
	}

	for _, thisCase = range testCases {
		t.Run(thisCase.Name, testFunc)
	}
}

func TestOneOf_BytesValue(t *testing.T) {
	assert := assert.New(t)

	cerealOpts := protoBson.NewMongoOpts().
		WithOneOfFields(new(messagesCereal_test.HasOneOfBytes))

	registryBuilder, err := protoBson.NewCerealRegistryBuilder(cerealOpts)
	if !assert.NoError(err, "create registry builder") {
		t.FailNow()
	}

	registry := registryBuilder.Build()

	message := &messagesCereal_test.HasOneOfBytes{
		Value: &messagesCereal_test.HasOneOfBytes_BytesValue{
			BytesValue: []byte("some bin data"),
		},
	}

	serialized, err := bson.MarshalWithRegistry(registry, message)
	if !assert.NoError(err, "serialize message") {
		t.FailNow()
	}

	document := bson.M{}
	err = bson.UnmarshalWithRegistry(registry, serialized, document)
	if !assert.NoError(err, "unmarshall to document") {
		t.FailNow()
	}

	log.Println("DOCUMENT:", document)
	if !assert.Contains(document, "value") {
		t.FailNow()
	}

	expectedValue := primitive.Binary{
		Subtype: 0x0,
		Data:    []byte("some bin data"),
	}

	if !assert.Equal(expectedValue, document["value"]) {
		t.FailNow()
	}

	deserialized := new(messagesCereal_test.HasOneOfBytes)
	err = bson.UnmarshalWithRegistry(registry, serialized, deserialized)
	if !assert.NoError(err, "unmarshall to protobuf") {
		t.FailNow()
	}

	assert.Equal(message, deserialized, "unmarshalled equals original")
}
