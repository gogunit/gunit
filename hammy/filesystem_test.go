package hammy_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

func Test_FileExists_success(t *testing.T) {
	assert := a.New(t)
	path := filepath.Join(t.TempDir(), "payload.txt")
	err := os.WriteFile(path, []byte("payload"), 0o600)
	assert.Is(a.NilError(err))

	assert.Is(a.FileExists(path))
}

func Test_FileExists_failure_missing(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.FileExists(filepath.Join(t.TempDir(), "missing.txt")))

	aSpy.HadErrorContaining(t, "wanted existing file")
}

func Test_FileExists_failure_directory(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.FileExists(t.TempDir()))

	aSpy.HadErrorContaining(t, "wanted existing file")
}

func Test_NoFileExists_success_missing(t *testing.T) {
	assert := a.New(t)

	assert.Is(a.NoFileExists(filepath.Join(t.TempDir(), "missing.txt")))
}

func Test_NoFileExists_success_directory(t *testing.T) {
	assert := a.New(t)

	assert.Is(a.NoFileExists(t.TempDir()))
}

func Test_NoFileExists_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	path := filepath.Join(t.TempDir(), "payload.txt")
	err := os.WriteFile(path, []byte("payload"), 0o600)
	assert.Is(a.NilError(err))

	assert.Is(a.NoFileExists(path))

	aSpy.HadErrorContaining(t, "wanted no file")
}

func Test_DirExists_success(t *testing.T) {
	assert := a.New(t)

	assert.Is(a.DirExists(t.TempDir()))
}

func Test_DirExists_failure_missing(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.DirExists(filepath.Join(t.TempDir(), "missing")))

	aSpy.HadErrorContaining(t, "wanted existing directory")
}

func Test_DirExists_failure_file(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	path := filepath.Join(t.TempDir(), "payload.txt")
	err := os.WriteFile(path, []byte("payload"), 0o600)
	assert.Is(a.NilError(err))

	assert.Is(a.DirExists(path))

	aSpy.HadErrorContaining(t, "wanted existing directory")
}

func Test_NoDirExists_success_missing(t *testing.T) {
	assert := a.New(t)

	assert.Is(a.NoDirExists(filepath.Join(t.TempDir(), "missing")))
}

func Test_NoDirExists_success_file(t *testing.T) {
	assert := a.New(t)
	path := filepath.Join(t.TempDir(), "payload.txt")
	err := os.WriteFile(path, []byte("payload"), 0o600)
	assert.Is(a.NilError(err))

	assert.Is(a.NoDirExists(path))
}

func Test_NoDirExists_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.NoDirExists(t.TempDir()))

	aSpy.HadErrorContaining(t, "wanted no directory")
}
