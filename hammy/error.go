package hammy

import "errors"

func NilError(err error) AssertionMessage {
	return Assert(err == nil, "got <%v>, want nil error", err)
}

func Error(err error) AssertionMessage {
	return Assert(err != nil, "got %v, want error", err)
}

func ErrorIs(err error, target error) AssertionMessage {
	return Assert(errors.Is(err, target), "got <%v>, want error matching <%v>", err, target)
}

func ErrorAs[T any](err error, target *T) AssertionMessage {
	return Assert(errors.As(err, target), "got <%v>, want error assignable to <%T>", err, target)
}
