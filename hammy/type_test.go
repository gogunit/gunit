package hammy_test

import (
	"testing"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

type testGreeter interface {
	Greet() string
}

type testGreeterImpl struct{}

func (testGreeterImpl) Greet() string {
	return "hello"
}

func Test_SamePointer_success(t *testing.T) {
	value := 42
	a.New(t).Is(a.Match(&value, a.SamePointer(&value)))
}

func Test_SamePointer_failure(t *testing.T) {
	aSpy := eye.Spy()
	left := 42
	right := 42
	a.New(aSpy).Is(a.Match(&left, a.SamePointer(&right)))
	aSpy.HadErrorContaining(t, "wanted same pointer")
}

func Test_TypeOf_success(t *testing.T) {
	var value any = examplePerson{Name: "Ada"}
	a.New(t).Is(a.Match(value, a.TypeOf[examplePerson]()))
}

func Test_AssignableTo_success(t *testing.T) {
	var value any = testGreeterImpl{}
	a.New(t).Is(a.Match(value, a.AssignableTo[testGreeter]()))
}

func Test_AssignableTo_failure_on_nil(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Match(any(nil), a.AssignableTo[testGreeter]()))
	aSpy.HadErrorContaining(t, "got <nil>, wanted assignable")
}
