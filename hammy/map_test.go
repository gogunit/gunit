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
	aSpy.HadErrorContaining(t, "got <[abc]>, wanted keys <[ghi]>")
}

func Test_Map_WithKeys_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]bool{"abc": true, "def": true}).WithKeys("abc", "def"))
}

func Test_Map_IsEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]bool{"abc": true, "def": true}).IsEmpty())
	aSpy.HadErrorContaining(t, "got len=<2>, wanted empty map")
}

func Test_Map_IsEmpty_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]int{}).IsEmpty())
}

func Test_Map_WithValues_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]int{"abc": 42, "def": 33}).WithValues(4, 3))
	aSpy.HadErrorContaining(t, "got <[]>, wanted values <[4 3]>")
}

func Test_Map_WithValues_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]int{"abc": 42, "def": 33}).WithValues(42, 33))
}

func Test_Map_WithoutKeys_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]int{"abc": 42, "def": 33}).WithoutKeys("abc", "def", "jkl"))
	aSpy.HadErrorContaining(t, "got keys <[abc def]>, wanted absent from map")
}

func Test_Map_WithoutKeys_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]int{"abc": 42, "def": 33}).WithoutKeys("ghi", "jkl"))
}

func Test_Map_Len_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]int{}).Len(2))
	aSpy.HadErrorContaining(t, "got len=<0>, wanted <2>")
}

func Test_Map_Len_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]int{"hi": 1}).Len(1))
}

func Test_Map_WithItem_value_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]int{"hi": 1}).WithItem("hi", 2))
	aSpy.HadErrorContaining(t, "got value=<1> for key=<hi>, wanted <2>")
}

func Test_Map_WithItem_key_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]int{"hi": 1}).WithItem("bye", 2))
	aSpy.HadErrorContaining(t, "got key absent, wanted value for key <bye>")
}

func Test_Map_HasItem_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]int{"hi": 1}).WithItem("hi", 1))
}
