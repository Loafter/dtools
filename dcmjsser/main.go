package main

import "log"
import "dtools/gdcmgobr"

func main() {
	var isOn bool
	isOn = gdcmgobr.CEcho("pacs.chaika.com", 104, "", "")
	log.Printf("dicom C-ECHO state is %v", isOn)
}
