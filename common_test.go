package main

import "testing"
import "log"
import "errors"
import "time"

type TestJobDispatcher struct {
}

type TestErrorDispatcher struct {
}

func DispatchTh(jobd interface{}, resultChan chan interface{}) {

}

func (testJobDispatcher *TestJobDispatcher) Dispatch(jobd interface{}, resultChan chan interface{}) error {
	job := jobd.(Job)

	if job.JobId == "erroid" {
		time.Sleep(time.Second * 2)
		resultChan <- FailedJob{JobId: "erroid", ErrorData: errors.New("generated error")}
	} else if job.JobId == "workid" {
		time.Sleep(time.Second * 1)
		resultChan <- DoneJob{JobId: "workid"}
	} else {
		resultChan <- FailedJob{JobId: job.JobId, ErrorData: errors.New("generated error")}
	}
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
	if err := jobBallancer.TerminateTakeJob(); err != nil {
		t.Errorf("error: terminate job failed " + err.Error())
	}

}

func TestJobBallancerNowait(t *testing.T) {
	testJobDispatcher := TestJobDispatcher{}
	testErrorDispatcher := TestErrorDispatcher{}
	jobBallancer := JobBallancer{}
	jobBallancer.Init(&testJobDispatcher, &testErrorDispatcher)

	jobBallancer.PushJob(Job{JobId: "erroid1"})
	jobBallancer.PushJob(Job{JobId: "erroid2"})
	jobBallancer.PushJob(Job{JobId: "erroid3"})
	jobBallancer.PushJob(Job{JobId: "erroid4"})
	jobBallancer.PushJob(Job{JobId: "erroid5"})
	jobBallancer.PushJob(Job{JobId: "erroid6"})
	jobBallancer.PushJob(Job{JobId: "erroid7"})
	jobBallancer.PushJob(Job{JobId: "erroid8"})
	jobBallancer.PushJob(Job{JobId: "erroid9"})
	if err := jobBallancer.TerminateTakeJob(); err != nil {
		t.Errorf("error: terminate job failed " + err.Error())
	}
}
