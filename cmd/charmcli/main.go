package main

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/charl/TaskAndGo/charmcli" // Import the charmcli package
	pb "github.com/charl/TaskAndGo/scheduler"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Connect to gRPC server.
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create gRPC client.
	client := pb.NewSchedulerClient(conn)

	// Run the Bubble Tea program using the charmcli package's InitialModel.
	program := tea.NewProgram(
		charmcli.InitialModel(client), // Use the function from the imported package
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := program.Run(); err != nil {
		fmt.Println("Error running Charm CLI:", err)
		os.Exit(1)
	}
}
