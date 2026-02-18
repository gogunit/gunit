package hammy_test

import (
	"testing"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

func Test_Slice_Len_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.Slice([]int{1, 2, 3}).Len(2))
	aSpy.HadErrorContaining(t, "got len()=3, wanted 2")
}

func Test_Slice_Len_success(t *testing.T) {
	assert := a.New(t)
	assert.Is(a.Slice([]int{1, 2, 3}).Len(3))
}

func Test_Slice_IsEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.Slice([]string{"hello"}).IsEmpty())
	aSpy.HadErrorContaining(t, "got len()=1, wanted 0")
}

func Test_Slice_IsEmpty_success(t *testing.T) {
	assert := a.New(t)
	assert.Is(a.Slice([]int{}).IsEmpty())
}

func Test_Slice_NotEmpty_success(t *testing.T) {
	assert := a.New(t)
	assert.Is(a.Slice([]int{1}).NotEmpty())
}

func Test_Slice_NotEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.Slice([]int{}).NotEmpty())
	aSpy.HadErrorContaining(t, "got len()=0, wanted > 0")
}

func Test_Slice_Contains_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.Slice([]int{1, 2, 3}).Contains(2, 3, 4))
	aSpy.HadErrorContaining(t, "got 1 unmatched items, wanted array containing the 3 items. Items at index 2 were missing")
}

func Test_Slice_Contains_success(t *testing.T) {
	assert := a.New(t)
	assert.Is(a.Slice([]int{1, 2, 3}).Contains(2, 3))
}

func Test_Slice_NotContains_success(t *testing.T) {
	assert := a.New(t)
	assert.Is(a.Slice([]int{1, 2, 3}).NotContains(4, 5))
}

func Test_Slice_NotContains_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.Slice([]int{1, 2, 3}).NotContains(2, 4))
	aSpy.HadErrorContaining(t, "got items at expected index 0 present in slice, wanted all absent")
}

func Test_Slice_EqualTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.Slice([]int{1, 2, 3}).EqualTo(2, 3, 4))
	aSpy.HadError(t)
}

func Test_Slice_EqualTo_string_success(t *testing.T) {
	assert := a.New(t)
	assert.Is(a.Slice([]string{"hello"}).EqualTo("hello"))
}

func Test_Slice_EqualTo_success(t *testing.T) {
	assert := a.New(t)
	assert.Is(a.Slice([]int{1, 2, 3}).EqualTo(1, 2, 3))
}

func Test_Slice_ContainsExactly_success(t *testing.T) {
	assert := a.New(t)
	assert.Is(a.Slice([]int{3, 2, 1}).ContainsExactly(1, 2, 3))
}

func Test_Slice_ContainsExactly_failure_with_differing_items(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.Slice([]int{2, 3, 4}).ContainsExactly(1, 2, 3))
	aSpy.HadError(t)
}

func Test_Slice_ContainsExactly_failure_with_differing_length(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.Slice([]int{2, 3, 4}).ContainsExactly(2, 3))
	aSpy.HadError(t)
}
