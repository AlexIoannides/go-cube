/*
Worker contains everything needed to manage the execution of tasks scheduled to a node.
*/
package worker

import (
	"fmt"

	"github.com/AlexIoannides/go-cube/internal/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

// Worker holds all of the information describing the current state of a worker.
type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

// CollectState collects statistics about the Worker.
func (w *Worker) CollectState() {
	fmt.Println("I will collect stats")
}

// RunTask identifies a task's current state and based on this starts/stops it.
func (w *Worker) RunTask() {
	fmt.Println("I will start or stop a task")
}

// StartTask will start a task.
func (w *Worker) StartTask() {
	fmt.Println("I will start a task")
}

// StopTask will stop a task
func (w *Worker) StopTask() {
	fmt.Println("I will stop a task")
}
