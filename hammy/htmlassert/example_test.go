package htmlassert_test

import (
	"fmt"

	"github.com/gogunit/gunit/hammy"
	"github.com/gogunit/gunit/hammy/htmlassert"
)

func Example() {
	body := `<article data-id="42"><h1>Hammy</h1><p>Semantic HTML assertions</p><a href="/docs">Docs</a></article>`
	assert := hammy.New(exampleT{})
	assert.Is(htmlassert.String(body).HasSelector("article > h1"))
	assert.Is(htmlassert.String(body).TextEqualTo("h1", "Hammy"))
	assert.Is(htmlassert.String(body).AttributeEqualTo("article", "data-id", "42"))
	assert.Is(htmlassert.String(body).SelectorCount("article a", 1))
	fmt.Println("all HTML assertions passed")
	// Output: all HTML assertions passed
}

type exampleT struct{}

func (exampleT) Helper()                           {}
func (exampleT) Errorf(format string, args ...any) { panic(fmt.Sprintf(format, args...)) }
func (exampleT) Fatalf(format string, args ...any) { panic(fmt.Sprintf(format, args...)) }
