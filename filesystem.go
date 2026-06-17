package gunit

import (
	"errors"
	"os"
)

func FileExists(t T, path string) {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Errorf("got path <%s> unavailable: %v, wanted existing file", path, err)
		return
	}
	if info.IsDir() {
		t.Errorf("got directory <%s>, wanted existing file", path)
	}
}
func NoFileExists(t T, path string) {
	t.Helper()
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	if err != nil {
		t.Errorf("got path <%s> unavailable: %v, wanted no existing file", path, err)
		return
	}
	if !info.IsDir() {
		t.Errorf("got existing file <%s>, wanted no file", path)
	}
}
func DirExists(t T, path string) {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Errorf("got path <%s> unavailable: %v, wanted existing directory", path, err)
		return
	}
	if !info.IsDir() {
		t.Errorf("got file <%s>, wanted existing directory", path)
	}
}
func NoDirExists(t T, path string) {
	t.Helper()
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	if err != nil {
		t.Errorf("got path <%s> unavailable: %v, wanted no existing directory", path, err)
		return
	}
	if info.IsDir() {
		t.Errorf("got existing directory <%s>, wanted no directory", path)
	}
}
