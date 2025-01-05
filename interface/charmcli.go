// interface/charmcli.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// Alias for the generated gRPC code.
	pb "github.com/charl/TaskAndGo/scheduler"

	// Bubble Tea and Bubbles for the UI components.
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Define tab indices for managing the active tab.
type tabIndex int

const (
	tabSubmit tabIndex = iota
	tabJobs
)

// model defines the state of the application.
type model struct {
	// Inputs for task submission.
	taskIDInput   textinput.Model
	priorityInput textinput.Model

	// GRPC client and error state.
	client pb.SchedulerClient
	err    error

	// Submission state.
	submitted bool
	status    string

	// Active tab.
	activeTab tabIndex

	// List of tasks for the Jobs tab.
	tasks []*pb.Task
}

// initialModel initializes the Bubble Tea model with input fields.
func initialModel(client pb.SchedulerClient) model {
	ti1 := textinput.New()
	ti1.Placeholder = "Enter Task ID"
	ti1.Focus()
	ti1.CharLimit = 32
	ti1.Width = 30

	ti2 := textinput.New()
	ti2.Placeholder = "Priority (low, medium, high)"
	ti2.CharLimit = 10
	ti2.Width = 30

	return model{
		taskIDInput:   ti1,
		priorityInput: ti2,
		client:        client,
		activeTab:     tabSubmit,
	}
}

// Init initializes the program, starting with blinking cursors.
func (m model) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles incoming messages and updates the model accordingly.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		// Quit the application.
		case "ctrl+c", "q":
			return m, tea.Quit

		// Switch tabs.
		case "tab":
			if m.activeTab == tabSubmit {
				m.activeTab = tabJobs
				// Fetch tasks when switching to Jobs tab.
				return m, fetchTasksCmd(m.client)
			} else {
				m.activeTab = tabSubmit
			}

		// Refresh tasks list in Jobs tab.
		case "r":
			if m.activeTab == tabJobs {
				return m, fetchTasksCmd(m.client)
			}

		// Handle Enter key based on active tab.
		case "enter":
			if m.activeTab == tabSubmit {
				if !m.submitted {
					// Submit the task.
					return m, submitTaskCmd(m)
				} else {
					// Check the status of the submitted task.
					return m, checkStatusCmd(m)
				}
			}
		}

	// Handle the list of tasks fetched from the server.
	case listTasksMsg:
		m.tasks = msg
		m.err = nil
		return m, nil

	// Handle submission success/failure.
	case submitTaskResponseMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, nil
		}
		if msg.response.Success {
			m.submitted = true
			m.status = "Submitted. Press Enter to check status."
		} else {
			m.err = fmt.Errorf("submission failed: %s", msg.response.Message)
		}
		return m, nil

	// Handle task status response.
	case taskStatusResponseMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, nil
		}
		if msg.response.Found {
			m.status = msg.response.CurrentStatus
		} else {
			m.status = "Task not found."
		}
		return m, nil

	// Handle errors from any command.
	case errMsg:
		m.err = msg
		return m, nil
	}

	// Update input fields only if on the Submit tab.
	if m.activeTab == tabSubmit {
		m.taskIDInput, cmd = m.taskIDInput.Update(msg)
		var cmd2 tea.Cmd
		m.priorityInput, cmd2 = m.priorityInput.Update(msg)
		return m, tea.Batch(cmd, cmd2)
	}

	return m, nil
}

// View renders the UI based on the current model state.
func (m model) View() string {
	// Display errors prominently.
	if m.err != nil {
		return fmt.Sprintf(
			"Error: %v\n\nPress 'q' to quit or 'tab' to switch tabs.\n",
			m.err,
		)
	}

	// Header
	s := "=== Charm CLI with Tabs ===\n\n"

	// Render based on the active tab.
	switch m.activeTab {
	case tabSubmit:
		s += "[ Task Submission ] | (Press TAB to switch to Jobs)\n\n"
		if !m.submitted {
			s += "Submit a new task:\n\n"
			s += fmt.Sprintf("Task ID: %s\n\n", m.taskIDInput.View())
			s += fmt.Sprintf("Priority: %s\n\n", m.priorityInput.View())
			s += "Press Enter to submit or 'q' to quit.\n"
		} else {
			s += fmt.Sprintf(
				"Task Submitted.\n\nCurrent Status: %s\n\nPress Enter to refresh status or 'q' to quit.\n",
				m.status,
			)
		}

	case tabJobs:
		s += "(Press TAB to switch to Submission) | [ Jobs List ]\n\n"
		s += "Press 'r' to refresh the tasks list.\n\n"
		if len(m.tasks) == 0 {
			s += "No tasks to display or still loading..."
		} else {
			s += "Current Tasks:\n\n"
			for _, t := range m.tasks {
				s += fmt.Sprintf(" • %s (%s) => Status: %s\n", t.TaskId, t.Priority, t.Status)
			}
		}
	}
	s += "\n"
	return s
}

//
// Commands and Helper Functions
//

// Custom message types for asynchronous operations.

// listTasksMsg carries the list of tasks fetched from the server.
type listTasksMsg []*pb.Task

// submitTaskResponseMsg carries the response from submitting a task.
type submitTaskResponseMsg struct {
	response *pb.SubmitTaskResponse
	err      error
}

// taskStatusResponseMsg carries the response from checking a task's status.
type taskStatusResponseMsg struct {
	response *pb.TaskStatusResponse
	err      error
}

// errMsg carries any errors from commands.
type errMsg error

// fetchTasksCmd retrieves the list of tasks from the server.
func fetchTasksCmd(client pb.SchedulerClient) tea.Cmd {
	return func() tea.Msg {
		res, err := client.ListTasks(context.Background(), &pb.TaskListRequest{})
		if err != nil {
			return errMsg(err)
		}
		return listTasksMsg(res.Tasks)
	}
}

// submitTaskCmd submits a new task to the server.
func submitTaskCmd(m model) tea.Cmd {
	return func() tea.Msg {
		taskID := m.taskIDInput.Value()
		priority := m.priorityInput.Value()

		if taskID == "" {
			return errMsg(fmt.Errorf("task ID cannot be empty"))
		}

		validPriorities := map[string]bool{
			"low":    true,
			"medium": true,
			"high":   true,
		}
		if !validPriorities[priority] {
			return errMsg(fmt.Errorf("invalid priority: %s", priority))
		}

		// Create the task object.
		task := &pb.Task{
			TaskId:   taskID,
			Priority: priority,
			Data:     "Sample Data",
			Status:   "QUEUED", // Initial status.
		}

		// Submit the task.
		res, err := m.client.SubmitTask(context.Background(), &pb.SubmitTaskRequest{
			Task: task,
		})

		return submitTaskResponseMsg{
			response: res,
			err:      err,
		}
	}
}

// checkStatusCmd checks the status of the submitted task.
func checkStatusCmd(m model) tea.Cmd {
	return func() tea.Msg {
		taskID := m.taskIDInput.Value()

		res, err := m.client.GetTaskStatus(context.Background(), &pb.TaskStatusRequest{
			TaskId: taskID,
		})

		return taskStatusResponseMsg{
			response: res,
			err:      err,
		}
	}
}

//
// Main Function
//

func main() {
	// Establish a secure gRPC connection.
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create the Scheduler client.
	client := pb.NewSchedulerClient(conn)

	// Initialize the Bubble Tea program with the initial model.
	p := tea.NewProgram(initialModel(client))

	// Run the program using Run() instead of the deprecated Start().
	_, err = p.Run()
	if err != nil {
		fmt.Println("Error running Charm CLI:", err)
		os.Exit(1)
	}

}
