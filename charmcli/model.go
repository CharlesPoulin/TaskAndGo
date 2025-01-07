// charmcli/model.go
package charmcli

import (
	pb "github.com/charl/TaskAndGo/scheduler"
	"github.com/charmbracelet/bubbles/textinput"
)

// tabIndex enumerates active tabs.
type tabIndex int

const (
	tabSubmit tabIndex = iota
	tabJobs
	tabResource
)

// Model defines the overall state.
type Model struct {
	focusIndex int
	inputs     []textinput.Model

	client    pb.SchedulerClient
	err       error
	submitted bool
	status    string
	activeTab tabIndex

	tasks     []*pb.Task
	resources []*pb.ResourceUsage
}

// InitialModel initializes the application state.
func InitialModel(client pb.SchedulerClient) Model {
	m := Model{
		inputs:     make([]textinput.Model, 2),
		focusIndex: 0,
		client:     client,
		activeTab:  tabSubmit,
	}

	// Initialize Task ID field.
	ti1 := textinput.New()
	ti1.Placeholder = "Enter Task ID"
	ti1.CharLimit = 32
	ti1.Width = 30
	ti1.Focus()

	// Initialize Priority field.
	ti2 := textinput.New()
	ti2.Placeholder = "Priority (low, medium, high)"
	ti2.CharLimit = 10
	ti2.Width = 30

	m.inputs[0] = ti1
	m.inputs[1] = ti2

	return m
}

// ResetSubmission clears inputs and resets submission state.
func (m *Model) ResetSubmission() {
	for i := range m.inputs {
		m.inputs[i].SetValue("")
		m.inputs[i].Blur()
	}
	m.inputs[0].Focus()
	m.focusIndex = 0
	m.submitted = false
	m.status = ""
	m.err = nil
}
