package main

import "log"
import "runtime"

//import "time"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	/*dcomClient := DClient{CallerAE_Title: "AE_DTOOLS"}
	disp := DDisp{dCln: dcomClient}
	for i := 0; i < 20; i++ {
		dicomCFindRequest := FindReq{ServerSet: EchoReq{Address: "213.165.94.158", Port: 104, ServerAE_Title: "GEPACS"}, PatientName: "A*"}
		go func() {
			if result, err := disp.Dispatch(dicomCFindRequest); err != nil {
				log.Println("error: Test stop fail %v", err)
			} else {
				log.Println(result)
			}
		}()
	}
	time.Sleep(time.Second * 100)*/

	djs := DJsServ{}
	if err := djs.Start(9978); err != nil {
		log.Println(err)
	}

}
