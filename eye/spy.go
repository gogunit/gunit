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
		t.Errorf("want Spy.Helper call, got ghosted")
	}
}

func (spy *TestSpy) HadError(t *testing.T) {
	t.Helper()

	spy.WasCalled(t)

	if !spy.ErrorCalled {
		t.Errorf("want Spy.Errorf call, got ghosted")
	}
}

func (spy *TestSpy) HadErrorContaining(t *testing.T, substr string) {
	t.Helper()

	spy.WasCalled(t)

	if !strings.Contains(spy.ErrorMessage, substr) {
		t.Errorf("want spy.Errorf call containing %v, not found", substr)
	}
}

func (spy *TestSpy) Helper() {
	spy.HelperCalled = true
}

func (spy *TestSpy) Errorf(format string, args ...any) {
	spy.ErrorCalled = true
	spy.ErrorMessage = fmt.Sprintf(format, args...)
}
