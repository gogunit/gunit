package gunit

import "github.com/google/go-cmp/cmp"

type Mappy[K comparable, V any] struct {
	T
	actual map[K]V
}

func Map[K comparable, V any](t T, actual map[K]V) *Mappy[K, V] {
	return &Mappy[K, V]{t, actual}
}

func (m *Mappy[K, V]) EqualTo(expected map[K]V) {
	m.Helper()
	if diff := cmp.Diff(expected, m.actual); diff != "" {
		m.Errorf("map mismatch (-want +got):\n%s", diff)
	}
}

func (m *Mappy[K, V]) IsEmpty() {
	m.Helper()
	if len(m.actual) != 0 {
		m.Errorf("want empty map, got len() = %v", len(m.actual))
	}
}

func (m *Mappy[K, V]) WithKeys(keys ...K) {
	m.Helper()
	var missing []K
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

func (m *Mappy[K, V]) WithoutKeys(keys ...K) {
	m.Helper()
	var present []K
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

func (m *Mappy[K, V]) WithValues(values ...V) {
	m.Helper()
	var missing []V
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

// TODO: Implement Includes() or Contains() for subset of pairs?

func (m *Mappy[K, V]) HasKey(key K) {
	m.Helper()
	if _, ok := m.actual[key]; !ok {
		m.Errorf("got key absent <%v>, wanted present in map", key)
	}
}

func (m *Mappy[K, V]) NotHasKey(key K) {
	m.Helper()
	if _, ok := m.actual[key]; ok {
		m.Errorf("got key present <%v>, wanted absent from map", key)
	}
}

func (m *Mappy[K, V]) KeysExactly(expected ...K) {
	m.Helper()
	actualKeys := make(map[K]bool, len(m.actual))
	for k := range m.actual {
		actualKeys[k] = true
	}
	expectedKeys := make(map[K]bool, len(expected))
	for _, k := range expected {
		expectedKeys[k] = true
	}
	var missing, extra []K
	for k := range expectedKeys {
		if !actualKeys[k] {
			missing = append(missing, k)
		}
	}
	for k := range actualKeys {
		if !expectedKeys[k] {
			extra = append(extra, k)
		}
	}
	if len(missing) > 0 || len(extra) > 0 {
		m.Errorf("got extra keys <%v> and missing keys <%v>, wanted exact key set", extra, missing)
	}
}

func (m *Mappy[K, V]) NotEmpty() {
	m.Helper()
	if len(m.actual) == 0 {
		m.Errorf("got len=<0>, wanted non-empty map")
	}
}

func (m *Mappy[K, V]) IsNotEmpty() { m.NotEmpty() }

func (m *Mappy[K, V]) NotContains(values ...V) {
	m.Helper()
	var present []V
	seen := make(map[int]bool)
	for _, actual := range m.actual {
		for i, expected := range values {
			if cmp.Equal(actual, expected) {
				if !seen[i] {
					present = append(present, expected)
				}
				seen[i] = true
			}
		}
	}
	if len(present) > 0 {
		m.Errorf("got values <%v>, wanted absent from map", present)
	}
}

func (m *Mappy[K, V]) Len(expected int) {
	m.Helper()
	if len(m.actual) != expected {
		m.Errorf("got len=<%v>, wanted <%v>", len(m.actual), expected)
	}
}

func (m *Mappy[K, V]) WithItem(k K, expected V) {
	m.Helper()
	actual, ok := m.actual[k]
	if !ok {
		m.Errorf("got key absent, wanted value for key <%v>", k)
		return
	}
	if !cmp.Equal(actual, expected) {
		m.Errorf("got value=<%v> for key=<%v>, wanted <%v>", actual, k, expected)
	}
}

func (m *Mappy[K, V]) WithItems(expected map[K]V) {
	m.Helper()
	var missing, mismatched []K
	for key, expectedValue := range expected {
		actualValue, ok := m.actual[key]
		if !ok {
			missing = append(missing, key)
			continue
		}
		if !cmp.Equal(actualValue, expectedValue) {
			mismatched = append(mismatched, key)
		}
	}
	if len(missing) > 0 || len(mismatched) > 0 {
		m.Errorf("got missing keys <%v> and mismatched keys <%v>, wanted entries <%v>", missing, mismatched, expected)
	}
}

func (m *Mappy[K, V]) WithoutItems(expected map[K]V) {
	m.Helper()
	for key, expectedValue := range expected {
		if actualValue, ok := m.actual[key]; ok && cmp.Equal(actualValue, expectedValue) {
			m.Errorf("got all entries <%v>, wanted at least one absent or different", expected)
			return
		}
	}
}
