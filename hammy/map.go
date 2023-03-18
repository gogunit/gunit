package hammy

import (
	"github.com/google/go-cmp/cmp"
)

func Map[K comparable, V any](actual map[K]V) *Mappy[K, V] {
	return &Mappy[K, V]{
		actual: actual,
	}
}

type Mappy[K comparable, V any] struct {
	actual map[K]V
}

func (m Mappy[K, V]) WithKeys(keys ...K) AssertionMessage {
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
	return Assert(len(missing) == 0, "want map with keys <%v>, but missing", missing)
}

func (m Mappy[K, V]) IsEmpty() AssertionMessage {
	return Assert(len(m.actual) == 0, "want empty map, got len() = %v", len(m.actual))
}

func (m Mappy[K, V]) WithValues(values ...V) AssertionMessage {
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
	return Assert(len(missing) == 0, "want map with values <%v>, but missing", missing)
}

func (m Mappy[K, V]) WithoutKeys(keys ...K) AssertionMessage {
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
	return Assert(len(present) == 0, "want map without keys <%v>, but present", present)
}

func (m Mappy[K, V]) Len(expected int) AssertionMessage {
	sz := len(m.actual)
	return Assert(sz == expected, "want len of <%v>, got <%v>", sz, expected)
}

func (m Mappy[K, V]) WithItem(k K, expected V) AssertionMessage {
	actual, ok := m.actual[k]
	if !ok {
		return Assert(false, "want key=<%v>, but was absent", k)
	}
	return Assert(cmp.Equal(actual, expected), "want value=<%v> for key=<hi>, got <%v>", expected, actual)
}

func (m Mappy[K, V]) EqualTo(expected map[K]V) AssertionMessage {
	diff := cmp.Diff(expected, m.actual)
	return Assert(diff == "", "Map mismatch (-want +got):\n%s", diff)
}
