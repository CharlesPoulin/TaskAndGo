// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: scheduler/scheduler.proto

package scheduler

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The Task message describes a task to be scheduled.
type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Priority      string                 `protobuf:"bytes,2,opt,name=priority,proto3" json:"priority,omitempty"` // e.g., "low", "medium", "high"
	Data          string                 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`         // Arbitrary data or payload
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Progress      int32                  `protobuf:"varint,6,opt,name=progress,proto3" json:"progress,omitempty"` // Percentage progress of the task (0-100)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_scheduler_scheduler_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *Task) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

func (x *Task) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Task) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Task) GetProgress() int32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

// New message representing a subtask.
type SubTask struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SubtaskId     string                 `protobuf:"bytes,1,opt,name=subtask_id,json=subtaskId,proto3" json:"subtask_id,omitempty"`
	ParentTaskId  string                 `protobuf:"bytes,2,opt,name=parent_task_id,json=parentTaskId,proto3" json:"parent_task_id,omitempty"`
	Data          string                 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"` // Additional fields: priority, status, progress for the subtask could be added here.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubTask) Reset() {
	*x = SubTask{}
	mi := &file_scheduler_scheduler_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubTask) ProtoMessage() {}

func (x *SubTask) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubTask.ProtoReflect.Descriptor instead.
func (*SubTask) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{1}
}

func (x *SubTask) GetSubtaskId() string {
	if x != nil {
		return x.SubtaskId
	}
	return ""
}

func (x *SubTask) GetParentTaskId() string {
	if x != nil {
		return x.ParentTaskId
	}
	return ""
}

func (x *SubTask) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// Request and response messages for submitting a subtask.
type SubmitSubTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Subtask       *SubTask               `protobuf:"bytes,1,opt,name=subtask,proto3" json:"subtask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitSubTaskRequest) Reset() {
	*x = SubmitSubTaskRequest{}
	mi := &file_scheduler_scheduler_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitSubTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitSubTaskRequest) ProtoMessage() {}

func (x *SubmitSubTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitSubTaskRequest.ProtoReflect.Descriptor instead.
func (*SubmitSubTaskRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitSubTaskRequest) GetSubtask() *SubTask {
	if x != nil {
		return x.Subtask
	}
	return nil
}

type SubmitSubTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitSubTaskResponse) Reset() {
	*x = SubmitSubTaskResponse{}
	mi := &file_scheduler_scheduler_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitSubTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitSubTaskResponse) ProtoMessage() {}

func (x *SubmitSubTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitSubTaskResponse.ProtoReflect.Descriptor instead.
func (*SubmitSubTaskResponse) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{3}
}

func (x *SubmitSubTaskResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SubmitSubTaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The SubmitTaskRequest encapsulates the Task object.
type SubmitTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          *Task                  `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitTaskRequest) Reset() {
	*x = SubmitTaskRequest{}
	mi := &file_scheduler_scheduler_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitTaskRequest) ProtoMessage() {}

func (x *SubmitTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitTaskRequest.ProtoReflect.Descriptor instead.
func (*SubmitTaskRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{4}
}

func (x *SubmitTaskRequest) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

// The SubmitTaskResponse returns an acknowledgment or error details.
type SubmitTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitTaskResponse) Reset() {
	*x = SubmitTaskResponse{}
	mi := &file_scheduler_scheduler_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitTaskResponse) ProtoMessage() {}

func (x *SubmitTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitTaskResponse.ProtoReflect.Descriptor instead.
func (*SubmitTaskResponse) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{5}
}

func (x *SubmitTaskResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SubmitTaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Request for retrieving a task’s status.
type TaskStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskStatusRequest) Reset() {
	*x = TaskStatusRequest{}
	mi := &file_scheduler_scheduler_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskStatusRequest) ProtoMessage() {}

