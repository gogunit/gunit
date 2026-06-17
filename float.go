package gunit

import "math"

type Floaty interface{ ~float32 | ~float64 }

func Float[F Floaty](t T, actual F) *Flt[F] { return &Flt[F]{T: t, actual: actual} }

type Flt[F Floaty] struct {
	T
	actual F
}

func (f *Flt[F]) CloseTo(expected, delta F) {
	f.Helper()
	if math.Abs(float64(f.actual-expected)) > float64(delta) {
		f.Errorf("got <%v>, wanted within <%v> of <%v>", f.actual, delta, expected)
	}
}

func (f *Flt[F]) IsNaN() {
	f.Helper()
	if !math.IsNaN(float64(f.actual)) {
		f.Errorf("got <%v>, wanted NaN", f.actual)
	}
}

func (f *Flt[F]) IsInf() {
	f.Helper()
	if !math.IsInf(float64(f.actual), 0) {
		f.Errorf("got <%v>, wanted infinity", f.actual)
	}
}

func (f *Flt[F]) IsInfSign(sign int) {
	f.Helper()
	if !math.IsInf(float64(f.actual), sign) {
		f.Errorf("got <%v>, wanted infinity with sign <%d>", f.actual, sign)
	}
}
