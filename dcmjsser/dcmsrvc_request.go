package main

import "encoding/json"
import "errors"
import "strconv"

// common dicom requests
type EchoReq struct {
	Address        string `json:"Address"`
	Port           int    `json:"Port,string"`
	ServerAE_Title string `json:"ServerAE_Title"`
}

func (ereq EchoReq) GetDescript() string {
	return "C-Echo request: " + ereq.Address + ":" + strconv.Itoa(ereq.Port) + " AE_TITLE:" + ereq.ServerAE_Title
}

func (ereq *EchoReq) InitFromJsonData(data []byte) error {
	err := json.Unmarshal(data, &ereq)
	if err != nil {
		return errors.New("error: Can't parse dicom cEcho request data")
	}
	return nil

}

type FindReq struct {
	ServerSet         EchoReq
	PatientName       string `json:"PatientName"`
	AccessionNumber   string `json:"AccessionNumber"`
	PatienDateOfBirth string `json:"PatienDateOfBirth"`
	StudyDate         string `json:"StudyDate"`
}

type EchoRes struct {
	IsAlive bool `json:"IsAlive"`
}

func (freq *FindReq) InitFromJsonData(data []byte) error {
	err := json.Unmarshal(data, &freq)
	if err != nil {
		return errors.New("error: Can't parse dicom cFind request data")
	}
	return nil

}

func (freq FindReq) GetDescript() string {
	st := freq.PatientName + " " + freq.AccessionNumber + " " + freq.PatienDateOfBirth + " " + freq.StudyDate
	return "C-Find request: " + freq.ServerSet.Address + ":" + strconv.Itoa(freq.ServerSet.Port) + " " + st
}

type FindRes struct {
	PatientName       string `json:"PatientName"`
	AccessionNumber   string `json:"AccessionNumber"`
	PatienDateOfBirth string `json:"PatienDateOfBirth"`
	StudyDate         string `json:"StudyDate"`
}

type CStorReq struct {
	ServerSet EchoReq
	File      string `json:"File"`
}

func (cstor *CStorReq) InitFromJsonData(data []byte) error {
	err := json.Unmarshal(data, &cstor)
	if err != nil {
		return errors.New("error: Can't parse dicom cStore request data")
	}
	return nil

}

func (cstor *CStorReq) GetDescript() string {
	return "C-Store request: " + cstor.ServerSet.Address + ":" + strconv.Itoa(cstor.ServerSet.Port) + " " + cstor.File
}
