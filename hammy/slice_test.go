package hammy_test

import (
	"github.com/nfisher/gunit/eye"
	"github.com/nfisher/gunit/hammy"
	"testing"
)

func Test_Slice_Len_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Slice([]int{1, 2, 3}).Len(2))
	aSpy.HadError(t)
}

func Test_Slice_Len_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Slice([]int{1, 2, 3}).Len(3))
}

func Test_Slice_Contains_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Slice([]int{1, 2, 3}).Contains(2, 3, 4))
	aSpy.HadError(t)
}

func Test_Slice_Contains_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.Slice([]int{1, 2, 3}).Contains(2, 3))
}
