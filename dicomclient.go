package main

type DCOMClient struct {
	Address        string
	Port           uint16
	ServerAE_Title string
	CallerAE_Title string
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

func (dcomClient *DCOMClient) CFindD() error {
	return nil
}
func (dcomClient *DCOMClient) CEcho(MessageID string, MessageIDBR string, AffectedSOPClassUID string, Status bool) error {
	return nil
}
