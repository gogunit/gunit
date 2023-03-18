package hammy_test

import (
	"testing"

	"github.com/gogunit/gunit/eye"
	"github.com/gogunit/gunit/hammy"
)

func Test_Map_EqualTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]bool{"def": true}).EqualTo(map[string]bool{"def": false}))
	aSpy.HadErrorContaining(t, "Map mismatch (-want +got):\n")
}

func Test_Map_EqualTo_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]bool{"abc": true}).EqualTo(map[string]bool{"abc": true}))
}

func Test_Map_WithKeys_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]bool{"abc": true, "def": true}).WithKeys("abc", "ghi"))
	aSpy.HadError(t)
}

func Test_Map_WithKeys_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]bool{"abc": true, "def": true}).WithKeys("abc", "def"))
}

func Test_Map_IsEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]bool{"abc": true, "def": true}).IsEmpty())
	aSpy.HadError(t)
}

func Test_Map_IsEmpty_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]int{}).IsEmpty())
}

func Test_Map_WithValues_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]int{"abc": 42, "def": 33}).WithValues(4, 3))
	aSpy.HadError(t)
}

func Test_Map_WithValues_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]int{"abc": 42, "def": 33}).WithValues(42, 33))
}

func Test_Map_WithoutKeys_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]int{"abc": 42, "def": 33}).WithoutKeys("abc", "def", "jkl"))
	aSpy.HadError(t)
}

func Test_Map_WithoutKeys_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]int{"abc": 42, "def": 33}).WithoutKeys("ghi", "jkl"))
}

func Test_Map_Len_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]int{}).Len(2))
	aSpy.HadError(t)
}

func Test_Map_Len_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]int{"hi": 1}).Len(1))
}

func Test_Map_HasItem_value_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]int{"hi": 1}).WithItem("hi", 2))
	aSpy.HadErrorContaining(t, "want value=<2> for key=<hi>, got <1>")
}

func Test_Map_HasItem_key_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]int{"hi": 1}).WithItem("bye", 2))
	aSpy.HadErrorContaining(t, "want key=<bye>, but was absent")
}

func Test_Map_HasItem_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]int{"hi": 1}).WithItem("hi", 1))
}
