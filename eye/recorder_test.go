package eye_test

import (
	"slices"
	"sync"
	"testing"

	"github.com/gogunit/gunit/eye"
)

func TestRecorderZeroValue(t *testing.T) {
	var recorder eye.Recorder[struct{}]

	if got := recorder.Len(); got != 0 {
		t.Errorf("Len() = %d, want 0", got)
	}
	if got := recorder.Calls(); len(got) != 0 {
		t.Errorf("Calls() = %v, want no calls", got)
	}

	recorder.Record(struct{}{})
	if got := recorder.Len(); got != 1 {
		t.Errorf("Len() = %d after recording a call, want 1", got)
	}
}

func TestRecorderPreservesArgumentsAndOrder(t *testing.T) {
	type arguments struct {
		name   string
		values []int
	}
	var recorder eye.Recorder[arguments]
	recorder.Record(arguments{name: "first", values: []int{1, 2}})
	recorder.Record(arguments{name: "second", values: []int{3}})

	calls := recorder.Calls()
	if len(calls) != 2 {
		t.Fatalf("Calls() = %#v, want two calls", calls)
	}
	if calls[0].name != "first" || !slices.Equal(calls[0].values, []int{1, 2}) {
		t.Errorf("first call = %#v, want first with values [1 2]", calls[0])
	}
	if calls[1].name != "second" || !slices.Equal(calls[1].values, []int{3}) {
		t.Errorf("second call = %#v, want second with values [3]", calls[1])
	}
}

func TestRecorderCallsReturnsIndependentSlice(t *testing.T) {
	var recorder eye.Recorder[string]
	recorder.Record("first")
	snapshot := recorder.Calls()

	snapshot[0] = "edited"
	recorder.Record("second")

	if got := recorder.Calls(); !slices.Equal(got, []string{"first", "second"}) {
		t.Errorf("Calls() = %q, want [first second]", got)
	}
}

func TestRecorderConcurrentUse(t *testing.T) {
	const callCount = 16
	var recorder eye.Recorder[int]
	var group sync.WaitGroup
	for call := range callCount {
		group.Go(func() {
			recorder.Record(call)
			recorder.Len()
			recorder.Calls()
		})
	}
	group.Wait()

	if got := recorder.Len(); got != callCount {
		t.Errorf("Len() = %d, want %d", got, callCount)
	}
	calls := recorder.Calls()
	slices.Sort(calls)
	for index, call := range calls {
		if call != index {
			t.Errorf("sorted calls[%d] = %d, want %d", index, call, index)
		}
	}
}
