package gunit_test

import (
	"github.com/gogunit/gunit"
	"github.com/gogunit/gunit/eye"
	"testing"
)

func Test_map_HasKey_success(t *testing.T) { gunit.Map(t, testMap()).HasKey("a") }
func Test_map_HasKey_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Map(aSpy, testMap()).HasKey("c")
	aSpy.HadErrorContaining(t, "wanted present in map")
}

func Test_map_NotHasKey_success(t *testing.T) { gunit.Map(t, testMap()).NotHasKey("c") }
func Test_map_NotHasKey_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Map(aSpy, testMap()).NotHasKey("a")
	aSpy.HadErrorContaining(t, "wanted absent from map")
}

func Test_map_KeysExactly_success(t *testing.T) { gunit.Map(t, testMap()).KeysExactly("a", "b") }
func Test_map_KeysExactly_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Map(aSpy, testMap()).KeysExactly("a", "c")
	aSpy.HadErrorContaining(t, "wanted exact key set")
}

func Test_map_NotEmpty_success(t *testing.T) { gunit.Map(t, testMap()).NotEmpty() }
func Test_map_NotEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Map(aSpy, map[string]int{}).NotEmpty()
	aSpy.HadErrorContaining(t, "wanted non-empty map")
}

func Test_map_IsNotEmpty_success(t *testing.T) { gunit.Map(t, testMap()).IsNotEmpty() }
func Test_map_IsNotEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Map(aSpy, map[string]int{}).IsNotEmpty()
	aSpy.HadErrorContaining(t, "wanted non-empty map")
}

func Test_map_NotContains_success(t *testing.T) { gunit.Map(t, testMap()).NotContains(3) }
func Test_map_NotContains_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Map(aSpy, testMap()).NotContains(1)
	aSpy.HadErrorContaining(t, "wanted absent from map")
}

func Test_map_Len_success(t *testing.T) { gunit.Map(t, testMap()).Len(2) }
func Test_map_Len_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Map(aSpy, testMap()).Len(3)
	aSpy.HadErrorContaining(t, "wanted <3>")
}

func Test_map_WithItem_success(t *testing.T) { gunit.Map(t, testMap()).WithItem("a", 1) }
func Test_map_WithItem_failure_for_missing_key(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Map(aSpy, testMap()).WithItem("c", 3)
	aSpy.HadErrorContaining(t, "got key absent")
}
func Test_map_WithItem_failure_for_mismatched_value(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Map(aSpy, testMap()).WithItem("a", 3)
	aSpy.HadErrorContaining(t, "wanted <3>")
}

func Test_map_WithItems_success(t *testing.T) {
	gunit.Map(t, testMap()).WithItems(map[string]int{"a": 1})
}
func Test_map_WithItems_failure_for_missing_key(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Map(aSpy, testMap()).WithItems(map[string]int{"c": 3})
	aSpy.HadErrorContaining(t, "missing keys")
}
func Test_map_WithItems_failure_for_mismatched_value(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Map(aSpy, testMap()).WithItems(map[string]int{"a": 3})
	aSpy.HadErrorContaining(t, "mismatched keys")
}

func Test_map_WithoutItems_success(t *testing.T) {
	gunit.Map(t, testMap()).WithoutItems(map[string]int{"a": 2})
}
func Test_map_WithoutItems_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Map(aSpy, testMap()).WithoutItems(map[string]int{"a": 1})
	aSpy.HadErrorContaining(t, "wanted at least one absent")
}
