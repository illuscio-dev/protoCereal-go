package protosql_test

import (
	"errors"
	"github.com/illuscio-dev/protoCereal-go/protosql"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"math"
	"testing"
)

func TestWrappersBool(t *testing.T) {
	cases := []*TestCaseRoundTrip{
		{
			Name:         "NotNull",
			Value:        protosql.Bool(wrapperspb.Bool(true)),
			Decoded:      new(protosql.BoolMarshaller),
			SQLFieldType: "boolean",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.BoolMarshaller)
				decoded := testCase.Decoded.(*protosql.BoolMarshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.Bool(nil),
			Decoded:      new(protosql.BoolMarshaller),
			SQLFieldType: "boolean",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*protosql.BoolMarshaller)
				assert.Nil(t, decoded.BoolValue)
			},
		},
		{
			Name:         "ErrType",
			Value:        11,
			Decoded:      new(protosql.BoolMarshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": expected type 'bool' for target value of type " +
				"'*wrapperspb.BoolValue', got 'int64'"),
			SubTest: nil,
		},
	}

	for _, thisCase := range cases {
		RunTestRoundTrip(t, thisCase)
	}
}

func TestWrappersBytes(t *testing.T) {
	cases := []*TestCaseRoundTrip{
		{
			Name:         "NotNull",
			Value:        protosql.Bytes(wrapperspb.Bytes([]byte("hello"))),
			Decoded:      new(protosql.BytesMarshaller),
			SQLFieldType: "blob",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.BytesMarshaller)
				decoded := testCase.Decoded.(*protosql.BytesMarshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.Bytes(nil),
			Decoded:      new(protosql.BytesMarshaller),
			SQLFieldType: "blob",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*protosql.BytesMarshaller)
				assert.Nil(t, decoded.BytesValue)
			},
		},
		{
			Name:         "ErrType",
			Value:        11,
			Decoded:      new(protosql.BytesMarshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": expected type '[]uint8' for target value of type " +
				"'*wrapperspb.BytesValue', got 'int64'"),
			SubTest: nil,
		},
	}

	for _, thisCase := range cases {
		RunTestRoundTrip(t, thisCase)
	}
}

func TestWrappersDouble(t *testing.T) {
	cases := []*TestCaseRoundTrip{
		{
			Name:         "NotNull",
			Value:        protosql.Double(wrapperspb.Double(64)),
			Decoded:      new(protosql.DoubleMarshaller),
			SQLFieldType: "double",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.DoubleMarshaller)
				decoded := testCase.Decoded.(*protosql.DoubleMarshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.Double(nil),
			Decoded:      new(protosql.DoubleMarshaller),
			SQLFieldType: "double",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*protosql.DoubleMarshaller)
				assert.Nil(t, decoded.DoubleValue)
			},
		},
		{
			Name:         "ErrType",
			Value:        11,
			Decoded:      new(protosql.DoubleMarshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": expected type 'float64' for target value of type " +
				"'*wrapperspb.DoubleValue', got 'int64'"),
			SubTest: nil,
		},
	}

	for _, thisCase := range cases {
		RunTestRoundTrip(t, thisCase)
	}
}

func TestWrappersFloat(t *testing.T) {
	cases := []*TestCaseRoundTrip{
		{
			Name:         "NotNull",
			Value:        protosql.Float(wrapperspb.Float(64)),
			Decoded:      new(protosql.FloatMarshaller),
			SQLFieldType: "float",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.FloatMarshaller)
				decoded := testCase.Decoded.(*protosql.FloatMarshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.Float(nil),
			Decoded:      new(protosql.FloatMarshaller),
			SQLFieldType: "float",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*protosql.FloatMarshaller)
				assert.Nil(t, decoded.FloatValue)
			},
		},
		{
			Name:         "Err_Type",
			Value:        11,
			Decoded:      new(protosql.FloatMarshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": expected type 'float64' for target value of type " +
				"'*wrapperspb.FloatValue', got 'int64'"),
			SubTest: nil,
		},
		{
			Name:         "Err_DecodeOverFlow",
			Value:        math.MaxFloat64,
			Decoded:      new(protosql.FloatMarshaller),
			SQLFieldType: "double",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": incoming float64 value overflows float32"),
			SubTest: nil,
		},
		{
			Name:         "Err_DecodeUnderFlow",
			Value:        math.MaxFloat64 * -1,
			Decoded:      new(protosql.FloatMarshaller),
			SQLFieldType: "double",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": incoming float64 value underflows float32"),
			SubTest: nil,
		},
	}

	for _, thisCase := range cases {
		RunTestRoundTrip(t, thisCase)
	}
}

func TestWrappersInt32(t *testing.T) {
	cases := []*TestCaseRoundTrip{
		{
			Name:         "NotNull",
			Value:        protosql.Int32(wrapperspb.Int32(42)),
			Decoded:      new(protosql.Int32Marshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.Int32Marshaller)
				decoded := testCase.Decoded.(*protosql.Int32Marshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.Int32(nil),
			Decoded:      new(protosql.Int32Marshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*protosql.Int32Marshaller)
				assert.Nil(t, decoded.Int32Value)
			},
		},
		{
			Name:         "Err_Type",
			Value:        "some value",
			Decoded:      new(protosql.Int32Marshaller),
			SQLFieldType: "text",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": expected type 'int64' for target value of type " +
				"'*wrapperspb.Int32Value', got 'string'"),
			SubTest: nil,
		},
		{
			Name:         "Err_DecodeOverflow",
			Value:        math.MaxInt32 + 1,
			Decoded:      new(protosql.Int32Marshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": received value overflows int32"),
			SubTest: nil,
		},
		{
			Name:         "Err_DecodeUnderflow",
			Value:        math.MinInt32 - 1,
			Decoded:      new(protosql.Int32Marshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": received value underflows int32"),
			SubTest: nil,
		},
	}

	for _, thisCase := range cases {
		RunTestRoundTrip(t, thisCase)
	}
}

