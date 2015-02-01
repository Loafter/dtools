#ifndef GDCMGOBR_H
#define GDCMGOBR_H

#include <string>
#include <gdcmCompositeNetworkFunctions.h>
#include <vector>
struct DicomCFindRequest 
 {
	std::string PatientName;     
	std::string StudyID;            
	std::string PatienDateOfBirth;
	std::string StudyDate; 
};
bool CEcho (const char *remote, int portno, const char *aetitle, const char *call);
bool CFind(const char* callingaetitle,const char* callaetitle,const char* hostname,int port ,DicomCFindRequest dicomCFindRequest,std::string* result);
#endif
