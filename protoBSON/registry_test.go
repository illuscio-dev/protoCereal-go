package protoBson_test

import (
	"fmt"
	protoBson "github.com/illuscio-dev/protoCereal-go/protoBSON"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
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
