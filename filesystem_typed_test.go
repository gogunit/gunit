package gunit_test

import (
	"github.com/gogunit/gunit"
	"github.com/gogunit/gunit/eye"
	"path/filepath"
	"testing"
)

func Test_filesystem_FileExists_success(t *testing.T) {
	_, file := makeTempDirAndFile(t)
	gunit.FileExists(t, file)
}
func Test_filesystem_FileExists_failure_for_missing(t *testing.T) {
	dir := t.TempDir()
	aSpy := eye.Spy()
	gunit.FileExists(aSpy, filepath.Join(dir, "missing.txt"))
	aSpy.HadErrorContaining(t, "wanted existing file")
}
func Test_filesystem_FileExists_failure_for_directory(t *testing.T) {
	dir := t.TempDir()
	aSpy := eye.Spy()
	gunit.FileExists(aSpy, dir)
	aSpy.HadErrorContaining(t, "wanted existing file")
}

func Test_filesystem_NoFileExists_success_for_missing(t *testing.T) {
	dir := t.TempDir()
	gunit.NoFileExists(t, filepath.Join(dir, "missing.txt"))
}
func Test_filesystem_NoFileExists_success_for_directory(t *testing.T) {
	gunit.NoFileExists(t, t.TempDir())
}
func Test_filesystem_NoFileExists_failure(t *testing.T) {
	_, file := makeTempDirAndFile(t)
	aSpy := eye.Spy()
	gunit.NoFileExists(aSpy, file)
	aSpy.HadErrorContaining(t, "wanted no file")
}

func Test_filesystem_DirExists_success(t *testing.T) {
	gunit.DirExists(t, t.TempDir())
}
func Test_filesystem_DirExists_failure_for_missing(t *testing.T) {
	dir := t.TempDir()
	aSpy := eye.Spy()
	gunit.DirExists(aSpy, filepath.Join(dir, "missing"))
	aSpy.HadErrorContaining(t, "wanted existing directory")
}
func Test_filesystem_DirExists_failure_for_file(t *testing.T) {
	_, file := makeTempDirAndFile(t)
	aSpy := eye.Spy()
	gunit.DirExists(aSpy, file)
	aSpy.HadErrorContaining(t, "wanted existing directory")
}

func Test_filesystem_NoDirExists_success_for_missing(t *testing.T) {
	dir := t.TempDir()
	gunit.NoDirExists(t, filepath.Join(dir, "missing"))
}
func Test_filesystem_NoDirExists_success_for_file(t *testing.T) {
	_, file := makeTempDirAndFile(t)
	gunit.NoDirExists(t, file)
}
func Test_filesystem_NoDirExists_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.NoDirExists(aSpy, t.TempDir())
	aSpy.HadErrorContaining(t, "wanted no directory")
}
