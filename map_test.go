package gunit_test

import (
	. "github.com/nfisher/gunit"
	. "github.com/nfisher/gunit/eye"
	"testing"
)

func Test_map_WithKeys_success(t *testing.T) {
	Map(t, map[string]bool{"abc": true, "def": true}).WithKeys("abc", "def")
}

func Test_map_WithKeys_failure(t *testing.T) {
	aSpy := Spy()
	Map(aSpy, map[string]bool{"abc": true, "def": true}).WithKeys("abc", "ghi")
	aSpy.HadError(t)
}

func Test_map_WithValues_success(t *testing.T) {
	Map(t, map[string]int{"abc": 42, "def": 33}).WithValues(42, 33)
}

func Test_map_WithValues_failure(t *testing.T) {
	aSpy := Spy()
	Map(aSpy, map[string]int{"abc": 42, "def": 33}).WithValues(4, 3)
	aSpy.HadError(t)
}

func Test_map_IsEmpty_success(t *testing.T) {
	Map(t, map[string]int{}).IsEmpty()
}

func Test_map_IsEmpty_failure(t *testing.T) {
	aSpy := Spy()
	Map(aSpy, map[string]int{"abc": 42, "def": 33}).IsEmpty()
	aSpy.HadError(t)
}

func Test_map_EqualTo_success(t *testing.T) {
	Map(t, map[string]int{"abc": 42, "def": 33}).EqualTo(map[string]int{"abc": 42, "def": 33})
}

func Test_map_EqualTo_failure(t *testing.T) {
	aSpy := Spy()
	Map(aSpy, map[string]int{"abc": 42, "def": 33}).EqualTo(map[string]int{})
	aSpy.HadError(t)
}

func Test_map_WithoutKeys_success(t *testing.T) {
	Map(t, map[string]int{"abc": 42, "def": 33}).WithoutKeys("ghi", "jkl")
}

func Test_map_WithoutKeys_failure(t *testing.T) {
	aSpy := Spy()
	Map(aSpy, map[string]int{"abc": 42, "def": 33}).WithoutKeys("abc", "def", "jkl")
	aSpy.HadError(t)
}
