# Hammy

[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/gogunit/gunit/hammy)

Hammy is a HamCrest inspired assertion library.
The aim is to provide terse compile-time oriented type checking.

## Getting Started

```go
package adder

import (
	"testing"

	a "github.com/gogunit/gunit/hammy"
)

func Test_calculator(t *testing.T) {
	assert := a.New(t)
	actual := Add(2, 3)
	assert.Is(a.Number(actual).EqualTo(5))
}
```

## Preferred Style

Prefer `Number`, `String`, `Slice`, `Map`, `Struct`, and `Float` for direct assertions and composed matcher checks:

```go
func Test_add_returns_small_positive_sum(t *testing.T) {
	assert := a.New(t)
	actual := Add(2, 3)

	assert.Is(a.Number(actual).Matches(a.AllOf(
		a.GreaterThan(0),
		a.LessThan(10),
	)))
}
```

Use `Match` when no typed wrapper fits or when the value is intentionally held as `any`.

## Dedicated Packages

Use `htmlassert` for semantic DOM equality, structural fragments, CSS selectors, normalized text, and element attributes; use `httpassert` for assertions on `*http.Response` values. HTTP, JSON, and YAML assertions all use the constructor-style API: wrap the actual value once, then call assertion methods on the returned struct.

```go
import ha "github.com/gogunit/gunit/hammy/httpassert"

response := ha.Response(resp)
assert.Is(response.HasStatus(http.StatusOK))
assert.Is(response.HeaderEqualTo("Content-Type", "application/json"))
assert.Is(response.HasBodyContaining(`"ok":true`))
```

Use `jsonassert` for semantic JSON equality that ignores object key order and insignificant whitespace:

```go
import ja "github.com/gogunit/gunit/hammy/jsonassert"

assert.Is(ja.String(actualJSON).EqualTo(expectedJSON))
assert.Is(ja.String(actualJSON).PathEqual("user.name", `"Ada"`))
assert.Is(ja.String(actualJSON).EqualToWithOptions(expectedJSON, ja.IgnorePaths("meta.request_id")))
```

Use `yamlassert` for semantic YAML equality, path checks, and multi-document streams:

```go
import ya "github.com/gogunit/gunit/hammy/yamlassert"

assert.Is(ya.String(actualYAML).EqualTo(expectedYAML))
assert.Is(ya.String(actualYAML).PathEqual("user.name", "Ada\n"))
assert.Is(ya.String(actualYAML).DocumentCount(2))
```

## Generic Matcher Core

```go
func Test_payload_has_expected_type(t *testing.T) {
	assert := a.New(t)
	var payload any = Response{Status: "ok"}

	assert.Is(a.Match(payload, a.TypeOf[Response]()))
}
```

## Map

* [x] EqualTo
* [x] HasEntry
* [x] HasEntries
* [x] HasKeyMatching
* [x] IsEmpty
* [x] Len
* [x] WithItem
* [x] WithItems
* [x] WithKeys
* [x] WithoutItems
* [x] WithoutKeys
* [x] WithValues
* [x] HasValueMatching

## Error

* [x] EqualError
* [x] Error
* [x] ErrorAs
* [x] ErrorContains
* [x] ErrorIs
* [x] ErrorMatchesRegexp
* [x] ErrorType
* [x] NilError
* [x] NotErrorAs
* [x] NotErrorIs

## Filesystem

* [x] DirExists
* [x] FileExists
* [x] NoDirExists
* [x] NoFileExists

## HTTP (`hammy/httpassert`)

HTTP keeps constructor-style wrappers (`Response(resp)`, `Request(req)`, and `Recorder(rec)`) so actual HTTP values live in assertion structs while methods remain chainable and backward-compatible.

* [x] BodyContains
* [x] BodyEqual
* [x] BodyEqualTo
* [x] BodyMatches
* [x] BodyMatchesRegexp
* [x] Header
* [x] HeaderContains
* [x] HeaderEqualTo
* [x] HasBodyContaining
* [x] HasHeader
* [x] HasHeaderContaining
* [x] HasHost
* [x] HasMethod
* [x] HasPath
* [x] HasQueryParam
* [x] HasStatus
* [x] HasStatusInRange
* [x] Host
* [x] Method
* [x] Path
* [x] QueryParam
* [x] Status
* [x] StatusInRange
* [x] URL
* [x] URLEqual
* [x] URLEqualTo

## HTML (`hammy/htmlassert`)

HTML assertions parse documents with the HTML5 parser, so they intentionally do not expose an `IsValid` assertion: the parser recovers from malformed markup by design. `String`, `Bytes`, and `Reader` provide the same focused assertion inventory. Selector-based text and attribute assertions require exactly one matching element.

