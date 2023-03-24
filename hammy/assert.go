package hammy

import "github.com/gogunit/gunit"

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

func (h *Hammy) That(msg string, a AssertionMessage) {
	h.Helper()
	if !a.IsSuccessful {
		h.Errorf("%s: %s", msg, a.Message)
	}
}

func Nil[T any](actual *T) AssertionMessage {
	return Assert(actual == nil, "got <%T>, wanted nil", actual)
}

func NotNil[T any](actual *T) AssertionMessage {
	return Assert(actual != nil, "got nil, wanted <%T>", actual)
}

func True(actual bool) AssertionMessage {
	return Assert(actual, "got false, wanted true")
}

func False(actual bool) AssertionMessage {
	return Assert(!actual, "got true, wanted false")
}
