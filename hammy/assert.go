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
