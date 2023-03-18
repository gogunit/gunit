package hammy_test

import (
	"testing"

	"github.com/gogunit/gunit/eye"
	"github.com/gogunit/gunit/hammy"
)

type aType struct {
	Id int
}

func Test_Struct_EqualTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Struct(aType{1}).EqualTo(aType{2}))
	aSpy.HadErrorContaining(t, "Struct mismatch (-want +got):\n")
}

func Test_Struct_EqualTo_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Struct(aType{1}).EqualTo(aType{1}))
}
