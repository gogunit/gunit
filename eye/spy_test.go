package eye_test

import (
	"github.com/gogunit/gunit/eye"
	"testing"
)

func Test_Spy_Helper_called(t *testing.T) {
	aSpy := eye.Spy()
	aSpy.Helper()
	if !aSpy.HelperCalled {
		t.Errorf("Spy.Helper not called, should have")
	}
}

func Test_Spy_Errorf_called(t *testing.T) {
	aSpy := eye.Spy()
	aSpy.Errorf("failure")
	if !aSpy.ErrorCalled {
		t.Errorf("Spy.Helper not called, should have")
	}
}
