//revive:disable:import-shadowing

package oneof

import (
	"github.com/illuscio-dev/protoCereal-go/cereal_test"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"reflect"
	"testing"
)

func TestMessageInterfaceType(t *testing.T) {
	assert := assert.New(t)

	var wizard interface{} = new(cereal_test.Wizard)

	_, ok := wizard.(proto.Message)
	assert.True(ok, "type assert valid")

	assert.True(
		reflect.TypeOf(wizard).Implements(protoMessageInterface),
		"wizard implements proto message",
	)
}
