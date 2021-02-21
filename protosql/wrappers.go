package protosql

/*
This file contains sql marshalling wrappers for protobuf well-known-type wrapper values.
*/

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"math"
	"reflect"
)

func newScanTypeErr(
	expectedSrc interface{}, receivedSrc interface{}, wrapperType interface{},
) error {
	return fmt.Errorf(
		"expected type '%v' for target value of type '%v', got '%v'",
		reflect.TypeOf(expectedSrc),
		reflect.TypeOf(wrapperType),
		reflect.TypeOf(receivedSrc),
	)
}

// BoolMarshaller marshals a wrapped *wrapperspb.BoolValue value, inserting
// null if the pointer is null, or the inner value if it is not.
type BoolMarshaller struct {
	*wrapperspb.BoolValue
}

// Value converts the value to a database-friendly format, or nil if the inner value is
// nil.
func (value BoolMarshaller) Value() (driver.Value, error) {
	if value.BoolValue == nil {
		return nil, nil
	}
	return value.BoolValue.Value, nil
}

// Scan converts a value from the database.
func (value *BoolMarshaller) Scan(src interface{}) error {
	if src == nil {
		*value = BoolMarshaller{}
		return nil
	}

	srcVal, ok := src.(bool)
	if !ok {
		return newScanTypeErr(
			false,
			src,
			value.BoolValue,
		)
	}

	value.BoolValue = wrapperspb.Bool(srcVal)
	return nil
}

// BytesMarshaller marshals a wrapped *wrapperspb.BytesValue value, inserting
// null if the pointer is null, or the inner value if it is not.
type BytesMarshaller struct {
	*wrapperspb.BytesValue
}

// Value converts the value to a database-friendly format, or nil if the inner value is
// nil.
func (value BytesMarshaller) Value() (driver.Value, error) {
	if value.BytesValue == nil {
		return nil, nil
	}
	return value.BytesValue.Value, nil
}

// Scan converts a value from the database.
func (value *BytesMarshaller) Scan(src interface{}) error {
	if src == nil {
		*value = BytesMarshaller{}
		return nil
	}

	srcVal, ok := src.([]byte)
	if !ok {
		return newScanTypeErr(
			[]byte{},
			src,
			value.BytesValue,
		)
	}

	// We need to copy the value here -- the driver owns the memory passed in here and
	// it may be re-used once this method exits.
	copiedVal := make([]byte, len(srcVal))
	copy(copiedVal, srcVal)

	value.BytesValue = wrapperspb.Bytes(copiedVal)
	return nil
}

// DoubleMarshaller marshals a wrapped *wrapperspb.DoubleValue value, inserting
// null if the pointer is null, or the inner value if it is not.
type DoubleMarshaller struct {
	*wrapperspb.DoubleValue
}

// Value converts the value to a database-friendly format, or nil if the inner value is
// nil.
func (value DoubleMarshaller) Value() (driver.Value, error) {
	if value.DoubleValue == nil {
		return nil, nil
	}
	return value.DoubleValue.Value, nil
}

// Scan converts a value from the database.
func (value *DoubleMarshaller) Scan(src interface{}) error {
	if src == nil {
		*value = DoubleMarshaller{}
		return nil
	}

	srcVal, ok := src.(float64)
	if !ok {
		return newScanTypeErr(
			float64(0),
			src,
			value.DoubleValue,
		)
	}

	value.DoubleValue = wrapperspb.Double(srcVal)
	return nil
}

// FloatMarshaller marshals a wrapped *wrapperspb.FloatValue value, inserting
// null if the pointer is null, or the inner value if it is not.
type FloatMarshaller struct {
	*wrapperspb.FloatValue
}

// Value converts the value to a database-friendly format, or nil if the inner value is
// nil.
func (value FloatMarshaller) Value() (driver.Value, error) {
	if value.FloatValue == nil {
		return nil, nil
	}
	// Cast inner value to float64 (the driver does not support float32)
	return float64(value.FloatValue.Value), nil
}

