package protoBson_test

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	protoBson "github.com/illuscio-dev/protoCereal-go/protoBson"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"strings"
	"testing"
)

func TestGenericWrapper(t *testing.T) {
	type testMessage struct {
		BoolVal     *wrappers.BoolValue
		BytesVal    *wrappers.BytesValue
		DoubleValue *wrappers.DoubleValue
		FloatValue  *wrappers.FloatValue
		Int32Value  *wrappers.Int32Value
		Int64Value  *wrappers.Int64Value
		StringValue *wrappers.StringValue
		UInt32Value *wrappers.UInt32Value
		UInt64Value *wrappers.UInt64Value
	}

	type TestCase struct {
		field           string
		value           interface{}
		serializedValue interface{}
	}

	testCases := []*TestCase{
		{
			field: "BoolVal",
			value: &wrappers.BoolValue{
				Value: true,
			},
			serializedValue: true,
		},
		{
			field:           "BoolVal",
			value:           nil,
			serializedValue: nil,
		},
		{
			field: "BytesVal",
			value: &wrappers.BytesValue{
				Value: []byte("some data"),
			},
			serializedValue: primitive.Binary{
				Subtype: 0,
				Data:    []byte("some data"),
			},
		},
		{
			field:           "BytesVal",
			value:           nil,
			serializedValue: nil,
		},
		{
			field: "DoubleValue",
			value: &wrappers.DoubleValue{
				Value: 0.0002688172043010753,
			},
			serializedValue: float64(0.0002688172043010753),
		},
		{
			field:           "DoubleValue",
			value:           nil,
			serializedValue: nil,
		},
		{
			field: "FloatValue",
			value: &wrappers.FloatValue{
				Value: 0.0002688172043010753,
			},
			serializedValue: float64(0.00026881720987148583),
		},
		{
			field:           "FloatValue",
			value:           nil,
			serializedValue: nil,
		},
		{
			field: "Int32Value",
			value: &wrappers.Int32Value{
				Value: 42,
			},
			serializedValue: int32(42),
		},
		{
			field:           "Int32Value",
			value:           nil,
			serializedValue: nil,
		},
		{
			field: "Int64Value",
			value: &wrappers.Int64Value{
				Value: 42,
			},
			serializedValue: int64(42),
		},
		{
			field:           "Int64Value",
			value:           nil,
			serializedValue: nil,
		},
		{
			field: "StringValue",
			value: &wrappers.StringValue{
				Value: "some text",
			},
			serializedValue: "some text",
		},
		{
			field:           "StringValue",
			value:           nil,
			serializedValue: nil,
		},
		{
			field: "UInt32Value",
			value: &wrappers.UInt32Value{
				Value: 42,
			},
			serializedValue: int64(42),
		},
		{
			field:           "UInt32Value",
			value:           nil,
			serializedValue: nil,
		},
		{
			field: "UInt64Value",
			value: &wrappers.UInt64Value{
				Value: 42,
			},
			serializedValue: int64(42),
		},
		{
			field:           "UInt64Value",
			value:           nil,
			serializedValue: nil,
		},
	}

	var thisCase *TestCase

	builder := bson.NewRegistryBuilder()
	err := protoBson.RegisterCerealCodecs(builder, nil)
	assert.NoError(t, err, "error building registry")
	if err != nil {
		t.FailNow()
	}
	registry := builder.Build()

	runTest := func(t *testing.T) {
		assert := assert.New(t)

		original := new(testMessage)
		value := reflect.ValueOf(original).Elem().FieldByName(thisCase.field)

		if thisCase.value == nil {
			value.Set(reflect.Zero(value.Type()))
		} else {
			value.Set(reflect.ValueOf(thisCase.value))
		}

		encoded, err := bson.MarshalWithRegistry(registry, original)
		if !assert.NoError(err, "error marshalling message") {
			t.FailNow()
		}

		mapData := make(bson.M)
		err = bson.UnmarshalWithRegistry(registry, encoded, mapData)
		if !assert.NoError(err, "error unmarshalling to map") {
			t.FailNow()
		}

		mapKey := strings.ToLower(thisCase.field)
		if !assert.Contains(mapData, mapKey, "map has correct key") {
			t.FailNow()
		}

		if !assert.Equal(
			thisCase.serializedValue,
			mapData[mapKey],
			"correct serialized value",
		) {
			t.FailNow()
		}

		decoded := new(testMessage)
		err = bson.UnmarshalWithRegistry(registry, encoded, decoded)
		if !assert.NoError(err, "error unmarshalling to proto") {
			t.FailNow()
		}

		assert.Equal(original, decoded, "unmarshal equality")
	}

	for _, thisCase = range testCases {
		name := thisCase.field
		if thisCase.value == nil {
			name = name + "_nil"
		}
		t.Run(name, runTest)
	}
}