The package uses `golang.org/x/net/html` for HTML5 parsing and Cascadia for CSS selector matching. Cascadia adds no transitive module beyond `golang.org/x/net`, which this package already requires for parsing; it is kept as a small, pinned dependency rather than maintaining a partial selector engine in this repository.

* [x] EqualTo
* [x] HasSelector
* [x] NoSelector
* [x] SelectorCount
* [x] TextEqualTo
* [x] TextContains
* [x] AttributeEqualTo
* [x] HasAttribute
* [x] NoAttribute
* [x] Contains

```go
import hta "github.com/gogunit/gunit/hammy/htmlassert"

page := hta.String(body)
assert.Is(page.HasSelector("main article"))
assert.Is(page.TextContains("h1", "Welcome"))
assert.Is(page.AttributeEqualTo("form", "method", "post"))
assert.Is(page.SelectorCount("nav a", 3))
```

## JSON (`hammy/jsonassert`)

* [x] String
  * [x] EqualTo
  * [x] EqualToWithOptions
  * [x] LinesEqualTo
  * [x] LinesEqualToWithOptions
  * [x] LinesContain
  * [x] LinesContainSubset
  * [x] IsValid
  * [x] Contains
  * [x] PathExists
  * [x] PathMissing
  * [x] PathEqual
  * [x] ArrayContains
* [x] Bytes
  * [x] EqualTo
  * [x] EqualToWithOptions
  * [x] LinesEqualTo
  * [x] LinesEqualToWithOptions
  * [x] IsValid
  * [x] Contains
  * [x] PathEqual
  * [x] ArrayContains
* [x] Reader
  * [x] EqualTo
  * [x] IsValid
* [x] IgnorePaths
* [x] UnorderedArraysAt

## YAML (`hammy/yamlassert`)

* [x] String
  * [x] EqualTo
  * [x] EqualToWithOptions
  * [x] IsValid
  * [x] Contains
  * [x] PathExists
  * [x] PathMissing
  * [x] PathEqual
  * [x] ArrayContains
  * [x] DocumentCount
  * [x] DocumentEqual
  * [x] DocumentContains
* [x] Bytes
  * [x] EqualTo
  * [x] EqualToWithOptions
  * [x] IsValid
  * [x] Contains
  * [x] PathEqual
  * [x] ArrayContains
  * [x] DocumentCount
  * [x] DocumentEqual
  * [x] DocumentContains
* [x] Reader
  * [x] EqualTo
  * [x] IsValid
* [x] IgnorePaths
* [x] UnorderedArraysAt

## Number

* [x] CloseTo (via `Float`)
* [x] EqualTo
* [x] GreaterThan
* [x] GreaterOrEqual
* [x] IsInf (via `Float`)
* [x] IsNaN (via `Float`)
* [x] IsZero
* [x] LessThan
* [x] LessOrEqual
* [x] Within

## Slice

* [x] Cap
* [x] Contains
* [x] ContainsAny
* [x] ContainsInAnyOrder
* [x] ContainsInOrder
* [x] EqualTo
* [x] Every
* [x] HasItem
* [x] Len
* [x] NotSubsetOf
* [x] SubsetOf

## String

* [x] EqualTo
* [x] Contains
* [x] EqualIgnoringCase
* [x] EqualIgnoringWhitespace
* [x] HasPrefix
* [x] HasPrefixIgnoringCase
* [x] HasSuffix
* [x] HasSuffixIgnoringCase
* [x] IsEmpty
* [x] MatchesRegexp
* [x] ToLowerEqualTo

## Struct

* [x] EqualTo
* [x] Having
* [x] HavingField

## Panic

* [x] NotPanics
* [x] PanicErrorIs
* [x] Panics
* [x] PanicsWithError
* [x] PanicsWithValue

## Polling

* [x] Consistently
* [x] Eventually
* [x] Never

## Time

* [x] After
* [x] AfterOrEqual
* [x] Before
* [x] BeforeOrEqual
* [x] EqualTo
* [x] Matches
* [x] WithinDuration
* [x] WithinRange

## Writing a Custom Matcher

```go
package model

import a "github.com/gogunit/gunit/hammy"

func Matcher(model Model) *Matchy {
	return &Matchy{
		model: model,
    }
}

type Matchy struct {
	model Model
}

func (m *Matcher) HasName(expected string) a.AssertionMessage {
	actual := m.model.Name
	return a.Assert(actual == expected, "want Name=<%v> equal to <%v>", actual, expected)
}
```

```go
assert.Is(Matcher(model).HasName("bar"))

// fails with the message
// want Name=<"foo"> equal to <"bar">
```
