package oneof

import (
	"errors"
	"github.com/illuscio-dev/protoCereal-go/cereal"
	"github.com/illuscio-dev/protoCereal-go/cereal_test"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestOneOf_ErrorNotImplemented(t *testing.T) {
	oneOfBuilder := NewCodecBuilder()
	oneOfBuilder.oneOfInterface = reflect.TypeOf(
		(*cereal_test.IsTestOneOfMultiMessageMage)(nil),
	).Elem()

	err := oneOfBuilder.validateValueWrapperType(
		reflect.TypeOf(new(cereal.Decimal)),
	)

	assert.EqualError(
		t,
		err,
		"oneof interface not implemented: "+
			"'*cereal.Decimal' does not implement "+
			"'cereal_test.isTestOneOfMultiMessage_Mage'",
	)

	assert.True(t, errors.Is(err, ErrWrongOneOfInterface))
}
