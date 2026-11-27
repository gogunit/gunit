package eye

import (
	"slices"
	"sync"
)

// Recorder records calls to a fake or callback. T is the value saved for each
// call. Use a struct for multiple arguments or struct{} for calls with no
// arguments.
//
// The zero value is ready to use. A Recorder is safe for concurrent use and
// must not be copied after first use.
type Recorder[T any] struct {
	mu    sync.Mutex
	calls []T
}

// Record saves one call.
func (r *Recorder[T]) Record(call T) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.calls = append(r.calls, call)
}

// Len returns the number of recorded calls.
func (r *Recorder[T]) Len() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.calls)
}

// Calls returns a snapshot in the order calls were recorded. The returned
// slice has independent storage. References within each call are copied
// shallowly.
func (r *Recorder[T]) Calls() []T {
	r.mu.Lock()
	defer r.mu.Unlock()
	return slices.Clone(r.calls)
}
