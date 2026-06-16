package hammy

import "net/url"

func URL(actual *url.URL) *URLAssert {
	if actual == nil {
		return &URLAssert{validation: invalidURL("got nil URL")}
	}
	return &URLAssert{actual: actual}
}

func ParseURL(actual string) *URLAssert {
	parsed, err := url.Parse(actual)
	if err != nil {
		return &URLAssert{validation: invalidURL("invalid URL <%s>: %v", actual, err)}
	}
	return URL(parsed)
}

type URLAssert struct {
	actual     *url.URL
	validation *AssertionMessage
}

func (u *URLAssert) Scheme(expected string) AssertionMessage {
	if msg := u.ready(); !msg.IsSuccessful {
		return msg
	}
	return Assert(u.actual.Scheme == expected, "got URL scheme <%s>, wanted <%s>", u.actual.Scheme, expected)
}

func (u *URLAssert) Host(expected string) AssertionMessage {
	if msg := u.ready(); !msg.IsSuccessful {
		return msg
	}
	return Assert(u.actual.Host == expected, "got URL host <%s>, wanted <%s>", u.actual.Host, expected)
}

func (u *URLAssert) Path(expected string) AssertionMessage {
	if msg := u.ready(); !msg.IsSuccessful {
		return msg
	}
	return Assert(u.actual.Path == expected, "got URL path <%s>, wanted <%s>", u.actual.Path, expected)
}

func (u *URLAssert) RawQuery(expected string) AssertionMessage {
	if msg := u.ready(); !msg.IsSuccessful {
		return msg
	}
	return Assert(u.actual.RawQuery == expected, "got URL raw query <%s>, wanted <%s>", u.actual.RawQuery, expected)
}

func (u *URLAssert) QueryParam(key, expected string) AssertionMessage {
	if msg := u.ready(); !msg.IsSuccessful {
		return msg
	}

	values, ok := u.actual.Query()[key]
	if !ok {
		return Assert(false, "missing URL query parameter <%s>, wanted <%s>", key, expected)
	}

	actual := values[0]
	return Assert(actual == expected, "got URL query parameter <%s> value <%s>, wanted <%s>", key, actual, expected)
}

func (u *URLAssert) NoQueryParam(key string) AssertionMessage {
	if msg := u.ready(); !msg.IsSuccessful {
		return msg
	}

	_, ok := u.actual.Query()[key]
	return Assert(!ok, "got URL query parameter <%s>, wanted it absent", key)
}

func (u *URLAssert) String(expected string) AssertionMessage {
	if msg := u.ready(); !msg.IsSuccessful {
		return msg
	}
	return Assert(u.actual.String() == expected, "got URL string <%s>, wanted <%s>", u.actual.String(), expected)
}

func invalidURL(str string, args ...any) *AssertionMessage {
	msg := Assert(false, str, args...)
	return &msg
}

func (u *URLAssert) ready() AssertionMessage {
	if u.validation != nil {
		return *u.validation
	}
	return Assert(true, "")
}
