

#include <iostream>
#include "gdcmgobr.h"
#include "gdcmPrinter.h"

bool CEcho (const char *remote, int portno, const char *aetitle, const char *call){
	return gdcm::CompositeNetworkFunctions::CEcho(remote,portno,aetitle,call);
}
bool CFind(bool findpatientroot, bool patientquery,bool seriesquery,
const char* callingaetitle,const char* callaetitle,bool imagequery, const char* hostname,int port){
	
	
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
    //the value in that tag corresponds to the query type
    std::vector<gdcm::DataSet> theDataSet;
    if( !gdcm::CompositeNetworkFunctions::CFind(hostname, (uint16_t)port, theQuery, theDataSet,
        callingaetitle, callaetitle) )
      {
      gdcmDebugMacro( "Problem in CFind." );
      return false;
      }
	return true;
}

