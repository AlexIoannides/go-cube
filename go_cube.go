package main

import (
	"fmt"
	"time"

	"github.com/AlexIoannides/go-cube/internal/manager"
	"github.com/AlexIoannides/go-cube/internal/node"
	"github.com/AlexIoannides/go-cube/internal/task"
	"github.com/AlexIoannides/go-cube/internal/worker"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {
	t := task.Task{
		Id:     uuid.New(),
		Name:   "Task-1",
		State:  task.Pending,
		Image:  "Image-1",
		Memory: 1024,
		Disk:   1,
	}

	te := task.TaskEvent{
		Id:        uuid.New(),
		State:     task.Pending,
		Timestamp: time.Now(),
		Task:      t,
	}

	fmt.Printf("task: %v\n", t)
	fmt.Printf("task event: %v\n", te)

	w := worker.Worker{
		Name:  "Worker-1",
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]*task.Task),
	}

	fmt.Printf("worker: %v\n", w)
	w.CollectState()
	w.RunTask()
	w.StartTask()
	w.StopTask()

	m := manager.Manager{
		Pending: *queue.New(),
		TaskDb:  make(map[string][]*task.Task),
		EventDb: make(map[string][]*task.TaskEvent),
		Workers: []string{w.Name},
	}

	fmt.Printf("manager: %v\n", m)
	m.SelectWorker()
	m.UpdateTasks()
	m.SendWork()

	n := node.Node{
		Name:            "Node-1",
		Ip:              "192.168.1.1",
		Cores:           4,
		Memory:          1024,
		MemoryAllocated: 0,
		Disk:            25,
		DiskAllocated:   0,
		Role:            "worker",
		TaskCount:       1,
	}

	fmt.Printf("manager: %v\n", n)
}
