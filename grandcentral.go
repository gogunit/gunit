package gunit

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
