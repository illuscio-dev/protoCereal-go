package protosql_test

import (
	"errors"
	"github.com/illuscio-dev/protoCereal-go/cereal_test"
	"github.com/illuscio-dev/protoCereal-go/protosql"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Tests that we can round-trip a timestamp
func TestMessageBlobMarshaller(t *testing.T) {
	assert := assert.New(t)

	testCases := []*TestCaseRoundTrip{
		{
			Name:         "NotNil",
			Value:        protosql.Message(&cereal_test.Wizard{Name: "Harry Potter"}),
			Decoded:      protosql.Message(new(cereal_test.Wizard)),
			SQLFieldType: "blob",
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(*protosql.MessageBlobMarshaller)
				decoded := testCase.Decoded.(*protosql.MessageBlobMarshaller)

				if !assert.NotNil(encoded.Message, "encoded not nil") {
					t.FailNow()
				}
				if !assert.IsType(
					new(cereal_test.Wizard),
					encoded.Message,
					"encoded wrong type",
				) {
					t.FailNow()
				}

				if !assert.NotNil(decoded.Message, "encoded not nil") {
					t.FailNow()
				}
				if !assert.IsType(
					new(cereal_test.Wizard),
					decoded.Message,
					"decoded wrong type",
				) {
					t.FailNow()
				}

				decodedMessage := decoded.Message.(*cereal_test.Wizard)
				assert.Equal(decodedMessage.Name, "Harry Potter")
			},
		},
		{
			Name:         "Nil",
			Value:        protosql.Message(nil),
			Decoded:      protosql.Message(new(cereal_test.Wizard)),
			SQLFieldType: "blob",
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*protosql.MessageBlobMarshaller)
				assert.Nil(decoded.Message)
			},
		},
		{
			Name:         "ScanErr_WrongType",
			Value:        "not a blob",
			Decoded:      protosql.Message(new(cereal_test.Wizard)),
			SQLFieldType: "string",
			DecodeErr: errors.New(
				"sql: Scan error on column index 0, name \"value\": expected" +
					" type '[]uint8' for target value of type" +
					" '*cereal_test.Wizard', got 'string'",
			),
		},
		{
			Name:         "ScanErr_UnmarshalErr",
			Value:        []byte("some bytes"),
			Decoded:      protosql.Message(new(cereal_test.Wizard)),
			SQLFieldType: "blob",
			DecodeErr: errors.New(
				"sql: Scan error on column index 0, name \"value\": error " +
					"unmarshalling message: proto: cannot parse reserved " +
					"wire type",
			),
		},
		{
			Name:         "ScanErr_NoTarget",
			Value:        protosql.Message(&cereal_test.Wizard{Name: "Harry Potter"}),
			Decoded:      protosql.Message(nil),
			SQLFieldType: "blob",
			DecodeErr: errors.New(
				"sql: Scan error on column index 0, name \"value\": " +
					"'MessageBlobMarshaller.Message' field is nil: target message " +
					"must be supplied",
			),
		},
	}

	for _, testCase := range testCases {
		RunTestRoundTrip(t, testCase)
	}
}

