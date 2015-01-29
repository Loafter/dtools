package main

import "errors"

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

func (dcomClient *DCOMClient) CFind(dicomCFindRequest DicomCFindRequest) error {
	{
		if err := dcomClient.checRequisites(); err != nil {
			return err
		}
		return nil
	}

}
func (dcomClient *DCOMClient) CEcho(dicomCEchoRequest DicomCEchoRequest) (interface{}, error) {
	if err := dcomClient.checRequisites(); err != nil {
		return "", err
	}
	//to do: write gdcm call
	return "", nil
}
