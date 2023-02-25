package gunit

type test interface {
	Helper()
	Errorf(format string, args ...any)
}
