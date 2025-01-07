// charmcli/view.go
package charmcli

import (
	"fmt"
)

// Minimal ANSI escape codes for coloring
const (
	colorTitle  = "\033[1;36m" // Bold Cyan
	colorHeader = "\033[1;33m" // Bold Yellow
	colorReset  = "\033[0m"
)

// View renders the current UI based on the active tab and state.
func (m Model) View() string {
	// Handle error first
	if m.err != nil {
		return fmt.Sprintf(
			"%sError:%s %v\n\nPress 'q' to quit or 'tab' to switch tabs.\n",
			colorHeader, colorReset, m.err,
		)
	}

	// Title line
	s := fmt.Sprintf("%s=== Charm CLI With Tabs ===%s\n\n", colorTitle, colorReset)

	switch m.activeTab {
	// ─────────────────────────────────────────
	// Tab 1: Submit
	// ─────────────────────────────────────────
	case tabSubmit:
		s += fmt.Sprintf(
			"%s[ TASK SUBMISSION ]%s | (Press TAB to switch) | Press 'q' to quit.\n\n",
			colorHeader, colorReset,
		)
		if !m.submitted {
			s += "Submit a new task:\n\n"
			s += "Task ID: " + m.inputs[0].View() + "\n\n"
			s += "Priority: " + m.inputs[1].View() + "\n\n"
			s += "Use TAB to move between fields.\n" +
				"Press Enter on the second field to submit.\n"
		} else {
			s += fmt.Sprintf("Task submitted.\n\nCurrent Status: %s\n\n", m.status)
			s += "Press Enter on the second field to refresh status.\n"
			s += "\nTask ID: " + m.inputs[0].View() + "\n\n"
			s += "Priority: " + m.inputs[1].View() + "\n"
		}

	// ─────────────────────────────────────────
	// Tab 2: Jobs
	// ─────────────────────────────────────────
	case tabJobs:
		s += fmt.Sprintf(
			"(Press TAB to switch) | %s[ JOBS LIST ]%s\n\n",
			colorHeader, colorReset,
		)
		s += "Press 'r' to refresh the tasks list.\n\n"

		if len(m.tasks) == 0 {
			s += "No tasks to display or still loading..."
		} else {
			s += "Current Tasks:\n\n"
			for _, t := range m.tasks {
				// If t.Progress is in your proto, show X%
				// Otherwise, adapt or remove this portion
				s += fmt.Sprintf(
					" • %s (Priority: %s) => Status: %s | Progress: %d%%\n",
					t.TaskId, t.Priority, t.Status, t.Progress,
				)
			}
		}

	// ─────────────────────────────────────────
	// Tab 3: Resource Usage
	// ─────────────────────────────────────────
	case tabResource:
		s += fmt.Sprintf(
			"(Press TAB to switch) | %s[ RESOURCE USAGE ]%s\n\n",
			colorHeader, colorReset,
		)
		s += "Press 'r' to refresh resource usage.\n\n"

		if len(m.resources) == 0 {
			s += "No resource data to display or still loading..."
		} else {
			s += "Node resource usage:\n\n"
			for _, r := range m.resources {
				// Adjust fields according to your pb.ResourceUsage struct
				s += fmt.Sprintf(
					" • Node: %s => CPU: %d%% | Mem: %d%%\n",
					r.NodeName, r.CpuPercent, r.MemPercent,
				)
			}
		}
	}

	s += "\n"
	return s
}
