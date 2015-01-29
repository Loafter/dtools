package main

import "io/ioutil"
import "net/http"
import "strconv"
import "log"
import "encoding/json"
import "errors"

type HttpResReq struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

//main service class
type DicomJsonService struct {
	jobBallancer    JobBallancer
	dicomDispatcher DicomDispatcher
}

//start and init service
func (service *DicomJsonService) Start(listenPort int) error {
	http.HandleFunc("/c-echo", service.cEcho)
	http.HandleFunc("/index.html", service.ServePage)
	if err := http.ListenAndServe(":"+strconv.Itoa(listenPort), nil); err != nil {
		return errors.New("error: can't start listen http server")
	}
	service.jobBallancer.Init(&service.dicomDispatcher, new(OnCompletedResp), new(OnErrorResp))
	return nil
}

//serve cEcho responce
func (service *DicomJsonService) cEcho(responseWriter http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	bodyData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		strErr := "error: Can't read http body data"
		responseWriter.Write([]byte(strErr))
		log.Println(strErr)
		return
	}
	var dicomCEchoRequest DicomCEchoRequest
	if err := json.Unmarshal(bodyData, &dicomCEchoRequest); err != nil {
		strErr := "error: Can't parse DicomCEchoRequest data"
		responseWriter.Write([]byte(strErr))
		log.Println(strErr)
	}
	httpResReq := HttpResReq{Request: request, ResponseWriter: responseWriter}
	service.jobBallancer.PushJob(dicomCEchoRequest, httpResReq, httpResReq)

}

//serve main page request
func (service *DicomJsonService) ServePage(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type: text/html", "*")
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		responseWriter.Write([]byte("error: Can't find start page \n"))
		return
	}
	responseWriter.Write(content)
}
