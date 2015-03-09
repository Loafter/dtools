#ifndef GDCMGOBR_H
#define GDCMGOBR_H




#ifdef __cplusplus
extern "C" {
#endif

int test();

int CEcho (const char*, int, const char*, const char* );
int CStore (const char*, int , const char*, const char*,const char*);
const char* CFind(const char*,const char*,const char* ,int,const char*,const char*,const char*,const char*,const char*);
int CGet(const char*,const char*,const char*,int,const char*,	const char*,const char*,const char* ,const char*,const char*);

#ifdef __cplusplus
}
#endif

#endif