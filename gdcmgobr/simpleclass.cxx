

#include <iostream>
#include "simpleclass.h"

bool CEcho (const char *remote, int portno, const char *aetitle, const char *call){
	return gdcm::CompositeNetworkFunctions::CEcho(remote,portno,aetitle,call);
}

