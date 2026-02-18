package hammy_test

import (
	"testing"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

func Test_Not_False_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.IsNot(a.False(false))
	aSpy.HadErrorContaining(t, "not(got true, wanted false)")
}

func Test_Not_False_success(t *testing.T) {
	assert := a.New(t)
	assert.IsNot(a.False(true))
}

func Test_Nil_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.Nil(t))
	aSpy.HadErrorContaining(t, "got <*testing.T>, wanted nil")
}

func Test_Nil_success(t *testing.T) {
	assert := a.New(t)
	var i *int = nil
	assert.Is(a.Nil(i))
}

func Test_NotNil_success(t *testing.T) {
	assert := a.New(t)
	var i = 1
	assert.Is(a.NotNil(&i))
}

func Test_NotNil_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	var i *int = nil
	assert.Is(a.NotNil(i))
	aSpy.HadErrorContaining(t, "got nil, wanted <*int>")
}

func Test_Nil_interface_success(t *testing.T) {
	assert := a.New(t)
	var value any
	assert.Is(a.Nil(value))
}

func Test_Nil_typed_interface_success(t *testing.T) {
	assert := a.New(t)
	var i *int
	var value any = i
	assert.Is(a.Nil(value))
}

func Test_Nil_map_success(t *testing.T) {
	assert := a.New(t)
	var value map[string]int
	assert.Is(a.Nil(value))
}

func Test_Nil_slice_success(t *testing.T) {
	assert := a.New(t)
	var value []int
	assert.Is(a.Nil(value))
}

func Test_Nil_chan_success(t *testing.T) {
	assert := a.New(t)
	var value chan int
	assert.Is(a.Nil(value))
}

func Test_Nil_func_success(t *testing.T) {
	assert := a.New(t)
	var value func()
	assert.Is(a.Nil(value))
}

func Test_NotNil_interface_success(t *testing.T) {
	assert := a.New(t)
	var value any = "hello"
	assert.Is(a.NotNil(value))
}

func Test_NotNil_map_success(t *testing.T) {
	assert := a.New(t)
	value := map[string]int{}
	assert.Is(a.NotNil(value))
}

func Test_NotNil_slice_success(t *testing.T) {
	assert := a.New(t)
	value := []int{}
	assert.Is(a.NotNil(value))
}

func Test_NotNil_chan_success(t *testing.T) {
	assert := a.New(t)
	value := make(chan int)
	assert.Is(a.NotNil(value))
}

func Test_NotNil_func_success(t *testing.T) {
	assert := a.New(t)
	value := func() {}
	assert.Is(a.NotNil(value))
}

func Test_True_success(t *testing.T) {
	assert := a.New(t)
	assert.Is(a.True(true))
}

func Test_True_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.True(false))
	aSpy.HadErrorContaining(t, "got false, wanted true")
}

func Test_False_success(t *testing.T) {
	assert := a.New(t)
	assert.Is(a.False(false))
}

func Test_False_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.False(true))
	aSpy.HadErrorContaining(t, "got true, wanted false")
}
