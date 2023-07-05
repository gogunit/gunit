package hammy_test

import (
	"errors"
	"testing"

	"github.com/gogunit/gunit/eye"
	hammy "github.com/gogunit/gunit/hammy"
)

func Test_is_error_success(t *testing.T) {
	var err = errors.New("hello world")
	hammy.New(t).Is(hammy.Error(err))
}

func Test_is_error_failure(t *testing.T) {
	aSpy := eye.Spy()
	hammy.New(aSpy).Is(hammy.Error(nil))
	aSpy.HadErrorContaining(t, "got <nil>, want error")
}

func Test_nil_error_success(t *testing.T) {
	hammy.New(t).Is(hammy.NilError(nil))
}

func Test_nil_error_failure(t *testing.T) {
	aSpy := eye.Spy()
	var err = errors.New("hello world")
	hammy.New(aSpy).Is(hammy.NilError(err))
	aSpy.HadErrorContaining(t, "got <hello world>, want nil error")
}
