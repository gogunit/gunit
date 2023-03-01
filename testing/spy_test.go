package testing_test

import (
	. "github.com/nfisher/gunit/testing"
	"testing"
)

func Test_Spy_Helper_called(t *testing.T) {
	aSpy := Spy()
	aSpy.Helper()
	if !aSpy.HelperCalled {
		t.Errorf("Spy.Helper not called, should have")
	}
}

func Test_Spy_Errorf_called(t *testing.T) {
	aSpy := Spy()
	aSpy.Errorf("failure")
	if !aSpy.ErrorCalled {
		t.Errorf("Spy.Helper not called, should have")
	}
}
