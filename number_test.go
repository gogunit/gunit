package gunit_test

import (
	. "github.com/nfisher/gunit"
	. "github.com/nfisher/gunit/testing"
	"testing"
)

func Test_int32_EqualTo(t *testing.T) {
	Number(t, int32(11)).EqualTo(11)
}

func Test_uint_EqualTo_fails(t *testing.T) {
	aSpy := Spy()
	Number(aSpy, uint(123)).EqualTo(234)
	aSpy.HadError(t)
}

func Test_int64_NotEqual(t *testing.T) {
	Number(t, int64(123)).NotEqualTo(234)
}

func Test_uint8_NotEqual_fails(t *testing.T) {
	aSpy := Spy()
	Number(aSpy, uint8(10)).NotEqualTo(10)
	aSpy.HadError(t)
}

func Test_int16_LessThan(t *testing.T) {
	Number(t, int16(10)).LessThan(11)
}

func Test_uint16_LessThan_fails(t *testing.T) {
	aSpy := Spy()
	Number(aSpy, uint16(11)).LessThan(10)
	aSpy.HadError(t)
}

func Test_int_GreaterThan_succeeds(t *testing.T) {
	Number(t, int(11)).GreaterThan(10)
}

func Test_uint32_GreaterThan_fails(t *testing.T) {
	aSpy := Spy()
	Number(aSpy, uint32(9)).GreaterThan(10)
	aSpy.HadError(t)
}

func Test_float64_LessOrEqual_equal_succeeds(t *testing.T) {
	Number(t, float64(10.0)).LessOrEqual(10.0)
}

func Test_float32_LessOrEqual_less_succeeds(t *testing.T) {
	Number(t, float32(9.0)).LessOrEqual(10.0)
}

func Test_uint64_LessOrEqual_greater_fails(t *testing.T) {
	aSpy := Spy()
	Number(aSpy, uint64(11)).LessOrEqual(10)
	aSpy.HadError(t)
}

func Test_float64_GreaterOrEqual_equal_succeeds(t *testing.T) {
	Number(t, float64(11.0)).GreaterOrEqual(11.0)
}

func Test_float32_GreaterOrEqual_greater_succeeds(t *testing.T) {
	Number(t, float32(11.0)).GreaterOrEqual(10.0)
}

func Test_int_GreaterOrEqual_fails(t *testing.T) {
	aSpy := Spy()
	Number(aSpy, int(9)).GreaterOrEqual(10)
	aSpy.HadError(t)
}

func Test_float64_Within_succeeds(t *testing.T) {
	Number(t, float64(11.0)).Within(11.1, 0.1)
}

func Test_float64_Within_over_fails(t *testing.T) {
	aSpy := Spy()
	Number(aSpy, 11.0).Within(10.0, 0.1)
	aSpy.HadError(t)
}

func Test_float64_Within_under_fails(t *testing.T) {
	aSpy := Spy()
	Number(aSpy, 9.0).Within(10.0, 0.1)
	aSpy.HadError(t)
}

func Test_int_IsZero_succeeds(t *testing.T) {
	Number(t, 0).IsZero()
}

func Test_int_IsZero_fails(t *testing.T) {
	aSpy := Spy()
	Number(aSpy, 1).IsZero()
	aSpy.HadError(t)
}

func Test_numeric_subtypes(t *testing.T) {
	td := map[string]struct {
		isAnswer func(t *testing.T)
	}{
		"int":     {isAnswer(Int(42))},
		"int8":    {isAnswer(Int8(42))},
		"int16":   {isAnswer(Int16(42))},
		"int32":   {isAnswer(Int32(42))},
		"int64":   {isAnswer(Int64(42))},
		"uint":    {isAnswer(Uint(42))},
		"uint8":   {isAnswer(Uint8(42))},
		"uint16":  {isAnswer(Uint16(42))},
		"uint32":  {isAnswer(Uint32(42))},
		"uint64":  {isAnswer(Uint64(42))},
		"float32": {isAnswer(Float32(42))},
		"float64": {isAnswer(Float64(42))},
	}
	for n, tc := range td {
		t.Run(n, func(t *testing.T) {
			tc.isAnswer(t)
		})
	}
}

func isAnswer[N Numeric](n N) func(t *testing.T) {
	return func(t *testing.T) { Number(t, n).EqualTo(42) }
}

type (
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Float32 float32
	Float64 float64
)
