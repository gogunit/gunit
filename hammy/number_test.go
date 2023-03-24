package hammy_test

import (
	"testing"

	"github.com/gogunit/gunit/eye"
	ham "github.com/gogunit/gunit/hammy"
)

func Test_int_EqualTo_success(t *testing.T) {
	assert := ham.New(t)
	assert.Is(ham.Number(123).EqualTo(123))
}

func Test_int8_EqualTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	ham.New(aSpy).Is(ham.Number(int8(123)).EqualTo(124))
	aSpy.HadError(t)
}

func Test_int16_NotEqual_success(t *testing.T) {
	ham.New(t).Is(ham.Number(uint16(42)).NotEqual(41))
}

func Test_int32_NotEqual_failure(t *testing.T) {
	aSpy := eye.Spy()
	ham.New(aSpy).Is(ham.Number(uint(42)).NotEqual(42))
	aSpy.HadError(t)
}

func Test_int_that_NotEqual_failure(t *testing.T) {
	aSpy := eye.Spy()
	ham.New(aSpy).That("answer to everything", ham.Number(42).NotEqual(42))
	aSpy.HadErrorContaining(t, "answer to everything")
}

func Test_int64_LessThan_success(t *testing.T) {
	ham.New(t).Is(ham.Number(int64(10)).LessThan(11))
}

func Test_uint_LessThan_failure(t *testing.T) {
	aSpy := eye.Spy()
	ham.New(aSpy).Is(ham.Number(uint(11)).LessThan(10))
	aSpy.HadError(t)
}

func Test_uint8_GreaterThan_success(t *testing.T) {
	ham.New(t).Is(ham.Number(uint8(8)).GreaterThan(7))
}

func Test_uint16_GreaterThan_failure(t *testing.T) {
	aSpy := eye.Spy()
	ham.New(aSpy).Is(ham.Number(uint16(9)).GreaterThan(10))
	aSpy.HadError(t)
}

func Test_uint32_LessOrEqual_success(t *testing.T) {
	ham.New(t).Is(ham.Number(uint32(10)).LessOrEqual(10))
}

func Test_uint64_LessOrEqual_failure(t *testing.T) {
	aSpy := eye.Spy()
	ham.New(aSpy).Is(ham.Number(uint64(11)).LessOrEqual(10))
	aSpy.HadError(t)
}

func Test_float32_GreaterOrEqual_success(t *testing.T) {
	ham.New(t).Is(ham.Number(uint32(10)).GreaterOrEqual(10))
}

func Test_float64_GreaterOrEqual_failure(t *testing.T) {
	aSpy := eye.Spy()
	ham.New(aSpy).Is(ham.Number(float64(10)).GreaterOrEqual(11))
	aSpy.HadError(t)
}

func Test_int_IsZero_success(t *testing.T) {
	ham.New(t).Is(ham.Number(0).IsZero())
}

func Test_int_IsZero_failure(t *testing.T) {
	aSpy := eye.Spy()
	ham.New(aSpy).Is(ham.Number(1).IsZero())
	aSpy.HadError(t)
}

func Test_float64_Within_success(t *testing.T) {
	ham.New(t).Is(ham.Number(float64(123)).Within(123.1, 0.2))
}

func Test_float32_Within_failure(t *testing.T) {
	aSpy := eye.Spy()
	ham.New(aSpy).Is(ham.Number(123).Within(140, 0.2))
	aSpy.HadError(t)
}
