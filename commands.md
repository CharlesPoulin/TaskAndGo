grpcurl -plaintext -d '{"task_id":"taskABC"}' \
localhost:50051 scheduler.Scheduler/GetTaskStatus

grpcurl -plaintext -d '{"task":{"task_id":"taskABC","priority":"high"}}' \
localhost:50051 scheduler.Scheduler/SubmitTask

grpcurl -plaintext \
localhost:50051 scheduler.Scheduler/ListTasks


curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64
chmod +x skaffold
sudo mv skaffold /usr/local/bin/




grpcurl -plaintext -d '{"task":{"task_id":"taskABC","priority":"high"}}' \
localhost:50051 scheduler.Scheduler/SubmitTask

{
  "success": true,
  "message": "Task taskABC submitted successfully"
}
grpcurl -plaintext \
localhost:50051 scheduler.Scheduler/ListTasks

{
  "tasks": [
    {
      "taskId": "taskABC",
      "status": "COMPLETED"
    }
  ]
}


{
  "task": {
    "task_id": "pi_python_1",
    "priority": "high",
    "data": "python:import decimal\n\ndef calc_pi(n=1000):\n    decimal.getcontext().prec = n\n    # Machin-like formula or other approach\n    # For brevity, here is a small snippet:\n    three = decimal.Decimal(3)\n    return three * (decimal.Decimal(1) - (decimal.Decimal(1)/decimal.Decimal(3* n)))\n\nprint(calc_pi())",
    "status": ""
  }
}


grpcurl -plaintext \
  -d '{
  "task_id": "pi_go_1",
  "priority": "high","data": "go:package main\n\nimport (\n\t\"fmt\"\n\t\"math/big\"\n)\n\nfunc main() {\n\tfmt.Println(piNilakantha(1000))\n}\n\nfunc piNilakantha(terms int) *big.Float {\n\t... etc\n}\n"}' \
  localhost:50051 \
  scheduler.Scheduler/SubmitTask
