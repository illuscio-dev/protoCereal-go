package protosql_test

import (
	"errors"
	"github.com/illuscio-dev/protoCereal-go/cereal_test"
	"github.com/illuscio-dev/protoCereal-go/protosql"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func checkEnums(
	t *testing.T, testCase *TestCaseRoundTrip, expectedEnumVal interface{},
) {
	encoded := testCase.Value.(*protosql.EnumStringMarshaller)
	decoded := testCase.Decoded.(*protosql.EnumStringMarshaller)

	assert := assert.New(t)
	if !assert.IsType(expectedEnumVal, encoded.Enum, "encoded is type") {
		t.FailNow()
	}
	if !assert.IsType(expectedEnumVal, decoded.Enum, "decoded is type") {
		t.FailNow()
	}

	assert.Equal(expectedEnumVal, decoded.Enum, "decoded is expected")
}

func TestEnums(t *testing.T) {
	cases := []*TestCaseRoundTrip{
		{
			Name:         "Houses_GRYFFINDOR",
			Value:        protosql.Enum(cereal_test.Houses_GRYFFINDOR),
			Decoded:      protosql.Enum(cereal_test.Houses_GRYFFINDOR),
			SQLFieldType: "text",
			EncodeErr:    nil,
			DecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				checkEnums(t, testCase, cereal_test.Houses_GRYFFINDOR)
			},
		},
		{
			Name:         "Houses_RAVENCLAW",
			Value:        protosql.Enum(cereal_test.Houses_RAVENCLAW),
			Decoded:      protosql.Enum(cereal_test.Houses_GRYFFINDOR),
			SQLFieldType: "text",
			EncodeErr:    nil,
			DecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				checkEnums(t, testCase, cereal_test.Houses_RAVENCLAW)
			},
		},
		{
			Name:         "Houses_HUFFLEPUFF",
			Value:        protosql.Enum(cereal_test.Houses_HUFFLEPUFF),
			Decoded:      protosql.Enum(cereal_test.Houses_GRYFFINDOR),
			SQLFieldType: "text",
			EncodeErr:    nil,
			DecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				checkEnums(t, testCase, cereal_test.Houses_HUFFLEPUFF)
			},
		},
		{
			Name:         "Houses_SLYTHERIN",
			Value:        protosql.Enum(cereal_test.Houses_SLYTHERIN),
			Decoded:      protosql.Enum(cereal_test.Houses_GRYFFINDOR),
			SQLFieldType: "text",
			EncodeErr:    nil,
			DecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				checkEnums(t, testCase, cereal_test.Houses_SLYTHERIN)
			},
		},
		// This tests decodes the value as a string, and ensures that it was encoded
		// correctly.
		{
			Name:         "Houses_SLYTHERIN_GetString",
			Value:        protosql.Enum(cereal_test.Houses_SLYTHERIN),
			Decoded:      new(protosql.StringMarshaller),
			SQLFieldType: "text",
			EncodeErr:    nil,
			DecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*protosql.StringMarshaller)

				assert.NotNil(
					t, decoded.StringValue, "string value not nil",
				)
				assert.Equal(
					t, cereal_test.Houses_SLYTHERIN.String(), decoded.GetValue(),
				)
			},
		},
		{
			Name:         "NilValue",
			Value:        protosql.Enum(nil),
			Decoded:      protosql.Enum(cereal_test.Houses_GRYFFINDOR),
			SQLFieldType: "text",
			EncodeErr:    nil,
			DecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*protosql.EnumStringMarshaller)
				assert.Nilf(t, decoded.Enum, "enum value nil")
			},
		},
		{
			Name:         "Err_Type",
			Value:        time.Now(),
			Decoded:      protosql.Enum(cereal_test.Houses_GRYFFINDOR),
			SQLFieldType: "datetime",
			EncodeErr:    nil,
			DecodeErr: errors.New(
				"sql: Scan error on column index 0, name \"value\": expected " +
					"type 'string' for target value of type " +
					"'cereal_test.Houses', got 'time.Time'",
			),
			SubTest: nil,
		},
		{
			Name:         "Err_UnknownString",
			Value:        "NotAHouse",
			Decoded:      protosql.Enum(cereal_test.Houses_GRYFFINDOR),
			SQLFieldType: "text",
			EncodeErr:    nil,
			DecodeErr: errors.New(
				"sql: Scan error on column index 0, name \"value\": received " +
					"string does not match any known Enum names for type " +
					"'cereal_test.Houses'",
			),
			SubTest: nil,
		},
		{
			Name:         "Err_Decode_NoPrototype",
			Value:        protosql.Enum(cereal_test.Houses_SLYTHERIN),
			Decoded:      protosql.Enum(nil),
			SQLFieldType: "text",
			EncodeErr:    nil,
			DecodeErr: errors.New(
				"sql: Scan error on column index 0, name \"value\": cannot " +
					"unmarshal proto Enum, Enum field is not set with concrete value " +
					"for type inspection",
			),
			SubTest: nil,
		},
	}

	for _, thisCase := range cases {
		RunTestRoundTrip(t, thisCase)
	}
}
