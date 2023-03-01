package testing

import "testing"

func Spy() *TestSpy {
	return &TestSpy{}
}

type TestSpy struct {
	HelperCalled bool
	ErrorCalled  bool
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

func (spy *TestSpy) Helper() {
	spy.HelperCalled = true
}

func (spy *TestSpy) Errorf(_ string, _ ...any) {
	spy.ErrorCalled = true
}
