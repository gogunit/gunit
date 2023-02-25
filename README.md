# gunit

[![Go](https://github.com/nfisher/gunit/actions/workflows/go.yml/badge.svg)](https://github.com/nfisher/gunit/actions/workflows/go.yml)
[![CodeQL](https://github.com/nfisher/gunit/actions/workflows/codeql.yml/badge.svg)](https://github.com/nfisher/gunit/actions/workflows/codeql.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/nfisher/gunit)
[![Go Report Card](https://goreportcard.com/badge/github.com/nfisher/gunit)](https://goreportcard.com/report/github.com/nfisher/gunit)

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
