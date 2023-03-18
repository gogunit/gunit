# gunit

[![Go](https://github.com/gogunit/gunit/actions/workflows/go.yml/badge.svg)](https://github.com/gogunit/gunit/actions/workflows/go.yml)
[![CodeQL](https://github.com/gogunit/gunit/actions/workflows/codeql.yml/badge.svg)](https://github.com/gogunit/gunit/actions/workflows/codeql.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gogunit/gunit)
[![Go Report Card](https://goreportcard.com/badge/github.com/gogunit/gunit)](https://goreportcard.com/report/github.com/gogunit/gunit)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/gogunit/gunit/)

Go unit test assertions library.

## Examples

```go
// direct assertion style
func Test_nine_plus_two_is_greater_than_ten(t *testing.T) {
	actual := 9 + 2
	expected := 10
	gunit.Number(t, actual).GreaterThan(expected)
}

// wrap testing.T struct
func Test_nine_plus_two_is_greater_than_ten(t *testing.T) {
	assert := gunit.New(t)
	actual := 9 + 2
	expected := 10
	assert.Int(actual).GreaterThan(expected)
}

```
