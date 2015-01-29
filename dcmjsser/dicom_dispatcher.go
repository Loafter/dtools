package main

import "errors"

//main dicom message dispatcher
type DicomDispatcher struct {
	dcomClient DCOMClient
}

func (dispatcher *DicomDispatcher) Dispatch(dicomRequest interface{}) (interface{}, error) {
	switch typedRequest := dicomRequest.(type) {
	case DicomCEchoRequest:
		dispatcher.dcomClient.CEcho(typedRequest)
	case DicomCFindRequest:
		dispatcher.dcomClient.CFind(typedRequest)
	}
	return nil, errors.New("error: can't dispatch non dicom request type")
}
