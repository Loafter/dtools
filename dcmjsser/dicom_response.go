package main

//import "net/http"
//import "encoding/json"
import "log"

//import "errors"

type OnErrorResp struct {
	responses map[string]HttpResReq
}

func (onErrorResp *OnErrorResp) Init(responses map[string]HttpResReq) {
	onErrorResp.responses = responses
}

func (*OnErrorResp) DispatchError(failedJob *FailedJob) error {
	log.Println("info: DispatchError")
	/*httpResReq, isType := failedJob.DataToError.(HttpResReq)
	if !isType {
		return errors.New("error: http responce and responce writer corrupted")
	}

	js, err := json.Marshal(failedJob.ErrorData)
	if err != nil {
		http.Error(httpResReq.ResponseWriter, err.Error(), http.StatusInternalServerError)
		stErr := "error: Can't create system error response"
		log.Println(stErr)
		httpResReq.ResponseWriter.Write(js)
		return errors.New(stErr)
	}
	httpResReq.ResponseWriter.Header().Set("Content-Type", "application/json")
	httpResReq.ResponseWriter.Write(js)*/
	return nil
}

type OnCompletedResp struct {
	responses map[string]HttpResReq
}

func (onCompletedResp *OnCompletedResp) Init(responses map[string]HttpResReq) {
	onCompletedResp.responses = responses
}
func (*OnCompletedResp) DispatchSuccess(completedJob *CompletedJob) error {
	/*httpResReq, isType := completedJob.DataToSuccess.(HttpResReq)
	if !isType {
		return errors.New("error: http responce and responce writer corrupted")
	}

	js, err := json.Marshal(completedJob.ResultData)
	if err != nil {
		http.Error(httpResReq.ResponseWriter, err.Error(), http.StatusInternalServerError)
		stErr := "error: Can't create complete job response"
		log.Println(stErr)
		httpResReq.ResponseWriter.Write(js)
		return errors.New(stErr)
	}
	httpResReq.ResponseWriter.Header().Set("Content-Type", "application/json")
	httpResReq.ResponseWriter.Write(js)*/
	return nil
}
