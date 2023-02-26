package gunit

func New(t T) *GrandCentral {
	return &GrandCentral{t: t}
}

type GrandCentral struct{ t T }

func (c *GrandCentral) Int(actual int) *Num[int] {
	return Number(c.t, actual)
}

func (c *GrandCentral) Int8(actual int8) *Num[int8] {
	return Number(c.t, actual)
}

func (c *GrandCentral) Int16(actual int16) *Num[int16] {
	return Number(c.t, actual)
}

func (c *GrandCentral) Int32(actual int32) *Num[int32] {
	return Number(c.t, actual)
}

func (c *GrandCentral) Int64(actual int64) *Num[int64] {
	return Number(c.t, actual)
}

func (c *GrandCentral) Float32(actual float32) *Num[float32] {
	return Number(c.t, actual)
}

func (c *GrandCentral) Float64(actual float64) *Num[float64] {
	return Number(c.t, actual)
}

func (c *GrandCentral) Uint(actual uint) *Num[uint] {
	return Number(c.t, actual)
}

func (c *GrandCentral) Uint8(actual uint8) *Num[uint8] {
	return Number(c.t, actual)
}

func (c *GrandCentral) Uint16(actual uint16) *Num[uint16] {
	return Number(c.t, actual)
}

func (c *GrandCentral) Uint32(actual uint32) *Num[uint32] {
	return Number(c.t, actual)
}

func (c *GrandCentral) Uint64(actual uint64) *Num[uint64] {
	return Number(c.t, actual)
}
