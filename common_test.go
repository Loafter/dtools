package main

import "testing"
import "log"
import "time"
import "errors"

//import "fmt"

type TestJobDispatcher struct {
}

func (*TestJobDispatcher) Dispatch(data interface{}) (interface{}, error) {
	time.Sleep(1 * time.Second)
	log.Printf("info: try dispatch data %v", data)
	return nil, errors.New("gen error")
}

type TestErrorDispatcher struct {
}

func (*TestErrorDispatcher) DispatchError(failedJob *FailedJob, data interface{}) error {
	log.Printf("DispatchError job %v job data %v \n", failedJob, data)
	return nil
}

type TestCompletedDispatcher struct {
}

func (*TestCompletedDispatcher) DispatchSuccess(completedJob *CompletedJob, data interface{}) error {
	log.Printf("TestCompletedDispatcher job %v job data %v \n", completedJob, data)
	return nil
}

func TestJobBallancer(t *testing.T) {
	testJobDispatcher := TestJobDispatcher{}
	testErrorDispatcher := TestErrorDispatcher{}
	testSuccessDispatcher := TestCompletedDispatcher{}
	jobBallancer := JobBallancer{}
	jobBallancer.Init(&testJobDispatcher, &testErrorDispatcher, &testSuccessDispatcher)

	jobBallancer.PushJob("is error", "dataToDispatchSuccess", "dataToDispatchError")
	time.Sleep(time.Second * 4)
	jobBallancer.TerminateTakeJob()

}

/*
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
*/
func TestDicomClient(t *testing.T) {
	/*DCOMClient:=DCOMClient{
			Address :
	Port           uint16
	ServerAE_Title string
	CallerAE_Title string
	}*/
}
