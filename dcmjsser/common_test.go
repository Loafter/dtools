package main

import "testing"
import "log"
import "time"

//import "strconv"

type TestJobDispatcher struct {
	i int
}

func (test *TestJobDispatcher) Dispatch(data interface{}) (interface{}, error) {
	time.Sleep(500 * time.Millisecond)
	test.i++

	log.Printf("info: try dispatch data %v", data)
	return data, nil
}

type TestErrorDispatcher struct {
}

func (*TestErrorDispatcher) DispatchError(failedJob FaJob) error {
	return nil
}

type TestCompletedDispatcher struct {
}

func (*TestCompletedDispatcher) DispatchSuccess(completedJob CompJob) error {
	return nil
}

/*
func TestJobBallancer(t *testing.T) {
	testJobDispatcher := TestJobDispatcher{}
	testErrorDispatcher := TestErrorDispatcher{}
	testSuccessDispatcher := TestCompletedDispatcher{}
	jobBallancer := JobBallancer{}
	jobBallancer.Init(&testJobDispatcher, &testSuccessDispatcher, &testErrorDispatcher)
	for i := 0; i < 40; i++ {
		jobBallancer.PushJob("data: " + strconv.Itoa(i))
	}

	jobBallancer.TerminateTakeJob()

}
*/

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
	disp := DDisp{dCln: dcomClient}
	dicomCFindRequest := FindReq{ServerSet: EchoReq{Address: "213.165.94.158", Port: 104, ServerAE_Title: "GEPACS"}, StudyDate: "*", StudyInstanceUID: "*", AccessionNumber: "*", PatienDateOfBirth: "*", PatientName: "a*"}
	go func() {
		if result, err := disp.Dispatch(dicomCFindRequest); err != nil {
			t.Errorf("error: Test stop fail %v", err)
		} else {
			log.Println(result)
		}
	}()
	time.Sleep(time.Second * 6)
}

func TestDicomCStoreClient(t *testing.T) {
	/*dicomCStoreRequest := CStorReq{ServerSet: EchoReq{Address: "213.165.94.158", Port: 104, ServerAE_Title: "GEPACS"}, File: "/home/andrew/Downloads/Dicom/ToSend/IM-0001-0041.dcm"}
	dcomClient := DClient{CallerAE_Title: "AE_DTOOLS"}
	if err := dcomClient.CStore(dicomCStoreRequest); err != nil {
		t.Errorf("error: Test stop failed %v", err)
	} else {
		log.Printf("info: cstore result ")
	}*/
}
