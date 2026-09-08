// Package htmlassert provides semantic assertions for HTML documents and fragments.
package htmlassert

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/andybalholm/cascadia"
	"github.com/gogunit/gunit/hammy"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// StringAssert wraps HTML supplied as a string.
type StringAssert struct{ actual string }

// BytesAssert wraps HTML supplied as bytes.
type BytesAssert struct{ actual []byte }

// ReaderAssert wraps HTML supplied through a reader. The reader is consumed by an assertion.
type ReaderAssert struct{ actual io.Reader }

func String(actual string) *StringAssert    { return &StringAssert{actual: actual} }
func Bytes(actual []byte) *BytesAssert      { return &BytesAssert{actual: actual} }
func Reader(actual io.Reader) *ReaderAssert { return &ReaderAssert{actual: actual} }

func (a *StringAssert) EqualTo(expected string) hammy.AssertionMessage {
	return equal([]byte(a.actual), []byte(expected))
}
func (a *BytesAssert) EqualTo(expected []byte) hammy.AssertionMessage {
	return equal(a.actual, expected)
}
func (a *ReaderAssert) EqualTo(expected io.Reader) hammy.AssertionMessage {
	actual, result := read("actual", a.actual)
	if !result.IsSuccessful {
		return result
	}
	want, result := read("expected", expected)
	if !result.IsSuccessful {
		return result
	}
	return equal(actual, want)
}

func equal(actual, expected []byte) hammy.AssertionMessage {
	got, err := parse(actual)
	if err != nil {
		return hammy.Assert(false, "actual HTML could not be parsed: %v", err)
	}
	want, err := parse(expected)
	if err != nil {
		return hammy.Assert(false, "expected HTML could not be parsed: %v", err)
	}
	g, w := canonical(got), canonical(want)
	return hammy.Assert(g == w, "HTML mismatch: got <%s>, expected <%s>", g, w)
}

func (a *StringAssert) HasSelector(selector string) hammy.AssertionMessage {
	return queryAssertion([]byte(a.actual), selector, queryHas)
}
func (a *BytesAssert) HasSelector(selector string) hammy.AssertionMessage {
	return queryAssertion(a.actual, selector, queryHas)
}
func (a *ReaderAssert) HasSelector(selector string) hammy.AssertionMessage {
	return a.with(func(b []byte) hammy.AssertionMessage { return queryAssertion(b, selector, queryHas) })
}
func (a *StringAssert) NoSelector(selector string) hammy.AssertionMessage {
	return queryAssertion([]byte(a.actual), selector, queryNone)
}
func (a *BytesAssert) NoSelector(selector string) hammy.AssertionMessage {
	return queryAssertion(a.actual, selector, queryNone)
}
func (a *ReaderAssert) NoSelector(selector string) hammy.AssertionMessage {
	return a.with(func(b []byte) hammy.AssertionMessage { return queryAssertion(b, selector, queryNone) })
}
func (a *StringAssert) SelectorCount(selector string, expected int) hammy.AssertionMessage {
	return selectorCount([]byte(a.actual), selector, expected)
}
func (a *BytesAssert) SelectorCount(selector string, expected int) hammy.AssertionMessage {
	return selectorCount(a.actual, selector, expected)
}
func (a *ReaderAssert) SelectorCount(selector string, expected int) hammy.AssertionMessage {
	return a.with(func(b []byte) hammy.AssertionMessage { return selectorCount(b, selector, expected) })
}

const (
	queryHas  = true
	queryNone = false
)

func queryAssertion(input []byte, selector string, want bool) hammy.AssertionMessage {
	nodes, result := selectNodes(input, selector)
	if !result.IsSuccessful {
		return result
	}
	ok := len(nodes) > 0
	if want {
		return hammy.Assert(ok, "selector <%s> matched <%d> nodes, expected at least one", selector, len(nodes))
	}
	return hammy.Assert(!ok, "selector <%s> matched <%d> nodes, expected none", selector, len(nodes))
}
func selectorCount(input []byte, selector string, expected int) hammy.AssertionMessage {
	nodes, result := selectNodes(input, selector)
	if !result.IsSuccessful {
		return result
	}
	return hammy.Assert(len(nodes) == expected, "selector <%s> matched <%d> nodes, expected <%d>", selector, len(nodes), expected)
}

