package gunit_test

import (
	"testing"

	"github.com/gogunit/gunit"
	"github.com/gogunit/gunit/eye"
)

type sizedItem struct {
	length int
}

func (s sizedItem) Len() int {
	return s.length
}

func Test_Length_EqualTo_success(t *testing.T) {
	gunit.Length(t, sizedItem{length: 3}).EqualTo(3)
}

func Test_Length_LessThan_success(t *testing.T) {
	gunit.Length(t, sizedItem{length: 3}).LessThan(4)
}

func Test_Length_GreaterThan_success(t *testing.T) {
	gunit.Length(t, sizedItem{length: 3}).GreaterThan(2)
}

func Test_Length_NotEqualTo_success(t *testing.T) {
	gunit.Length(t, sizedItem{length: 3}).NotEqualTo(2)
}

func Test_Length_GreaterOrEqual_success(t *testing.T) {
	gunit.Length(t, sizedItem{length: 3}).GreaterOrEqual(3)
}

func Test_Length_LessOrEqual_success(t *testing.T) {
	gunit.Length(t, sizedItem{length: 3}).LessOrEqual(3)
}

func Test_Length_Within_success(t *testing.T) {
	gunit.Length(t, sizedItem{length: 3}).Within(4, 1)
}

func Test_Length_IsZero_success(t *testing.T) {
	gunit.Length(t, sizedItem{}).IsZero()
}

func Test_Length_GreaterThan_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Length(aSpy, sizedItem{length: 3}).GreaterThan(4)
	aSpy.HadErrorContaining(t, "got len()=3, wanted greater than 4")
}

func Test_Length_nil_failure(t *testing.T) {
	aSpy := eye.Spy()
	var actual *pointerSizedItem

	gunit.Length(aSpy, actual).EqualTo(3)
	aSpy.HadErrorContaining(t, "got nil length-bounded item, wanted len()=3")
}

type pointerSizedItem struct {
	length int
}

func (s *pointerSizedItem) Len() int {
	return s.length
}
