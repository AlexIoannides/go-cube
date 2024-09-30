/*
Task contains everything needed by workers to execute and manage a workload.
*/
package task

// State describes a task at a specific moment in time.
type State byte

const (
	Pending State = iota
	Scheduled
	Running
	Failed
	Completed
)
