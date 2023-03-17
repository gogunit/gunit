package hammy_test

import (
	"github.com/nfisher/gunit/eye"
	"github.com/nfisher/gunit/hammy"
	"testing"
)

func Test_string_EqualTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.String("hi").EqualTo("by"))
	aSpy.HadError(t)
}

func Test_string_EqualTo_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.String("hi").EqualTo("hi"))
}
