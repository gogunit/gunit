package gunit

type T interface {
	Helper()
	Errorf(format string, args ...any)
}
