# Distributed Task Scheduler

A lightweight, distributed task scheduler built in Go for running prioritized jobs across multiple nodes in a Kubernetes cluster. This project provides an open-source, scalable alternative to proprietary task scheduling tools.

1. Splitting a Single Task Across Multiple Nodes
example:


2. Sending Batches of Tasks to Available Nodes
example:

--- 

Roadmap

https://miro.com/app/board/uXjVLw2tOPU=/?share_link_id=323582765789 


---

## **Features**
- **gRPC-Based Communication**: Efficient communication between nodes for task delegation and status updates.
- **Priority-Based Queue**: Ensures critical tasks are executed first.
- **Monitoring Dashboard**: Real-time task status and resource usage via a simple frontend or CLI.
- **Kubernetes Deployment**: Fully Dockerized and deployable on Kubernetes with auto-scaling capabilities.

---

## **Use Cases**
- Distributed data processing pipelines.
- CI/CD workflows and DevOps automation.
- Machine learning task distribution.
- General-purpose job scheduling in distributed systems.

---

## **Technologies**
- **Programming Language**: Go
- **Communication Protocol**: gRPC
- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **Monitoring**: CLI for observability

---

## **Setup and Installation**
### Prerequisites
1. Install Go (v1.20+)
2. Install Docker
3. Install Kubernetes (Minikube or any cluster)

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/distributed-task-scheduler.git
   cd distributed-task-scheduler



   # Usage Guide for the Distributed Task Scheduler ```

## **Submit Tasks**
You can submit tasks to the scheduler via CLI or API using a `POST` request. Replace `<scheduler-node-ip>` and `<port>` with the appropriate values for your setup.

### Example Command:
```bash
curl -X POST http://<scheduler-node-ip>:<port>/submit-task \
-d '{"task_id": "123", "priority": "high", "data": "task details"}'

# TaskAndGo Scheduler

This project implements a gRPC-based task scheduler that can distribute tasks among nodes in different ways. The scheduler supports two primary “scheduling strategies”:

1. **Split Strategy** – Splits a single task into parallel chunks across multiple nodes (useful for big, parallelizable tasks).
2. **Batch Strategy** – Sends entire tasks to one node at a time for simpler “batch” processing.

## Architecture

- **scheduler/strategy.go** – Contains the `SchedulingStrategy` interface and two concrete strategies:
  - `SplitStrategy`
  - `BatchStrategy`
- **scheduler/scheduler_server.go** – Defines the `SchedulerServer` gRPC service implementation.  
  It uses a `TaskStore` (in-memory) to track tasks and a `SchedulingStrategy` to decide how those tasks are dispatched.
- **cmd/server/main.go** – The entry point for running the gRPC server. You can set an environment variable (`SCHEDULER_STRATEGY=batch` or `SCHEDULER_STRATEGY=split`) to choose which strategy the server will use at runtime.

## Running the Server

```bash
# Default: uses split strategy
go run cmd/server/main.go

# Force batch strategy
SCHEDULER_STRATEGY=batch go run cmd/server/main.go
