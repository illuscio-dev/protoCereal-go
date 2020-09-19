package protoBson_test

import (
	"fmt"
	"github.com/illuscio-dev/protoCereal-go/messagesCereal"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestCodec_Decimal_BsonVal(t *testing.T) {
	assert := assert.New(t)

	type hasDecimal struct {
		Value *messagesCereal.Decimal
	}

	decimalVal, err := primitive.ParseDecimal128("10.25")
	assert.NoError(err, "error parsing decimal")

	high, low := decimalVal.GetBytes()

	original := &hasDecimal{
		Value: &messagesCereal.Decimal{
			High: high,
			Low:  low,
		},
	}

	encoded, err := bson.MarshalWithRegistry(testRegistry, original)
	assert.NoError(err, "encoding err")

	decoded := make(map[string]interface{})
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, decoded)
	assert.NoError(err, "decoding error")

	for key, value := range decoded {
		fmt.Printf("%v: %v\n", key, value)
	}

	assert.IsType(primitive.Decimal128{}, decoded["value"], "decimal")
	binValue, ok := decoded["value"].(primitive.Decimal128)
	if !ok {
		t.FailNow()
	}

	assert.Equal(decimalVal, binValue, "values equal")
}

func TestCodec_Decimal_RoundTrip(t *testing.T) {
	assert := assert.New(t)

	type hasDecimal struct {
		Value *messagesCereal.Decimal
	}

	decimalVal, err := primitive.ParseDecimal128("10.25")
	assert.NoError(err, "error parsing decimal")

	high, low := decimalVal.GetBytes()

	original := &hasDecimal{
		Value: &messagesCereal.Decimal{
			High: high,
			Low:  low,
		},
	}

	encoded, err := bson.MarshalWithRegistry(testRegistry, original)
	assert.NoError(err, "encoding err")

	decoded := new(hasDecimal)
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, decoded)
	assert.NoError(err, "decoding error")

	assert.Equal(original.Value, decoded.Value, "decimal val")
}

func TestCodec_Decimal_RoundTrip_Null(t *testing.T) {
	assert := assert.New(t)

	type hasDecimal struct {
		Value *messagesCereal.Decimal
	}

	original := &hasDecimal{
		Value: nil,
	}

	encoded, err := bson.MarshalWithRegistry(testRegistry, original)
	assert.NoError(err, "encoding err")

	decoded := new(hasDecimal)
	err = bson.UnmarshalWithRegistry(testRegistry, encoded, decoded)
	assert.NoError(err, "decoding error")

	assert.Nil(original.Value, "uuid nil")
}
