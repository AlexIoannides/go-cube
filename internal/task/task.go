/*
Task contains everything needed by workers to execute and manage a workload.
*/
package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

// State describes a task at a specific moment in time.
type State byte

const (
	Pending State = iota
	Scheduled
	Running
	Failed
	Completed
)

// Task holds all information about a task scheduled on a Worker.
type Task struct {
	Id            uuid.UUID
	Name          string
	State         State
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}

// TaskEvent represents desired transition states for scheduled tasks
type TaskEvent struct {
	Id        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}
