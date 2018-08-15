package msg 

import (
    "bytes"
    "encoding/binary"
    "unsafe"
)

func Str2bytes(s string) []byte {
    x := (*[2]uintptr)(unsafe.Pointer(&s))
    h := [3]uintptr{x[0], x[1], x[1]}
    return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2str(b []byte) string {
    return *(*string)(unsafe.Pointer(&b))
} 
type msg_c2s_2000 struct {
 Pid uint32
 Roomid uint32
 Data uint32
 FrameIdx uint32
 }
 func (this *msg_c2s_2000) Size() uint32{
 sizev := uint32(0)
 sizev = sizev + 4
 sizev = sizev + 4
 sizev = sizev + 4
 sizev = sizev + 4
 return sizev 
}
 func (this *msg_c2s_2000) Pack() []byte {
 outbuff := bytes.NewBuffer([]byte{})
                if err := binary.Write(outbuff, binary.LittleEndian, this.Pid); err != nil {
                        panic(err)
                    }                   
                
                if err := binary.Write(outbuff, binary.LittleEndian, this.Roomid); err != nil {
                        panic(err)
                    }                   
                
                if err := binary.Write(outbuff, binary.LittleEndian, this.Data); err != nil {
                        panic(err)
                    }                   
                
                if err := binary.Write(outbuff, binary.LittleEndian, this.FrameIdx); err != nil {
                        panic(err)
                    }                   
                
 return outbuff.Bytes() 
}
 func (this *msg_c2s_2000) Parse(_data []byte) {
 outbuff := bytes.NewBuffer(_data)
                if err := binary.Read(outbuff, binary.LittleEndian, &this.Pid); err != nil {
                    panic(err)
                }                   
                
                if err := binary.Read(outbuff, binary.LittleEndian, &this.Roomid); err != nil {
                    panic(err)
                }                   
                
                if err := binary.Read(outbuff, binary.LittleEndian, &this.Data); err != nil {
                    panic(err)
                }                   
                
                if err := binary.Read(outbuff, binary.LittleEndian, &this.FrameIdx); err != nil {
                    panic(err)
                }                   
                
 } 
type msg_s2c_2001 struct {
 Frameid uint32
 Datalist []uint32
 Difflist []uint32
 }
 func (this *msg_s2c_2001) Size() uint32{
 sizev := uint32(0)
 sizev = sizev + 4
 sizev = sizev + 4 + 4*uint32(len(this.Datalist)) 
 sizev = sizev + 4 + 4*uint32(len(this.Difflist)) 
 return sizev 
}
 func (this *msg_s2c_2001) Pack() []byte {
 outbuff := bytes.NewBuffer([]byte{})
                if err := binary.Write(outbuff, binary.LittleEndian, this.Frameid); err != nil {
                        panic(err)
                    }                   
                
                if err := binary.Write(outbuff, binary.LittleEndian, uint32(len(this.Datalist)) ); err != nil {
                    panic(err)
                }
                for _,val := range this.Datalist {
                    if err := binary.Write(outbuff, binary.LittleEndian, val); err != nil {
                        panic(err)
                    }                   
                }
                if err := binary.Write(outbuff, binary.LittleEndian, uint32(len(this.Difflist)) ); err != nil {
                    panic(err)
                }
                for _,val := range this.Difflist {
                    if err := binary.Write(outbuff, binary.LittleEndian, val); err != nil {
                        panic(err)
                    }                   
                }
 return outbuff.Bytes() 
}
 func (this *msg_s2c_2001) Parse(_data []byte) {
 outbuff := bytes.NewBuffer(_data)
                if err := binary.Read(outbuff, binary.LittleEndian, &this.Frameid); err != nil {
                    panic(err)
                }                   
                
                {
                lenv := 0
                if err := binary.Read(outbuff, binary.LittleEndian, &lenv ); err != nil {
                    panic(err)
                }
                for i:=0;i<lenv;i++ {
                    var _tmp uint32
                    if err := binary.Read(outbuff, binary.LittleEndian, &_tmp); err != nil {
                        panic(err)
                    }
                    this.Datalist = append(this.Datalist, _tmp)                 
                }
                }
                {
                lenv := 0
                if err := binary.Read(outbuff, binary.LittleEndian, &lenv ); err != nil {
                    panic(err)
                }
                for i:=0;i<lenv;i++ {
                    var _tmp uint32
                    if err := binary.Read(outbuff, binary.LittleEndian, &_tmp); err != nil {
                        panic(err)
                    }
                    this.Difflist = append(this.Difflist, _tmp)                 
                }
                }
 } 
type msg_c2s_2002 struct {
 Ids []uint32
 }
 func (this *msg_c2s_2002) Size() uint32{
 sizev := uint32(0)
 sizev = sizev + 4 + 4*uint32(len(this.Ids)) 
 return sizev 
}
 func (this *msg_c2s_2002) Pack() []byte {
 outbuff := bytes.NewBuffer([]byte{})
                if err := binary.Write(outbuff, binary.LittleEndian, uint32(len(this.Ids)) ); err != nil {
                    panic(err)
                }
                for _,val := range this.Ids {
                    if err := binary.Write(outbuff, binary.LittleEndian, val); err != nil {
                        panic(err)
                    }                   
                }
 return outbuff.Bytes() 
}
 func (this *msg_c2s_2002) Parse(_data []byte) {
 outbuff := bytes.NewBuffer(_data)
                {
                lenv := 0
                if err := binary.Read(outbuff, binary.LittleEndian, &lenv ); err != nil {
                    panic(err)
                }
                for i:=0;i<lenv;i++ {
                    var _tmp uint32
                    if err := binary.Read(outbuff, binary.LittleEndian, &_tmp); err != nil {
                        panic(err)
                    }
                    this.Ids = append(this.Ids, _tmp)                 
                }
                }
 } 
type msg_s2c_2002 struct {
 FrameList []msg_s2c_2001
 }
 func (this *msg_s2c_2002) Size() uint32{
 sizev := uint32(0)
 sizev = sizev + 4
for _,val := range this.FrameList {
sizev = sizev + val.Size() 
} 
 return sizev 
}
 func (this *msg_s2c_2002) Pack() []byte {
 outbuff := bytes.NewBuffer([]byte{})
                if err := binary.Write(outbuff, binary.LittleEndian, uint32(len(this.FrameList)) ); err != nil {
                    panic(err)
                }
                for _,val := range this.FrameList {
                    if err := binary.Write(outbuff, binary.LittleEndian, val.Pack() ); err != nil {
                        panic(err)
                    }
                }
 return outbuff.Bytes() 
}
 func (this *msg_s2c_2002) Parse(_data []byte) {
 outbuff := bytes.NewBuffer(_data)
                if err := binary.Write(outbuff, binary.LittleEndian, uint32(len(this.FrameList)) ); err != nil {
                    panic(err)
                }
                for _,val := range this.FrameList {
                    if err := binary.Write(outbuff, binary.LittleEndian, val.Pack() ); err != nil {
                        panic(err)
                    }
                }
 } 
type msg_c2s_2004 struct {
 FrameIdx uint64
 }
 func (this *msg_c2s_2004) Size() uint32{
 sizev := uint32(0)
 sizev = sizev + 8
 return sizev 
}
 func (this *msg_c2s_2004) Pack() []byte {
 outbuff := bytes.NewBuffer([]byte{})
                if err := binary.Write(outbuff, binary.LittleEndian, this.FrameIdx); err != nil {
                        panic(err)
                    }                   
                
 return outbuff.Bytes() 
}
 func (this *msg_c2s_2004) Parse(_data []byte) {
 outbuff := bytes.NewBuffer(_data)
                if err := binary.Read(outbuff, binary.LittleEndian, &this.FrameIdx); err != nil {
                    panic(err)
                }                   
                
 }