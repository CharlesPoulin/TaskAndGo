package charmcli

import (
	"context"
	"fmt"

	pb "github.com/charl/TaskAndGo/scheduler"
	tea "github.com/charmbracelet/bubbletea"
)

// fetchTasksCmd fetches tasks from the server.
func fetchTasksCmd(client pb.SchedulerClient) tea.Cmd {
	return func() tea.Msg {
		res, err := client.ListTasks(context.Background(), &pb.TaskListRequest{})
		if err != nil {
			return errMsg(err)
		}
		return listTasksMsg(res.Tasks)
	}
}

// submitTaskCmd submits a new task.
func submitTaskCmd(m Model) tea.Cmd {
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

		task := &pb.Task{
			TaskId:   taskID,
			Priority: priority,
			Data:     "Sample Data",
			Status:   "QUEUED",
		}

		res, err := m.client.SubmitTask(context.Background(), &pb.SubmitTaskRequest{
			Task: task,
		})

		return submitTaskResponseMsg{
			response: res,
			err:      err,
		}
	}
}

// checkStatusCmd checks the status of a submitted task.
func checkStatusCmd(m Model) tea.Cmd {
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
