// charmcli.go
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

// tabIndex enumerates which tab is active in the UI.
type tabIndex int

const (
	tabSubmit tabIndex = iota
	tabJobs
)

// model defines the state of the entire application (both tabs).
type model struct {
	// We’ll keep the two text inputs in a slice for easier focus management.
	// focusIndex indicates which input is currently "active."
	focusIndex int
	inputs     []textinput.Model

	// gRPC client and error state.
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

// --- Initialization ---

func initialModel(client pb.SchedulerClient) model {
	// Prepare our text inputs: Task ID and Priority.
	// We’ll store them in a slice so we can cycle focus with Tab.
	m := model{
		inputs:     make([]textinput.Model, 2),
		focusIndex: 0,
		client:     client,
		activeTab:  tabSubmit,
	}

	// Task ID field.
	ti1 := textinput.New()
	ti1.Placeholder = "Enter Task ID"
	ti1.CharLimit = 32
	ti1.Width = 30
	ti1.Focus() // by default, set the first input as focused

	// Priority field.
	ti2 := textinput.New()
	ti2.Placeholder = "Priority (low, medium, high)"
	ti2.CharLimit = 10
	ti2.Width = 30

	m.inputs[0] = ti1
	m.inputs[1] = ti2

	return m
}

// Init is the standard Bubble Tea initialization function.
func (m model) Init() tea.Cmd {
	// Start the cursor blinking in whichever field is focused.
	return textinput.Blink
}

// --- Update ---

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		// Quit the entire application.
		case "ctrl+c", "q":
			return m, tea.Quit

		// Press tab to switch between the "Submit" tab and the "Jobs" tab.
		// This frees up the Tab key for switching between input fields.
		case "tab":
			if m.activeTab == tabSubmit {
				m.activeTab = tabJobs
				// Fetch tasks when switching to the Jobs tab.
				return m, fetchTasksCmd(m.client)
			} else {
				m.activeTab = tabSubmit
			}

		// Refresh tasks if we are in the Jobs tab.
		case "r":
			if m.activeTab == tabJobs {
				return m, fetchTasksCmd(m.client)
			}

		// Handle navigation and Enter key in the Submit tab.
		case "shift+tab", "up", "down", "enter":
			if m.activeTab == tabSubmit {
				s := msg.String()

				// If the user presses Enter and we have not yet submitted,
				// check if we are on the last input or not.
				if s == "enter" {
					// If we are at the last input and haven't submitted yet, submit.
					if m.focusIndex == len(m.inputs)-1 && !m.submitted {
						return m, submitTaskCmd(m)
					}

					// If we already submitted, pressing Enter checks status again.
					if m.focusIndex == len(m.inputs)-1 && m.submitted {
						return m, checkStatusCmd(m)
					}

					// Otherwise, move focus to the next input if there is one.
					m.focusIndex++
					if m.focusIndex > len(m.inputs)-1 {
						m.focusIndex = len(m.inputs) - 1
					}
				} else {
					// Cycle focus with tab/up/down
					if s == "shift+tab" || s == "up" {
						m.focusIndex--
					} else {
						m.focusIndex++
					}
					if m.focusIndex < 0 {
						m.focusIndex = 0
					} else if m.focusIndex > len(m.inputs)-1 {
						m.focusIndex = len(m.inputs) - 1
					}
				}

				// Update focus states accordingly.
				cmds = make([]tea.Cmd, len(m.inputs))
				for i := range m.inputs {
					if i == m.focusIndex {
						cmds[i] = m.inputs[i].Focus()
					} else {
						m.inputs[i].Blur()
					}
				}
				return m, tea.Batch(cmds...)
			}

		} // end switch msg.String()

	// Tasks fetched from the server.
	case listTasksMsg:
		m.tasks = msg
		m.err = nil
		return m, nil

	// Submission response (success or failure).
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

	// Task status response.
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

	// Error from any command.
	case errMsg:
		m.err = msg
		return m, nil
	}

	// If we’re on the Jobs tab, we don’t allow text editing in the inputs.
	// But if we’re on the Submit tab, forward character input to whichever input is focused.
	if m.activeTab == tabSubmit {
		cmd := m.updateInputs(msg)
		return m, cmd
	}

	return m, nil
}

// updateInputs updates whichever text input is currently focused.
func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		// Only the focused text input will actually accept keystrokes.
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return tea.Batch(cmds...)
}

// --- View ---

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf(
			"Error: %v\n\nPress 'q' to quit or 'tab' to switch tabs.\n",
			m.err,
		)
	}

	s := "=== Charm CLI with Tabs ===\n\n"

	switch m.activeTab {
	case tabSubmit:
		s += "[ Task Submission ] | (Press tab to switch to Jobs)\n\n"
		if !m.submitted {
			// If not yet submitted, show the two text inputs
			s += "Submit a new task:\n\n"
			s += "Task ID: " + m.inputs[0].View() + "\n\n"
			s += "Priority: " + m.inputs[1].View() + "\n\n"
			s += "Use Tab to move between fields. Press Enter on the second field to submit.\n"
			s += "Press 'q' to quit.\n"
		} else {
			// If already submitted, show the single set of "Task Submitted" + "Current Status".
			s += fmt.Sprintf(
				"Task Submitted.\n\nCurrent Status: %s\n\n",
				m.status,
			)
			s += "Press Enter on the second field to refresh status or 'q' to quit.\n"
			// If you want to keep displaying the input fields, you can show them, but not repeated messages:
			s += "\nTask ID: " + m.inputs[0].View() + "\n\n"
			s += "Priority: " + m.inputs[1].View() + "\n"
		}

	case tabJobs:
		s += "(Press tab  to switch to Submission) | [ Jobs List ]\n\n"
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

// --- Custom message types for asynchronous operations ---

type listTasksMsg []*pb.Task

type submitTaskResponseMsg struct {
	response *pb.SubmitTaskResponse
	err      error
}

type taskStatusResponseMsg struct {
	response *pb.TaskStatusResponse
	err      error
}

type errMsg error

// --- Commands ---

func fetchTasksCmd(client pb.SchedulerClient) tea.Cmd {
	return func() tea.Msg {
		res, err := client.ListTasks(context.Background(), &pb.TaskListRequest{})
		if err != nil {
			return errMsg(err)
		}
		return listTasksMsg(res.Tasks)
	}
}

func submitTaskCmd(m model) tea.Cmd {
	return func() tea.Msg {
		taskID := m.inputs[0].Value()
		priority := m.inputs[1].Value()

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

func checkStatusCmd(m model) tea.Cmd {
	return func() tea.Msg {
		taskID := m.inputs[0].Value()

		res, err := m.client.GetTaskStatus(context.Background(), &pb.TaskStatusRequest{
			TaskId: taskID,
		})

		return taskStatusResponseMsg{
			response: res,
			err:      err,
		}
	}
}

// --- Main ---

func main() {
	// Connect to gRPC server (insecure for demonstration).
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

	// Initialize the Bubble Tea program with our model.
	p := tea.NewProgram(initialModel(client))

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running Charm CLI:", err)
		os.Exit(1)
	}
}
