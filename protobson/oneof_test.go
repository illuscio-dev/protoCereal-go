package protobson_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/illuscio-dev/protoCereal-go/cereal"
	"github.com/illuscio-dev/protoCereal-go/cereal_test"
	protoBson "github.com/illuscio-dev/protoCereal-go/protobson"
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

	cerealOpts.WithOneOfFields(new(cereal_test.TestOneOfFirst))

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
		value           cereal_test.IsTestOneOfFirstSomeValue
		serializedValue interface{}
	}

	testCases := []testCase{
		{
			name: "int32",
			value: &cereal_test.TestOneOfFirst_FieldInt32{
				FieldInt32: 42,
			},
			serializedValue: int32(42),
		},
		{
			name: "string",
			value: &cereal_test.TestOneOfFirst_FieldString{
				FieldString: "Some Data",
			},
			serializedValue: "Some Data",
		},
		{
			name: "decimal",
			value: &cereal_test.TestOneOfFirst_FieldDecimal{
				FieldDecimal: &cereal.Decimal{
					High: 42,
					Low:  77,
				},
			},
			serializedValue: primitive.NewDecimal128(42, 77),
		},
		{
			name: "uuid",
			value: &cereal_test.TestOneOfFirst_FieldUuid{
				FieldUuid: cereal.UUIDFromGoogle(uuidVal),
			},
			serializedValue: primitive.Binary{
				Subtype: bsontype.BinaryUUID,
				Data:    uuidVal[:],
			},
		},
		{
			name: "raw",
			value: &cereal_test.TestOneOfFirst_FieldRaw{
				FieldRaw: &cereal.RawData{Data: []byte("some bin data")},
			},
			serializedValue: primitive.Binary{
				Subtype: bsontype.BinaryUserDefined,
				Data:    []byte("some bin data"),
			},
		},
		{
			name: "wizard",
			value: &cereal_test.TestOneOfFirst_FieldWizard{
				FieldWizard: &cereal_test.Wizard{
					Name: "Harry Potter",
				},
			},
			serializedValue: primitive.M{
				"name": "Harry Potter",
			},
		},
		{
			name: "interior nil",
			value: &cereal_test.TestOneOfFirst_FieldDecimal{
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

		original := &cereal_test.TestOneOfFirst{
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

		unmarshalled := new(cereal_test.TestOneOfFirst)
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
		new(cereal_test.TestOneOfMultiMessage),
	)

	err := protoBson.RegisterCerealCodecs(builder, cerealOpts)
	if !assert.NoError(t, err, "cereal registration") {
		t.FailNow()
	}

	registry := builder.Build()

	type hasMessage struct {
		Message *cereal_test.TestOneOfMultiMessage
	}

	assert := assert.New(t)

	original := &hasMessage{
		Message: &cereal_test.TestOneOfMultiMessage{
			Mage: &cereal_test.TestOneOfMultiMessage_Wizard{
				Wizard: &cereal_test.Wizard{Name: "Harry Potter"},
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
		WithOneOfFields(new(cereal_test.TestOneOfFirst))

	_ = protoBson.RegisterCerealCodecs(builder, cerealOpts)
}

func TestOneOf_CustomMapping(t *testing.T) {
	cerealOpts := protoBson.
		NewMongoOpts().
		WithCustomWrappers(
			new(cereal_test.DecimalList),
		).
		WithOneOfElementMapping(
			new(cereal_test.DecimalList),
			bsontype.Array,
			0x0,
		).
		WithOneOfFields(new(cereal_test.HasCustomOneOf))

	builder, err := protoBson.NewCerealRegistryBuilder(cerealOpts)
	if !assert.NoError(t, err, "create registry") {
		t.FailNow()
	}

	registry := builder.Build()

	type testCase struct {
		Name            string
		ElementValue    cereal_test.IsHasCustomOneofList
		SerializedValue interface{}
	}

	var thisCase *testCase

	testCases := []*testCase{
		{
			Name: "String",
			ElementValue: &cereal_test.HasCustomOneOf_StringValue{
				StringValue: "some value",
			},
			SerializedValue: "some value",
		},
		{
			Name: "Decimal",
			ElementValue: &cereal_test.HasCustomOneOf_DecimalValue{
				DecimalValue: &cereal.Decimal{
					High: 47,
					Low:  101,
				},
			},
			SerializedValue: primitive.NewDecimal128(47, 101),
		},
		{
			Name: "DecimalList",
			ElementValue: &cereal_test.HasCustomOneOf_DecimalList{
				DecimalList: &cereal_test.DecimalList{
					Value: []*cereal.Decimal{
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

		message := &cereal_test.HasCustomOneOf{
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

		deserialized := new(cereal_test.HasCustomOneOf)
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
		WithOneOfFields(new(cereal_test.HasOneOfBytes))

	registryBuilder, err := protoBson.NewCerealRegistryBuilder(cerealOpts)
	if !assert.NoError(err, "create registry builder") {
		t.FailNow()
	}

	registry := registryBuilder.Build()

	message := &cereal_test.HasOneOfBytes{
		Value: &cereal_test.HasOneOfBytes_BytesValue{
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

	deserialized := new(cereal_test.HasOneOfBytes)
	err = bson.UnmarshalWithRegistry(registry, serialized, deserialized)
	if !assert.NoError(err, "unmarshall to protobuf") {
		t.FailNow()
	}

	assert.Equal(message, deserialized, "unmarshalled equals original")
}

func TestMarshallOneOfInMap(t *testing.T) {
	assert := assert.New(t)

	cerealOpts := protoBson.NewMongoOpts().
		WithOneOfFields(new(cereal_test.TestOneOfFirst))

	registryBuilder, err := protoBson.NewCerealRegistryBuilder(cerealOpts)
	if !assert.NoError(err, "create registry builder") {
		t.FailNow()
	}

	registry := registryBuilder.Build()

	original := bson.M{
		"field": &cereal_test.TestOneOfFirst_FieldBool{
			FieldBool: true,
		},
	}

	serialized, err := bson.MarshalWithRegistry(registry, original)
	if !assert.NoError(err, "serialization") {
		t.FailNow()
	}

	document := bson.M{}
	err = bson.UnmarshalWithRegistry(registry, serialized, document)
	if !assert.NoError(err, "document deserialization") {
		t.FailNow()
	}

	assert.Equal(bson.M{"field": true}, document)

	deserialized := make(map[string]*cereal_test.TestOneOfFirst_FieldBool)
	err = bson.UnmarshalWithRegistry(registry, serialized, deserialized)
	if !assert.NoError(err, "deserialization") {
		t.FailNow()
	}

	assert.Equal(true, deserialized["field"].FieldBool)
}

func TestOneOf_Wrapper(t *testing.T) {
	assert := assert.New(t)

	cerealOpts := protoBson.NewMongoOpts().
		WithCustomWrappers(new(cereal_test.TestOneOfFirst)).
		WithOneOfFields(new(cereal_test.TestOneOfFirst))

	registryBuilder, err := protoBson.NewCerealRegistryBuilder(cerealOpts)
	if !assert.NoError(err, "create registry builder") {
		t.FailNow()
	}

	registry := registryBuilder.Build()

	type hasOneOfWrapper struct {
		Field *cereal_test.TestOneOfFirst
	}

	message := &hasOneOfWrapper{
		Field: &cereal_test.TestOneOfFirst{
			SomeValue: &cereal_test.TestOneOfFirst_FieldBool{
				FieldBool: true,
			},
		},
	}

	serialized, err := bson.MarshalWithRegistry(registry, message)
	if !assert.NoError(err, "serialize message") {
		t.FailNow()
	}

	document := bson.M{}
	err = bson.UnmarshalWithRegistry(registry, serialized, document)
	if !assert.NoError(err, "error de-serializing to document") {
		t.FailNow()
	}

	assert.Equal(bson.M{"field": true}, document)

	deserialized := new(hasOneOfWrapper)
	err = bson.UnmarshalWithRegistry(registry, serialized, deserialized)
	if !assert.NoError(err, "error de-serializing into message") {
		t.FailNow()
	}

	assert.Equal(message, deserialized)
}
