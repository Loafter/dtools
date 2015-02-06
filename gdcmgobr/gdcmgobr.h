#ifndef GDCMGOBR_H
#define GDCMGOBR_H

#include <string>
#include <gdcmCompositeNetworkFunctions.h>
#include <vector>

bool CEcho (std::string, int portno, std::string aetitle, std::string call);
std::string CFind(std::string callingaetitle,std::string callaetitle,std::string hostname,int port ,
			std::string PatientName,std::string AccessionNumber,std::string PatienDateOfBirth,
			std::string StudyDate);
#endif
