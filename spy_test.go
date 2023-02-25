package gunit_test

import "testing"

func spy() *testSpy {
	return &testSpy{}
}

type testSpy struct {
	helperCalled bool
	errorCalled  bool
}

func (t *testSpy) Helper() {
	t.helperCalled = true
}

func (t *testSpy) Errorf(_ string, _ ...any) {
	t.errorCalled = true
}

func Test_Spy_Helper_called(t *testing.T) {
	aSpy := spy()
	aSpy.Helper()
	if !aSpy.helperCalled {
		t.Errorf("spy.Helper not called, should have")
	}
}

func Test_Spy_Errorf_called(t *testing.T) {
	aSpy := spy()
	aSpy.Errorf("failure")
	if !aSpy.errorCalled {
		t.Errorf("spy.Helper not called, should have")
	}
}
