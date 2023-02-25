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

func New(t T) *GrandCentral {
	return &GrandCentral{t: t}
}

type GrandCentral struct{ t T }

func (c *GrandCentral) Int(actual int) *Num[int] {
	return Number(c.t, actual)
}

func (c *GrandCentral) Uint(actual uint) *Num[uint] {
	return Number(c.t, actual)
}

func (c *GrandCentral) Float64(actual float64) *Num[float64] {
	return Number(c.t, actual)
}
