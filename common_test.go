package main

import "testing"
import "log"

import "time"
import "errors"

//import "sync"

type TestJobDispatcher struct {
}

type TestErrorDispatcher struct {
}

func DispatchTh(jobd interface{}, resultChan chan interface{}) {
	job := jobd.(Job)
	for i := 0; i < 9999; i++ {

	}
	if job.JobId == "erroid" {

		resultChan <- FailedJob{JobId: "erroid", ErrorData: errors.New("generated error")}
	}
	if job.JobId == "workid" {
		resultChan <- DoneJob{JobId: "workid"}
	}
}

func (testJobDispatcher *TestJobDispatcher) Dispatch(jobd interface{}, resultChan chan interface{}) error {
	go DispatchTh(jobd, resultChan)
	return nil
}

func (testErrorDispatcher *TestErrorDispatcher) DispatchError(failedJob *FailedJob) error {
	log.Print("success dispatch error")
	return nil
}

func TestJobBallancer(t *testing.T) {
	testJobDispatcher := TestJobDispatcher{}
	testErrorDispatcher := TestErrorDispatcher{}
	jobBallancer := JobBallancer{}
	jobBallancer.Init(&testJobDispatcher, &testErrorDispatcher)

	if err := jobBallancer.PushJob(Job{JobId: "workid"}); err != nil {
		t.Errorf("error: push err job failed " + err.Error())
		return
	}

	if err := jobBallancer.PushJob(Job{JobId: "erroid"}); err != nil {
		t.Errorf("error: push err job failed " + err.Error())
		return
	}
	time.Sleep(time.Second)
	if err := jobBallancer.TerminateTakeJob(); err != nil {
		t.Errorf("error: terminate job failed " + err.Error())
	}

}
