package main

import "errors"
import "dtools/gdcmgobr"

type DCOMClient struct {
	CallerAE_Title string
}

func (dcomClient *DCOMClient) checRequisites() error {
	if len(dcomClient.CallerAE_Title) == 0 {
		return errors.New("error: CallerAE_Title is empty")
	}
	return nil
}

func (dcomClient *DCOMClient) CStore() error {
	return nil
}

func (dcomClient *DCOMClient) CGet() error {
	return nil
}

func (dcomClient *DCOMClient) CMove() error {
	return nil
}

func (dcomClient *DCOMClient) CFind(dicomCFindRequest DicomCFindRequest) (interface{}, error) {
	//dataSet := gdcmgobr.GenCfindResult()
	//keys := gdcmgobr.GenCfindKeys()
	//e := gdcmgobr.
	//	gdcmgobr.CFind("AE_TITLE", "GE_PASC", "pacs.chaika.com", 104, keys, dataSet)
	//datatrr := gdcmgobr.Std_vector_Sl_gdcm_DataSet_Sg_

	//result := gdcmgobr.SwigcptrStd_vector_Sl_gdcm_DataSet_Sg_.Swigcptr()
	//

	return "", nil
}
func (dcomClient *DCOMClient) CEcho(dicomCEchoRequest DicomCEchoRequest) (DicomCEchoResult, error) {
	if err := dcomClient.checRequisites(); err != nil {
		return DicomCEchoResult{}, err
	}
	isAlive := gdcmgobr.CEcho(dicomCEchoRequest.Address, dicomCEchoRequest.Port, dicomCEchoRequest.ServerAE_Title, dcomClient.CallerAE_Title)
	dicomCEchoResult := DicomCEchoResult{IsAlive: isAlive}
	return dicomCEchoResult, nil
}