// Scan converts a value from the database.
func (value *FloatMarshaller) Scan(src interface{}) error {
	if src == nil {
		*value = FloatMarshaller{}
		return nil
	}

	srcVal, ok := src.(float64)
	if !ok {
		return newScanTypeErr(
			float64(0),
			src,
			value.FloatValue,
		)
	}

	// Check that we aren't going to overflow the value.
	if srcVal > math.MaxFloat32 {
		return errors.New("incoming float64 value overflows float32")
	} else if srcVal < math.MaxFloat32*-1 {
		return errors.New("incoming float64 value underflows float32")
	}

	value.FloatValue = wrapperspb.Float(float32(srcVal))
	return nil
}

// Int32Marshaller marshals a wrapped *wrapperspb.Int32Value value, inserting
// null if the pointer is null, or the inner value if it is not.
type Int32Marshaller struct {
	*wrapperspb.Int32Value
}

// Value converts the value to a database-friendly format, or nil if the inner value is
// nil.
func (value Int32Marshaller) Value() (driver.Value, error) {
	if value.Int32Value == nil {
		return nil, nil
	}
	// The sql driver only supports int64 so we need to case it.
	return int64(value.Int32Value.Value), nil
}

// Scan converts a value from the database.
func (value *Int32Marshaller) Scan(src interface{}) error {
	if src == nil {
		*value = Int32Marshaller{}
		return nil
	}

	srcVal, ok := src.(int64)
	if !ok {
		return newScanTypeErr(
			int64(0),
			src,
			value.Int32Value,
		)
	}

	if srcVal > math.MaxInt32 {
		return errors.New("received value overflows int32")
	} else if srcVal < math.MinInt32 {
		return errors.New("received value underflows int32")
	}

	value.Int32Value = wrapperspb.Int32(int32(srcVal))
	return nil
}

// Int64Marshaller marshals a wrapped *wrapperspb.Int64Value value, inserting
// null if the pointer is null, or the inner value if it is not.
type Int64Marshaller struct {
	*wrapperspb.Int64Value
}

// Value converts the value to a database-friendly format, or nil if the inner value is
// nil.
func (value Int64Marshaller) Value() (driver.Value, error) {
	if value.Int64Value == nil {
		return nil, nil
	}
	return value.Int64Value.Value, nil
}

// Scan converts a value from the database.
func (value *Int64Marshaller) Scan(src interface{}) error {
	if src == nil {
		*value = Int64Marshaller{}
		return nil
	}

	srcVal, ok := src.(int64)
	if !ok {
		return newScanTypeErr(
			int64(0),
			src,
			value.Int64Value,
		)
	}

	value.Int64Value = wrapperspb.Int64(srcVal)
	return nil
}

// StringMarshaller marshals a wrapped *wrapperspb.StringValue value, inserting
// null if the pointer is null, or the inner value if it is not.
type StringMarshaller struct {
	*wrapperspb.StringValue
}

// Value converts the value to a database-friendly format, or nil if the inner value is
// nil.
func (value StringMarshaller) Value() (driver.Value, error) {
	if value.StringValue == nil {
		return nil, nil
	}
	return value.StringValue.Value, nil
}

// Scan converts a value from the database.
func (value *StringMarshaller) Scan(src interface{}) error {
	if src == nil {
		*value = StringMarshaller{}
		return nil
	}

	srcVal, ok := src.(string)
	if !ok {
		return newScanTypeErr(
			"",
			src,
			value.StringValue,
		)
	}

	value.StringValue = wrapperspb.String(srcVal)
	return nil
}

// UInt32Marshaller marshals a wrapped *wrapperspb.UInt32Value value, inserting
// null if the pointer is null, or the inner value if it is not.
type UInt32Marshaller struct {
	*wrapperspb.UInt32Value
}

// Value converts the value to a database-friendly format, or nil if the inner value is
// nil.
func (value UInt32Marshaller) Value() (driver.Value, error) {
	if value.UInt32Value == nil {
		return nil, nil
	}
	// The sql driver only supports int64 so we need to case it.
	return int64(value.UInt32Value.Value), nil
}

