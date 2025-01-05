# Distributed Task Scheduler

A lightweight, distributed task scheduler built in Go for running prioritized jobs across multiple nodes in a Kubernetes cluster. This project provides an open-source, scalable alternative to proprietary task scheduling tools.

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
- **Monitoring**: Web-based frontend or CLI for observability

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
