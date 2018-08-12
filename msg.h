#include <iostream>
#include <vector>
#include <string.h>
using namespace std;
#define UINT32 unsigned int
 
class C2s_2000 {
public:
 UINT32 pid;
 UINT32 roomid;
 vector<UINT32> data;
 vector<UINT32> frameidx;
 string name;
 size_t size();
 size_t pack(char*);
 size_t parse(char* , size_t );
 static C2s_2000 *m_pInst;
 static C2s_2000* getInst(){  if(m_pInst==NULL){ m_pInst = new C2s_2000;} return m_pInst;}
}; 
class C2s_2001 {
public:
 vector<C2s_2000*> lst;
 C2s_2000 single;
 size_t size();
 size_t pack(char*);
 size_t parse(char* , size_t );
 static C2s_2001 *m_pInst;
 static C2s_2001* getInst(){  if(m_pInst==NULL){ m_pInst = new C2s_2001;} return m_pInst;}
};