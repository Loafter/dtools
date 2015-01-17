package main

import "net/http"

type Tasker interface {
	TaskId() string
	TaskType() int
	IsConflict(tasker Tasker) bool
}
type BasicTask struct {
	Type int    `json:"Type,string,omitempty"`
	Id   string `json:"Id,string,omitempty"`
}

func (basicTask *BasicTask) TaskType() int {
	return basicTask.Type
}

func (basicTask *BasicTask) IsConflict(tasker Tasker) bool {
	return basicTask.Type == tasker.TaskType()
}

type TerminateDispatchJob struct {
}

type TaskResponse struct {
	Tasker         Tasker
	responseWriter http.ResponseWriter
	httpRequest    *http.Request
}

type TaskError struct {
	Error          error
	Tasker         Tasker
	responseWriter http.ResponseWriter
	httpRequest    *http.Request
}
