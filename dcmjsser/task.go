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
	JobId string
	Data  interface{}
}

type FailedJob struct {
	Job       Job
	ErrorData interface{}
}

type CompletedJob struct {
	Job        Job
	ResultData interface{}
}

type TerminateDispatchJob struct {
}
