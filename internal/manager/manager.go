/*
Manager contains everything needed to manage the execution of tasks on workers.
*/
package manager

import (
	"fmt"

	"github.com/AlexIoannides/go-cube/internal/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

// Manager holds all information about submitted tasks and the workers managing their execution.
type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

// SelectWorker selects an appropriate worker for task submission.
func (m *Manager) SelectWorker() {
	fmt.Println("I will select an appropriate worker")
}

// UpdateTasks will update the manager about the latest state of all tasks.
func (m *Manager) UpdateTasks() {
	fmt.Println("I will update the manager with the latest tasks info")
}

// SendWork sends tasks to workers for execution.
func (m *Manager) SendWork() {
	fmt.Println("I will send tasks to workers")
}
