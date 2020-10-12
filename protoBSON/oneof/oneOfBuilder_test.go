package oneof

import (
	"errors"
	"github.com/illuscio-dev/protoCereal-go/cerealMessages"
	"github.com/illuscio-dev/protoCereal-go/cerealMessages_test"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestOneOf_ErrorNotImplemented(t *testing.T) {
	oneOfBuilder := NewCodecBuilder()
	oneOfBuilder.oneOfInterface = reflect.TypeOf(
		(*cerealMessages_test.IsTestOneOfMultiMessageMage)(nil),
	).Elem()

	err := oneOfBuilder.validateValueWrapperType(
		reflect.TypeOf(new(cerealMessages.Decimal)),
	)

	assert.EqualError(
		t,
		err,
		"oneof interface not implemented: "+
			"'*cerealMessages.Decimal' does not implement "+
			"'cerealMessages_test.isTestOneOfMultiMessage_Mage'",
	)

	assert.True(t, errors.Is(err, ErrWrongOneOfInterface))
}
