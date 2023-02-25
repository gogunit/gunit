package gunit_test

import (
	. "github.com/nfisher/gunit"
	"testing"
)

func Test_Int(t *testing.T) {
	aSpy := spy()
	assert := New(aSpy)
	assert.Int(-64).EqualTo(-64)
	aSpy.WasCalled(t)
}

func Test_Unit(t *testing.T) {
	aSpy := spy()
	assert := New(aSpy)
	assert.Uint(1).EqualTo(1)
}

func Test_Float64(t *testing.T) {
	aSpy := spy()
	assert := New(aSpy)
	assert.Float64(11.0).EqualTo(11.0)
	aSpy.WasCalled(t)
}