func (x *TaskStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskStatusRequest.ProtoReflect.Descriptor instead.
func (*TaskStatusRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{6}
}

func (x *TaskStatusRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

// TaskStatusResponse returns the status of a given task.
type TaskStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Found         bool                   `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`                // Whether the task was found
	CurrentStatus string                 `protobuf:"bytes,2,opt,name=currentStatus,proto3" json:"currentStatus,omitempty"` // e.g., "QUEUED", "RUNNING", "COMPLETED"
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`             // Additional info (e.g., logs)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskStatusResponse) Reset() {
	*x = TaskStatusResponse{}
	mi := &file_scheduler_scheduler_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskStatusResponse) ProtoMessage() {}

func (x *TaskStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskStatusResponse.ProtoReflect.Descriptor instead.
func (*TaskStatusResponse) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{7}
}

func (x *TaskStatusResponse) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

func (x *TaskStatusResponse) GetCurrentStatus() string {
	if x != nil {
		return x.CurrentStatus
	}
	return ""
}

func (x *TaskStatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Request for listing all tasks.
type TaskListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskListRequest) Reset() {
	*x = TaskListRequest{}
	mi := &file_scheduler_scheduler_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListRequest) ProtoMessage() {}

func (x *TaskListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListRequest.ProtoReflect.Descriptor instead.
func (*TaskListRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{8}
}

// Response containing a list of tasks.
type TaskListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*Task                `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskListResponse) Reset() {
	*x = TaskListResponse{}
	mi := &file_scheduler_scheduler_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListResponse) ProtoMessage() {}

func (x *TaskListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListResponse.ProtoReflect.Descriptor instead.
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{9}
}

func (x *TaskListResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type RegisterNodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NodeId        string                 `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Capacity      string                 `protobuf:"bytes,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterNodeRequest) Reset() {
	*x = RegisterNodeRequest{}
	mi := &file_scheduler_scheduler_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNodeRequest) ProtoMessage() {}

func (x *RegisterNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNodeRequest.ProtoReflect.Descriptor instead.
func (*RegisterNodeRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{10}
}

func (x *RegisterNodeRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *RegisterNodeRequest) GetCapacity() string {
	if x != nil {
		return x.Capacity
	}
	return ""
}

type RegisterNodeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterNodeResponse) Reset() {
	*x = RegisterNodeResponse{}
	mi := &file_scheduler_scheduler_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNodeResponse) ProtoMessage() {}

func (x *RegisterNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNodeResponse.ProtoReflect.Descriptor instead.
func (*RegisterNodeResponse) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{11}
}

func (x *RegisterNodeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RegisterNodeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Request for getting resource usage.
type ResourceUsageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceUsageRequest) Reset() {
	*x = ResourceUsageRequest{}
	mi := &file_scheduler_scheduler_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceUsageRequest) ProtoMessage() {}

func (x *ResourceUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceUsageRequest.ProtoReflect.Descriptor instead.
func (*ResourceUsageRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{12}
}

// Representation of resource usage for a single node.
type ResourceUsage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NodeName      string                 `protobuf:"bytes,1,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	CpuPercent    int32                  `protobuf:"varint,2,opt,name=cpuPercent,proto3" json:"cpuPercent,omitempty"`
	MemPercent    int32                  `protobuf:"varint,3,opt,name=memPercent,proto3" json:"memPercent,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceUsage) Reset() {
	*x = ResourceUsage{}
	mi := &file_scheduler_scheduler_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceUsage) ProtoMessage() {}

func (x *ResourceUsage) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceUsage.ProtoReflect.Descriptor instead.
func (*ResourceUsage) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{13}
}

func (x *ResourceUsage) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *ResourceUsage) GetCpuPercent() int32 {
	if x != nil {
		return x.CpuPercent
	}
	return 0
}

func (x *ResourceUsage) GetMemPercent() int32 {
	if x != nil {
		return x.MemPercent
	}
	return 0
}

