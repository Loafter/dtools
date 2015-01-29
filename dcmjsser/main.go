package main

import "log"

//import "dtools/gdcmgobr"

func main() {
	dicomJsonService := DicomJsonService{}
	if err := dicomJsonService.Start(9978); err != nil {
		log.Println(err)
	}
}
