// schedulerimpl/strategy.go
package schedulerimpl

import (
	"log"
	"time"

	pb "github.com/charl/TaskAndGo/scheduler" // Proto package
)

// SchedulingStrategy describes how to distribute tasks/subtasks.
type SchedulingStrategy interface {
	DistributeTask(task *pb.Task, store *TaskStore)
	DistributeSubTask(subtask *pb.SubTask, store *TaskStore)
}

// ---------------------- SPLIT STRATEGY ----------------------
type SplitStrategy struct{}

func (s *SplitStrategy) DistributeTask(task *pb.Task, store *TaskStore) {
	// Mark as RUNNING
	store.mu.Lock()
	store.tasks[task.TaskId] = "RUNNING"
	store.mu.Unlock()

	// Simulate parallel distribution
	go func() {
		time.Sleep(3 * time.Second)
		store.mu.Lock()
		store.tasks[task.TaskId] = "COMPLETED"
		store.mu.Unlock()
		log.Printf("[SplitStrategy] Task %s completed (split across multiple nodes)", task.TaskId)
	}()
}

func (s *SplitStrategy) DistributeSubTask(subtask *pb.SubTask, store *TaskStore) {
	log.Printf("[SplitStrategy] Subtask %s distributed among nodes", subtask.SubtaskId)
}

// ---------------------- BATCH STRATEGY ----------------------
type BatchStrategy struct{}

func (b *BatchStrategy) DistributeTask(task *pb.Task, store *TaskStore) {
	// Mark as RUNNING
	store.mu.Lock()
	store.tasks[task.TaskId] = "RUNNING"
	store.mu.Unlock()

	// Simulate “batch” distribution
	go func() {
		time.Sleep(2 * time.Second)
		store.mu.Lock()
		store.tasks[task.TaskId] = "COMPLETED"
		store.mu.Unlock()
		log.Printf("[BatchStrategy] Task %s completed (single node batch)", task.TaskId)
	}()
}

func (b *BatchStrategy) DistributeSubTask(subtask *pb.SubTask, store *TaskStore) {
	log.Printf("[BatchStrategy] Subtask %s assigned to a single node", subtask.SubtaskId)
}
