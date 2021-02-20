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
			Name: "NotNull",
			Value: protosql.BoolMarshaller{
				BoolValue: wrapperspb.Bool(true),
			},
			Decoded:      new(protosql.BoolMarshaller),
			SQLFieldType: "boolean",
			EncodeErr:    nil,
			DecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.BoolMarshaller)
				decoded := testCase.Decoded.(*protosql.BoolMarshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.BoolMarshaller{BoolValue: nil},
			Decoded:      new(protosql.BoolMarshaller),
			SQLFieldType: "boolean",
			EncodeErr:    nil,
			DecodeErr:    nil,
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
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name " +
				"\"value\": expected type 'bool' for target value of type " +
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
			Name: "NotNull",
			Value: protosql.BytesMarshaller{
				BytesValue: wrapperspb.Bytes([]byte("hello")),
			},
			Decoded:      new(protosql.BytesMarshaller),
			SQLFieldType: "blob",
			EncodeErr:    nil,
			DecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.BytesMarshaller)
				decoded := testCase.Decoded.(*protosql.BytesMarshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.BytesMarshaller{BytesValue: nil},
			Decoded:      new(protosql.BytesMarshaller),
			SQLFieldType: "blob",
			EncodeErr:    nil,
			DecodeErr:    nil,
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
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name " +
				"\"value\": expected type '[]uint8' for target value of type " +
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
			Name: "NotNull",
			Value: protosql.DoubleMarshaller{
				DoubleValue: wrapperspb.Double(64),
			},
			Decoded:      new(protosql.DoubleMarshaller),
			SQLFieldType: "double",
			EncodeErr:    nil,
			DecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.DoubleMarshaller)
				decoded := testCase.Decoded.(*protosql.DoubleMarshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.DoubleMarshaller{DoubleValue: nil},
			Decoded:      new(protosql.DoubleMarshaller),
			SQLFieldType: "double",
			EncodeErr:    nil,
			DecodeErr:    nil,
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
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name " +
				"\"value\": expected type 'float64' for target value of type " +
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
			Name: "NotNull",
			Value: protosql.FloatMarshaller{
				FloatValue: wrapperspb.Float(64),
			},
			Decoded:      new(protosql.FloatMarshaller),
			SQLFieldType: "float",
			EncodeErr:    nil,
			DecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.FloatMarshaller)
				decoded := testCase.Decoded.(*protosql.FloatMarshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.FloatMarshaller{FloatValue: nil},
			Decoded:      new(protosql.FloatMarshaller),
			SQLFieldType: "float",
			EncodeErr:    nil,
			DecodeErr:    nil,
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
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name " +
				"\"value\": expected type 'float64' for target value of type " +
				"'*wrapperspb.FloatValue', got 'int64'"),
			SubTest: nil,
		},
		{
			Name:         "Err_DecodeOverFlow",
			Value:        math.MaxFloat64,
			Decoded:      new(protosql.FloatMarshaller),
			SQLFieldType: "double",
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name " +
				"\"value\": incoming float64 value overflows float32"),
			SubTest: nil,
		},
		{
			Name:         "Err_DecodeUnderFlow",
			Value:        math.MaxFloat64 * -1,
			Decoded:      new(protosql.FloatMarshaller),
			SQLFieldType: "double",
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name " +
				"\"value\": incoming float64 value underflows float32"),
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
			Name: "NotNull",
			Value: protosql.Int32Marshaller{
				Int32Value: wrapperspb.Int32(42),
			},
			Decoded:      new(protosql.Int32Marshaller),
			SQLFieldType: "integer",
			EncodeErr:    nil,
			DecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.Int32Marshaller)
				decoded := testCase.Decoded.(*protosql.Int32Marshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.Int32Marshaller{Int32Value: nil},
			Decoded:      new(protosql.Int32Marshaller),
			SQLFieldType: "integer",
			EncodeErr:    nil,
			DecodeErr:    nil,
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
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name " +
				"\"value\": expected type 'int64' for target value of type " +
				"'*wrapperspb.Int32Value', got 'string'"),
			SubTest: nil,
		},
		{
			Name:         "Err_DecodeOverflow",
			Value:        math.MaxInt32 + 1,
			Decoded:      new(protosql.Int32Marshaller),
			SQLFieldType: "integer",
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name " +
				"\"value\": received value overflows int32"),
			SubTest: nil,
		},
		{
			Name:         "Err_DecodeUnderflow",
			Value:        math.MinInt32 - 1,
			Decoded:      new(protosql.Int32Marshaller),
			SQLFieldType: "integer",
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name " +
				"\"value\": received value underflows int32"),
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
			Name: "NotNull",
			Value: protosql.Int64Marshaller{
				Int64Value: wrapperspb.Int64(42),
			},
			Decoded:      new(protosql.Int64Marshaller),
			SQLFieldType: "integer",
			EncodeErr:    nil,
			DecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.Int64Marshaller)
				decoded := testCase.Decoded.(*protosql.Int64Marshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.Int64Marshaller{Int64Value: nil},
			Decoded:      new(protosql.Int64Marshaller),
			SQLFieldType: "integer",
			EncodeErr:    nil,
			DecodeErr:    nil,
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
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name " +
				"\"value\": expected type 'int64' for target value of type " +
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
			Name: "NotNull",
			Value: protosql.StringMarshaller{
				StringValue: wrapperspb.String("some value"),
			},
			Decoded:      new(protosql.StringMarshaller),
			SQLFieldType: "text",
			EncodeErr:    nil,
			DecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.StringMarshaller)
				decoded := testCase.Decoded.(*protosql.StringMarshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name: "Null",
			Value: protosql.StringMarshaller{
				StringValue: nil,
			},
			Decoded:      new(protosql.StringMarshaller),
			SQLFieldType: "text",
			EncodeErr:    nil,
			DecodeErr:    nil,
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
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name " +
				"\"value\": expected type 'string' for target value of type " +
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
			Name: "NotNull",
			Value: protosql.UInt32Marshaller{
				UInt32Value: wrapperspb.UInt32(42),
			},
			Decoded:      new(protosql.UInt32Marshaller),
			SQLFieldType: "integer",
			EncodeErr:    nil,
			DecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.UInt32Marshaller)
				decoded := testCase.Decoded.(*protosql.UInt32Marshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.UInt32Marshaller{UInt32Value: nil},
			Decoded:      new(protosql.UInt32Marshaller),
			SQLFieldType: "integer",
			EncodeErr:    nil,
			DecodeErr:    nil,
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
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name " +
				"\"value\": expected type 'int64' for target value of type " +
				"'*wrapperspb.UInt32Value', got 'string'"),
			SubTest: nil,
		},
		{
			Name:         "Err_DecodeOverflow",
			Value:        uint64(math.MaxUint32 + 1),
			Decoded:      new(protosql.UInt32Marshaller),
			SQLFieldType: "integer",
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name " +
				"\"value\": received value overflows uint32"),
			SubTest: nil,
		},
		{
			Name:         "Err_DecodeUnderFlow",
			Value:        -1,
			Decoded:      new(protosql.UInt32Marshaller),
			SQLFieldType: "integer",
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name " +
				"\"value\": received value underflows uint32"),
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
			Name: "NotNull",
			Value: protosql.UInt64Marshaller{
				UInt64Value: wrapperspb.UInt64(42),
			},
			Decoded:      new(protosql.UInt64Marshaller),
			SQLFieldType: "integer",
			EncodeErr:    nil,
			DecodeErr:    nil,
			SubTest: func(t *testing.T, testCase *TestCaseRoundTrip) {
				encoded := testCase.Value.(protosql.UInt64Marshaller)
				decoded := testCase.Decoded.(*protosql.UInt64Marshaller)

				assert.Equal(t, encoded.GetValue(), decoded.GetValue())
			},
		},
		{
			Name:         "Null",
			Value:        protosql.UInt64Marshaller{UInt64Value: nil},
			Decoded:      new(protosql.UInt64Marshaller),
			SQLFieldType: "integer",
			EncodeErr:    nil,
			DecodeErr:    nil,
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
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name " +
				"\"value\": expected type 'int64' for target value of type " +
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
			EncodeErr: errors.New("sql: converting argument $1 type: cannot " +
				"encode uint64: overflows int64"),
			DecodeErr: nil,
			SubTest:   nil,
		},
		{
			Name:         "Err_DecodeUnderflow",
			Value:        -1,
			Decoded:      new(protosql.UInt64Marshaller),
			SQLFieldType: "integer",
			EncodeErr:    nil,
			DecodeErr: errors.New("sql: Scan error on column index 0, name" +
				" \"value\": received value underflows uint64"),
			SubTest: nil,
		},
	}

	for _, thisCase := range cases {
		RunTestRoundTrip(t, thisCase)
	}
}
