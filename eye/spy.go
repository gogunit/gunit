package eye

import (
	"fmt"
	"strings"
	"testing"
)

func Spy() *TestSpy {
	return &TestSpy{}
}

type TestSpy struct {
	HelperCalled bool
	ErrorCalled  bool
	ErrorMessage string
}

func (spy *TestSpy) WasCalled(t *testing.T) {
	t.Helper()
	if !spy.HelperCalled {
		t.Errorf("got ghosted, wanted call to Spy.Helper")
	}
}

func (spy *TestSpy) HadError(t *testing.T) {
	t.Helper()

	spy.WasCalled(t)
	if !spy.ErrorCalled {
		t.Errorf("got ghosted, wanted call to Spy.Errorf")
	}
}

func (spy *TestSpy) HadErrorContaining(t *testing.T, substr string) {
	t.Helper()

	spy.WasCalled(t)
	if !strings.Contains(spy.ErrorMessage, substr) {
		t.Errorf("got:\n%s\nwanted call to spy.Errorf containing:\n%s\n", spy.ErrorMessage, substr)
	}
}

func (spy *TestSpy) Helper() {
	spy.HelperCalled = true
}

func (spy *TestSpy) Errorf(format string, args ...any) {
	spy.ErrorCalled = true
	spy.ErrorMessage = fmt.Sprintf(format, args...)
}
