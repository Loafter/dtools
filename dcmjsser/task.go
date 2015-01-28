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
	JobId         string
	Data          interface{}
	DataToSuccess interface{}
	DataToError   interface{}
}

type FailedJob struct {
	JobId     string
	ErrorData interface{}
}

type CompletedJob struct {
	JobId      string
	ResultData interface{}
}

type TerminateDispatchJob struct {
}