func TestWrappersInt64(t *testing.T) {
	cases := []*TestCaseRoundTrip{
		{
			Name:         "NotNull",
			Value:        protosql.Int64(wrapperspb.Int64(42)),
			Decoded:      new(protosql.Int64Marshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.Int64Marshaller)
				decoded := testCase.Decoded.(*protosql.Int64Marshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.Int64(nil),
			Decoded:      new(protosql.Int64Marshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*protosql.Int64Marshaller)
				assert.Nil(t, decoded.Int64Value)
			},
		},
		{
			Name:         "Err_Type",
			Value:        "some value",
			Decoded:      new(protosql.Int64Marshaller),
			SQLFieldType: "text",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": expected type 'int64' for target value of type " +
				"'*wrapperspb.Int64Value', got 'string'"),
			SubTest: nil,
		},
	}

	for _, thisCase := range cases {
		RunTestRoundTrip(t, thisCase)
	}
}

func TestWrappersString(t *testing.T) {
	cases := []*TestCaseRoundTrip{
		{
			Name:         "NotNull",
			Value:        protosql.String(wrapperspb.String("some value")),
			Decoded:      new(protosql.StringMarshaller),
			SQLFieldType: "text",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.StringMarshaller)
				decoded := testCase.Decoded.(*protosql.StringMarshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.String(nil),
			Decoded:      new(protosql.StringMarshaller),
			SQLFieldType: "text",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*protosql.StringMarshaller)
				assert.Nil(t, decoded.StringValue)
			},
		},
		{
			Name:         "Err_Type",
			Value:        11,
			Decoded:      new(protosql.StringMarshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": expected type 'string' for target value of type " +
				"'*wrapperspb.StringValue', got 'int64'"),
			SubTest: nil,
		},
	}

	for _, thisCase := range cases {
		RunTestRoundTrip(t, thisCase)
	}
}

func TestWrappersUInt32(t *testing.T) {
	cases := []*TestCaseRoundTrip{
		{
			Name:         "NotNull",
			Value:        protosql.UInt32(wrapperspb.UInt32(42)),
			Decoded:      new(protosql.UInt32Marshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.UInt32Marshaller)
				decoded := testCase.Decoded.(*protosql.UInt32Marshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.UInt32(nil),
			Decoded:      new(protosql.UInt32Marshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*protosql.UInt32Marshaller)
				assert.Nil(t, decoded.UInt32Value)
			},
		},
		{
			Name:         "ErrType",
			Value:        "some value",
			Decoded:      new(protosql.UInt32Marshaller),
			SQLFieldType: "text",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": expected type 'int64' for target value of type " +
				"'*wrapperspb.UInt32Value', got 'string'"),
			SubTest: nil,
		},
		{
			Name:         "Err_DecodeOverflow",
			Value:        uint64(math.MaxUint32 + 1),
			Decoded:      new(protosql.UInt32Marshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": received value overflows uint32"),
			SubTest: nil,
		},
		{
			Name:         "Err_DecodeUnderFlow",
			Value:        -1,
			Decoded:      new(protosql.UInt32Marshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": received value underflows uint32"),
			SubTest: nil,
		},
	}

	for _, thisCase := range cases {
		RunTestRoundTrip(t, thisCase)
	}
}

func TestWrappersUInt64(t *testing.T) {
	cases := []*TestCaseRoundTrip{
		{
			Name:         "NotNull",
			Value:        protosql.UInt64(wrapperspb.UInt64(42)),
			Decoded:      new(protosql.UInt64Marshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.UInt64Marshaller)
				decoded := testCase.Decoded.(*protosql.UInt64Marshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.UInt64(nil),
			Decoded:      new(protosql.UInt64Marshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				decoded := testCase.Decoded.(*protosql.UInt64Marshaller)
				assert.Nil(t, decoded.UInt64Value)
			},
		},
		{
			Name:         "Err_Type",
			Value:        "some value",
			Decoded:      new(protosql.UInt64Marshaller),
			SQLFieldType: "text",
			ExpectedEncodeErr:    nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": expected type 'int64' for target value of type " +
				"'*wrapperspb.UInt64Value', got 'string'"),
			SubTest: nil,
		},
		{
			Name: "Err_EncodeOverflow",
			Value: protosql.UInt64Marshaller{
				UInt64Value: wrapperspb.UInt64(math.MaxUint64),
			},
			Decoded:      nil,
			SQLFieldType: "integer",
			ExpectedEncodeErr: errors.New("sql: converting argument $1 type:" +
				" cannot encode uint64: overflows int64"),
			ExpectedDecodeErr: nil,
			SubTest:   nil,
		},
		{
			Name:         "Err_DecodeUnderflow",
			Value:        -1,
			Decoded:      new(protosql.UInt64Marshaller),
			SQLFieldType: "integer",
			ExpectedEncodeErr: nil,
			ExpectedDecodeErr: errors.New("sql: Scan error on column index 0, " +
				"name \"value\": received value underflows uint64"),
			SubTest: nil,
		},
	}

	for _, thisCase := range cases {
		RunTestRoundTrip(t, thisCase)
	}
}
