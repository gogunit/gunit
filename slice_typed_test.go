package gunit_test

import (
	"github.com/gogunit/gunit"
	"github.com/gogunit/gunit/eye"
	"testing"
)

func Test_slice_Contains_success(t *testing.T) {
	gunit.Slice(t, []int{1, 2, 3}).Contains(2, 3)
}
func Test_slice_Contains_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Slice(aSpy, []int{1}).Contains(2)
	aSpy.HadErrorContaining(t, "missing items")
}

func Test_slice_ContainsAny_success(t *testing.T) {
	gunit.Slice(t, []int{1, 2, 3}).ContainsAny(4, 3)
}
func Test_slice_ContainsAny_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Slice(aSpy, []int{1}).ContainsAny(2)
	aSpy.HadErrorContaining(t, "no matching item")
}

func Test_slice_NotContains_success(t *testing.T) {
	gunit.Slice(t, []int{1, 2, 3}).NotContains(4)
}
func Test_slice_NotContains_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Slice(aSpy, []int{1}).NotContains(1)
	aSpy.HadErrorContaining(t, "present in slice")
}

func Test_slice_EqualTo_success(t *testing.T) {
	gunit.Slice(t, []int{1, 2, 3}).EqualTo(1, 2, 3)
}
func Test_slice_EqualTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Slice(aSpy, []int{1}).EqualTo(2)
	aSpy.HadErrorContaining(t, "slice mismatch")
}

func Test_slice_Len_success(t *testing.T) {
	gunit.Slice(t, make([]int, 2, 4)).Len(2)
}
func Test_slice_Len_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Slice(aSpy, []int{1}).Len(2)
	aSpy.HadErrorContaining(t, "got len()=1")
}

func Test_slice_Cap_success(t *testing.T) {
	gunit.Slice(t, make([]int, 2, 4)).Cap(4)
}
func Test_slice_Cap_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Slice(aSpy, make([]int, 1, 2)).Cap(1)
	aSpy.HadErrorContaining(t, "got cap()=2")
}

func Test_slice_IsEmpty_success(t *testing.T) {
	gunit.Slice(t, []int{}).IsEmpty()
}
func Test_slice_IsEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Slice(aSpy, []int{1}).IsEmpty()
	aSpy.HadErrorContaining(t, "wanted 0")
}

func Test_slice_IsNotEmpty_success(t *testing.T) {
	gunit.Slice(t, []int{1}).IsNotEmpty()
}
func Test_slice_IsNotEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Slice(aSpy, []int{}).IsNotEmpty()
	aSpy.HadErrorContaining(t, "wanted > 0")
}

func Test_slice_NotEmpty_success(t *testing.T) {
	gunit.Slice(t, []int{1}).NotEmpty()
}
func Test_slice_NotEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Slice(aSpy, []int{}).NotEmpty()
	aSpy.HadErrorContaining(t, "wanted > 0")
}

func Test_slice_ContainsExactly_success(t *testing.T) {
	gunit.Slice(t, []int{1, 2}).ContainsExactly(2, 1)
}
func Test_slice_ContainsExactly_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Slice(aSpy, []int{1}).ContainsExactly(1, 2)
	aSpy.HadErrorContaining(t, "length mismatch")
}

func Test_slice_SubsetOf_success(t *testing.T) {
	gunit.Slice(t, []int{1, 2}).SubsetOf(1, 2, 3)
}
func Test_slice_SubsetOf_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Slice(aSpy, []int{1, 4}).SubsetOf(1, 2, 3)
	aSpy.HadErrorContaining(t, "outside expected set")
}

func Test_slice_NotSubsetOf_success(t *testing.T) {
	gunit.Slice(t, []int{1, 4}).NotSubsetOf(1, 2, 3)
}
func Test_slice_NotSubsetOf_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Slice(aSpy, []int{1, 2}).NotSubsetOf(1, 2, 3)
	aSpy.HadErrorContaining(t, "got subset")
}
