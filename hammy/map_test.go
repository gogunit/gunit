package hammy_test

import (
	"github.com/gogunit/gunit/eye"
	"github.com/gogunit/gunit/hammy"
	"testing"
)

func Test_map_WithKeys_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]bool{"abc": true, "def": true}).WithKeys("abc", "ghi"))
	aSpy.HadError(t)
}

func Test_map_WithKeys_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]bool{"abc": true, "def": true}).WithKeys("abc", "def"))
}

func Test_map_IsEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]bool{"abc": true, "def": true}).IsEmpty())
	aSpy.HadError(t)
}

func Test_map_IsEmpty_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]int{}).IsEmpty())
}

func Test_map_WithValues_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]int{"abc": 42, "def": 33}).WithValues(4, 3))
	aSpy.HadError(t)
}

func Test_map_WithValues_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]int{"abc": 42, "def": 33}).WithValues(42, 33))
}

func Test_map_WithoutKeys_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]int{"abc": 42, "def": 33}).WithoutKeys("abc", "def", "jkl"))
	aSpy.HadError(t)
}

func Test_map_WithoutKeys_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]int{"abc": 42, "def": 33}).WithoutKeys("ghi", "jkl"))
}

func Test_map_Len_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Map(map[string]int{}).Len(2))
	aSpy.HadError(t)
}

func Test_map_Len_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Map(map[string]int{"hi": 1}).Len(1))
}
