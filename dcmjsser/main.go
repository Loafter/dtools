package main

import "log"
import "runtime"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	djs := DJsServ{}
	log.Println("info: gui located at address http://127.0.0.1:9978/index.html, for recive study use scsc_port=11112, aetitle=AE_DTLS")
	if err := djs.Start(9978); err != nil {
		log.Println(err)
	}

}
