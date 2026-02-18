package hammy_test

import (
	"testing"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

type aType struct {
	Id int
}

func Test_Struct_EqualTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.Struct(aType{1}).EqualTo(aType{2}))
	aSpy.HadErrorContaining(t, "Structs are not equal (+got -want):\n")
}

func Test_Struct_EqualTo_success(t *testing.T) {
	assert := a.New(t)
	assert.Is(a.Struct(aType{1}).EqualTo(aType{1}))
}
