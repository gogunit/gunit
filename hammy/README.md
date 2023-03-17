# Hammy

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

* [ ] EqualTo
* [ ] IsEmpty
* [ ] WithKeys
* [ ] WithoutKeys
* [ ] WithValues

## Number

* [ ] EqualTo
* [ ] GreaterThan
* [ ] GreaterThanOrEqualTo
* [ ] IsZero
* [ ] Len
* [ ] LessThan
* [ ] LessThanOrEqualTo
* [ ] WithIn

## Slice

* [ ] Contains
* [ ] EqualTo
* [ ] Len

## String

* [ ] EqualTo
* [ ] LowerCaseEqualTo
* [ ] Contains
* [ ] HasPrefix
* [ ] HasSuffix
* [ ] IsEmpty

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