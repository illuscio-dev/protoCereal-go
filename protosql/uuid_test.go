package protosql_test

import (
	"errors"
	"github.com/illuscio-dev/protoCereal-go/cereal"
	"github.com/illuscio-dev/protoCereal-go/protosql"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Tests that we can round-trip a uuid
func TestRoundTrip_UUID(t *testing.T) {
	assert := assert.New(t)

	stringVal := ""

	testCases := []*TestCaseRoundTrip{
		{
			Name:         "NotNil",
			Value:        cereal.MustUUIDRandom(),
			Decoded:      new(cereal.UUID),
			SQLFieldType: "blob(16)",
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(*cereal.UUID)
				decoded := testCase.Decoded.(*cereal.UUID)

				assert.Equal(
					encoded.MustGoogle().String(),
					decoded.MustGoogle().String(),
					"UUIDs equal",
				)
			},
		},
		{
			Name:         "Nil",
			Value:        nil,
			Decoded:      new(cereal.UUID),
			SQLFieldType: "blob(16)",
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*cereal.UUID)
				assert.Nil(decoded.Bin)
			},
		},
		{
			Name:         "HexMarshaller",
			Value:        protosql.UUIDHexMarshaller{UUID: cereal.MustUUIDRandom()},
			Decoded:      new(cereal.UUID),
			SQLFieldType: "string",
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.UUIDHexMarshaller)
				decoded := testCase.Decoded.(*cereal.UUID)

				assert.Equal(
					encoded.MustGoogle().String(),
					decoded.MustGoogle().String(),
					"UUIDs equal",
				)
			},
		},
		{
			Name:         "HexMarshaller_Nil",
			Value:        protosql.UUIDHexMarshaller{UUID: nil},
			Decoded:      new(cereal.UUID),
			SQLFieldType: "string",
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*cereal.UUID)

				assert.Nil(
					decoded.Bin,
					"UUID nil",
				)
			},
		},
		{
			Name:         "HexMarshaller_RoundTrip",
			Value:        protosql.UUIDHexMarshaller{UUID: cereal.MustUUIDRandom()},
			Decoded:      new(protosql.UUIDHexMarshaller),
			SQLFieldType: "string",
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.UUIDHexMarshaller)
				decoded := testCase.Decoded.(*protosql.UUIDHexMarshaller)

				assert.Equal(
					encoded.MustGoogle().String(),
					decoded.MustGoogle().String(),
					"UUIDs equal",
				)
			},
		},
		{
			Name:         "HexMarshaller_RoundTrip_Nil",
			Value:        protosql.UUIDHex(nil),
			Decoded:      new(protosql.UUIDHexMarshaller),
			SQLFieldType: "string",
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*protosql.UUIDHexMarshaller)

				assert.Nil(
					decoded.UUID,
					"UUID nil",
				)
			},
		},
		// Test that the hex string is actually being stored and not the blob.
		{
			Name:         "HexMarshaller_StringVal",
			Value:        protosql.UUIDHex(cereal.MustUUIDRandom()),
			Decoded:      &stringVal,
			SQLFieldType: "string",
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.UUIDHexMarshaller)

				assert.Equal(
					encoded.MustGoogle().String(),
					stringVal,
					"UUIDs equal",
				)
			},
		},
		// Test that the hex string is actually being stored and not the blob.
		{
			Name: "HexMarshaller_ErrBadEncode",
			Value: protosql.UUIDHexMarshaller{
				UUID: &cereal.UUID{Bin: []byte("not a uuid")},
			},
			Decoded:      &stringVal,
			SQLFieldType: "string",
			ExpectedEncodeErr: errors.New(
				"sql: converting argument $1 type: error converting uuid bytes: " +
					"proto uuid message must be 16 bytes: 10 bytes found",
			),
		},
		{
			Name:         "ScanErr_BadBytes",
			Value:        []byte("not a uuid"),
			Decoded:      new(cereal.UUID),
			SQLFieldType: "blob(16)",
			ExpectedDecodeErr: errors.New(
				"sql: Scan error on column index 0, name \"value\": proto uuid" +
					" message must be 16 bytes: 10 bytes found",
			),
		},
		{
			Name:         "ScanErr_BadString",
			Value:        "not a uuid",
			Decoded:      new(cereal.UUID),
			SQLFieldType: "string",
			ExpectedDecodeErr: errors.New(
				"sql: Scan error on column index 0, name \"value\": error " +
					"parsing uuid string: invalid UUID length: 10",
			),
		},
	}

	for _, testCase := range testCases {
		RunTestRoundTrip(t, testCase)
	}
}
