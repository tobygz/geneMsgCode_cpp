#include "msg.h"
 
msg_c2s_2000* msg_c2s_2000::m_pInst = new msg_c2s_2000 ;
 size_t msg_c2s_2000::size(){
 size_t sizev = 0;
 sizev = sizev + 4;
 sizev = sizev + 4;
 sizev = sizev + 4;
 sizev = sizev + 4; return sizev;
}
 size_t msg_c2s_2000::pack(char* buf){
 char* p = buf;
size_t len = 0, sz = 0; memcpy(p,&Pid,4); p+=4; memcpy(p,&Roomid,4); p+=4; memcpy(p,&Data,4); p+=4; memcpy(p,&FrameIdx,4); p+=4;
 return p-buf;
}
 size_t msg_c2s_2000::parse(char* buf, size_t _len){
 char* p = buf;
 size_t len = 0, len1=0;
 memcpy(&Pid,p,4); p+=4;
 memcpy(&Roomid,p,4); p+=4;
 memcpy(&Data,p,4); p+=4;
 memcpy(&FrameIdx,p,4); p+=4;
 return p-buf;
}
 
msg_s2c_2001* msg_s2c_2001::m_pInst = new msg_s2c_2001 ;
 size_t msg_s2c_2001::size(){
 size_t sizev = 0;
 sizev = sizev + 4;
 sizev = sizev + 4 + 4 * Datalist.size();
 sizev = sizev + 4 + 4 * Difflist.size(); return sizev;
}
 size_t msg_s2c_2001::pack(char* buf){
 char* p = buf;
size_t len = 0, sz = 0; memcpy(p,&Frameid,4); p+=4;
 len = Datalist.size(); memcpy(p,&len,4); p+=4;
 for(int i=0;i<Datalist.size(); i++){ memcpy(p, &Datalist[i], 4); p+=4;}
 len = Difflist.size(); memcpy(p,&len,4); p+=4;
 for(int i=0;i<Difflist.size(); i++){ memcpy(p, &Difflist[i], 4); p+=4;}
 return p-buf;
}
 size_t msg_s2c_2001::parse(char* buf, size_t _len){
 char* p = buf;
 size_t len = 0, len1=0;
 memcpy(&Frameid,p,4); p+=4;
 memcpy(&len,p,4); p+=4; for(int i=0; i<len; i++){UINT32 u32v = 0; memcpy(&u32v,p, 4); p+=4;Datalist[i]=u32v; }
 memcpy(&len,p,4); p+=4; for(int i=0; i<len; i++){UINT32 u32v = 0; memcpy(&u32v,p, 4); p+=4;Difflist[i]=u32v; }
 return p-buf;
}
 
msg_c2s_2002* msg_c2s_2002::m_pInst = new msg_c2s_2002 ;
 size_t msg_c2s_2002::size(){
 size_t sizev = 0;
 sizev = sizev + 4 + 4 * Ids.size(); return sizev;
}
 size_t msg_c2s_2002::pack(char* buf){
 char* p = buf;
size_t len = 0, sz = 0;
 len = Ids.size(); memcpy(p,&len,4); p+=4;
 for(int i=0;i<Ids.size(); i++){ memcpy(p, &Ids[i], 4); p+=4;}
 return p-buf;
}
 size_t msg_c2s_2002::parse(char* buf, size_t _len){
 char* p = buf;
 size_t len = 0, len1=0;
 memcpy(&len,p,4); p+=4; for(int i=0; i<len; i++){UINT32 u32v = 0; memcpy(&u32v,p, 4); p+=4;Ids[i]=u32v; }
 return p-buf;
}
 
msg_s2c_2002* msg_s2c_2002::m_pInst = new msg_s2c_2002 ;
 size_t msg_s2c_2002::size(){
 size_t sizev = 0;
 sizev += 4; for(int i=0; i<FrameList.size(); i++){ sizev += FrameList[i]->size(); } return sizev;
}
 size_t msg_s2c_2002::pack(char* buf){
 char* p = buf;
size_t len = 0, sz = 0;
 len = FrameList.size(); memcpy(p, &len,4); p+=4; for(int i=0;i<len;i++){ p += FrameList[i]->pack(p); }
 return p-buf;
}
 size_t msg_s2c_2002::parse(char* buf, size_t _len){
 char* p = buf;
 size_t len = 0, len1=0;
 memcpy(&len,p,4); p+=4; for(int i=0;i<len;i++){ p+= FrameList[i]->parse(p,_len-(p-buf)); }
 return p-buf;
}
 
msg_c2s_2004* msg_c2s_2004::m_pInst = new msg_c2s_2004 ;
 size_t msg_c2s_2004::size(){
 size_t sizev = 0;
 sizev = sizev + 8; return sizev;
}
 size_t msg_c2s_2004::pack(char* buf){
 char* p = buf;
size_t len = 0, sz = 0; memcpy(p,&FrameIdx,8); p+=8;
 return p-buf;
}
 size_t msg_c2s_2004::parse(char* buf, size_t _len){
 char* p = buf;
 size_t len = 0, len1=0;
 memcpy(&FrameIdx,p,8); p+=8;
 return p-buf;
}
