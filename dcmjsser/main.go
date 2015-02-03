package main

import "log"
import "runtime"

//import "dtools/gdcmgobr"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	dicomJsonService := DJsServ{}
	if err := dicomJsonService.Start(9978); err != nil {
		log.Println(err)
	}
}
