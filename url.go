package gunit

import "net/url"

func URL(t T, actual *url.URL) *URLAssert { return &URLAssert{T: t, actual: actual} }
func ParseURL(t T, actual string) *URLAssert {
	parsed, err := url.Parse(actual)
	if err != nil {
		return &URLAssert{T: t, validation: err, raw: actual}
	}
	return URL(t, parsed)
}

type URLAssert struct {
	T
	actual     *url.URL
	validation error
	raw        string
}

func (u *URLAssert) Scheme(expected string) {
	if !u.ready() {
		return
	}
	if u.actual.Scheme != expected {
		u.Errorf("got URL scheme <%s>, wanted <%s>", u.actual.Scheme, expected)
	}
}

func (u *URLAssert) Host(expected string) {
	if !u.ready() {
		return
	}
	if u.actual.Host != expected {
		u.Errorf("got URL host <%s>, wanted <%s>", u.actual.Host, expected)
	}
}

func (u *URLAssert) Path(expected string) {
	if !u.ready() {
		return
	}
	if u.actual.Path != expected {
		u.Errorf("got URL path <%s>, wanted <%s>", u.actual.Path, expected)
	}
}

func (u *URLAssert) RawQuery(expected string) {
	if !u.ready() {
		return
	}
	if u.actual.RawQuery != expected {
		u.Errorf("got URL raw query <%s>, wanted <%s>", u.actual.RawQuery, expected)
	}
}

func (u *URLAssert) QueryParam(key, expected string) {
	if !u.ready() {
		return
	}
	values, ok := u.actual.Query()[key]
	if !ok {
		u.Errorf("missing URL query parameter <%s>, wanted <%s>", key, expected)
		return
	}
	if values[0] != expected {
		u.Errorf("got URL query parameter <%s> value <%s>, wanted <%s>", key, values[0], expected)
	}
}

func (u *URLAssert) NoQueryParam(key string) {
	if !u.ready() {
		return
	}
	if _, ok := u.actual.Query()[key]; ok {
		u.Errorf("got URL query parameter <%s>, wanted it absent", key)
	}
}

func (u *URLAssert) String(expected string) {
	if !u.ready() {
		return
	}
	if u.actual.String() != expected {
		u.Errorf("got URL string <%s>, wanted <%s>", u.actual.String(), expected)
	}
}

func (u *URLAssert) ready() bool {
	u.Helper()
	if u.validation != nil {
		u.Errorf("invalid URL <%s>: %v", u.raw, u.validation)
		return false
	}
	if u.actual == nil {
		u.Errorf("got nil URL")
		return false
	}
	return true
}
