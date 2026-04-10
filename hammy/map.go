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

func (m Mappy[K, V]) WithKeys(expected ...K) AssertionMessage {
	var found []K
	var missing []K
	has := make(map[K]bool, len(m.actual))
	hasNot := make(map[K]bool)
	for k := range m.actual {
		has[k] = true
	}
	for _, k := range expected {
		if !has[k] {
			if !hasNot[k] {
				missing = append(missing, k)
			}
			hasNot[k] = true
			continue
		}
		found = append(found, k)
	}
	return Assert(len(missing) == 0, "got <%v>, wanted keys <%v>", found, missing)
}

func (m Mappy[K, V]) HasKey(key K) AssertionMessage {
	_, ok := m.actual[key]
	return Assert(ok, "got key absent <%v>, wanted present in map", key)
}

func (m Mappy[K, V]) NotHasKey(key K) AssertionMessage {
	_, ok := m.actual[key]
	return Assert(!ok, "got key present <%v>, wanted absent from map", key)
}

func (m Mappy[K, V]) KeysExactly(expected ...K) AssertionMessage {
	actualKeys := make(map[K]bool, len(m.actual))
	for k := range m.actual {
		actualKeys[k] = true
	}

	expectedKeys := make(map[K]bool, len(expected))
	for _, k := range expected {
		expectedKeys[k] = true
	}

	var missing []K
	for k := range expectedKeys {
		if !actualKeys[k] {
			missing = append(missing, k)
		}
	}

	var extra []K
	for k := range actualKeys {
		if !expectedKeys[k] {
			extra = append(extra, k)
		}
	}

	return Assert(len(missing) == 0 && len(extra) == 0, "got extra keys <%v> and missing keys <%v>, wanted exact key set", extra, missing)
}

func (m Mappy[K, V]) IsEmpty() AssertionMessage {
	return Assert(len(m.actual) == 0, "got len=<%d>, wanted empty map", len(m.actual))
}

func (m Mappy[K, V]) NotEmpty() AssertionMessage {
	return Assert(len(m.actual) > 0, "got len=<%d>, wanted non-empty map", len(m.actual))
}

func (m Mappy[K, V]) WithValues(values ...V) AssertionMessage {
	var found []V
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
			continue
		}
		found = append(found, v)
	}
	return Assert(len(missing) == 0, "got <%v>, wanted values <%v>", found, missing)
}

func (m Mappy[K, V]) NotContains(values ...V) AssertionMessage {
	var present []V
	has := make(map[int]bool)
	for _, actual := range m.actual {
		for i, expected := range values {
			if cmp.Equal(actual, expected) {
				if !has[i] {
					present = append(present, expected)
				}
				has[i] = true
			}
		}
	}
	return Assert(len(present) == 0, "got values <%v>, wanted absent from map", present)
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
	return Assert(len(present) == 0, "got keys <%v>, wanted absent from map", present)
}

func (m Mappy[K, V]) Len(expected int) AssertionMessage {
	sz := len(m.actual)
	return Assert(sz == expected, "got len=<%v>, wanted <%v>", sz, expected)
}

func (m Mappy[K, V]) WithItem(k K, expected V) AssertionMessage {
	actual, ok := m.actual[k]
	if !ok {
		return Assert(false, "got key absent, wanted value for key <%v>", k)
	}
	return Assert(cmp.Equal(actual, expected), "got value=<%v> for key=<hi>, wanted <%v>", actual, expected)
}

func (m Mappy[K, V]) EqualTo(expected map[K]V) AssertionMessage {
	diff := cmp.Diff(expected, m.actual)
	return Assert(diff == "", "Map mismatch (-want +got):\n%s", diff)
}

func (m Mappy[K, V]) HasEntry(keyMatcher Matcher[K], valueMatcher Matcher[V]) AssertionMessage {
	return HasEntry[K, V](keyMatcher, valueMatcher).Match(m.actual)
}

func (m Mappy[K, V]) HasKeyMatching(matcher Matcher[K]) AssertionMessage {
	return HasKeyMatching[K, V](matcher).Match(m.actual)
}

func (m Mappy[K, V]) HasValueMatching(matcher Matcher[V]) AssertionMessage {
	return HasValueMatching[K, V](matcher).Match(m.actual)
}
