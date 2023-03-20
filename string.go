package gunit

import "strings"

type Stringy interface {
	~string
}

func String[S Stringy](t T, actual S) *Str[S] {
	return &Str[S]{t, actual}
}

type Str[S Stringy] struct {
	T
	actual S
}

func (s *Str[S]) EqualTo(expected S) {
	s.Helper()
	if s.actual != expected {
		s.Errorf("wanted <%v> equal to <%v>", s.actual, expected)
	}
}

func (s *Str[S]) Contains(needle S) {
	s.Helper()
	if !strings.Contains(string(s.actual), string(needle)) {
		s.Errorf("want <%v> to contain <%v>", s.actual, needle)
	}
}

func (s *Str[S]) HasPrefix(prefix S) {
	s.Helper()
	if !strings.HasPrefix(string(s.actual), string(prefix)) {
		s.Errorf("want <%v> to have prefix <%v>", s.actual, prefix)
	}
}

func (s *Str[S]) HasSuffix(suffix S) {
	s.Helper()
	if !strings.HasSuffix(string(s.actual), string(suffix)) {
		s.Errorf("want <%v> to have prefix <%v>", s.actual, suffix)
	}
}

func (s *Str[S]) IsEmpty() {
	s.Helper()
	if s.actual != "" {
		s.Errorf("want <%v> to be empty, was not", s.actual)
	}
}

func (s *Str[S]) IsNotEmpty() {
	s.Helper()
	if s.actual == "" {
		s.Errorf("want <%v> to be empty, was not", s.actual)
	}
}
