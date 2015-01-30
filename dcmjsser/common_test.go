package main

import "testing"
import "log"
import "time"
import "errors"
import "math"
import "strconv"

//import "fmt"

type TestJobDispatcher struct {
	i int
}

func (test *TestJobDispatcher) Dispatch(data interface{}) (interface{}, error) {
	time.Sleep(500 * time.Duration(test.i) * time.Millisecond)
	test.i++

	log.Printf("info: try dispatch data %v", data)
	if math.Mod(float64(test.i), 2) == 0.0 {
		return nil, errors.New("gen error")
	} else {
		return data, nil
	}
}

type TestErrorDispatcher struct {
}

func (*TestErrorDispatcher) DispatchError(failedJob FailedJob) error {
	log.Printf("info: TestErrorDispatcher job %v job data %v \n", failedJob)
	return nil
}

type TestCompletedDispatcher struct {
}

func (*TestCompletedDispatcher) DispatchSuccess(completedJob CompletedJob) error {
	log.Printf("info: TestCompletedDispatcher job %v job data %v \n", completedJob)
	return nil
}

func TestJobBallancer(t *testing.T) {
	testJobDispatcher := TestJobDispatcher{}
	testErrorDispatcher := TestErrorDispatcher{}
	testSuccessDispatcher := TestCompletedDispatcher{}
	jobBallancer := JobBallancer{}
	jobBallancer.Init(&testJobDispatcher, &testSuccessDispatcher, &testErrorDispatcher)
	for i := 0; i < 10; i++ {
		jobBallancer.PushJob("is error" + strconv.Itoa(i))
	}

	jobBallancer.TerminateTakeJob()

}

func TestDicomClient(t *testing.T) {
	/*DCOMClient:=DCOMClient{
			Address :
	Port           uint16
	ServerAE_Title string
	CallerAE_Title string
	}*/
}
