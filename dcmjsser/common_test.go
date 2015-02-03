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

func (*TestErrorDispatcher) DispatchError(failedJob FaJob) error {
	//log.Printf("info: TestErrorDispatcher job %v job data %v \n", failedJob)
	return nil
}

type TestCompletedDispatcher struct {
}

func (*TestCompletedDispatcher) DispatchSuccess(completedJob CompJob) error {
	//log.Printf("info: TestCompletedDispatcher job %v job data %v \n", completedJob)
	return nil
}

func TestJobBallancer(t *testing.T) {
	testJobDispatcher := TestJobDispatcher{}
	testErrorDispatcher := TestErrorDispatcher{}
	testSuccessDispatcher := TestCompletedDispatcher{}
	jobBallancer := JobBallancer{}
	jobBallancer.Init(&testJobDispatcher, &testSuccessDispatcher, &testErrorDispatcher)
	for i := 0; i < 5; i++ {
		jobBallancer.PushJob("is error" + strconv.Itoa(i))
	}

	jobBallancer.TerminateTakeJob()

}

func TestDicomCEchoClient(t *testing.T) {
	dicomCEchoRequest := EchoReq{Address: "213.165.94.158", Port: 104, ServerAE_Title: "GEPACS"}
	dcomClient := DClient{CallerAE_Title: "AE_DTOOLS"}
	if pingRes, err := dcomClient.CEcho(dicomCEchoRequest); err != nil {
		t.Errorf("error: Test stop failed %v", err)
	} else {
		log.Printf("info: ping result %v", pingRes)
	}
}

func TestDicomCFindClient(t *testing.T) {
	dcomClient := DClient{CallerAE_Title: "AE_DTOOLS"}
	dicomCFindRequest := FindReq{EchoReq: EchoReq{Address: "213.165.94.158", Port: 104, ServerAE_Title: "GEPACS"}, PatientName: "Ab*"}
	if result, err := dcomClient.CFind(dicomCFindRequest); err != nil {
		t.Errorf("error: Test stop fail %v", err)
	} else {
		log.Println(result)
	}
	//time.Sleep(5000 * time.Millisecond)

}
