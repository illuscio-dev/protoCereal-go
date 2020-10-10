package oneof

import (
	"errors"
	"github.com/illuscio-dev/protoCereal-go/messagesCereal"
	"github.com/illuscio-dev/protoCereal-go/messagesCereal_test"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestOneOf_ErrorNotImplemented(t *testing.T) {
	oneOfBuilder := NewCodecBuilder()
	oneOfBuilder.oneOfInterface = reflect.TypeOf(
		(*messagesCereal_test.IsTestOneOfMultiMessageMage)(nil),
	).Elem()

	err := oneOfBuilder.validateValueWrapperType(
		reflect.TypeOf(new(messagesCereal.Decimal)),
	)

	assert.EqualError(
		t,
		err,
		"oneof interface not implemented: "+
			"'*messagesCereal.Decimal' does not implement "+
			"'messagesCereal_test.isTestOneOfMultiMessage_Mage'",
	)

	assert.True(t, errors.Is(err, ErrWrongOneOfInterface))
}
