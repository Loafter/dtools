package main

const (
	Started = iota
	Failed
	Done
)

type IsVerifiable interface {
	IsConflict(IsVerifiable) bool
}
type Job struct {
	JobId   string
	JobData interface{}
}

type FailedJob struct {
	JobId     string
	ErrorData error
	Data      interface{}
}

type DoneJob struct {
	JobId string
	Data  interface{}
}

type TerminateDispatchJob struct {
}
