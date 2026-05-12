package hammy

import (
	"errors"

	"github.com/google/go-cmp/cmp"
)

func Panics(fn func()) AssertionMessage {
	value, panicked := capturePanic(fn)
	return Assert(panicked, "got no panic, wanted panic; recovered value <%v>", value)
}

func NotPanics(fn func()) AssertionMessage {
	value, panicked := capturePanic(fn)
	return Assert(!panicked, "got panic <%v>, wanted no panic", value)
}

func PanicsWithValue(expected any, fn func()) AssertionMessage {
	value, panicked := capturePanic(fn)
	if !panicked {
		return Assert(false, "got no panic, wanted panic value <%v>", expected)
	}
	return Assert(cmp.Equal(value, expected), "got panic value <%v>, wanted <%v>", value, expected)
}

func PanicsWithError(expected string, fn func()) AssertionMessage {
	value, panicked := capturePanic(fn)
	if !panicked {
		return Assert(false, "got no panic, wanted panic error <%s>", expected)
	}

	err, ok := value.(error)
	if !ok {
		return Assert(false, "got panic value <%v>, wanted error panic <%s>", value, expected)
	}
	return Assert(err.Error() == expected, "got panic error <%v>, wanted <%s>", err, expected)
}

func PanicErrorIs(target error, fn func()) AssertionMessage {
	value, panicked := capturePanic(fn)
	if !panicked {
		return Assert(false, "got no panic, wanted panic error matching <%v>", target)
	}

	err, ok := value.(error)
	if !ok {
		return Assert(false, "got panic value <%v>, wanted error matching <%v>", value, target)
	}
	return Assert(errors.Is(err, target), "got panic error <%v>, wanted matching <%v>", err, target)
}

func capturePanic(fn func()) (value any, panicked bool) {
	panicked = true
	defer func() {
		value = recover()
	}()
	fn()
	panicked = false
	return nil, false
}
