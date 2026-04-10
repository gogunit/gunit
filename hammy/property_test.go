package hammy_test

import (
	"testing"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

type profile struct {
	Name string
	Meta meta
}

type meta struct {
	Age int
}

func Test_Having_success(t *testing.T) {
	person := profile{Name: "Ada", Meta: meta{Age: 37}}
	a.New(t).Is(a.Match(person, a.Having(func(actual profile) int {
		return actual.Meta.Age
	}, a.GreaterThan(30))))
}

func Test_HavingField_failure_reports_field_name(t *testing.T) {
	aSpy := eye.Spy()
	person := profile{Name: "Ada", Meta: meta{Age: 37}}
	a.New(aSpy).Is(a.Match(person, a.HavingField("Name", func(actual profile) string {
		return actual.Name
	}, a.EqualTo("Grace"))))
	aSpy.HadErrorContaining(t, "field Name")
}

func Test_Having_composes_for_nested_projection(t *testing.T) {
	person := profile{Name: "Ada", Meta: meta{Age: 37}}
	a.New(t).Is(a.Match(person, a.Having(func(actual profile) meta {
		return actual.Meta
	}, a.HavingField("Age", func(actual meta) int {
		return actual.Age
	}, a.GreaterOrEqual(37)))))
}
