

#include <iostream>
#include "gdcmgobr.h"
#include <sstream>
#include "gdcmPrinter.h"

std::vector<bool> TestVec()
{
	std::vector<bool> cfindResult;
	return cfindResult;
}

bool CEcho (const char *remote, int portno, const char *aetitle, const char *call){
	return gdcm::CompositeNetworkFunctions::CEcho(remote,portno,aetitle,call);
}

bool CFind(const char* callingaetitle,const char* callaetitle,const char* hostname,int port,DicomCFindRequest dicomCFindRequest,std::string* result)
{	
	std::vector< std::pair<gdcm::Tag, std::string> > keys;
    gdcm::ERootType theRoot = gdcm::eStudyRootType;
    gdcm::EQueryLevel theLevel = gdcm::eStudy;
 
    gdcm::SmartPointer<gdcm::BaseRootQuery> theQuery = gdcm::CompositeNetworkFunctions::ConstructQuery(theRoot, theLevel ,keys);

    if (!theQuery)
      {
      std::cerr << "Query construction failed." <<std::endl;
      return false;
      }


    //doing a non-strict query, the second parameter there.
    //look at the base query comments
    if (!theQuery->ValidateQuery(false))
      {
      std::cerr << "You have not constructed a valid find query."
        " Please try again." << std::endl;
      return false;
      } 
	 std::vector<gdcm::DataSet> theDataSet;
    if( !gdcm::CompositeNetworkFunctions::CFind(hostname, (uint16_t)port, theQuery, theDataSet, callingaetitle, callaetitle) )
      {
      gdcmDebugMacro( "Problem in CFind." );
      return false;
      }
	gdcm::Printer p;
	std::stringstream ss;
    std::ostream &os = ss;
    for( std::vector<gdcm::DataSet>::iterator itor
      = theDataSet.begin(); itor != theDataSet.end(); itor++)
      {
      os << "Find Response: " << (itor - theDataSet.begin() + 1) << std::endl;
      p.PrintDataSet( *itor, os );
      os << std::endl;
      }
	std::string retStr(ss.str());
	return true;
}

