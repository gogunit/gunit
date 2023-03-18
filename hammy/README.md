# Hammy

[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/gogunit/gunit/hammy)

Hammy is a HamCrest inspired assertion library.
The aim is to provide terse compile-time oriented type checking.

## Getting Started

```go
package adder

import (
	"github.com/gogunit/gunit/hammy"
)

func Test_calculator(t *testing.T) {
	assert := hammy.New(t)
	actual := Add(2, 3)
	assert.Is(hammy.Number(actual).EqualTo(5))
}
```

## Map

* [x] EqualTo
* [x] IsEmpty
* [x] Len
* [x] WithItem
* [x] WithKeys
* [x] WithoutKeys
* [x] WithValues

## Number

* [x] EqualTo
* [x] GreaterThan
* [x] GreaterOrEqual
* [x] IsZero
* [x] LessThan
* [x] LessOrEqual
* [x] Within

## Slice

* [x] Contains
* [x] EqualTo
* [x] Len

## String

* [x] EqualTo
* [x] ToLowerEqualTo
* [x] Contains
* [x] HasPrefix
* [x] HasSuffix
* [x] IsEmpty

## Struct

* [x] EqualTo

## Writing a Custom Matcher

```go
package model

import "github.com/gogunit/gunit/hammy"

func Matcher(model Model) *Matchy {
	return &Matchy{
		model: model,
    }
}

type Matchy struct {
	model Model
}

func (m *Matcher) HasName(expected string) hammy.AssertionMessage {
	actual := m.model.Name
	return hammy.Assert(actual == expected, "want Name=<%v> equal to <%v>", actual, expected)
}
```

```go
assert.Is(Matcher(model).HasName("bar"))

// fails with the message
// want Name=<"foo"> equal to <"bar">
```