package worker

import (
	"go-concur/internal/request"
	"sync"
)

// Task encapsulates a work item that should go in a work
// pool.
type Task struct {
	// Err holds an error that occurred during a task. Its
	// result is only meaningful after Run has been called
	// for the pool that holds it.
	Err error

	f func(s request.Apis, n string) error
	s request.Apis
	n string
}

// NewTask initializes a new task based on a given work
// function.
func NewTask(f func(request.Apis, string) error, s request.Apis, n string) *Task {
	return &Task{f: f, s: s, n: n}
}

// Run runs a Task and does appropriate accounting via a
// given sync.WorkGroup.
func (t *Task) Run(wg *sync.WaitGroup) {
	t.Err = t.f(t.s, t.n)
	wg.Done()
}
