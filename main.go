package main

import (
	"context"
	"fmt"
	"net"
	"sync"

	pb "github.com/charl/TaskAndGo/scheduler" // Import path to your generated code
	"google.golang.org/grpc"
)

// SchedulerService implements all the RPC methods defined in scheduler.proto.
type SchedulerService struct {
	pb.UnimplementedSchedulerServer
	// We'll store tasks in an in-memory sync.Map (key=taskId, value=*pb.Task).
	tasks sync.Map
}

// SubmitTask handles creating or updating a task in our store.
func (s *SchedulerService) SubmitTask(ctx context.Context, req *pb.SubmitTaskRequest) (*pb.SubmitTaskResponse, error) {
	// Extract the task from the request.
	task := req.GetTask()
	if task == nil {
		return &pb.SubmitTaskResponse{
			Success: false,
			Message: "No task provided",
		}, nil
	}

	// Store the task in our in-memory map (overwrites if the key already exists).
	s.tasks.Store(task.GetTaskId(), task)

	return &pb.SubmitTaskResponse{
		Success: true,
		Message: fmt.Sprintf("Task %s submitted successfully", task.GetTaskId()),
	}, nil
}

// GetTaskStatus retrieves the status of a given task by ID.
func (s *SchedulerService) GetTaskStatus(ctx context.Context, req *pb.TaskStatusRequest) (*pb.TaskStatusResponse, error) {
	taskID := req.GetTaskId()

	// Attempt to load the task from our map.
	val, ok := s.tasks.Load(taskID)
	if !ok {
		// Not found.
		return &pb.TaskStatusResponse{
			Found:   false,
			Message: fmt.Sprintf("Task %s not found", taskID),
		}, nil
	}

	// The task exists, so cast back to *pb.Task.
	task := val.(*pb.Task)
	return &pb.TaskStatusResponse{
		Found:         true,
		CurrentStatus: task.GetStatus(),
		Message:       fmt.Sprintf("Task %s status: %s", taskID, task.GetStatus()),
	}, nil
}

// ListTasks returns all tasks currently in our in-memory store.
func (s *SchedulerService) ListTasks(ctx context.Context, req *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	var allTasks []*pb.Task

	// Range over all tasks in the sync.Map.
	s.tasks.Range(func(key, value interface{}) bool {
		if t, ok := value.(*pb.Task); ok {
			allTasks = append(allTasks, t)
		}
		return true // keep iterating
	})

	return &pb.TaskListResponse{
		Tasks: allTasks,
	}, nil
}

func main() {
	// Listen on TCP port 50051.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	// Create a new gRPC server.
	grpcServer := grpc.NewServer()

	// Register our SchedulerService implementation.
	pb.RegisterSchedulerServer(grpcServer, &SchedulerService{})

	fmt.Println("Running server on :50051")
	// Block the main goroutine and start serving.
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
