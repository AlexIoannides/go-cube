/*
Node contains everything needed to represent the physical nodes that run workers.
*/
package node

// Node holds all the information about the current state of a physical node in the orchestration cluster.
type Node struct {
	Name            string
	Ip              string
	Cores           int
	Memory          int
	MemoryAllocated int
	Disk            int
	DiskAllocated   int
	Role            string
	TaskCount       int
}
