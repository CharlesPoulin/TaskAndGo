# Distributed Task Scheduler

A lightweight, distributed task scheduler built in Go for running prioritized jobs across multiple nodes in a Kubernetes cluster. This project provides an open‐source, scalable alternative to proprietary task scheduling tools.

> **Note:** A task’s `data` field contains the code to be executed. The code must be prefixed with a language identifier (e.g. `"python:"` or `"go:"`). The scheduler’s strategy extracts the prefix, writes the code to a temporary file (or directory), and executes it with the appropriate interpreter.

---

## Key Concepts

### 1. Splitting a Single Task Across Multiple Nodes

**Example:**  
A large data processing job is divided into several subtasks, each dispatched to a different node for parallel processing.

### 2. Sending Batches of Tasks to Available Nodes

**Example:**  
A set of independent jobs (e.g. code execution tasks) is submitted as a batch. Each task is assigned to a single node for execution.

---

## Roadmap

A visual representation of the design and task flows can be found on our [Miro board](https://miro.com/app/board/uXjVLw2tOPU=/?share_link_id=323582765789).

---

## Features

- **gRPC-Based Communication:** Efficient, low-latency communication between nodes.
- **Priority-Based Queue:** Critical tasks are executed first.
- **Monitoring Dashboard:** Real-time task status and resource usage via a CLI or simple frontend.
- **Kubernetes Deployment:** Fully Dockerized and deployable on Kubernetes with auto-scaling.
- **Code Execution Tasks:** Supports executing code snippets in Go or Python. The `data` field is used to store the code (prefixed with the language).

---

## Use Cases

- Distributed data processing pipelines.
- CI/CD workflows and DevOps automation.
- Machine learning task distribution.
- General-purpose job scheduling in distributed systems.
- Running code tasks (e.g. calculating 1000 decimal digits of π with Go or Python).

---

## Technologies

- **Programming Language:** Go
- **Communication Protocol:** gRPC
- **Containerization:** Docker
- **Orchestration:** Kubernetes
- **Monitoring:** CLI / Custom Dashboard

---

## Setup and Installation

### Prerequisites

- **Go** (v1.20+)
- **Docker**
- **Kubernetes** (e.g., Minikube)
- *(Optional)* **grpcurl** for testing

### Clone the Repository

```bash
git clone https://github.com/yourusername/distributed-task-scheduler.git
cd distributed-task-scheduler
```
> **Note:** Tasks are submitted via gRPC. Each task payload includes:
> - **task_id:** Unique identifier.
> - **priority:** e.g., `"high"`.
> - **data:** The code to execute, prefixed with a language identifier:
>   - Python Tasks: `"python:<python code>"`
>   - Go Tasks: `"go:<go code>"`

---

## Submitting Tasks

Tasks are submitted via gRPC. Each task payload includes the fields mentioned above.

### Example Using grpcurl

Submit a Go task to calculate π:

```bash
grpcurl -plaintext -d '{
  "task": {
    "task_id": "pi_go_1",
    "priority": "high",
    "data": "go:package main\n\nimport (\n\t\"fmt\"\n\t\"math/big\"\n)\n\nfunc main() {\n\tpi := piNilakantha(1000)\n\tfmt.Println(pi.Text(\'f\', 1000))\n}\n\nfunc piNilakantha(terms int) *big.Float {\n\tpi := big.NewFloat(3.0)\n\tsign := 1.0\n\tindex := 2.0\n\tfor i := 0; i < terms; i++ {\n\t\td1 := big.NewFloat(index)\n\t\td2 := big.NewFloat(index + 1)\n\t\td3 := big.NewFloat(index + 2)\n\t\td := new(big.Float).Mul(d1, d2)\n\t\td.Mul(d, d3)\n\t\tfrac := new(big.Float).Quo(big.NewFloat(4.0), d)\n\t\tif i%%2 == 0 {\n\t\t\tpi.Add(pi, frac)\n\t\t} else {\n\t\t\tpi.Sub(pi, frac)\n\t\t}\n\t\tindex += 2.0\n\t}\n\treturn pi\n}"
  }
}' localhost:50051 scheduler.Scheduler/SubmitTask


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
export SCHEDULER_STRATEGY=batch
go run cmd/server/main.go
```



docker build -t cpoulin/taskandgo:latest .
docker run -p 50051:50051 cpoulin/taskandgo:latest


Kubernetes

minikube start
it should ike that : cpoulin@DESKTOP-053J5O1:~$ minikube start
😄  minikube v1.35.0 on Ubuntu 22.04 (amd64)
✨  Automatically selected the docker driver. Other choices: none, ssh
📌  Using Docker driver with root privileges
👍  Starting "minikube" primary control-plane node in "minikube" cluster
🚜  Pulling base image v0.0.46 ...
💾  Downloading Kubernetes v1.32.0 preload ...
    > gcr.io/k8s-minikube/kicbase...:  500.31 MiB / 500.31 MiB  100.00% 7.91 Mi
    > preloaded-images-k8s-v18-v1...:  333.57 MiB / 333.57 MiB  100.00% 4.32 Mi
🔥  Creating docker container (CPUs=2, Memory=3900MB) ...
🐳  Preparing Kubernetes v1.32.0 on Docker 27.4.1 ...
    ▪ Generating certificates and keys ...
    ▪ Booting up control plane ...
    ▪ Configuring RBAC rules ...
🔗  Configuring bridge CNI (Container Networking Interface) ...
🔎  Verifying Kubernetes components...
    ▪ Using image gcr.io/k8s-minikube/storage-provisioner:v5
🌟  Enabled addons: storage-provisioner, default-storageclass
💡  kubectl not found. If you need it, try: 'minikube kubectl -- get pods -A'
🏄  Done! kubectl is now configured to use "minikube" cluster and "default" namespace by default

kubectl apply -f k8s-deployment.yaml

sudo snap install kubectl --classic

alias kubectl="minikube kubectl --"


kubectl get pods
kubectl logs <name_of_pod>

service/taskandgo-service created
cpoulin@DESKTOP-053J5O1:~/project/TaskAndGo$ kubectl get pods
NAME                                    READY   STATUS    RESTARTS   AGE
taskandgo-deployment-7d59dfdc7d-hdmfj   1/1     Running   0          13s
cpoulin@DESKTOP-053J5O1:~/project/TaskAndGo$ kubectl logs taskandgo-deployment-7d59dfdc7d-hdmfj
Server started with strategy: *schedulerimpl.BatchStrategy on port 50051
1


# Forward local port 50051 to the service in your cluster
kubectl port-forward deployment/taskandgo-deployment 50051:50051

# Test the SubmitTask RPC
grpcurl -plaintext -d '{"task":{"task_id":"task123","priority":"high"}}' \
  localhost:50051 scheduler.Scheduler/SubmitTask









Where i am 

cpoulin@DESKTOP-053J5O1:~/project/TaskAndGo$ grpcurl -plaintext -d '{
  "task": {
    "task_id": "pi_go_2",
    "priority": "high",
    "data": "go:package main\n\nimport (\n\t\"fmt\"\n\t\"math/big\"\n)\n\nfunc main() {\n\tfmt.Println(piNilakantha
(1000))\n}\n\nfunc piNilakantha(terms int) *big.Float {\n\tpi := big.NewFloat(3.0)\n\tone := big.NewFloat(1.0)\n\tt
wo := big.NewFloat(2.0)\n\tfour := big.NewFloat(4.0)\n\tsign := one\n\tdenom := big.NewFloat(2.0)\n\n\tfor i := 0; 
i < terms; i++ {\n\t\tterm := new(big.Float).Quo(four, new(big.Float).Mul(denom, new(big.Float).Mul(new(big.Float).
Add(denom, one), new(big.Float).Add(denom, two))))\n\t\tpi.Add(pi, new(big.Float).Mul(sign, term))\n\n\t\tsign.Neg(
sign)\n\t\tdenom.Add(denom, two)\n\t}\n\n\treturn pi\n}\n"
  }
}' localhost:50051 scheduler.Scheduler/SubmitTask
{
  "success": true,
  "message": "Task pi_go_2 submitted successfully"
}

[taskandgo-container] 2025/02/19 17:04:59 Received task pi_go_2 with priority high
[taskandgo-container] 2025/02/19 17:04:59 [BatchStrategy] Detected language go for task pi_go_2
[taskandgo-container] 2025/02/19 17:04:59 [BatchStrategy] Output for task pi_go_2:
[taskandgo-container] 3.1172378336938116
[taskandgo-container] 2025/02/19 17:05:00 [BatchStrategy] Task pi_go_2 completed
