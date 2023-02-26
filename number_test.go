package gunit_test

import (
	. "github.com/nfisher/gunit"
	"testing"
)

func TestNumber_subtypes(t *testing.T) {
	type (
		MyInt     int
		MyInt8    int8
		MyInt16   int16
		MyInt32   int32
		MyInt64   int64
		MyUint    uint
		MyUint8   uint8
		MyUint16  uint16
		MyUint32  uint32
		MyUint64  uint64
		MyFloat32 float32
		MyFloat64 float64
	)
	Number(t, MyInt(42)).EqualTo(42)
	Number(t, MyInt8(42)).EqualTo(42)
	Number(t, MyInt16(42)).EqualTo(42)
	Number(t, MyInt32(42)).EqualTo(42)
	Number(t, MyInt64(42)).EqualTo(42)
	Number(t, MyUint(42)).EqualTo(42)
	Number(t, MyUint8(42)).EqualTo(42)
	Number(t, MyUint16(42)).EqualTo(42)
	Number(t, MyUint32(42)).EqualTo(42)
	Number(t, MyUint64(42)).EqualTo(42)
	Number(t, MyFloat32(42)).EqualTo(42)
	Number(t, MyFloat64(42)).EqualTo(42)
}

func Test_int32_EqualTo(t *testing.T) {
	Number(t, int32(11)).EqualTo(11)
}

func Test_uint_EqualTo_fails(t *testing.T) {
	aSpy := spy()
	Number(aSpy, uint(123)).EqualTo(234)
	aSpy.HadError(t)
}

func Test_int64_NotEqual(t *testing.T) {
	Number(t, int64(123)).NotEqualTo(234)
}

func Test_uint8_NotEqual_fails(t *testing.T) {
	aSpy := spy()
	Number(aSpy, uint8(10)).NotEqualTo(10)
	aSpy.HadError(t)
}

func Test_int16_LessThan(t *testing.T) {
	Number(t, int16(10)).LessThan(11)
}

func Test_uint16_LessThan_fails(t *testing.T) {
	aSpy := spy()
	Number(aSpy, uint16(11)).LessThan(10)
	aSpy.HadError(t)
}

func Test_int_GreaterThan_succeeds(t *testing.T) {
	Number(t, int(11)).GreaterThan(10)
}

func Test_uint32_GreaterThan_fails(t *testing.T) {
	aSpy := spy()
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
	aSpy := spy()
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
	aSpy := spy()
	Number(aSpy, int(9)).GreaterOrEqual(10)
	aSpy.HadError(t)
}

func Test_float64_Within_succeeds(t *testing.T) {
	Number(t, float64(11.0)).Within(11.1, 0.1)
}

func Test_float64_Within_over_fails(t *testing.T) {
	aSpy := spy()
	Number(aSpy, 11.0).Within(10.0, 0.1)
	aSpy.HadError(t)
}

func Test_float64_Within_under_fails(t *testing.T) {
	aSpy := spy()
	Number(aSpy, 9.0).Within(10.0, 0.1)
	aSpy.HadError(t)
}
