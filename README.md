# gunit

[![Go](https://github.com/nfisher/gunit/actions/workflows/go.yml/badge.svg)](https://github.com/nfisher/gunit/actions/workflows/go.yml)
[![CodeQL](https://github.com/nfisher/gunit/actions/workflows/codeql.yml/badge.svg)](https://github.com/nfisher/gunit/actions/workflows/codeql.yml)

Go unit test assertions library.

## Example

```go
func Test_int_GreaterThan_succeeds(t *testing.T) {
  actual := 9 + 2
  expected := 10
	Number(t, actual).GreaterThan(expected)
}
```
