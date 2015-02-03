package main

import "log"
import "runtime"

//import "dtools/gdcmgobr"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	djs := DJsServ{}
	if err := djs.Start(9978); err != nil {
		log.Println(err)
	}
}
