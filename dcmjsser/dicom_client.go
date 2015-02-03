package main

import "errors"
import "dtools/gdcmgobr"
import "encoding/json"

type DicomClient struct {
	CallerAE_Title string
}

func (dcomClient *DicomClient) checRequisites() error {
	if len(dcomClient.CallerAE_Title) == 0 {
		return errors.New("error: CallerAE_Title is empty")
	}
	return nil
}

func (dcomClient *DicomClient) CStore() error {
	return nil
}

func (dcomClient *DicomClient) CGet() error {
	return nil
}

func (dcomClient *DicomClient) CMove() error {
	return nil
}

func (dcomClient *DicomClient) CFind(dicomCFindRequest DicomCFindRequest) (interface{}, error) {
	cfindResult := gdcmgobr.CFind(dcomClient.CallerAE_Title, dicomCFindRequest.ServerAE_Title, dicomCFindRequest.Address,
		dicomCFindRequest.Port, dicomCFindRequest.PatientName, dicomCFindRequest.AccessionNumber,
		dicomCFindRequest.PatienDateOfBirth, dicomCFindRequest.StudyDate)
	var cfindData []DicomCFindResult
	err := json.Unmarshal([]byte(cfindResult), &cfindData)
	if err != nil {
		return nil, errors.New("error: Can't parse dicom cFind result data")
	}
	return cfindData, nil
}
func (dcomClient *DicomClient) CEcho(dicomCEchoRequest DicomCEchoRequest) (DicomCEchoResult, error) {
	if err := dcomClient.checRequisites(); err != nil {
		return DicomCEchoResult{}, err
	}
	isAlive := gdcmgobr.CEcho(dicomCEchoRequest.Address, dicomCEchoRequest.Port, dicomCEchoRequest.ServerAE_Title, dcomClient.CallerAE_Title)
	dicomCEchoResult := DicomCEchoResult{IsAlive: isAlive}
	return dicomCEchoResult, nil
}
