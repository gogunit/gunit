package hammy_test

import (
	"testing"

	"github.com/gogunit/gunit/eye"
	"github.com/gogunit/gunit/hammy"
)

func Test_Slice_Len_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Slice([]int{1, 2, 3}).Len(2))
	aSpy.HadErrorContaining(t, "got len()=3, wanted 2")
}

func Test_Slice_Len_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Slice([]int{1, 2, 3}).Len(3))
}

func Test_Slice_IsEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Slice([]string{"hello"}).IsEmpty())
	aSpy.HadErrorContaining(t, "got len()=1, wanted 0")
}

func Test_Slice_IsEmpty_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Slice([]int{}).IsEmpty())
}

func Test_Slice_Contains_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Slice([]int{1, 2, 3}).Contains(2, 3, 4))
	aSpy.HadErrorContaining(t, "got 1 unmatched items, wanted array containing the 3 items. Items at index 2 were missing")
}

func Test_Slice_Contains_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Slice([]int{1, 2, 3}).Contains(2, 3))
}

func Test_Slice_EqualTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Slice([]int{1, 2, 3}).EqualTo(2, 3, 4))
	aSpy.HadError(t)
}

func Test_Slice_EqualTo_string_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Slice([]string{"hello"}).EqualTo("hello"))
}

func Test_Slice_EqualTo_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Slice([]int{1, 2, 3}).EqualTo(1, 2, 3))
}
