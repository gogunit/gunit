package hammy

import (
	"reflect"

	"github.com/gogunit/gunit"
)

func New(t gunit.T) *Hammy {
	return &Hammy{t}
}

type Hammy struct {
	gunit.T
}

func (h *Hammy) Is(a AssertionMessage) {
	h.Helper()
	if !a.IsSuccessful {
		h.Errorf(a.Message)
	}
}

func (h *Hammy) IsNot(a AssertionMessage) {
	h.Helper()
	if a.IsSuccessful {
		h.Errorf("not(" + a.Message + ")")
	}
}

func (h *Hammy) That(msg string, a AssertionMessage) {
	h.Helper()
	if !a.IsSuccessful {
		h.Errorf("%s: %s", msg, a.Message)
	}
}

func Nil(actual any) AssertionMessage {
	return Assert(isNil(actual), "got <%T>, wanted nil", actual)
}

func NotNil(actual any) AssertionMessage {
	return Assert(!isNil(actual), "got nil, wanted <%T>", actual)
}

func True(actual bool) AssertionMessage {
	return Assert(actual, "got false, wanted true")
}

func False(actual bool) AssertionMessage {
	return Assert(!actual, "got true, wanted false")
}

func isNil(actual any) bool {
	if actual == nil {
		return true
	}

	value := reflect.ValueOf(actual)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return value.IsNil()
	default:
		return false
	}
}