func (a *StringAssert) TextEqualTo(selector, expected string) hammy.AssertionMessage {
	return textAssert([]byte(a.actual), selector, expected, false)
}
func (a *BytesAssert) TextEqualTo(selector, expected string) hammy.AssertionMessage {
	return textAssert(a.actual, selector, expected, false)
}
func (a *ReaderAssert) TextEqualTo(selector, expected string) hammy.AssertionMessage {
	return a.with(func(b []byte) hammy.AssertionMessage { return textAssert(b, selector, expected, false) })
}
func (a *StringAssert) TextContains(selector, expected string) hammy.AssertionMessage {
	return textAssert([]byte(a.actual), selector, expected, true)
}
func (a *BytesAssert) TextContains(selector, expected string) hammy.AssertionMessage {
	return textAssert(a.actual, selector, expected, true)
}
func (a *ReaderAssert) TextContains(selector, expected string) hammy.AssertionMessage {
	return a.with(func(b []byte) hammy.AssertionMessage { return textAssert(b, selector, expected, true) })
}
func textAssert(input []byte, selector, expected string, contains bool) hammy.AssertionMessage {
	node, result := one(input, selector)
	if !result.IsSuccessful {
		return result
	}
	actual, want := normalizedText(node), normalize(expected)
	if contains {
		return hammy.Assert(strings.Contains(actual, want), "text for selector <%s> was <%s>, expected it to contain <%s>", selector, actual, want)
	}
	return hammy.Assert(actual == want, "text for selector <%s> was <%s>, expected <%s>", selector, actual, want)
}

func (a *StringAssert) AttributeEqualTo(selector, name, expected string) hammy.AssertionMessage {
	return attributeAssert([]byte(a.actual), selector, name, expected, attrEqual)
}
func (a *BytesAssert) AttributeEqualTo(selector, name, expected string) hammy.AssertionMessage {
	return attributeAssert(a.actual, selector, name, expected, attrEqual)
}
func (a *ReaderAssert) AttributeEqualTo(selector, name, expected string) hammy.AssertionMessage {
	return a.with(func(b []byte) hammy.AssertionMessage { return attributeAssert(b, selector, name, expected, attrEqual) })
}
func (a *StringAssert) HasAttribute(selector, name string) hammy.AssertionMessage {
	return attributeAssert([]byte(a.actual), selector, name, "", attrHas)
}
func (a *BytesAssert) HasAttribute(selector, name string) hammy.AssertionMessage {
	return attributeAssert(a.actual, selector, name, "", attrHas)
}
func (a *ReaderAssert) HasAttribute(selector, name string) hammy.AssertionMessage {
	return a.with(func(b []byte) hammy.AssertionMessage { return attributeAssert(b, selector, name, "", attrHas) })
}
func (a *StringAssert) NoAttribute(selector, name string) hammy.AssertionMessage {
	return attributeAssert([]byte(a.actual), selector, name, "", attrNone)
}
func (a *BytesAssert) NoAttribute(selector, name string) hammy.AssertionMessage {
	return attributeAssert(a.actual, selector, name, "", attrNone)
}
func (a *ReaderAssert) NoAttribute(selector, name string) hammy.AssertionMessage {
	return a.with(func(b []byte) hammy.AssertionMessage { return attributeAssert(b, selector, name, "", attrNone) })
}

type attrMode int

const (
	attrEqual attrMode = iota
	attrHas
	attrNone
)

func attributeAssert(input []byte, selector, name, expected string, mode attrMode) hammy.AssertionMessage {
	node, result := one(input, selector)
	if !result.IsSuccessful {
		return result
	}
	actual, found := attribute(node, name)
	switch mode {
	case attrHas:
		return hammy.Assert(found, "selector <%s> had no attribute <%s>", selector, name)
	case attrNone:
		return hammy.Assert(!found, "selector <%s> had attribute <%s> with value <%s>, expected it to be absent", selector, name, actual)
	default:
		return hammy.Assert(found && actual == expected, "attribute <%s> for selector <%s> was <%s> (present: %t), expected <%s>", name, selector, actual, found, expected)
	}
}

func (a *StringAssert) Contains(expected string) hammy.AssertionMessage {
	return contains([]byte(a.actual), []byte(expected))
}
func (a *BytesAssert) Contains(expected []byte) hammy.AssertionMessage {
	return contains(a.actual, expected)
}
func (a *ReaderAssert) Contains(expected io.Reader) hammy.AssertionMessage {
	actual, result := read("actual", a.actual)
	if !result.IsSuccessful {
		return result
	}
	want, result := read("expected", expected)
	if !result.IsSuccessful {
		return result
	}
	return contains(actual, want)
}
func contains(actual, expected []byte) hammy.AssertionMessage {
	doc, err := parse(actual)
	if err != nil {
		return hammy.Assert(false, "actual HTML could not be parsed: %v", err)
	}
	fragments, err := html.ParseFragment(bytes.NewReader(expected), &html.Node{Type: html.ElementNode, Data: "body", DataAtom: atom.Body})
	if err != nil {
		return hammy.Assert(false, "expected HTML fragment could not be parsed: %v", err)
	}
	want := canonicalForest(fragments)
	if want == "" {
		return hammy.Assert(true, "HTML contains empty fragment")
	}
	if structuralContains(doc, want) {
		return hammy.Assert(true, "HTML contains expected fragment <%s>", want)
	}
	return hammy.Assert(false, "HTML did not contain expected fragment <%s>", want)
}

