package hammy

import "math"

func Float[F Floaty](actual F) *Flt[F] {
	return &Flt[F]{actual: actual}
}

type Flt[F Floaty] struct {
	actual F
}

func (f *Flt[F]) CloseTo(expected, delta F) AssertionMessage {
	return CloseTo(expected, delta).Match(f.actual)
}

func (f *Flt[F]) IsNaN() AssertionMessage {
	return IsNaN[F]().Match(f.actual)
}

func (f *Flt[F]) IsInf() AssertionMessage {
	return IsInf[F]().Match(f.actual)
}

func (f *Flt[F]) IsInfSign(sign int) AssertionMessage {
	return IsInfSign[F](sign).Match(f.actual)
}

func CloseTo[F Floaty](expected, delta F) Matcher[F] {
	return MatchFunc(func(actual F) AssertionMessage {
		diff := math.Abs(float64(actual - expected))
		return Assert(diff <= float64(delta), "got <%v>, wanted within <%v> of <%v>", actual, delta, expected)
	})
}

func IsNaN[F Floaty]() Matcher[F] {
	return MatchFunc(func(actual F) AssertionMessage {
		return Assert(math.IsNaN(float64(actual)), "got <%v>, wanted NaN", actual)
	})
}

func IsInf[F Floaty]() Matcher[F] {
	return MatchFunc(func(actual F) AssertionMessage {
		return Assert(math.IsInf(float64(actual), 0), "got <%v>, wanted infinity", actual)
	})
}

func IsInfSign[F Floaty](sign int) Matcher[F] {
	return MatchFunc(func(actual F) AssertionMessage {
		return Assert(math.IsInf(float64(actual), sign), "got <%v>, wanted infinity with sign <%d>", actual, sign)
	})
}
