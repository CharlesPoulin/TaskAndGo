// schedulerimpl/strategy.go
package schedulerimpl

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	pb "github.com/charl/TaskAndGo/scheduler" // Proto package
)

// SchedulingStrategy describes how to distribute tasks/subtasks.
type SchedulingStrategy interface {
	DistributeTask(task *pb.Task, store *TaskStore)
	DistributeSubTask(subtask *pb.SubTask, store *TaskStore)
}

// ---------------------- SPLIT STRATEGY ----------------------
type SplitStrategy struct{}

func (s *SplitStrategy) DistributeTask(task *pb.Task, store *TaskStore) {
	// Mark as RUNNING
	store.mu.Lock()
	store.tasks[task.TaskId] = "RUNNING"
	store.mu.Unlock()

	// Simulate parallel distribution
	go func() {
		time.Sleep(3 * time.Second)
		store.mu.Lock()
		store.tasks[task.TaskId] = "COMPLETED"
		store.mu.Unlock()
		log.Printf("[SplitStrategy] Task %s completed (split across multiple nodes)", task.TaskId)
	}()
}

func (s *SplitStrategy) DistributeSubTask(subtask *pb.SubTask, store *TaskStore) {
	log.Printf("[SplitStrategy] Subtask %s distributed among nodes", subtask.SubtaskId)
}

// ---------------------- BATCH STRATEGY ----------------------
type BatchStrategy struct{}

func (b *BatchStrategy) DistributeTask(task *pb.Task, store *TaskStore) {
	// Mark as RUNNING
	store.mu.Lock()
	store.tasks[task.TaskId] = "RUNNING"
	store.mu.Unlock()

	go func() {
		codeParts := strings.SplitN(task.Data, ":", 2)
		if len(codeParts) != 2 {
			log.Printf("[BatchStrategy] Task %s: invalid code format", task.TaskId)
			markCompleted(task, store)
			return
		}
		lang, code := codeParts[0], codeParts[1]
		log.Printf("[BatchStrategy] Detected language %s for task %s", lang, task.TaskId)

		var cmd *exec.Cmd

		switch lang {
		case "python":
			// Write the Python code to a temporary file
			tmpFile, err := os.CreateTemp("", "py-task-*.py")
			if err != nil {
				log.Printf("[BatchStrategy] Error creating temp file: %v", err)
				markCompleted(task, store)
				return
			}
			defer os.Remove(tmpFile.Name())
			if _, err := tmpFile.WriteString(code); err != nil {
				log.Printf("[BatchStrategy] Error writing to temp file: %v", err)
				tmpFile.Close()
				markCompleted(task, store)
				return
			}
			tmpFile.Close()
			cmd = exec.Command("python3", tmpFile.Name())

		case "go":
			// Create a temporary directory for the Go module
			tmpDir, err := os.MkdirTemp("", "go-task-")
			if err != nil {
				log.Printf("[BatchStrategy] Error creating temp directory: %v", err)
				markCompleted(task, store)
				return
			}
			defer os.RemoveAll(tmpDir)

			// Write a minimal go.mod file to enable module mode.
			goModContent := "module task\n\ngo 1.20\n"
			err = ioutil.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte(goModContent), 0644)
			if err != nil {
				log.Printf("[BatchStrategy] Error writing go.mod: %v", err)
				markCompleted(task, store)
				return
			}

			// Write the provided Go code to main.go
			err = ioutil.WriteFile(filepath.Join(tmpDir, "main.go"), []byte(code), 0644)
			if err != nil {
				log.Printf("[BatchStrategy] Error writing main.go: %v", err)
				markCompleted(task, store)
				return
			}
			// Run the Go code in the temporary directory
			cmd = exec.Command("go", "run", ".")
			cmd.Dir = tmpDir

		default:
			log.Printf("[BatchStrategy] Unknown language %s for task %s", lang, task.TaskId)
			markCompleted(task, store)
			return
		}

		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("[BatchStrategy] Error executing code for task %s: %v\nOutput:\n%s", task.TaskId, err, string(output))
		} else {
			log.Printf("[BatchStrategy] Output for task %s:\n%s", task.TaskId, string(output))
		}

		markCompleted(task, store)
	}()
}

func (b *BatchStrategy) DistributeSubTask(subtask *pb.SubTask, store *TaskStore) {
	log.Printf("[BatchStrategy] Subtask %s assigned to a single node", subtask.SubtaskId)
}

func markCompleted(task *pb.Task, store *TaskStore) {
	time.Sleep(1 * time.Second)
	store.mu.Lock()
	store.tasks[task.TaskId] = "COMPLETED"
	store.mu.Unlock()
	log.Printf("[BatchStrategy] Task %s completed", task.TaskId)
}
