

#include <iostream>
#include "gdcmgobr.h"
#include <sstream>
#include "gdcmPrinter.h"
#include <string>

std::vector<bool> TestVec()
{
	std::vector<bool> cfindResult;
	return cfindResult;
}

bool CEcho (std::string remote, int portno, std::string aetitle,std::string call){
	return gdcm::CompositeNetworkFunctions::CEcho(remote.c_str(),portno,aetitle.c_str(),call.c_str());
}
 std::string GetStringValueFromTag(const gdcm::Tag t, const gdcm::DataSet ds)
{
  std::string buffer;
  if( ds.FindDataElement( t ) )
    {
    const gdcm::DataElement& de = ds.GetDataElement(t );
    const gdcm::ByteValue *bv = de.GetByteValue();
    if( bv ) // Can be Type 2
      {
      buffer = std::string( bv->GetPointer(), bv->GetLength() );
      }
    }

  // Since return is a const char* the very first \0 will be considered
  return buffer;
}


bool CGet(std::string aetitle,std::string call,std::string hostname,int port ,std::string  StudyInstanceUID,
			std::string PatientName,std::string AccessionNumber,std::string PatienDateOfBirth,
			std::string StudyDate,std::string SFolder)
{	
    std::vector< std::pair<gdcm::Tag, std::string> > keys;
	if(StudyInstanceUID.size()!=0)
    	keys.push_back(std::make_pair(gdcm::Tag(0x0020,0x000D),StudyInstanceUID));
/*	if(PatientName.size()!=0)
    	keys.push_back(std::make_pair(gdcm::Tag(0x0010,0x0010),PatientName));
	if(AccessionNumber.size()!=0)
		keys.push_back(std::make_pair(gdcm::Tag(0x0008,0x0050),AccessionNumber)); 
	if(PatienDateOfBirth.size()!=0)
		keys.push_back(std::make_pair(gdcm::Tag(0x0010,0x0030),PatienDateOfBirth));
	if(StudyDate.size()!=0)
		keys.push_back(std::make_pair(gdcm::Tag(0x0008,0x0020),StudyDate));
*/
	std::cout<<StudyInstanceUID;
    gdcm::ERootType theRoot = gdcm::eStudyRootType;
    gdcm::EQueryLevel theLevel = gdcm::eStudy;
 
    gdcm::SmartPointer<gdcm::BaseRootQuery> theQuery = gdcm::CompositeNetworkFunctions::ConstructQuery(theRoot, theLevel ,keys,true);

    if (!theQuery)
      {
      std::cerr << "Query construction failed." <<std::endl;
      return false;
      }

	;
    //doing a non-strict query, the second parameter there.
    //look at the base query comments
    if (!theQuery->ValidateQuery(false))
      {
      std::cerr << "You have not constructed a valid find query. Please try again." << std::endl;
      return false;
      } 
   return gdcm::CompositeNetworkFunctions::CMove(hostname.c_str(), (uint16_t)port,theQuery,11112, aetitle.c_str(),call.c_str(),SFolder.c_str());
  
}


std::string CFind(std::string callingaetitle,std::string callaetitle,std::string hostname,int port ,std::string  StudyInstanceUID,std::string PatientName,std::string AccessionNumber,std::string PatienDateOfBirth,std::string StudyDate)
{
    std::vector< std::pair<gdcm::Tag, std::string> > keys;
	if(StudyInstanceUID.size()!=0)
    	keys.push_back(std::make_pair(gdcm::Tag(0x0020,0x000D),StudyInstanceUID));
	if(PatientName.size()!=0)
    	keys.push_back(std::make_pair(gdcm::Tag(0x0010,0x0010),PatientName));
	if(AccessionNumber.size()!=0)
		keys.push_back(std::make_pair(gdcm::Tag(0x0008,0x0050),AccessionNumber)); 
	if(PatienDateOfBirth.size()!=0)
		keys.push_back(std::make_pair(gdcm::Tag(0x0010,0x0030),PatienDateOfBirth));
	if(StudyDate.size()!=0)
		keys.push_back(std::make_pair(gdcm::Tag(0x0008,0x0020),StudyDate));

    gdcm::ERootType theRoot = gdcm::eStudyRootType;
    gdcm::EQueryLevel theLevel = gdcm::eStudy;
 
    gdcm::SmartPointer<gdcm::BaseRootQuery> theQuery = gdcm::CompositeNetworkFunctions::ConstructQuery(theRoot, theLevel ,keys);

    if (!theQuery)
      {
      std::cerr << "Query construction failed." <<std::endl;
      return "";
      }

	;
    //doing a non-strict query, the second parameter there.
    //look at the base query comments
    if (!theQuery->ValidateQuery(false))
      {
      std::cerr << "You have not constructed a valid find query."
        " Please try again." << std::endl;
      return "";
      } 

	 std::vector<gdcm::DataSet> theDataSet;
	//std::cout<<"cfb"<<std::endl;
    if( !gdcm::CompositeNetworkFunctions::CFind(hostname.c_str(), (uint16_t)port, theQuery, theDataSet, callingaetitle.c_str(),callaetitle.c_str()) )
      {
        return "[]";
      }
	//std::cout<<"cfe"<<std::endl;
	std::string reqRes="[";
    for(int i=0;i<theDataSet.size();i++)
      {
		gdcm::DataSet dat=theDataSet[i];
		reqRes=reqRes+"{\"StudyInstanceUID\":\""+GetStringValueFromTag(gdcm::Tag(0x0020,0x000D),dat)+"\",";
		reqRes=reqRes+"\"PatientName\":\""+GetStringValueFromTag(gdcm::Tag(0x0010,0x0010),dat)+"\",";
		reqRes=reqRes+"\"AccessionNumber\":\""+GetStringValueFromTag(gdcm::Tag(0x0008,0x0050),dat)+"\",";
		reqRes=reqRes+"\"PatienDateOfBirth\":\""+GetStringValueFromTag(gdcm::Tag(0x0008,0x0020),dat)+"\",";
		reqRes=reqRes+"\"StudyDate\":\""+GetStringValueFromTag(gdcm::Tag(0x0008,0x0020),dat);
		if(i==theDataSet.size()-1)
			reqRes=reqRes+"\"}";
		else
			reqRes=reqRes+"\"}, \n";
      }
	reqRes=reqRes+"] \n";
	return reqRes;
}

std::vector<std::string> &split(const std::string &s, char delim, std::vector<std::string> &elems) {
    std::stringstream ss(s);
    std::string item;
    while (std::getline(ss, item, delim)) {
        elems.push_back(item);
    }
    return elems;
}

std::vector<std::string> split(const std::string &s, char delim) {
    std::vector<std::string> elems;
    split(s, delim, elems);
    return elems;
}

bool CStore (std::string remote, int portno, std::string aetitle, std::string call,std::string file)	
{
	gdcm::Directory::FilenamesType thefiles;
	if (file.empty())
	 return false;
	char l=*file.end();
      if(gdcm::System::FileIsDirectory(file.c_str()))
        {
        gdcm::Directory::FilenamesType files;
        gdcm::Directory dir;
        dir.Load(file, true);
        files = dir.GetFilenames();
        thefiles.insert(thefiles.end(), files.begin(), files.end());
        }
      else
        {
        // This is a file simply add it
        thefiles.push_back(file);
     }
    bool didItWork = gdcm::CompositeNetworkFunctions::CStore(remote.c_str(), (uint16_t)portno, thefiles, aetitle.c_str(), call.c_str());
    gdcmDebugMacro( (didItWork ? "Store was successful." : "Store failed.") );
    return didItWork;
}
