package hammy_test

import (
	"math"
	"testing"

	a "github.com/gogunit/gunit/hammy"
)

func Test_Number_Matches_success(t *testing.T) {
	a.New(t).Is(a.Number(5).Matches(a.AllOf(
		a.GreaterThan(0),
		a.LessThan(10),
	)))
}

func Test_String_Matches_success(t *testing.T) {
	a.New(t).Is(a.String("hello").Matches(a.EqualIgnoringCase("HELLO")))
}

func Test_Slice_Matches_success(t *testing.T) {
	a.New(t).Is(a.Slice([]int{2, 1}).Matches(a.ContainsInAnyOrder(
		a.EqualTo(1),
		a.EqualTo(2),
	)))
}

func Test_Map_Matches_success(t *testing.T) {
	a.New(t).Is(a.Map(map[string]int{"alpha": 1, "beta": 2}).Matches(a.HasEntry(
		a.EqualTo("beta"),
		a.GreaterThan(1),
	)))
}

func Test_Struct_Matches_success(t *testing.T) {
	type sample struct {
		Age int
	}

	a.New(t).Is(a.Struct(sample{Age: 37}).Matches(a.HavingField("Age", func(actual sample) int {
		return actual.Age
	}, a.GreaterThan(30))))
}

func Test_Float_Matches_success(t *testing.T) {
	a.New(t).Is(a.Float(math.NaN()).Matches(a.IsNaN[float64]()))
}
