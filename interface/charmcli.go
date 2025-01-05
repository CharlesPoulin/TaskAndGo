// interface/charmcli.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/charl/TaskAndGo/scheduler" // Updated import path
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	"google.golang.org/grpc"
)

type model struct {
	taskIDInput   textinput.Model
	priorityInput textinput.Model
	err           error
	client        pb.SchedulerClient
	submitted     bool
	status        string
}

func initialModel(client pb.SchedulerClient) model {
	ti1 := textinput.New()
	ti1.Placeholder = "Enter Task ID"
	ti1.Focus()
	ti1.CharLimit = 32

	ti2 := textinput.New()
	ti2.Placeholder = "Priority (low, medium, high)"
	ti2.CharLimit = 10

	return model{
		taskIDInput:   ti1,
		priorityInput: ti2,
		client:        client,
	}
}

func (m model) Init() bubbletea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg bubbletea.Msg) (bubbletea.Model, bubbletea.Cmd) {
	var cmd bubbletea.Cmd

	switch msg := msg.(type) {
	case bubbletea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, bubbletea.Quit
		case "enter":
			// Submit the task
			if !m.submitted {
				taskID := m.taskIDInput.Value()
				priority := m.priorityInput.Value()

				res, err := m.client.SubmitTask(context.Background(), &pb.SubmitTaskRequest{
					Task: &pb.Task{
						TaskId:   taskID,
						Priority: priority,
						Data:     "Sample Data",
					},
				})
				if err != nil {
					m.err = err
				} else {
					if res.Success {
						m.submitted = true
						m.status = "Submitted. Checking status..."
					} else {
						m.err = fmt.Errorf("Submission failed: %s", res.Message)
					}
				}
			} else {
				// After submission, retrieve status
				taskID := m.taskIDInput.Value()
				statusRes, err := m.client.GetTaskStatus(context.Background(), &pb.TaskStatusRequest{
					TaskId: taskID,
				})
				if err != nil {
					m.err = err
				} else if statusRes.Found {
					m.status = statusRes.CurrentStatus
				} else {
					m.status = "Task not found"
				}
			}
		}
	}

	// Update text inputs
	m.taskIDInput, cmd = m.taskIDInput.Update(msg)
	m.priorityInput, _ = m.priorityInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n\nPress q to quit.\n", m.err)
	}

	if !m.submitted {
		return fmt.Sprintf(
			"Submit a task:\n\nTask ID: %s\n\nPriority: %s\n\nPress Enter to submit or q to quit.\n",
			m.taskIDInput.View(),
			m.priorityInput.View(),
		)
	}

	return fmt.Sprintf(
		"Task Submitted.\n\nCurrent Status: %s\n\nPress Enter to refresh status or q to quit.\n",
		m.status,
	)
}

func main() {
	// Dial gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewSchedulerClient(conn)
	p := bubbletea.NewProgram(initialModel(client))
	if err := p.Start(); err != nil {
		fmt.Println("Error running Charm CLI:", err)
		os.Exit(1)
	}
}
