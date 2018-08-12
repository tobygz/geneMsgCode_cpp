#include "msg.h"
 
C2s_2000* C2s_2000::m_pInst = new C2s_2000 ;
 size_t C2s_2000::size(){
 size_t sizev = 0;
 sizev = sizev + 4;
 sizev = sizev + 4;
 sizev = sizev + 4 + 4 * data.size();
 sizev = sizev + 4 + 4 * frameidx.size();
 sizev = sizev + 4 + name.size(); return sizev;
}
 size_t C2s_2000::pack(char* buf){
 char* p = buf;
size_t len = 0, sz = 0; memcpy(p,&pid,4); p+=4; memcpy(p,&roomid,4); p+=4;
 len = data.size(); memcpy(p,&len,4); p+=4;
 for(int i=0;i<data.size(); i++){ memcpy(p, &data[i], 4); p+=4;}
 len = frameidx.size(); memcpy(p,&len,4); p+=4;
 for(int i=0;i<frameidx.size(); i++){ memcpy(p, &frameidx[i], 4); p+=4;}
 len=name.size();
 memcpy(p,&len,4); p+=4;
 memcpy(p,name.c_str(),len); p+=len;
 return p-buf;
}
 size_t C2s_2000::parse(char* buf, size_t _len){
 char* p = buf;
 size_t len = 0, len1=0;
 memcpy(&pid,p,4); p+=4;
 memcpy(&roomid,p,4); p+=4;
 memcpy(&len,p,4); p+=4; for(int i=0; i<len; i++){UINT32 u32v = 0; memcpy(&u32v,p, 4); p+=4;data[i]=u32v; }
 memcpy(&len,p,4); p+=4; for(int i=0; i<len; i++){UINT32 u32v = 0; memcpy(&u32v,p, 4); p+=4;frameidx[i]=u32v; }
 len1 =name.size();
 memcpy(p,&len1,4); p+=4;
 memcpy(p,name.c_str(),len1); p+=len1;
 return p-buf;
}
 
C2s_2001* C2s_2001::m_pInst = new C2s_2001 ;
 size_t C2s_2001::size(){
 size_t sizev = 0;
 sizev += 4; for(int i=0; i<lst.size(); i++){ sizev += lst[i]->size(); }
 sizev += single.size(); return sizev;
}
 size_t C2s_2001::pack(char* buf){
 char* p = buf;
size_t len = 0, sz = 0;
 len = lst.size(); memcpy(p, &len,4); p+=4; for(int i=0;i<len;i++){ p += lst[i]->pack(p); }
 p += single.pack(p);
 return p-buf;
}
 size_t C2s_2001::parse(char* buf, size_t _len){
 char* p = buf;
 size_t len = 0, len1=0;
 memcpy(&len,p,4); p+=4; for(int i=0;i<len;i++){ p+= lst[i]->parse(p,_len-(p-buf)); }
 p += single.parse(p, _len-(p-buf));
 return p-buf;
}
