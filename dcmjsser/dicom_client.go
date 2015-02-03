package main

import "errors"
import "dtools/gdcmgobr"
import "encoding/json"

type DClient struct {
	CallerAE_Title string
}

func (dc *DClient) checRequisites() error {
	if len(dc.CallerAE_Title) == 0 {
		return errors.New("error: CallerAE_Title is empty")
	}
	return nil
}

func (dc *DClient) CStore() error {
	return nil
}

func (dc *DClient) CGet() error {
	return nil
}

func (dc *DClient) CMove() error {
	return nil
}

func (dc *DClient) CFind(freq FindReq) (interface{}, error) {
	cfindResult := gdcmgobr.CFind(dc.CallerAE_Title, freq.ServerAE_Title, freq.Address,
		freq.Port, freq.PatientName, freq.AccessionNumber,
		freq.PatienDateOfBirth, freq.StudyDate)
	var fdat []FindRes
	err := json.Unmarshal([]byte(cfindResult), &fdat)
	if err != nil {
		return nil, errors.New("error: Can't parse dicom cFind result data")
	}
	return fdat, nil
}
func (dc *DClient) CEcho(dicomCEchoRequest EchoReq) (EchoRes, error) {
	if err := dc.checRequisites(); err != nil {
		return EchoRes{}, err
	}
	isAlive := gdcmgobr.CEcho(dicomCEchoRequest.Address, dicomCEchoRequest.Port, dicomCEchoRequest.ServerAE_Title, dc.CallerAE_Title)
	dicomCEchoResult := EchoRes{IsAlive: isAlive}
	return dicomCEchoResult, nil
}
