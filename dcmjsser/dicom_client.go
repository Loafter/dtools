package main

import "errors"
import "dtools/gdcmgobr"
import "encoding/json"
import "strings"

type DClient struct {
}

func (dc *DClient) CStore(cStorReq CStorReq) (CStorReq, error) {
	cae := cStorReq.ServerSet.ClientAE_Title
	sae := cStorReq.ServerSet.ServerAE_Title
	ip := cStorReq.ServerSet.Address
	port := cStorReq.ServerSet.Port
	fls := cStorReq.File
	isStore := gdcmgobr.CStore(ip, port, cae, sae, fls)
	if !isStore {
		return CStorReq{}, errors.New("error: can't store dicom file " + fls)
	}
	return cStorReq, nil
}

func (dc *DClient) CGet(cgt CGetReq) (CGetReq, error) {
	cae := cgt.FindReq.ServerSet.ClientAE_Title
	sae := cgt.FindReq.ServerSet.ServerAE_Title
	ip := cgt.FindReq.ServerSet.Address
	port := cgt.FindReq.ServerSet.Port
	pn := cgt.FindReq.PatientName
	an := cgt.FindReq.AccessionNumber
	bd := cgt.FindReq.PatienDateOfBirth
	sd := cgt.FindReq.StudyDate
	stid := cgt.FindReq.StudyInstanceUID
	pid := cgt.FindReq.PatientID
	fp := cgt.Folder
	cget := gdcmgobr.CGet(cae, sae, ip, port, stid, pn, an, bd, sd, pid, fp)
	if !cget {
		return CGetReq{}, errors.New("error: can't cget dicom file " + pn)
	}
	return cgt, nil
}

func (dc *DClient) CFind(freq FindReq) ([]FindRes, error) {

	cae := freq.ServerSet.ClientAE_Title
	sae := freq.ServerSet.ServerAE_Title
	ip := freq.ServerSet.Address
	port := freq.ServerSet.Port
	pn := freq.PatientName
	an := freq.AccessionNumber
	bd := freq.PatienDateOfBirth
	sd := freq.StudyDate
	stid := freq.StudyInstanceUID
	pid := freq.PatientID
	cfindResult := gdcmgobr.CFind(cae, sae, ip, port, stid, pn, an, bd, sd, pid)
	cfindResult = strings.Replace(cfindResult, string(0), "", -1)
	//cfindResult = strings.Replace(cfindResult, "\"", " ", -1)
	var fdat []FindRes
	err := json.Unmarshal([]byte(cfindResult), &fdat)
	if err != nil {
		return nil, errors.New("error: can't parse dicom cFind result data " + err.Error() + cfindResult)
	}
	return fdat, nil
}
func (dc *DClient) CEcho(dicomCEchoRequest EchoReq) (EchoRes, error) {
	isAlive := gdcmgobr.CEcho(dicomCEchoRequest.Address, dicomCEchoRequest.Port, dicomCEchoRequest.ClientAE_Title, dicomCEchoRequest.ServerAE_Title)
	dicomCEchoResult := EchoRes{IsAlive: isAlive}
	return dicomCEchoResult, nil
}
