package hammy_test

import (
	"math"
	"testing"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

func Test_Float_CloseTo_success(t *testing.T) {
	a.New(t).Is(a.Float(10.0).CloseTo(10.1, 0.2))
}

func Test_Float_CloseTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Float(10.0).CloseTo(10.5, 0.2))
	aSpy.HadErrorContaining(t, "wanted within <0.2> of <10.5>")
}

func Test_Float_IsNaN_success(t *testing.T) {
	a.New(t).Is(a.Float(math.NaN()).IsNaN())
}

func Test_Float_IsInf_success(t *testing.T) {
	a.New(t).Is(a.Float(math.Inf(1)).IsInf())
}

func Test_Float_IsInfSign_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Float(math.Inf(-1)).IsInfSign(1))
	aSpy.HadErrorContaining(t, "wanted infinity with sign <1>")
}
