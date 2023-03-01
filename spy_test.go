package gunit_test

import "testing"
import . "github.com/nfisher/gunit"

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
