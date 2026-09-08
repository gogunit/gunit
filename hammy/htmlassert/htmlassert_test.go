package htmlassert

import (
	"errors"
	"strings"
	"testing"
)

func TestSemanticAssertions(t *testing.T) {
	actual := `<!doctype html><html><body><main id="content" class="page"><h1>Hello &amp; welcome</h1><p>Nested <strong> text </strong> here</p><input disabled></main></body></html>`
	tests := []struct {
		name    string
		success bool
	}{
		{"equal document", String(actual).EqualTo(`<!DOCTYPE html><html><body><main class="page" id="content"><h1>Hello &amp; welcome</h1><p>Nested<strong>text</strong>here</p><input disabled=""></main></body></html>`).IsSuccessful},
		{"selector", String(actual).HasSelector(`main.page > h1`).IsSuccessful},
		{"no selector", String(actual).NoSelector(`footer`).IsSuccessful},
		{"count", String(actual).SelectorCount(`main > *`, 3).IsSuccessful},
		{"nested text", String(actual).TextEqualTo(`p`, "Nested text here").IsSuccessful},
		{"escaped entity", String(actual).TextContains(`h1`, "& welcome").IsSuccessful},
		{"attribute", String(actual).AttributeEqualTo(`main`, "id", "content").IsSuccessful},
		{"boolean attribute", String(actual).HasAttribute(`input`, "disabled").IsSuccessful},
		{"missing attribute", String(actual).NoAttribute(`input`, "checked").IsSuccessful},
		{"fragment", String(actual).Contains(`<h1>Hello &amp; welcome</h1>`).IsSuccessful},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.success {
				t.Fatal("assertion failed")
			}
		})
	}
}

func TestBytesAndReader(t *testing.T) {
	if !Bytes([]byte(`<p>x</p>`)).HasSelector("p").IsSuccessful {
		t.Fatal("bytes failed")
	}
	if !Reader(strings.NewReader(`<p>x</p>`)).TextEqualTo("p", "x").IsSuccessful {
		t.Fatal("reader failed")
	}
	if !Reader(strings.NewReader(`<p>x</p>`)).EqualTo(strings.NewReader(`<p>x</p>`)).IsSuccessful {
		t.Fatal("reader equality failed")
	}
}

func TestFailuresAreDescriptive(t *testing.T) {
	html := String(`<div class=x></div><div class=x></div>`)
	failures := []struct {
		result            bool
		message, contains string
	}{
		{html.HasSelector("[").IsSuccessful, html.HasSelector("[").Message, "malformed selector <[>"},
		{html.HasSelector("p").IsSuccessful, html.HasSelector("p").Message, "selector <p>"},
		{html.NoSelector("div").IsSuccessful, html.NoSelector("div").Message, "matched <2>"},
		{html.SelectorCount("div", 3).IsSuccessful, html.SelectorCount("div", 3).Message, "expected <3>"},
		{html.TextEqualTo(".x", "value").IsSuccessful, html.TextEqualTo(".x", "value").Message, "expected exactly one"},
		{html.AttributeEqualTo("div:first-child", "id", "wanted").IsSuccessful, html.AttributeEqualTo("div:first-child", "id", "wanted").Message, "expected <wanted>"},
		{html.Contains(`<span></span>`).IsSuccessful, html.Contains(`<span></span>`).Message, "did not contain"},
	}
	for _, f := range failures {
		if f.result || !strings.Contains(f.message, f.contains) {
			t.Fatalf("failure %q did not contain %q", f.message, f.contains)
		}
	}
}

func TestEmptyInputIsParsedAsAnHTMLDocument(t *testing.T) {
	if !String("").EqualTo("").IsSuccessful {
		t.Fatal("empty documents should compare equal")
	}
	if !String("").NoSelector("body > *").IsSuccessful {
		t.Fatal("empty body should have no elements")
	}
}

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) { return 0, errors.New("boom") }
func TestReaderFailures(t *testing.T) {
	for _, result := range []struct {
		ok      bool
		message string
	}{
		{Reader(nil).HasSelector("p").IsSuccessful, Reader(nil).HasSelector("p").Message},
		{Reader(failingReader{}).HasSelector("p").IsSuccessful, Reader(failingReader{}).HasSelector("p").Message},
		{Reader(strings.NewReader("<p>x</p>")).EqualTo(nil).IsSuccessful, Reader(strings.NewReader("<p>x</p>")).EqualTo(nil).Message},
	} {
		if result.ok || result.message == "" {
			t.Fatal("expected descriptive reader failure")
		}
	}
}
