package charmcli

import (
	"fmt"
)

// View renders the current UI based on the active tab and state.
func (m Model) View() string {
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
			s += "Submit a new task:\n\n"
			s += "Task ID: " + m.inputs[0].View() + "\n\n"
			s += "Priority: " + m.inputs[1].View() + "\n\n"
			s += "Use Tab to move between fields. Press Enter on the second field to submit.\n"
			s += "Press 'q' to quit.\n"
		} else {
			s += fmt.Sprintf(
				"Task Submitted.\n\nCurrent Status: %s\n\n",
				m.status,
			)
			s += "Press Enter on the second field to refresh status or 'q' to quit.\n"
			s += "\nTask ID: " + m.inputs[0].View() + "\n\n"
			s += "Priority: " + m.inputs[1].View() + "\n"
		}

	case tabJobs:
		s += "(Press tab to switch to Submission) | [ Jobs List ]\n\n"
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
