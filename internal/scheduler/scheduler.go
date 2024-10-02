/*
Scheduler contains everything needed to select an appropriate worker for sending a task to.
*/
package scheduler

// Scheduler defines the interface for algorithms that can select appropriate workers for a given task.
type Scheduler interface {
	SelectCandidateNodes()
	Score()
	Pick()
}
