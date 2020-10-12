package oneof

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/illuscio-dev/protoCereal-go/cerealMessages"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
)

var decimalType = reflect.TypeOf(new(cerealMessages.Decimal))
var timestampType = reflect.TypeOf(new(timestamppb.Timestamp))
var uuidType = reflect.TypeOf(new(cerealMessages.UUID))
var rawDataType = reflect.TypeOf(new(cerealMessages.RawData))
var wrapperBoolType = reflect.TypeOf(new(wrappers.BoolValue))
var wrapperBytesType = reflect.TypeOf(new(wrappers.BytesValue))
var wrapperDoubleType = reflect.TypeOf(new(wrappers.DoubleValue))
var wrapperFloatType = reflect.TypeOf(new(wrappers.FloatValue))
var wrapperInt32Type = reflect.TypeOf(new(wrappers.Int32Value))
var wrapperInt64Type = reflect.TypeOf(new(wrappers.Int64Value))
var wrapperStringType = reflect.TypeOf(new(wrappers.StringValue))
var wrapperUInt32Type = reflect.TypeOf(new(wrappers.UInt32Value))
var wrapperUInt64Type = reflect.TypeOf(new(wrappers.UInt64Value))
var protoMessageInterface = reflect.TypeOf((*proto.Message)(nil)).Elem()
var bytesFieldType = reflect.TypeOf(make([]uint8, 0))
