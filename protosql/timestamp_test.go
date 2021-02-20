package protosql_test

import (
	"errors"
	"github.com/illuscio-dev/protoCereal-go/protosql"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

// Tests that we can round-trip a timestamp
func TestRoundTrip_DatetimeField(t *testing.T) {
	assert := assert.New(t)

	testCases := []*TestCaseRoundTrip{
		{
			Name: "NotNil",
			Value: protosql.TimestampMarshaller{
				Timestamp: timestamppb.New(time.Now().UTC()),
			},
			Decoded:      &protosql.TimestampMarshaller{},
			SQLFieldType: "datetime",
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.TimestampMarshaller)
				decoded := testCase.Decoded.(*protosql.TimestampMarshaller)

				assert.True(
					decoded.AsTime().Equal(encoded.AsTime()),
					"%v equals %v",
					decoded.AsTime(),
					encoded.AsTime(),
				)
			},
		},
		{
			Name: "Nil",
			Value: protosql.TimestampMarshaller{
				Timestamp: nil,
			},
			Decoded:      &protosql.TimestampMarshaller{},
			SQLFieldType: "datetime",
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*protosql.TimestampMarshaller)
				assert.Nil(decoded.Timestamp)
			},
		},
		{
			Name:         "ScanErr",
			Value:        "not a timestamp",
			Decoded:      &protosql.TimestampMarshaller{},
			SQLFieldType: "string",
			DecodeErr: errors.New(
				"sql: Scan error on column index 0, name \"value\": " +
					"unsupported Scan, storing driver.Value type string into type" +
					" *time.Time",
			),
		},
	}

	for _, testCase := range testCases {
		RunTestRoundTrip(t, testCase)
	}
}
