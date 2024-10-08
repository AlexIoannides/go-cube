/*
Task contains everything needed by workers to execute and manage a workload.
*/
package task

import (
	"time"

	"github.com/docker/docker/client"
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

// Task holds all high-level information about a task scheduled on a Worker.
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

// TaskEvent represents desired transition states for scheduled tasks.
type TaskEvent struct {
	Id        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}

// Config holds all of the container configuration required to run a task's workload.
type Config struct {
	Name          string
	AttachStdin   bool
	AttachStdout  bool
	AttachStderr  bool
	ExposedPorts  nat.PortSet
	Cmd           []string
	Image         string
	Cpu           float64
	Memory        int64
	Disk          int64
	Env           []string
	RestartPolicy string
}

// Docker enables workers to manage tasks using Docker.
type Docker struct {
	Client *client.Client
	Config Config
}

// DockerResult hold information about requests issues to the Docker client.
type DockerResult struct {
	Error       error
	Action      string
	ContainerId string
	Result      string
}
