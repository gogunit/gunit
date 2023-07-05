package hammy

func NilError(err error) AssertionMessage {
	return Assert(err == nil, "got <%v>, want nil error", err)
}

func Error(err error) AssertionMessage {
	return Assert(err != nil, "got %v, want error", err)
}