// Response containing a list of resource usages.
type ResourceUsageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Usages        []*ResourceUsage       `protobuf:"bytes,1,rep,name=usages,proto3" json:"usages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceUsageResponse) Reset() {
	*x = ResourceUsageResponse{}
	mi := &file_scheduler_scheduler_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceUsageResponse) ProtoMessage() {}

func (x *ResourceUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_scheduler_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceUsageResponse.ProtoReflect.Descriptor instead.
func (*ResourceUsageResponse) Descriptor() ([]byte, []int) {
	return file_scheduler_scheduler_proto_rawDescGZIP(), []int{14}
}

func (x *ResourceUsageResponse) GetUsages() []*ResourceUsage {
	if x != nil {
		return x.Usages
	}
	return nil
}

var File_scheduler_scheduler_proto protoreflect.FileDescriptor

var file_scheduler_scheduler_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22, 0x83, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x62, 0x0a, 0x07,
	0x53, 0x75, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x44, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x53, 0x75, 0x62, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x74,
	0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x4b, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x53, 0x75, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x38, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x48, 0x0a,
	0x12, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x6a, 0x0a, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x11, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22,
	0x4a, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x22, 0x4a, 0x0a, 0x14, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x6b, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x70, 0x75, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x63, 0x70, 0x75, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x6d, 0x65, 0x6d, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x22, 0x49, 0x0a, 0x15,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x75, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x06, 0x75, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0x95, 0x03, 0x0a, 0x09, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x52, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x53, 0x75, 0x62, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x1f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x53, 0x75, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x53, 0x75, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12,
	0x1a, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x16, 0x5a, 0x14, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x3b, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scheduler_scheduler_proto_rawDescOnce sync.Once
	file_scheduler_scheduler_proto_rawDescData = file_scheduler_scheduler_proto_rawDesc
)

func file_scheduler_scheduler_proto_rawDescGZIP() []byte {
	file_scheduler_scheduler_proto_rawDescOnce.Do(func() {
		file_scheduler_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_scheduler_scheduler_proto_rawDescData)
	})
	return file_scheduler_scheduler_proto_rawDescData
}

var file_scheduler_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_scheduler_scheduler_proto_goTypes = []any{
	(*Task)(nil),                  // 0: scheduler.Task
	(*SubTask)(nil),               // 1: scheduler.SubTask
	(*SubmitSubTaskRequest)(nil),  // 2: scheduler.SubmitSubTaskRequest
	(*SubmitSubTaskResponse)(nil), // 3: scheduler.SubmitSubTaskResponse
	(*SubmitTaskRequest)(nil),     // 4: scheduler.SubmitTaskRequest
	(*SubmitTaskResponse)(nil),    // 5: scheduler.SubmitTaskResponse
	(*TaskStatusRequest)(nil),     // 6: scheduler.TaskStatusRequest
	(*TaskStatusResponse)(nil),    // 7: scheduler.TaskStatusResponse
	(*TaskListRequest)(nil),       // 8: scheduler.TaskListRequest
	(*TaskListResponse)(nil),      // 9: scheduler.TaskListResponse
	(*RegisterNodeRequest)(nil),   // 10: scheduler.RegisterNodeRequest
	(*RegisterNodeResponse)(nil),  // 11: scheduler.RegisterNodeResponse
	(*ResourceUsageRequest)(nil),  // 12: scheduler.ResourceUsageRequest
	(*ResourceUsage)(nil),         // 13: scheduler.ResourceUsage
	(*ResourceUsageResponse)(nil), // 14: scheduler.ResourceUsageResponse
}
var file_scheduler_scheduler_proto_depIdxs = []int32{
	1,  // 0: scheduler.SubmitSubTaskRequest.subtask:type_name -> scheduler.SubTask
	0,  // 1: scheduler.SubmitTaskRequest.task:type_name -> scheduler.Task
	0,  // 2: scheduler.TaskListResponse.tasks:type_name -> scheduler.Task
	13, // 3: scheduler.ResourceUsageResponse.usages:type_name -> scheduler.ResourceUsage
	4,  // 4: scheduler.Scheduler.SubmitTask:input_type -> scheduler.SubmitTaskRequest
	2,  // 5: scheduler.Scheduler.SubmitSubTask:input_type -> scheduler.SubmitSubTaskRequest
	6,  // 6: scheduler.Scheduler.GetTaskStatus:input_type -> scheduler.TaskStatusRequest
	8,  // 7: scheduler.Scheduler.ListTasks:input_type -> scheduler.TaskListRequest
	12, // 8: scheduler.Scheduler.GetResourceUsage:input_type -> scheduler.ResourceUsageRequest
	5,  // 9: scheduler.Scheduler.SubmitTask:output_type -> scheduler.SubmitTaskResponse
	3,  // 10: scheduler.Scheduler.SubmitSubTask:output_type -> scheduler.SubmitSubTaskResponse
	7,  // 11: scheduler.Scheduler.GetTaskStatus:output_type -> scheduler.TaskStatusResponse
	9,  // 12: scheduler.Scheduler.ListTasks:output_type -> scheduler.TaskListResponse
	14, // 13: scheduler.Scheduler.GetResourceUsage:output_type -> scheduler.ResourceUsageResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_scheduler_scheduler_proto_init() }
func file_scheduler_scheduler_proto_init() {
	if File_scheduler_scheduler_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scheduler_scheduler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scheduler_scheduler_proto_goTypes,
		DependencyIndexes: file_scheduler_scheduler_proto_depIdxs,
		MessageInfos:      file_scheduler_scheduler_proto_msgTypes,
	}.Build()
	File_scheduler_scheduler_proto = out.File
	file_scheduler_scheduler_proto_rawDesc = nil
	file_scheduler_scheduler_proto_goTypes = nil
	file_scheduler_scheduler_proto_depIdxs = nil
}
