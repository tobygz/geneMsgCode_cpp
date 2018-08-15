#include <iostream>
#include <vector>
#include <string.h>
using namespace std;
#define UINT32 unsigned int
#define UINT64 unsigned long long 
class msg_c2s_2000 {
public:
 UINT32 Pid;
 UINT32 Roomid;
 UINT32 Data;
 UINT32 FrameIdx;
 size_t size();
 size_t pack(char*);
 size_t parse(char* , size_t );
 static msg_c2s_2000 *m_pInst;
 static msg_c2s_2000* getInst(){  if(m_pInst==NULL){ m_pInst = new msg_c2s_2000;} return m_pInst;}
}; 
class msg_s2c_2001 {
public:
 UINT32 Frameid;
 vector<UINT32> Datalist;
 vector<UINT32> Difflist;
 size_t size();
 size_t pack(char*);
 size_t parse(char* , size_t );
 static msg_s2c_2001 *m_pInst;
 static msg_s2c_2001* getInst(){  if(m_pInst==NULL){ m_pInst = new msg_s2c_2001;} return m_pInst;}
}; 
class msg_c2s_2002 {
public:
 vector<UINT32> Ids;
 size_t size();
 size_t pack(char*);
 size_t parse(char* , size_t );
 static msg_c2s_2002 *m_pInst;
 static msg_c2s_2002* getInst(){  if(m_pInst==NULL){ m_pInst = new msg_c2s_2002;} return m_pInst;}
}; 
class msg_s2c_2002 {
public:
 vector<msg_s2c_2001*> FrameList;
 size_t size();
 size_t pack(char*);
 size_t parse(char* , size_t );
 static msg_s2c_2002 *m_pInst;
 static msg_s2c_2002* getInst(){  if(m_pInst==NULL){ m_pInst = new msg_s2c_2002;} return m_pInst;}
}; 
class msg_c2s_2004 {
public:
 UINT64 FrameIdx;
 size_t size();
 size_t pack(char*);
 size_t parse(char* , size_t );
 static msg_c2s_2004 *m_pInst;
 static msg_c2s_2004* getInst(){  if(m_pInst==NULL){ m_pInst = new msg_c2s_2004;} return m_pInst;}
};