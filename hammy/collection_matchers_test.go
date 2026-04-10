package hammy_test

import (
	"testing"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

func Test_Slice_Every_success(t *testing.T) {
	a.New(t).Is(a.Slice([]int{2, 4, 6}).Every(a.GreaterThan(1)))
}

func Test_Slice_Every_failure_reports_index(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Slice([]int{2, 0, 6}).Every(a.GreaterThan(1)))
	aSpy.HadErrorContaining(t, "item at index 1")
}

func Test_Slice_HasItem_success(t *testing.T) {
	a.New(t).Is(a.Slice([]string{"alpha", "beta"}).HasItem(a.HasPrefix("bet")))
}

func Test_Slice_HasItem_failure_reports_sample(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Slice([]string{"alpha", "beta"}).HasItem(a.EqualTo("gamma")))
	aSpy.HadErrorContaining(t, "got no matching item in slice")
}

func Test_Slice_ContainsInOrder_success(t *testing.T) {
	a.New(t).Is(a.Slice([]string{"alpha", "beta"}).ContainsInOrder(
		a.EqualTo("alpha"),
		a.HasSuffix("ta"),
	))
}

func Test_Slice_ContainsInOrder_failure_reports_index(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Slice([]int{1, 2}).ContainsInOrder(
		a.EqualTo(1),
		a.GreaterThan(3),
	))
	aSpy.HadErrorContaining(t, "item at index 1")
}

func Test_Slice_ContainsInAnyOrder_handles_duplicates(t *testing.T) {
	a.New(t).Is(a.Slice([]int{2, 1, 2}).ContainsInAnyOrder(
		a.EqualTo(2),
		a.EqualTo(2),
		a.EqualTo(1),
	))
}

func Test_Slice_ContainsInAnyOrder_failure_reports_matcher_index(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Slice([]int{1, 2, 3}).ContainsInAnyOrder(
		a.EqualTo(1),
		a.EqualTo(2),
		a.EqualTo(4),
	))
	aSpy.HadErrorContaining(t, "matcher at index 2")
}

func Test_Map_HasEntry_success(t *testing.T) {
	a.New(t).Is(a.Map(map[string]int{"alpha": 1, "beta": 2}).HasEntry(
		a.EqualTo("beta"),
		a.GreaterThan(1),
	))
}

func Test_Map_HasEntry_failure_reports_samples(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Map(map[string]int{"alpha": 1}).HasEntry(
		a.EqualTo("beta"),
		a.GreaterThan(1),
	))
	aSpy.HadErrorContaining(t, "got no matching map entry")
}

func Test_Map_HasKeyMatching_success(t *testing.T) {
	a.New(t).Is(a.Map(map[string]int{"alpha": 1}).HasKeyMatching(a.HasSuffix("pha")))
}

func Test_Map_HasValueMatching_success(t *testing.T) {
	a.New(t).Is(a.Map(map[string]int{"alpha": 3}).HasValueMatching(a.GreaterThan(2)))
}
