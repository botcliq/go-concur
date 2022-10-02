package worker

import "sync"

// Task encapsulates a work item that should go in a work
// pool.
type Task struct {
	// Err holds an error that occurred during a task. Its
	// result is only meaningful after Run has been called
	// for the pool that holds it.
	Err error

	f func(s string, n int) error
	s string
	n int
}

// NewTask initializes a new task based on a given work
// function.
func NewTask(f func(string, int) error, s string, n int) *Task {
	return &Task{f: f, s: s, n: n}
}

// Run runs a Task and does appropriate accounting via a
// given sync.WorkGroup.
func (t *Task) Run(wg *sync.WaitGroup) {
	t.Err = t.f(s, n)
	wg.Done()
}
