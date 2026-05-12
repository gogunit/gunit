package httpassert

import (
	"bytes"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/gogunit/gunit/hammy"
)

func Status(resp *http.Response, expected int) hammy.AssertionMessage {
	if resp == nil {
		return hammy.Assert(false, "got nil response, wanted status <%d>", expected)
	}
	return hammy.Assert(resp.StatusCode == expected, "got status <%d>, wanted <%d>", resp.StatusCode, expected)
}

func StatusInRange(resp *http.Response, min, max int) hammy.AssertionMessage {
	if min > max {
		return hammy.Assert(false, "got invalid status range <%d..%d>", min, max)
	}
	if resp == nil {
		return hammy.Assert(false, "got nil response, wanted status in range <%d..%d>", min, max)
	}
	return hammy.Assert(resp.StatusCode >= min && resp.StatusCode <= max, "got status <%d>, wanted in range <%d..%d>", resp.StatusCode, min, max)
}

func Header(resp *http.Response, key, expected string) hammy.AssertionMessage {
	if resp == nil {
		return hammy.Assert(false, "got nil response, wanted header <%s> equal to <%s>", key, expected)
	}
	actual := resp.Header.Get(key)
	return hammy.Assert(actual == expected, "got header <%s>=<%s>, wanted <%s>", key, actual, expected)
}

func HeaderContains(resp *http.Response, key, expected string) hammy.AssertionMessage {
	if resp == nil {
		return hammy.Assert(false, "got nil response, wanted header <%s> containing <%s>", key, expected)
	}
	actual := resp.Header.Get(key)
	return hammy.Assert(strings.Contains(actual, expected), "got header <%s>=<%s>, wanted containing <%s>", key, actual, expected)
}

func BodyEqual(resp *http.Response, expected string) hammy.AssertionMessage {
	actual, result := readBody(resp)
	if !result.IsSuccessful {
		return result
	}
	return hammy.Assert(actual == expected, "got body <%s>, wanted equal to <%s>", actual, expected)
}

func BodyContains(resp *http.Response, expected string) hammy.AssertionMessage {
	actual, result := readBody(resp)
	if !result.IsSuccessful {
		return result
	}
	return hammy.Assert(strings.Contains(actual, expected), "got body <%s>, wanted containing <%s>", actual, expected)
}

func BodyMatchesRegexp(resp *http.Response, pattern string) hammy.AssertionMessage {
	actual, result := readBody(resp)
	if !result.IsSuccessful {
		return result
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return hammy.Assert(false, "invalid regexp <%s>: %v", pattern, err)
	}
	return hammy.Assert(re.MatchString(actual), "got body <%s>, wanted regexp <%s>", actual, pattern)
}

func readBody(resp *http.Response) (string, hammy.AssertionMessage) {
	if resp == nil {
		return "", hammy.Assert(false, "got nil response, wanted response body")
	}
	if resp.Body == nil {
		resp.Body = io.NopCloser(bytes.NewReader(nil))
		return "", hammy.Assert(true, "read body")
	}

	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewReader(body))
	if err != nil {
		return "", hammy.Assert(false, "got body read error: %v", err)
	}
	return string(body), hammy.Assert(true, "read body")
}
