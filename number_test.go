package gunit_test

import (
	. "github.com/nfisher/gunit"
	"testing"
)

func Test_int32_EqualTo(t *testing.T) {
	Number(t, int32(11)).EqualTo(11)
}

func Test_uint_EqualTo_fails(t *testing.T) {
	aSpy := spy()
	Number(aSpy, uint(123)).EqualTo(234)
	if !aSpy.errorCalled {
		t.Errorf("spy.Errorf not called, should have")
	}
}

func Test_int64_NotEqual(t *testing.T) {
	Number(t, int64(123)).NotEqualTo(234)
}

func Test_uint8_NotEqual_fails(t *testing.T) {
	aSpy := spy()
	Number(aSpy, uint8(10)).NotEqualTo(10)
	if !aSpy.errorCalled {
		t.Errorf("spy.Errorf not called, should have")
	}
}

func Test_int16_LessThan(t *testing.T) {
	Number(t, int16(10)).LessThan(11)
}

func Test_uint16_LessThan_fails(t *testing.T) {
	aSpy := spy()
	Number(aSpy, uint16(11)).LessThan(10)
	if !aSpy.errorCalled {
		t.Errorf("spy.Errorf not called, should have")
	}
}

func Test_int_GreaterThan_succeeds(t *testing.T) {
	Number(t, int(11)).GreaterThan(10)
}

func Test_uint32_GreaterThan_fails(t *testing.T) {
	aSpy := spy()
	Number(aSpy, uint32(9)).GreaterThan(10)
	if !aSpy.errorCalled {
		t.Errorf("spy.Errorf not called, should have")
	}
}

func Test_float64_LessOrEqual_equal_succeeds(t *testing.T) {
	Number(t, float64(10.0)).LessOrEqual(10.0)
}

func Test_float32_LessOrEqual_less_succeeds(t *testing.T) {
	Number(t, float32(9.0)).LessOrEqual(10.0)
}

func Test_uint64_LessOrEqual_greater_fails(t *testing.T) {
	aSpy := spy()
	Number(aSpy, uint64(11)).LessOrEqual(10)
	if !aSpy.errorCalled {
		t.Errorf("spy.Errorf not called, should have")
	}
}

func Test_float64_GreaterOrEqual_equal_succeeds(t *testing.T) {
	Number(t, float64(11.0)).GreaterOrEqual(11.0)
}

func Test_float32_GreaterOrEqual_greater_succeeds(t *testing.T) {
	Number(t, float32(11.0)).GreaterOrEqual(10.0)
}

func Test_int_GreaterOrEqual_fails(t *testing.T) {
	aSpy := spy()
	Number(aSpy, int(9)).GreaterOrEqual(10)
	if !aSpy.errorCalled {
		t.Errorf("spy.Errorf not called, should have")
	}
}

func Test_float64_Within_succeeds(t *testing.T) {
	Number(t, float64(11.0)).Within(11.1, 0.1)
}

func Test_float64_Within_over_fails(t *testing.T) {
	aSpy := spy()
	Number(aSpy, 11.0).Within(10.0, 0.1)
	if !aSpy.errorCalled {
		t.Errorf("spy.Errorf not called, should have")
	}
}

func Test_float64_Within_under_fails(t *testing.T) {
	aSpy := spy()
	Number(aSpy, 9.0).Within(10.0, 0.1)
	if !aSpy.errorCalled {
		t.Errorf("spy.Errorf not called, should have")
	}
}
