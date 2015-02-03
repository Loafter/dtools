package main

import "encoding/json"
import "errors"

// common dicom requests
type EchoReq struct {
	Address        string `json:"Address"`
	Port           int    `json:"Port,string"`
	ServerAE_Title string `json:"ServerAE_Title"`
}

func (ereq *EchoReq) InitFromJsonData(data []byte) error {
	err := json.Unmarshal(data, &ereq)
	if err != nil {
		return errors.New("error: Can't parse dicom cEcho request data")
	}
	return nil

}

type FindReq struct {
	EchoReq
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

type FindRes struct {
	PatientName       string `json:"PatientName"`
	AccessionNumber   string `json:"AccessionNumber"`
	PatienDateOfBirth string `json:"PatienDateOfBirth"`
	StudyDate         string `json:"StudyDate"`
}
