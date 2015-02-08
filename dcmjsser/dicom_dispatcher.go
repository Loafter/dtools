package main

import "errors"

//main dicom message dispatcher
type DDisp struct {
	dCln DClient
}

func (dsp *DDisp) Dispatch(dreq interface{}) (interface{}, error) {

	switch tr := dreq.(type) {
	case CStorReq:
		return nil, dsp.dCln.CStore(tr)
	case EchoReq:
		return dsp.dCln.CEcho(tr)
	case FindReq:
		return dsp.dCln.CFind(tr)

	}
	return nil, errors.New("error: can't dispatch non dicom request type")

}
