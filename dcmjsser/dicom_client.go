package main

import "errors"
import "encoding/json"

//import "strings"
import "log"

// #cgo CPPFLAGS:  -I/usr/include/c++/4.9/ -I/usr/local/include/gdcm-2.4/ -I/usr/include/x86_64-linux-gnu/c++/4.9/
// #cgo LDFLAGS: -L /usr/local/lib/ -lgdcmMSFF -lgdcmMEXD -lsocketxx -lgdcmMSFF -lgdcmjpeg8 -lgdcmjpeg12 -lgdcmjpeg16 -lgdcmopenjpeg -lgdcmcharls -lgdcmuuid -lgdcmDICT -lgdcmIOD -lgdcmexpat -lgdcmDSED -lgdcmCommon -lgdcmzlib -ldl
// #include "gdcmgobr.h"
import "C"

type DClient struct {
	CallerAE_Title string
}

func (dc *DClient) checRequisites() error {
	if len(dc.CallerAE_Title) == 0 {
		return errors.New("error: CallerAE_Title is empty")
	}
	return nil
}

func (dc *DClient) CStore(cStorReq CStorReq) (CStorReq, error) {
	if err := dc.checRequisites(); err != nil {
		return CStorReq{}, err
	}
	cae := dc.CallerAE_Title
	sae := cStorReq.ServerSet.ServerAE_Title
	ip := cStorReq.ServerSet.Address
	port := cStorReq.ServerSet.Port
	fls := cStorReq.File
	isStore := C.CStore(C.CString(ip), C.int(port), C.CString(cae), C.CString(sae), C.CString(fls))

	if int(isStore) == 0 {
		return CStorReq{}, errors.New("error: can't store dicom file " + fls)
	}
	return cStorReq, nil
}

func (dc *DClient) CGet(cgt CGetReq) (CGetReq, error) {
	if err := dc.checRequisites(); err != nil {
		return CGetReq{}, err
	}
	cae := dc.CallerAE_Title
	sae := cgt.FindReq.ServerSet.ServerAE_Title
	ip := cgt.FindReq.ServerSet.Address
	port := cgt.FindReq.ServerSet.Port
	pn := cgt.FindReq.PatientName
	an := cgt.FindReq.AccessionNumber
	bd := cgt.FindReq.PatienDateOfBirth
	sd := cgt.FindReq.StudyDate
	stid := cgt.FindReq.StudyInstanceUID
	fp := cgt.Folder
	cget := C.CGet(C.CString(cae), C.CString(sae), C.CString(ip), C.int(port), C.CString(stid), C.CString(pn), C.CString(an), C.CString(bd), C.CString(sd), C.CString(fp))
	if cget == 0 {
		return CGetReq{}, errors.New("error: can't cget dicom file " + pn)
	}
	return cgt, nil
}

func (dc *DClient) CFind(freq FindReq) ([]FindRes, error) {
	if err := dc.checRequisites(); err != nil {
		return nil, err
	}
	cae := dc.CallerAE_Title
	sae := freq.ServerSet.ServerAE_Title
	ip := freq.ServerSet.Address
	port := freq.ServerSet.Port
	pn := freq.PatientName
	an := freq.AccessionNumber
	bd := freq.PatienDateOfBirth
	sd := freq.StudyDate
	stid := freq.StudyInstanceUID
	cfindResultC := C.CFind(C.CString(cae), C.CString(sae), C.CString(ip), C.int(port), C.CString(stid), C.CString(pn), C.CString(an), C.CString(bd), C.CString(sd))
	cfindResult := C.GoString(cfindResultC)
	//cfindResult = strings.Replace(cfindResult, string(0), "", -1)
	log.Println(cfindResultC)
	var fdat []FindRes
	err := json.Unmarshal([]byte(cfindResult), &fdat)
	if err != nil {
		return nil, errors.New("error: can't parse dicom cFind result data " + err.Error() + cfindResult)
	}
	return fdat, nil
}
func (dc *DClient) CEcho(dicomCEchoRequest EchoReq) (EchoRes, error) {
	C.test()
	if err := dc.checRequisites(); err != nil {
		return EchoRes{}, err
	}
	isAlive := C.CEcho(C.CString(dicomCEchoRequest.Address), C.int(dicomCEchoRequest.Port), C.CString(dc.CallerAE_Title), C.CString(dicomCEchoRequest.ServerAE_Title))
	dicomCEchoResult := EchoRes{IsAlive: isAlive == 1}
	return dicomCEchoResult, nil
}
