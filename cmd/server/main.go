// cmd/server/main.go
package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// The proto package is "github.com/charl/TaskAndGo/scheduler"
	pb "github.com/charl/TaskAndGo/scheduler"

	// Our new sub-package that holds the actual service impl & strategies
	"github.com/charl/TaskAndGo/scheduler/schedulerimpl"
)

func main() {
	// 1) Create store
	store := schedulerimpl.NewTaskStore()

	// 2) Choose which strategy to use
	strategyChoice := os.Getenv("SCHEDULER_STRATEGY") // "batch" or "split"
	var chosenStrategy schedulerimpl.SchedulingStrategy
	switch strategyChoice {
	case "batch":
		chosenStrategy = &schedulerimpl.BatchStrategy{}
	default:
		chosenStrategy = &schedulerimpl.SplitStrategy{} // fallback
	}

	// 3) Create your gRPC service implementation
	svc := schedulerimpl.NewSchedulerService(store, chosenStrategy)

	// 4) Listen on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// 5) Create & register a gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterSchedulerServer(grpcServer, svc)

	reflection.Register(grpcServer)

	fmt.Printf("Server started with strategy: %T on port 50051\n", chosenStrategy)

	// 6) Serve
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
