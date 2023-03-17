package hammy

import (
	"fmt"
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
	return AssertionMessage{
		IsSuccessful: len(missing) == 0,
		Message:      fmt.Sprintf("want map with keys <%v>, but missing", missing),
	}
}

func (m Mappy[K, V]) IsEmpty() AssertionMessage {
	return AssertionMessage{
		IsSuccessful: len(m.actual) == 0,
		Message:      fmt.Sprintf("want empty map, got len() = %v", len(m.actual)),
	}
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
	return AssertionMessage{
		IsSuccessful: len(missing) == 0,
		Message:      fmt.Sprintf("want map with values <%v>, but missing", missing),
	}

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
	return AssertionMessage{
		IsSuccessful: len(present) == 0,
		Message:      fmt.Sprintf("want map without keys <%v>, but present", present),
	}
}
