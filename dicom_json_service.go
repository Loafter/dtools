package main

import "io/ioutil"
import "net/http"
import "strconv"
import "log"
import "encoding/json"

//main service class
type DicomJsonService struct {
	jobBallancer *JobBallancer
}

//start and init service
func (service *DicomJsonService) Start(listenPort int) error {
	http.HandleFunc("/c-echo", service.cEcho)
	http.HandleFunc("/index.html", service.ServePage)
	retVal := http.ListenAndServe(":"+strconv.Itoa(listenPort), nil)
	service.jobBallancer = &JobBallancer{}
	/*dcomClient := DCOMClient{CallerAE_Title: "AE_DTOOLS"}
	errorDispatcher := DicomErrorDispatcher{}
	//service.jobBallancer.Init(&DicomDispatcher{dcomClient: dcomClient}, &errorDispatcher)
	*/return retVal
}

//serve cEcho responce
func (service *DicomJsonService) cEcho(responseWriter http.ResponseWriter, request *http.Request) {
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
	service.jobBallancer.PushJob(dicomCEchoRequest, nil, nil)

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
