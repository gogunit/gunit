package gunit

import "github.com/google/go-cmp/cmp"

type Mappy[K comparable, V any] struct {
	T
	actual map[K]V
}

func Map[K comparable, V any](t T, actual map[K]V) *Mappy[K, V] {
	return &Mappy[K, V]{t, actual}
}

func (m *Mappy[K, V]) WithKeys(keys ...K) {
	m.Helper()
	missing := []K{}
	has := make(map[K]bool, len(m.actual))
	hasNot := make(map[K]bool)
	for k := range m.actual {
		has[k] = true
	}
	for _, k := range keys {
		if !has[k] {
			if !hasNot[k] {
				missing = append(missing, k)
			}
			hasNot[k] = true
		}
	}
	if len(missing) > 0 {
		m.Errorf("want map with keys <%v>, but missing", missing)
	}
}

func (m *Mappy[K, V]) WithValues(values ...V) {
	m.Helper()
	missing := []V{}
	present := make(map[int]bool)
	for _, actual := range m.actual {
		for i, expected := range values {
			if cmp.Equal(actual, expected) {
				present[i] = true
				break
			}
		}
	}
	for i, v := range values {
		if !present[i] {
			missing = append(missing, v)
		}
	}
	if len(missing) > 0 {
		m.Errorf("want map with values <%v>, but missing", missing)
	}
}

func (m *Mappy[K, V]) IsEmpty() {
	m.Helper()
	if len(m.actual) != 0 {
		m.Errorf("want empty map, got len() = %v", len(m.actual))
	}
}

func (m *Mappy[K, V]) EqualTo(expected map[K]V) {
	m.Helper()
	if diff := cmp.Diff(expected, m.actual); diff != "" {
		m.Errorf("map mismatch (-want +got):\n%s", diff)
	}
}

func (m *Mappy[K, V]) WithoutKeys(keys ...K) {
	m.Helper()
	present := []K{}
	has := make(map[K]bool, len(m.actual))
	added := make(map[K]bool)
	for k := range m.actual {
		has[k] = true
	}
	for _, k := range keys {
		if has[k] {
			if !added[k] {
				present = append(present, k)
			}
			added[k] = true
		}
	}
	if len(present) > 0 {
		m.Errorf("want map without keys <%v>, but present", present)
	}
}

// TODO: Implement Includes() or Contains() for subset of pairs?
