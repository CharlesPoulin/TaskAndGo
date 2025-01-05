// Test/scheduler_test.go
package main

import (
	"context"
	"net"
	"testing"

	pb "github.com/charl/TaskAndGo/scheduler" // Updated import path
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

type mockSchedulerServer struct {
	pb.UnimplementedSchedulerServer
	// Add fields if needed for mocking
}

func (s *mockSchedulerServer) SubmitTask(ctx context.Context, req *pb.SubmitTaskRequest) (*pb.SubmitTaskResponse, error) {
	// Mock implementation
	return &pb.SubmitTaskResponse{
		Success: true,
		Message: "Task submitted successfully",
	}, nil
}

func (s *mockSchedulerServer) GetTaskStatus(ctx context.Context, req *pb.TaskStatusRequest) (*pb.TaskStatusResponse, error) {
	// Mock implementation
	return &pb.TaskStatusResponse{
		Found:         true,
		CurrentStatus: "QUEUED",
		Message:       "Task is in QUEUED state",
	}, nil
}

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterSchedulerServer(s, &mockSchedulerServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			// Handle error
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestScheduler(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewSchedulerClient(conn)

	// Test SubmitTask
	res, err := client.SubmitTask(ctx, &pb.SubmitTaskRequest{
		Task: &pb.Task{
			TaskId:   "test-task-001",
			Priority: "high",
			Data:     "Test data",
		},
	})
	if err != nil {
		t.Fatalf("SubmitTask RPC failed: %v", err)
	}
	if !res.Success {
		t.Errorf("Expected success, got failure. Message: %s", res.Message)
	}

	// Test GetTaskStatus
	statusRes, err := client.GetTaskStatus(ctx, &pb.TaskStatusRequest{
		TaskId: "test-task-001",
	})
	if err != nil {
		t.Fatalf("GetTaskStatus RPC failed: %v", err)
	}
	if !statusRes.Found {
		t.Errorf("Expected task to be found")
	}
	if statusRes.CurrentStatus != "QUEUED" && statusRes.CurrentStatus != "RUNNING" {
		t.Errorf("Expected task status to be QUEUED or RUNNING, got: %s", statusRes.CurrentStatus)
	}
}
