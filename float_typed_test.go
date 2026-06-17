package gunit_test

import (
	"github.com/gogunit/gunit"
	"github.com/gogunit/gunit/eye"
	"math"
	"testing"
)

func Test_float_CloseTo_success(t *testing.T) {
	gunit.Float(t, 1.1).CloseTo(1.0, 0.2)
}
func Test_float_CloseTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Float(aSpy, 1.3).CloseTo(1.0, 0.2)
	aSpy.HadErrorContaining(t, "wanted within")
}

func Test_float_IsNaN_success(t *testing.T) {
	gunit.Float(t, math.NaN()).IsNaN()
}
func Test_float_IsNaN_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Float(aSpy, 1.0).IsNaN()
	aSpy.HadErrorContaining(t, "wanted NaN")
}

func Test_float_IsInf_success(t *testing.T) {
	gunit.Float(t, math.Inf(1)).IsInf()
}
func Test_float_IsInf_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Float(aSpy, 1.0).IsInf()
	aSpy.HadErrorContaining(t, "wanted infinity")
}

func Test_float_IsInfSign_success(t *testing.T) {
	gunit.Float(t, math.Inf(-1)).IsInfSign(-1)
}
func Test_float_IsInfSign_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Float(aSpy, math.Inf(1)).IsInfSign(-1)
	aSpy.HadErrorContaining(t, "wanted infinity with sign")
}
