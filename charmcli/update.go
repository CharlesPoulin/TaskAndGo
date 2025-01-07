package charmcli

import (
	"fmt"

	pb "github.com/charl/TaskAndGo/scheduler"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

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

// Init starts cursor blinking.
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles all incoming messages and updates the model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "tab":
			if m.activeTab == tabSubmit {
				m.ResetSubmission()
				m.activeTab = tabJobs
				return m, fetchTasksCmd(m.client)
			} else {
				m.ResetSubmission()
				m.activeTab = tabSubmit
			}

		case "r":
			if m.activeTab == tabJobs {
				return m, fetchTasksCmd(m.client)
			}

		case "shift+tab", "up", "down", "enter":
			if m.activeTab == tabSubmit {
				s := msg.String()
				if s == "enter" {
					if m.focusIndex == len(m.inputs)-1 && !m.submitted {
						return m, submitTaskCmd(m)
					}
					if m.focusIndex == len(m.inputs)-1 && m.submitted {
						return m, checkStatusCmd(m)
					}
					m.focusIndex++
					if m.focusIndex > len(m.inputs)-1 {
						m.focusIndex = len(m.inputs) - 1
					}
				} else {
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

		}

	case listTasksMsg:
		m.tasks = msg
		m.err = nil
		return m, nil

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

	case errMsg:
		m.err = msg
		return m, nil
	}

	if m.activeTab == tabSubmit {
		return m, m.updateInputs(msg)
	}

	return m, nil
}

// updateInputs forwards character input to the focused text input.
func (m *Model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return tea.Batch(cmds...)
}
