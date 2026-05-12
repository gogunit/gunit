package hammy

import (
	"errors"
	"os"
)

func FileExists(path string) AssertionMessage {
	info, err := os.Stat(path)
	if err != nil {
		return Assert(false, "got path <%s> unavailable: %v, wanted existing file", path, err)
	}
	return Assert(!info.IsDir(), "got directory <%s>, wanted existing file", path)
}

func NoFileExists(path string) AssertionMessage {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return Assert(true, "got no file at <%s>", path)
	}
	if err != nil {
		return Assert(false, "got path <%s> unavailable: %v, wanted no existing file", path, err)
	}
	return Assert(info.IsDir(), "got existing file <%s>, wanted no file", path)
}

func DirExists(path string) AssertionMessage {
	info, err := os.Stat(path)
	if err != nil {
		return Assert(false, "got path <%s> unavailable: %v, wanted existing directory", path, err)
	}
	return Assert(info.IsDir(), "got file <%s>, wanted existing directory", path)
}

func NoDirExists(path string) AssertionMessage {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return Assert(true, "got no directory at <%s>", path)
	}
	if err != nil {
		return Assert(false, "got path <%s> unavailable: %v, wanted no existing directory", path, err)
	}
	return Assert(!info.IsDir(), "got existing directory <%s>, wanted no directory", path)
}
