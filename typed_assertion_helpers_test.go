package gunit_test

import (
	"github.com/gogunit/gunit"
	"net/url"
	"os"
	"path/filepath"
	"testing"
)

func makeTempDirAndFile(t *testing.T) (string, string) {
	t.Helper()
	dir := t.TempDir()
	file := filepath.Join(dir, "file.txt")
	if err := os.WriteFile(file, []byte("hello"), 0o600); err != nil {
		t.Fatal(err)
	}
	return dir, file
}

func testURL(t gunit.T) *gunit.URLAssert {
	parsed, err := url.Parse("https://example.com/path?q=go")
	if err != nil {
		t.Fatalf("failed to parse test URL: %v", err)
	}
	return gunit.URL(t, parsed)
}

func testMap() map[string]int { return map[string]int{"a": 1, "b": 2} }
