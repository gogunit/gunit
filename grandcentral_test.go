package gunit_test

import (
	. "github.com/nfisher/gunit"
	. "github.com/nfisher/gunit/testing"
	"testing"
)

func Test_Int(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.Int(-64).EqualTo(-64)
	aSpy.WasCalled(t)
}

func Test_Int8(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.Int8(-64).EqualTo(-64)
	aSpy.WasCalled(t)
}

func Test_Int16(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.Int16(-64).EqualTo(-64)
	aSpy.WasCalled(t)
}

func Test_Int32(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.Int32(-64).EqualTo(-64)
	aSpy.WasCalled(t)
}

func Test_Int64(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.Int64(-64).EqualTo(-64)
	aSpy.WasCalled(t)
}

func Test_Unit(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.Uint(1).EqualTo(1)
	aSpy.WasCalled(t)
}

func Test_Unit8(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.Uint8(1).EqualTo(1)
	aSpy.WasCalled(t)
}

func Test_Unit16(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.Uint16(1).EqualTo(1)
	aSpy.WasCalled(t)
}

func Test_Unit32(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.Uint32(1).EqualTo(1)
	aSpy.WasCalled(t)
}

func Test_Unit64(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.Uint64(1).EqualTo(1)
	aSpy.WasCalled(t)
}

func Test_Float32(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.Float32(11.0).EqualTo(11.0)
	aSpy.WasCalled(t)
}

func Test_Float64(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.Float64(11.0).EqualTo(11.0)
	aSpy.WasCalled(t)
}
