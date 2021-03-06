package protobson_test

import (
	"fmt"
	"github.com/illuscio-dev/protoCereal-go/cereal_test"
	protoBson "github.com/illuscio-dev/protoCereal-go/protobson"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"google.golang.org/protobuf/reflect/protoreflect"
	"testing"
)

var testRegistry *bsoncodec.Registry

func init() {
	testRegistryBuilder := bsoncodec.NewRegistryBuilder()

	err := protoBson.RegisterCerealCodecs(testRegistryBuilder, nil)
	if err != nil {
		panic(fmt.Errorf("error building test registry: %w", err))
	}

	testRegistry = testRegistryBuilder.Build()
}

func TestOpts_WithCustomWrappers(t *testing.T) {
	assert := assert.New(t)
	type hasWrapper struct {
		Info *cereal_test.ListValue
	}

	registryBuilder := bsoncodec.NewRegistryBuilder()
	opts := protoBson.NewMongoOpts().
		WithCustomWrappers(new(cereal_test.ListValue))

	err := protoBson.RegisterCerealCodecs(registryBuilder, opts)
	if !assert.NoError(err, "create registry") {
		t.FailNow()
	}

	registry := registryBuilder.Build()

	innerValue := []string{"some", "values"}
	original := &hasWrapper{
		Info: &cereal_test.ListValue{
			Value: innerValue,
		},
	}

	serialized, err := bson.MarshalWithRegistry(registry, original)
	if !assert.NoError(err, "marshal message") {
		t.FailNow()
	}

	// Unmarshall into a document and check that the inner field of our custom
	// wrapper type was extracted.
	document := bson.M{}
	err = bson.UnmarshalWithRegistry(registry, serialized, &document)
	if !assert.NoError(err, "unmarshall into document") {
		t.FailNow()
	}

	if !assert.Contains(document, "info", "document key exists") {
		t.FailNow()
	}

	docValue := bson.A{"some", "values"}
	if !assert.Equal(docValue, document["info"], "key value") {
		t.FailNow()
	}

	unmarshalled := new(hasWrapper)
	err = bson.UnmarshalWithRegistry(registry, serialized, unmarshalled)
	if !assert.NoError(err, "unmarshall to struct") {
		t.FailNow()
	}

	assert.Equal(original, unmarshalled, "message match")
}

func TestRegisterCustomWrapper_WithMultiplePublic(t *testing.T) {
	registryBuilder := bsoncodec.NewRegistryBuilder()
	opts := protoBson.NewMongoOpts().
		WithCustomWrappers(new(cereal_test.TestProto))

	err := protoBson.RegisterCerealCodecs(registryBuilder, opts)
	assert.EqualError(
		t,
		err,
		"error creating custom wrapper codec: wrapper expected to have"+
			" exactly 1 public field, found 2 public fields for type"+
			" '*cereal_test.TestProto'",
	)
}

type badMessage struct{}

func (m badMessage) ProtoReflect() protoreflect.Message {
	return (protoreflect.Message)(nil)
}

func TestRegisterCustomWrapper_NonStructPointerWrapper(t *testing.T) {

	registryBuilder := bsoncodec.NewRegistryBuilder()
	opts := protoBson.NewMongoOpts().
		WithCustomWrappers(badMessage{})

	err := protoBson.RegisterCerealCodecs(registryBuilder, opts)
	assert.EqualError(
		t,
		err,
		"error creating custom wrapper codec: wrapper codec expected"+
			" pointer to struct, got 'protobson_test.badMessage'",
	)
}
