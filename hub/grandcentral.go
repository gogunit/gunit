package hub

import "github.com/gogunit/gunit"

func New(t gunit.T) *GrandCentral {
	return &GrandCentral{t: t}
}

type GrandCentral struct{ t gunit.T }

func (c *GrandCentral) Int(actual int) *gunit.Num[int] {
	return gunit.Number(c.t, actual)
}

func (c *GrandCentral) Int8(actual int8) *gunit.Num[int8] {
	return gunit.Number(c.t, actual)
}

func (c *GrandCentral) Int16(actual int16) *gunit.Num[int16] {
	return gunit.Number(c.t, actual)
}

func (c *GrandCentral) Int32(actual int32) *gunit.Num[int32] {
	return gunit.Number(c.t, actual)
}

func (c *GrandCentral) Int64(actual int64) *gunit.Num[int64] {
	return gunit.Number(c.t, actual)
}

func (c *GrandCentral) Float32(actual float32) *gunit.Num[float32] {
	return gunit.Number(c.t, actual)
}

func (c *GrandCentral) Float64(actual float64) *gunit.Num[float64] {
	return gunit.Number(c.t, actual)
}

func (c *GrandCentral) Uint(actual uint) *gunit.Num[uint] {
	return gunit.Number(c.t, actual)
}

func (c *GrandCentral) Uint8(actual uint8) *gunit.Num[uint8] {
	return gunit.Number(c.t, actual)
}

func (c *GrandCentral) Uint16(actual uint16) *gunit.Num[uint16] {
	return gunit.Number(c.t, actual)
}

func (c *GrandCentral) Uint32(actual uint32) *gunit.Num[uint32] {
	return gunit.Number(c.t, actual)
}

func (c *GrandCentral) Uint64(actual uint64) *gunit.Num[uint64] {
	return gunit.Number(c.t, actual)
}
