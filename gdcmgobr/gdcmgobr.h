#ifndef GDCMGOBR_H
#define GDCMGOBR_H

#include <string>
#include <gdcmCompositeNetworkFunctions.h>
#include <vector>

bool CEcho (const char *remote, int portno, const char *aetitle, const char *call);
std::string CFind(const char* callingaetitle,const char* callaetitle,const char* hostname,int port ,
			std::string PatientName,std::string AccessionNumber,std::string PatienDateOfBirth,
			std::string StudyDate);
#endif