// Scan converts a value from the database.
func (value *UInt32Marshaller) Scan(src interface{}) error {
	if src == nil {
		*value = UInt32Marshaller{}
		return nil
	}

	srcVal, ok := src.(int64)
	if !ok {
		return newScanTypeErr(
			int64(0),
			src,
			value.UInt32Value,
		)
	}

	if srcVal > math.MaxUint32 {
		return errors.New("received value overflows uint32")
	} else if srcVal < 0 {
		return errors.New("received value underflows uint32")
	}

	value.UInt32Value = wrapperspb.UInt32(uint32(srcVal))
	return nil
}

// UInt64Marshaller marshals a wrapped *wrapperspb.UInt64Value value, inserting
// null if the pointer is null, or the inner value if it is not.
type UInt64Marshaller struct {
	*wrapperspb.UInt64Value
}

// Value converts the value to a database-friendly format, or nil if the inner value is
// nil.
func (value UInt64Marshaller) Value() (driver.Value, error) {
	if value.UInt64Value == nil {
		return nil, nil
	}

	if value.UInt64Value.Value > math.MaxInt64 {
		return nil, errors.New("cannot encode uint64: overflows int64")
	}

	// The sql driver only supports int64 so we need to case it.
	return int64(value.UInt64Value.Value), nil
}

// Scan converts a value from the database.
func (value *UInt64Marshaller) Scan(src interface{}) error {
	if src == nil {
		*value = UInt64Marshaller{}
		return nil
	}

	srcVal, ok := src.(int64)
	if !ok {
		return newScanTypeErr(
			int64(0),
			src,
			value.UInt64Value,
		)
	}

	if srcVal < 0 {
		return errors.New(
			"received value underflows uint64",
		)
	}

	value.UInt64Value = wrapperspb.UInt64(uint64(srcVal))
	return nil
}

// Bool returns a BoolMarshaller with an inner value set to wrapper.
func Bool(wrapper *wrapperspb.BoolValue) BoolMarshaller {
	return BoolMarshaller{BoolValue: wrapper}
}

// Bytes returns a BytesMarshaller with an inner value set to wrapper.
func Bytes(wrapper *wrapperspb.BytesValue) BytesMarshaller {
	return BytesMarshaller{BytesValue: wrapper}
}

// Double returns a DoubleMarshaller with an inner value set to wrapper.
func Double(wrapper *wrapperspb.DoubleValue) DoubleMarshaller {
	return DoubleMarshaller{DoubleValue: wrapper}
}

// Float returns a FloatMarshaller with an inner value set to wrapper.
func Float(wrapper *wrapperspb.FloatValue) FloatMarshaller {
	return FloatMarshaller{FloatValue: wrapper}
}

// Int32 returns a Int32Marshaller with an inner value set to wrapper.
func Int32(wrapper *wrapperspb.Int32Value) Int32Marshaller {
	return Int32Marshaller{Int32Value: wrapper}
}

// Int64 returns a Int64Marshaller with an inner value set to wrapper.
func Int64(wrapper *wrapperspb.Int64Value) Int64Marshaller {
	return Int64Marshaller{Int64Value: wrapper}
}

// String returns a StringMarshaller with an inner value set to wrapper.
func String(wrapper *wrapperspb.StringValue) StringMarshaller {
	return StringMarshaller{StringValue: wrapper}
}

// UInt32 returns a UInt32Marshaller with an inner value set to wrapper.
func UInt32(wrapper *wrapperspb.UInt32Value) UInt32Marshaller {
	return UInt32Marshaller{UInt32Value: wrapper}
}

// UInt64 returns a UInt64Marshaller with an inner value set to wrapper.
func UInt64(wrapper *wrapperspb.UInt64Value) UInt64Marshaller {
	return UInt64Marshaller{UInt64Value: wrapper}
}
