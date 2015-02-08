package main

import "errors"
import "dtools/gdcmgobr"
import "encoding/json"

//import "log"

type DClient struct {
	CallerAE_Title string
}

func (dc *DClient) checRequisites() error {
	if len(dc.CallerAE_Title) == 0 {
		return errors.New("error: CallerAE_Title is empty")
	}
	return nil
}

func (dc *DClient) CStore(cStorReq CStorReq) error {
	cae := dc.CallerAE_Title
	sae := cStorReq.ServerSet.ServerAE_Title
	ip := cStorReq.ServerSet.Address
	port := cStorReq.ServerSet.Port
	fls := cStorReq.File
	isStore := gdcmgobr.CStore(ip, port, sae, cae, fls)
	if !isStore {
		return errors.New("error: Can't store dicom file " + fls)
	}
	return nil
}

func (dc *DClient) CGet() error {
	return nil
}

func (dc *DClient) CMove() error {
	return nil
}

func (dc *DClient) CFind(freq FindReq) (interface{}, error) {
	cae := dc.CallerAE_Title
	sae := freq.ServerSet.ServerAE_Title
	ip := freq.ServerSet.Address
	port := freq.ServerSet.Port
	pn := freq.PatientName
	an := freq.AccessionNumber
	bd := freq.PatienDateOfBirth
	sd := freq.StudyDate
	cfindResult := gdcmgobr.CFind(cae, sae, ip, port, pn, an, bd, sd)
	var fdat []FindRes
	err := json.Unmarshal([]byte(cfindResult), &fdat)
	if err != nil {
		return nil, errors.New("error: Can't parse dicom cFind result data " + err.Error() + cfindResult)
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
