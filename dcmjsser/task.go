package main

const (
	Started = iota
	Failed
	Done
)

type Descriptable interface {
	GetDescript() string
}

type IsVerifiable interface {
	IsConflict(IsVerifiable) bool
}
type Job struct {
	JobId string
	Data  interface{}
}

type FaJob struct {
	Job       Job
	ErrorData interface{}
}

type CompJob struct {
	Job        Job
	ResultData interface{}
}

type TermJob struct {
}
