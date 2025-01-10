// schedulerimpl/service.go
package schedulerimpl

import (
	"context"
	"fmt"
	"log"
	"sync"

	pb "github.com/charl/TaskAndGo/scheduler" // The generated proto code
)

// TaskStore is a simple in-memory store for tasks, subtasks, and nodes.
type TaskStore struct {
	mu       sync.Mutex
	tasks    map[string]string      // e.g. "taskID" -> "RUNNING"/"COMPLETED"
	subtasks map[string]*pb.SubTask // e.g. "subtaskID" -> SubTask object
	nodes    map[string]string      // e.g. "nodeID" -> capacity
}

// NewTaskStore constructs a blank store.
func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks:    make(map[string]string),
		subtasks: make(map[string]*pb.SubTask),
		nodes:    make(map[string]string),
	}
}

// SchedulerService implements the proto interface "SchedulerServer".
type SchedulerService struct {
	pb.UnimplementedSchedulerServer

	store    *TaskStore
	strategy SchedulingStrategy
}

// NewSchedulerService is a constructor that injects a TaskStore + Strategy.
func NewSchedulerService(store *TaskStore, strategy SchedulingStrategy) *SchedulerService {
	return &SchedulerService{
		store:    store,
		strategy: strategy,
	}
}

// -------------- RPC METHODS --------------

// SubmitTask handles creating a new task in memory and delegates distribution.
func (s *SchedulerService) SubmitTask(ctx context.Context, req *pb.SubmitTaskRequest) (*pb.SubmitTaskResponse, error) {
	task := req.GetTask()
	if task == nil {
		return &pb.SubmitTaskResponse{
			Success: false,
			Message: "No task provided",
		}, nil
	}

	// Prevent duplicates
	s.store.mu.Lock()
	defer s.store.mu.Unlock()

	if _, exists := s.store.tasks[task.TaskId]; exists {
		return &pb.SubmitTaskResponse{
			Success: false,
			Message: fmt.Sprintf("Task %s already exists", task.TaskId),
		}, nil
	}

	s.store.tasks[task.TaskId] = "QUEUED"
	log.Printf("Received task %s with priority %s", task.TaskId, task.Priority)

	// Hand off to the scheduling strategy asynchronously
	go s.strategy.DistributeTask(task, s.store)

	return &pb.SubmitTaskResponse{
		Success: true,
		Message: fmt.Sprintf("Task %s submitted successfully", task.TaskId),
	}, nil
}

// SubmitSubTask handles subtask creation and delegates distribution.
func (s *SchedulerService) SubmitSubTask(ctx context.Context, req *pb.SubmitSubTaskRequest) (*pb.SubmitSubTaskResponse, error) {
	subtask := req.GetSubtask()
	if subtask == nil {
		return &pb.SubmitSubTaskResponse{
			Success: false,
			Message: "No subtask provided",
		}, nil
	}

	s.store.mu.Lock()
	s.store.subtasks[subtask.SubtaskId] = subtask
	s.store.mu.Unlock()

	log.Printf("Received subtask %s for parent task %s", subtask.SubtaskId, subtask.ParentTaskId)

	// Hand off to the scheduling strategy
	go s.strategy.DistributeSubTask(subtask, s.store)

	return &pb.SubmitSubTaskResponse{
		Success: true,
		Message: fmt.Sprintf("Subtask %s submitted for task %s", subtask.SubtaskId, subtask.ParentTaskId),
	}, nil
}

// GetTaskStatus returns a Task's status from the in-memory store.
func (s *SchedulerService) GetTaskStatus(ctx context.Context, req *pb.TaskStatusRequest) (*pb.TaskStatusResponse, error) {
	taskID := req.GetTaskId()

	s.store.mu.Lock()
	status, ok := s.store.tasks[taskID]
	s.store.mu.Unlock()

	if !ok {
		return &pb.TaskStatusResponse{
			Found:   false,
			Message: fmt.Sprintf("Task %s not found", taskID),
		}, nil
	}

	return &pb.TaskStatusResponse{
		Found:         true,
		CurrentStatus: status,
		Message:       fmt.Sprintf("Task %s is in %s state", taskID, status),
	}, nil
}

// ListTasks returns all tasks from the store.
func (s *SchedulerService) ListTasks(ctx context.Context, req *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	var allTasks []*pb.Task

	s.store.mu.Lock()
	defer s.store.mu.Unlock()

	for taskID, status := range s.store.tasks {
		// For brevity, we only fill TaskId + Status
		t := &pb.Task{
			TaskId: taskID,
			Status: status,
		}
		allTasks = append(allTasks, t)
	}

	return &pb.TaskListResponse{Tasks: allTasks}, nil
}

// GetResourceUsage is just a stub here. Return an empty response or fill it in.
func (s *SchedulerService) GetResourceUsage(ctx context.Context, req *pb.ResourceUsageRequest) (*pb.ResourceUsageResponse, error) {
	// Not implemented, just return empty
	return &pb.ResourceUsageResponse{}, nil
}

// RegisterNode is also optional if you’d like to store node info.
func (s *SchedulerService) RegisterNode(ctx context.Context, req *pb.RegisterNodeRequest) (*pb.RegisterNodeResponse, error) {
	nodeID := req.GetNodeId()
	capacity := req.GetCapacity()

	if nodeID == "" {
		return &pb.RegisterNodeResponse{
			Success: false,
			Message: "Node ID cannot be empty",
		}, nil
	}

	s.store.mu.Lock()
	s.store.nodes[nodeID] = capacity
	s.store.mu.Unlock()

	log.Printf("Registered node %s with capacity %s", nodeID, capacity)

	return &pb.RegisterNodeResponse{
		Success: true,
		Message: fmt.Sprintf("Node %s registered successfully", nodeID),
	}, nil
}