func parse(input []byte) (*html.Node, error) { return html.Parse(bytes.NewReader(input)) }
func selectNodes(input []byte, selector string) ([]*html.Node, hammy.AssertionMessage) {
	matcher, err := cascadia.Parse(selector)
	if err != nil {
		return nil, hammy.Assert(false, "malformed selector <%s>: %v", selector, err)
	}
	doc, err := parse(input)
	if err != nil {
		return nil, hammy.Assert(false, "actual HTML could not be parsed: %v", err)
	}
	return cascadia.QueryAll(doc, matcher), hammy.Assert(true, "selector <%s> parsed", selector)
}
func one(input []byte, selector string) (*html.Node, hammy.AssertionMessage) {
	nodes, result := selectNodes(input, selector)
	if !result.IsSuccessful {
		return nil, result
	}
	if len(nodes) == 0 {
		return nil, hammy.Assert(false, "selector <%s> matched no elements", selector)
	}
	if len(nodes) != 1 {
		return nil, hammy.Assert(false, "selector <%s> matched <%d> elements, expected exactly one", selector, len(nodes))
	}
	return nodes[0], hammy.Assert(true, "selector <%s> matched one element", selector)
}
func (a *ReaderAssert) with(fn func([]byte) hammy.AssertionMessage) hammy.AssertionMessage {
	b, result := read("actual", a.actual)
	if !result.IsSuccessful {
		return result
	}
	return fn(b)
}
func read(label string, reader io.Reader) ([]byte, hammy.AssertionMessage) {
	if reader == nil {
		return nil, hammy.Assert(false, "%s HTML reader was nil", label)
	}
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, hammy.Assert(false, "could not read %s HTML: %v", label, err)
	}
	return b, hammy.Assert(true, "read %s HTML", label)
}
func normalize(s string) string { return strings.Join(strings.Fields(s), " ") }
func normalizedText(n *html.Node) string {
	var b strings.Builder
	var walk func(*html.Node)
	walk = func(x *html.Node) {
		if x.Type == html.TextNode {
			b.WriteString(x.Data)
		}
		for c := x.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return normalize(b.String())
}
func attribute(n *html.Node, name string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == name {
			return a.Val, true
		}
	}
	return "", false
}

func canonical(n *html.Node) string { var b strings.Builder; writeCanonical(&b, n); return b.String() }
func canonicalForest(nodes []*html.Node) string {
	var b strings.Builder
	for _, n := range nodes {
		writeCanonical(&b, n)
	}
	return b.String()
}
func writeCanonical(b *strings.Builder, n *html.Node) {
	switch n.Type {
	case html.DocumentNode:
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			writeCanonical(b, c)
		}
	case html.ElementNode:
		b.WriteByte('<')
		b.WriteString(n.Namespace)
		b.WriteByte(':')
		b.WriteString(n.Data)
		attrs := append([]html.Attribute(nil), n.Attr...)
		sort.Slice(attrs, func(i, j int) bool {
			if attrs[i].Namespace != attrs[j].Namespace {
				return attrs[i].Namespace < attrs[j].Namespace
			}
			if attrs[i].Key != attrs[j].Key {
				return attrs[i].Key < attrs[j].Key
			}
			return attrs[i].Val < attrs[j].Val
		})
		for _, a := range attrs {
			fmt.Fprintf(b, " %s:%s=%q", a.Namespace, a.Key, a.Val)
		}
		b.WriteByte('>')
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			writeCanonical(b, c)
		}
		fmt.Fprintf(b, "</%s:%s>", n.Namespace, n.Data)
	case html.TextNode:
		if text := normalize(n.Data); text != "" {
			fmt.Fprintf(b, "#%q", text)
		}
	case html.CommentNode:
		fmt.Fprintf(b, "<!--%s-->", n.Data)
	case html.DoctypeNode:
		fmt.Fprintf(b, "<!DOCTYPE %s>", n.Data)
	}
}
func structuralContains(n *html.Node, want string) bool {
	if canonical(n) == want {
		return true
	}
	for first := n.FirstChild; first != nil; first = first.NextSibling {
		var nodes []*html.Node
		for x := first; x != nil; x = x.NextSibling {
			nodes = append(nodes, x)
			if canonicalForest(nodes) == want {
				return true
			}
		}
		if structuralContains(first, want) {
			return true
		}
	}
	return false
}
