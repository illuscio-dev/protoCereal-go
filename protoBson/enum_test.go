package protoBson_test

import (
	"github.com/illuscio-dev/protoCereal-go/cereal_test"
	protoBson "github.com/illuscio-dev/protoCereal-go/protoBson"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"testing"
)

func TestOpts_WithEnumStrings(t *testing.T) {
	builder := bsoncodec.NewRegistryBuilder()

	cerealOpts := protoBson.NewMongoOpts().WithEnumStrings(true)
	err := protoBson.RegisterCerealCodecs(builder, cerealOpts)
	if !assert.NoError(t, err, "error registering codecs") {
		t.FailNow()
	}

	registry := builder.Build()

	type testCase struct {
		House cereal_test.Houses
	}

	testCases := []testCase{
		{
			House: cereal_test.Houses_GRYFFINDOR,
		},
		{
			House: cereal_test.Houses_RAVENCLAW,
		},
		{
			House: cereal_test.Houses_HUFFLEPUFF,
		},
		{
			House: cereal_test.Houses_SLYTHERIN,
		},
	}

	var thisCase testCase

	runTest := func(t *testing.T) {
		assert := assert.New(t)

		message := &cereal_test.EnumTest{
			House: thisCase.House,
		}

		encoded, err := bson.MarshalWithRegistry(registry, message)
		if !assert.NoError(err, "marshal message") {
			t.FailNow()
		}

		dataMap := bson.M{}
		err = bson.UnmarshalWithRegistry(registry, encoded, dataMap)
		if !assert.NoError(err, "unmarshal to map") {
			t.FailNow()
		}

		if !assert.Contains(dataMap, "house") {
			t.FailNow()
		}

		if !assert.Equal(dataMap["house"], thisCase.House.String()) {
			t.FailNow()
		}

		decoded := new(cereal_test.EnumTest)
		err = bson.UnmarshalWithRegistry(registry, encoded, decoded)
		if !assert.NoError(err, "unmarshal to message") {
			t.FailNow()
		}

		assert.Equal(message, decoded)
	}

	for _, thisCase = range testCases {
		t.Run(thisCase.House.String(), runTest)
	}
}
