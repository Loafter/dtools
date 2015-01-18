package main

import "testing"
import "errors"
import "log"
import "time"

type TestJobDispatcher struct {
}

type TestErrorDispatcher struct {
}

func (testJobDispatcher *TestJobDispatcher) Dispatch(taskResponse TaskResponse, resultChan chan interface{}) error {
	go func() {
		if taskResponse.Tasker.TaskId() == "erroid" {
			resultChan <- errors.New("error catched")
		} else if taskResponse.Tasker.TaskId() == "workid" {
			resultChan <- taskResponse.Tasker
		} else if taskResponse.Tasker.TaskId() == "errnotif" {
			resultChan <- TaskError{Error: errors.New("error catched"), Tasker: taskResponse.Tasker}
		}
	}()
	return nil
}

func (testErrorDispatcher *TestErrorDispatcher) NotifyError(taskError TaskError) error {
	log.Println(taskError.Error.Error())
	return nil
}
func (testErrorDispatcher *TestErrorDispatcher) DispatchError(err error) error {
	log.Println(err.Error())
	return nil
}

func TestJobBallancer(t *testing.T) {
	testJobDispatcher := TestJobDispatcher{}
	testErrorDispatcher := TestErrorDispatcher{}
	jobBallancer := JobBallancer{}
	jobBallancer.Init(&testJobDispatcher, &testErrorDispatcher)
	err := jobBallancer.PushJob(TaskResponse{Tasker: &BasicTask{Id: "workid"}, responseWriter: nil, httpRequest: nil})
	if err != nil {
		log.Println("error: push err job failed " + err.Error())
		t.Errorf("error: push normal job failed ")
		return
	}

	err = jobBallancer.PushJob(TaskResponse{Tasker: &BasicTask{Id: "erroid"}, responseWriter: nil, httpRequest: nil})
	if err != nil {
		log.Println("error: push err job failed " + err.Error())
		t.Errorf("error: push err job failed ")
		return
	}

	err = jobBallancer.PushJob(TaskResponse{Tasker: &BasicTask{Id: "errnotif"}, responseWriter: nil, httpRequest: nil})
	if err != nil {
		log.Println("error: push err job failed " + err.Error())
		t.Errorf("error: push errnot job failed ")
		return
	}
	jobBallancer.TerminateTakeJob()

	time.Sleep(5 * time.Second)

}
