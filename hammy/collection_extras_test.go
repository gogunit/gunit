package hammy_test

import (
	"testing"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

func Test_OneOf_success(t *testing.T) {
	assert := a.New(t)

	assert.Is(a.Match("beta", a.OneOf("alpha", "beta", "gamma")))
}

func Test_OneOf_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.Match("delta", a.OneOf("alpha", "beta", "gamma")))

	aSpy.HadErrorContaining(t, "wanted one of <[alpha beta gamma]>")
}

func Test_Slice_Cap_success(t *testing.T) {
	assert := a.New(t)
	actual := make([]int, 0, 4)

	assert.Is(a.Slice(actual).Cap(4))
}

func Test_Slice_Cap_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	actual := make([]int, 0, 4)

	assert.Is(a.Slice(actual).Cap(3))

	aSpy.HadErrorContaining(t, "got cap()=4, wanted 3")
}

func Test_Slice_ContainsAny_success(t *testing.T) {
	assert := a.New(t)

	assert.Is(a.Slice([]int{1, 2, 3}).ContainsAny(5, 3))
}

func Test_Slice_ContainsAny_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.Slice([]int{1, 2, 3}).ContainsAny(4, 5))

	aSpy.HadErrorContaining(t, "got no matching item, wanted any of <[4 5]>")
}

func Test_Slice_SubsetOf_success(t *testing.T) {
	assert := a.New(t)

	assert.Is(a.Slice([]int{1, 3}).SubsetOf(1, 2, 3))
}

func Test_Slice_SubsetOf_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.Slice([]int{1, 4}).SubsetOf(1, 2, 3))

	aSpy.HadErrorContaining(t, "got items outside expected set <[index 1: 4]>")
}

func Test_Slice_NotSubsetOf_success(t *testing.T) {
	assert := a.New(t)

	assert.Is(a.Slice([]int{1, 4}).NotSubsetOf(1, 2, 3))
}

func Test_Slice_NotSubsetOf_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.Slice([]int{1, 3}).NotSubsetOf(1, 2, 3))

	aSpy.HadErrorContaining(t, "wanted at least one item outside <[1 2 3]>")
}

func Test_Map_WithItems_success(t *testing.T) {
	assert := a.New(t)
	actual := map[string]int{"alpha": 1, "beta": 2, "gamma": 3}

	assert.Is(a.Map(actual).WithItems(map[string]int{"alpha": 1, "gamma": 3}))
}

func Test_Map_WithItems_failure_missing_key(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	actual := map[string]int{"alpha": 1}

	assert.Is(a.Map(actual).WithItems(map[string]int{"beta": 2}))

	aSpy.HadErrorContaining(t, "got missing keys <[beta]>")
}

func Test_Map_WithItems_failure_mismatched_value(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	actual := map[string]int{"alpha": 1}

	assert.Is(a.Map(actual).WithItems(map[string]int{"alpha": 2}))

	aSpy.HadErrorContaining(t, "mismatched keys <[alpha]>")
}

func Test_Map_WithoutItems_success_missing_key(t *testing.T) {
	assert := a.New(t)
	actual := map[string]int{"alpha": 1}

	assert.Is(a.Map(actual).WithoutItems(map[string]int{"beta": 2}))
}

func Test_Map_WithoutItems_success_different_value(t *testing.T) {
	assert := a.New(t)
	actual := map[string]int{"alpha": 1}

	assert.Is(a.Map(actual).WithoutItems(map[string]int{"alpha": 2}))
}

func Test_Map_WithoutItems_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	actual := map[string]int{"alpha": 1, "beta": 2}

	assert.Is(a.Map(actual).WithoutItems(map[string]int{"alpha": 1}))

	aSpy.HadErrorContaining(t, "got all entries <map[alpha:1]>")
}
