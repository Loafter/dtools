package main

import "errors"

//import "log"

//main dicom message dispatcher
type DicomDispatcher struct {
	dcomClient DicomClient
}

func (dispatcher *DicomDispatcher) Dispatch(dicomRequest interface{}) (interface{}, error) {
	switch typedRequest := dicomRequest.(type) {
	case DicomCEchoRequest:
		return dispatcher.dcomClient.CEcho(typedRequest)
	case DicomCFindRequest:
		return dispatcher.dcomClient.CFind(typedRequest)

	}
	return nil, errors.New("error: can't dispatch non dicom request type")

}
