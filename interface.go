package gunit

type T interface {
	Helper()
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
}
