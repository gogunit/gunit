package gunit_test

import (
	. "github.com/nfisher/gunit"
	"testing"
)

func Test_map_HasKeys_success(t *testing.T) {
	Map(t, map[string]bool{"abc": true, "def": true}).HasKeys("abc", "def")
}

func Test_map_HasKeys_failure(t *testing.T) {
	aSpy := spy()
	Map(aSpy, map[string]bool{"abc": true, "def": true}).HasKeys("abc", "ghi")
	aSpy.HadError(t)
}

func Test_map_HasValues_success(t *testing.T) {
	Map(t, map[string]int{"abc": 42, "def": 33}).HasValues(42, 33)
}

func Test_map_HasValues_failure(t *testing.T) {
	aSpy := spy()
	Map(aSpy, map[string]int{"abc": 42, "def": 33}).HasValues(4, 3)
	aSpy.HadError(t)
}

func Test_map_IsEmpty_success(t *testing.T) {
	Map(t, map[string]int{}).IsEmpty()
}

func Test_map_IsEmpty_failure(t *testing.T) {
	aSpy := spy()
	Map(aSpy, map[string]int{"abc": 42, "def": 33}).IsEmpty()
	aSpy.HadError(t)
}

func Test_map_EqualTo_success(t *testing.T) {
	Map(t, map[string]int{"abc": 42, "def": 33}).EqualTo(map[string]int{"abc": 42, "def": 33})
}

func Test_map_EqualTo_failure(t *testing.T) {
	aSpy := spy()
	Map(aSpy, map[string]int{"abc": 42, "def": 33}).EqualTo(map[string]int{})
	aSpy.HadError(t)
}
