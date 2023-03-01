package gunit

import "strings"

func String(t T, actual string) *Str {
	return &Str{t, actual}
}

type Str struct {
	T
	actual string
}

func (s *Str) EqualTo(expected string) {
	s.Helper()
	if s.actual != expected {
		s.Errorf("want <%v> equal to <%v>", s.actual, expected)
	}
}

func (s *Str) Contains(needle string) {
	s.Helper()
	if !strings.Contains(s.actual, needle) {
		s.Errorf("want <%v> to contain <%v>", s.actual, needle)
	}
}

func (s *Str) HasPrefix(prefix string) {
	s.Helper()
	if !strings.HasPrefix(s.actual, prefix) {
		s.Errorf("want <%v> to have prefix <%v>", s.actual, prefix)
	}
}

func (s *Str) HasSuffix(suffix string) {
	s.Helper()
	if !strings.HasSuffix(s.actual, suffix) {
		s.Errorf("want <%v> to have prefix <%v>", s.actual, suffix)
	}
}

func (s *Str) IsEmpty() {
	s.Helper()
	if s.actual != "" {
		s.Errorf("want <%v> to be empty, was not", s.actual)
	}
}

func (s *Str) IsNotEmpty() {
	s.Helper()
	if s.actual == "" {
		s.Errorf("want <%v> to be empty, was not", s.actual)
	}
}
