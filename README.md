# gunit

[![Go](https://github.com/gogunit/gunit/actions/workflows/go.yml/badge.svg)](https://github.com/gogunit/gunit/actions/workflows/go.yml)
[![CodeQL](https://github.com/gogunit/gunit/actions/workflows/codeql.yml/badge.svg)](https://github.com/gogunit/gunit/actions/workflows/codeql.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gogunit/gunit)
[![Go Report Card](https://goreportcard.com/badge/github.com/gogunit/gunit)](https://goreportcard.com/report/github.com/gogunit/gunit)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/gogunit/gunit/)

gounit is test assertions library for Go. It was developed to address the shortcoming of many assertion frameworks that employ assertion of types at runtime rather than compile time.

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

## Hammy Examples

```go
package adder

import (
	"testing"

	a "github.com/gogunit/gunit/hammy"
)

func Test_add_returns_expected_sum(t *testing.T) {
	assert := a.New(t)
	actual := Add(2, 3)
	assert.Is(a.Number(actual).EqualTo(5))
}
```

```go
package service

import (
	"errors"
	"fmt"
	"testing"

	a "github.com/gogunit/gunit/hammy"
)

var errTimeout = errors.New("timeout")

func Test_run_wraps_timeout_error(t *testing.T) {
	assert := a.New(t)
	err := fmt.Errorf("request failed: %w", errTimeout)
	assert.Is(a.ErrorIs(err, errTimeout))
}
```

```go
func Test_people_are_sorted_by_name(t *testing.T) {
	assert := a.New(t)
	people := []Person{{Name: "Ada"}, {Name: "Linus"}}

	assert.Is(a.Slice(people).ContainsInOrder(
		a.HavingField("Name", func(person Person) string { return person.Name }, a.EqualTo("Ada")),
		a.HavingField("Name", func(person Person) string { return person.Name }, a.EqualTo("Linus")),
	))
}
```
