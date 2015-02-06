

#include <iostream>
#include "gdcmgobr.h"
#include <sstream>
#include "gdcmPrinter.h"

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

std::string CFind(std::string callingaetitle,std::string callaetitle,std::string hostname,int port ,
			std::string PatientName,std::string AccessionNumber,std::string PatienDateOfBirth,
			std::string StudyDate)
{	
    std::vector< std::pair<gdcm::Tag, std::string> > keys;
    		keys.push_back(std::make_pair(gdcm::Tag(0x0010,0x0010),PatientName));
		keys.push_back(std::make_pair(gdcm::Tag(0x0008,0x0050),AccessionNumber)); 
		keys.push_back(std::make_pair(gdcm::Tag(0x0010,0x0030),PatienDateOfBirth));
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
    if( !gdcm::CompositeNetworkFunctions::CFind(hostname.c_str(), (uint16_t)port, theQuery, theDataSet,  callingaetitle.c_str(),callaetitle.c_str()) )
      {
		std::cout<<"cfind err";
        return "[]";
      }

	std::string reqRes="[";
    for(int i=0;i<theDataSet.size();i++)
      {
		gdcm::DataSet dat=theDataSet[i];
		reqRes=reqRes+"{\"PatientName\":\""+GetStringValueFromTag(gdcm::Tag(0x0010,0x0010),dat)+"\",";
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

