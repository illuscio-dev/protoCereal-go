package protoBson

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestOpts_WithAddDefaultCodecs(t *testing.T) {
	opts := NewMongoOpts().WithAddDefaultCodecs(false)
	assert.False(t, opts.addDefaultCodecs)

	registryBuilder := bson.NewRegistryBuilder()

	err := RegisterCerealCodecs(registryBuilder, opts)
	if !assert.NoError(t, err, "build registry") {
		t.FailNow()
	}
}
