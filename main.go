package main

import (
	"fmt"
	"os"
	"reflect"
)

type msg_c2s_2000 struct {
	Pid      uint32
	Roomid   uint32
	Data     uint32
	FrameIdx uint32
}

type msg_s2c_2001 struct {
	Frameid  uint32
	Datalist uint32 `repeat`
	Difflist uint32 `repeat`
}

type msg_c2s_2002 struct {
	Ids uint32 `repeat`
}

type msg_s2c_2002 struct {
	FrameList msg_s2c_2001 `repeat`
}

type msg_c2s_2004 struct {
	FrameIdx uint64
}

func geneCppClass_v1_header(t reflect.Type) string {
	retval := fmt.Sprintf("class %s {\npublic:", t.Name())
	//gene class

	for i := 0; i < t.NumField(); i++ {
		nowf := t.Field(i)
		if nowf.Type.Kind() == reflect.Uint64 {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n vector<UINT64> %s;", retval, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n UINT64 %s;", retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.Uint32 {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n vector<UINT32> %s;", retval, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n UINT32 %s;", retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.String {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n vector<string> %s;", retval, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n string %s;", retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.Struct {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n vector<%s*> %s;", retval, nowf.Type.Name(), nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n %s %s;", retval, nowf.Type.Name(), nowf.Name)
			}
		}
	}

	retval = fmt.Sprintf("%s\n size_t size();", retval)
	retval = fmt.Sprintf("%s\n size_t pack(char*);", retval)
	retval = fmt.Sprintf("%s\n size_t parse(char* , size_t );", retval)

	//generate static instance
	retval = fmt.Sprintf("%s\n static %s *m_pInst;", retval, t.Name())
	retval = fmt.Sprintf("%s\n static %s* getInst(){  if(m_pInst==NULL){ m_pInst = new %s;} return m_pInst;}\n};", retval, t.Name(), t.Name())

	return retval
}

func geneCppClass_v1_cpp(t reflect.Type) string {

	retval := fmt.Sprintf("%s* %s::m_pInst = new %s ;", t.Name(), t.Name(), t.Name())

	//gene size func
	retval = fmt.Sprintf("%s\n size_t %s::size(){", retval, t.Name())
	retval = fmt.Sprintf("%s\n size_t sizev = 0;", retval)

	for i := 0; i < t.NumField(); i++ {
		nowf := t.Field(i)
		if nowf.Type.Kind() == reflect.Uint64 {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n sizev = sizev + 4 + 8 * %s.size();", retval, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n sizev = sizev + 8;", retval)
			}
		} else if nowf.Type.Kind() == reflect.Uint32 {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n sizev = sizev + 4 + 4 * %s.size();", retval, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n sizev = sizev + 4;", retval)
			}
		} else if nowf.Type.Kind() == reflect.String {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n sizev += %s.size(); for(int i=0; i<%s.size();i++){ sizev +=4; sizev += %s.size(); }", retval, nowf.Name, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n sizev = sizev + 4 + %s.size();", retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.Struct {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n sizev += 4; for(int i=0; i<%s.size(); i++){ sizev += %s[i]->size(); }", retval, nowf.Name, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n sizev += %s.size();", retval, nowf.Name)
			}
		}
	}
	retval = fmt.Sprintf("%s return sizev;\n}", retval)

	//gene pack func
	retval = fmt.Sprintf("%s\n size_t %s::pack(char* buf){", retval, t.Name())
	retval = fmt.Sprintf("%s\n char* p = buf;\nsize_t len = 0, sz = 0;", retval)

	for i := 0; i < t.NumField(); i++ {
		nowf := t.Field(i)
		if nowf.Type.Kind() == reflect.Uint64 {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n len = %s.size(); memcpy(p,&len,4); p+=4;", retval, nowf.Name)
				retval = fmt.Sprintf("%s\n for(int i=0;i<%s.size(); i++){ memcpy(p, &%s[i], 8); p+=8;}", retval, nowf.Name, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s memcpy(p,&%s,8); p+=8;", retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.Uint32 {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n len = %s.size(); memcpy(p,&len,4); p+=4;", retval, nowf.Name)
				retval = fmt.Sprintf("%s\n for(int i=0;i<%s.size(); i++){ memcpy(p, &%s[i], 4); p+=4;}", retval, nowf.Name, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s memcpy(p,&%s,4); p+=4;", retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.String {
			if nowf.Tag == "repeat" {
				//memcpy(p, &%s.size(), 4); p+=4; for(int i=0; i<%s.size();i++){ memcpy(p,&%s[i].size(),4); p+= 4;  memcpy(p,%s[i].c_str(),%s[i].size() ); }
				retval = fmt.Sprintf("%s\n sz = %s.size(); memcpy(p, &sz, 4); p+=4; for(int i=0; i<%s.size();i++){sz =%s[i].size();  memcpy(p,&sz,4); p+= 4;  memcpy(p,%s[i].c_str(),%s[i].size() ); }",
					retval, nowf.Name, nowf.Name, nowf.Name, nowf.Name, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n len=%s.size();", retval, nowf.Name)
				retval = fmt.Sprintf("%s\n memcpy(p,&len,4); p+=4;", retval)
				retval = fmt.Sprintf("%s\n memcpy(p,%s.c_str(),len); p+=len;", retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.Struct {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n len = %s.size(); memcpy(p, &len,4); p+=4; for(int i=0;i<len;i++){ p += %s[i]->pack(p); }", retval, nowf.Name, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n p += %s.pack(p);", retval, nowf.Name)
			}
		}
	}
	retval = fmt.Sprintf("%s\n return p-buf;\n}", retval)

	//gene parse func
	retval = fmt.Sprintf("%s\n size_t %s::parse(char* buf, size_t _len){", retval, t.Name())
	retval = fmt.Sprintf("%s\n char* p = buf;\n size_t len = 0, len1=0;", retval)

	for i := 0; i < t.NumField(); i++ {
		nowf := t.Field(i)
		if nowf.Type.Kind() == reflect.Uint64 {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n memcpy(&len,p,4); p+=4; for(int i=0; i<len; i++){UINT64 u64v = 0; memcpy(&u64v,p, 8); p+=8;%s[i]=u64v; }", retval, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n memcpy(&%s,p,8); p+=8;", retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.Uint32 {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n memcpy(&len,p,4); p+=4; for(int i=0; i<len; i++){UINT32 u32v = 0; memcpy(&u32v,p, 4); p+=4;%s[i]=u32v; }", retval, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n memcpy(&%s,p,4); p+=4;", retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.String {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n memcpy(&len,p,4); p+=4; for(int i=0; i<len; i++){ memcpy(&len1,p, 4); p+=4; string str(p,len1); %s[i] = str; }", retval, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n len1 =%s.size();", retval, nowf.Name)
				retval = fmt.Sprintf("%s\n memcpy(p,&len1,4); p+=4;", retval)
				retval = fmt.Sprintf("%s\n memcpy(p,%s.c_str(),len1); p+=len1;", retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.Struct {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n memcpy(&len,p,4); p+=4; for(int i=0;i<len;i++){ p+= %s[i]->parse(p,_len-(p-buf)); }", retval, nowf.Name)
			} else {
				//p += %s->parse(p);
				retval = fmt.Sprintf("%s\n p += %s.parse(p, _len-(p-buf));", retval, nowf.Name)
			}
		}
	}
	retval = fmt.Sprintf("%s\n return p-buf;\n}\n", retval)

	return retval
}

func geneCpp() {
	work_lst := make([]reflect.Type, 0)
	work_lst = append(work_lst, reflect.TypeOf(msg_c2s_2000{}))
	work_lst = append(work_lst, reflect.TypeOf(msg_s2c_2001{}))
	work_lst = append(work_lst, reflect.TypeOf(msg_c2s_2002{}))
	work_lst = append(work_lst, reflect.TypeOf(msg_s2c_2002{}))
	work_lst = append(work_lst, reflect.TypeOf(msg_c2s_2004{}))

	header := fmt.Sprintf("#include <iostream>\n#include <vector>\n#include <string.h>\nusing namespace std;\n#define UINT32 unsigned int\n#define UINT64 unsigned long long")
	for _, elem := range work_lst {
		ret := geneCppClass_v1_header(elem)
		header = fmt.Sprintf("%s \n%s", header, ret)
	}

	body := "#include \"msg.h\"\n"
	for _, elem := range work_lst {
		ret := geneCppClass_v1_cpp(elem)
		body = fmt.Sprintf("%s \n%s", body, ret)
	}

	f, err := os.OpenFile("msg.h", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	f.WriteString(header)
	f.Sync()
	f.Close()

	f1, err1 := os.OpenFile("msg.cpp", os.O_RDWR|os.O_CREATE, 0755)
	if err1 != nil {
		panic(err1)
	}
	f1.WriteString(body)
	f1.Sync()
	f1.Close()
}

func geneGolang() {
	work_lst := make([]reflect.Type, 0)
	work_lst = append(work_lst, reflect.TypeOf(msg_c2s_2000{}))
	work_lst = append(work_lst, reflect.TypeOf(msg_s2c_2001{}))
	work_lst = append(work_lst, reflect.TypeOf(msg_c2s_2002{}))
	work_lst = append(work_lst, reflect.TypeOf(msg_s2c_2002{}))
	work_lst = append(work_lst, reflect.TypeOf(msg_c2s_2004{}))

	header := `package msg 

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
}`
	for _, elem := range work_lst {
		ret := geneGolang_elem(elem)
		header = fmt.Sprintf("%s \n%s", header, ret)
	}

	f, err := os.OpenFile("msg.go", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	f.WriteString(header)
	f.Sync()
	f.Close()
}

func geneGolang_elem(t reflect.Type) string {
	//t := reflect.TypeOf(msg_c2s_2000{})
	//gene struct desc
	retval := fmt.Sprintf("type %s struct {", t.Name())
	for i := 0; i < t.NumField(); i++ {
		nowf := t.Field(i)
		if nowf.Type.Kind() == reflect.Uint64 {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n %s []uint64", retval, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n %s uint64", retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.Uint32 {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n %s []uint32", retval, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n %s uint32", retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.String {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n %s []string", retval, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n %s string", retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.Struct {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n %s []%s", retval, nowf.Name, nowf.Type.Name())
			} else {
				retval = fmt.Sprintf("%s\n %s %s", retval, nowf.Name, nowf.Type.Name())
			}
		}
	}
	retval = fmt.Sprintf("%s\n }", retval)

	//gene size func
	retval = fmt.Sprintf("%s\n func (this *%s) Size() uint32{\n sizev := uint32(0)", retval, t.Name())
	for i := 0; i < t.NumField(); i++ {
		nowf := t.Field(i)
		if nowf.Type.Kind() == reflect.Uint64 {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n sizev = sizev + 4 + 8*uint32(len(this.%s)) ", retval, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n sizev = sizev + 8", retval)
			}
		} else if nowf.Type.Kind() == reflect.Uint32 {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n sizev = sizev + 4 + 4*uint32(len(this.%s)) ", retval, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n sizev = sizev + 4", retval)
			}
		} else if nowf.Type.Kind() == reflect.String {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n sizev = sizev + 4\nfor _,val := range this.%s {\nsizev = sizev + 4 + uint32(len(val))\n} ", retval, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n sizev = sizev + 4 + uint32(len(this.%s))", retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.Struct {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf("%s\n sizev = sizev + 4\nfor _,val := range this.%s {\nsizev = sizev + val.Size() \n} ", retval, nowf.Name)
			} else {
				retval = fmt.Sprintf("%s\n sizev = sizev + %s.Size()", retval, nowf.Name)
			}
		}
	}
	retval = fmt.Sprintf("%s\n return sizev \n}", retval)

	//gene pack
	retval = fmt.Sprintf("%s\n func (this *%s) Pack() []byte {\n outbuff := bytes.NewBuffer([]byte{})", retval, t.Name())
	for i := 0; i < t.NumField(); i++ {
		nowf := t.Field(i)
		if nowf.Type.Kind() == reflect.Uint64 || nowf.Type.Kind() == reflect.Uint32 {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf(`%s
                if err := binary.Write(outbuff, binary.LittleEndian, uint32(len(this.%s)) ); err != nil {
                    panic(err)
                }
                for _,val := range this.%s {
                    if err := binary.Write(outbuff, binary.LittleEndian, val); err != nil {
                        panic(err)
                    }                   
                }`, retval, nowf.Name, nowf.Name)
			} else {
				retval = fmt.Sprintf(`%s
                if err := binary.Write(outbuff, binary.LittleEndian, this.%s); err != nil {
                        panic(err)
                    }                   
                `, retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.String {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf(`%s
                if err := binary.Write(outbuff, binary.LittleEndian, uint32(len(this.%s)) ); err != nil {
                    panic(err)
                }
                for _,val := range this.%s {
                    if err := binary.Write(outbuff, binary.LittleEndian, uint32(len(val))); err != nil {
                        panic(err)
                    }
                    if err := binary.Write(outbuff, binary.LittleEndian, Str2bytes(val) ); err != nil {
                        panic(err)
                    }
                }`, retval, nowf.Name, nowf.Name)

			} else {
				retval = fmt.Sprintf(`%s
                    if err := binary.Write(outbuff, binary.LittleEndian, uint32(len(this.%s))); err != nil {
                        panic(err)
                    }
                    if err := binary.Write(outbuff, binary.LittleEndian, Str2bytes(this.%s) ); err != nil {
                        panic(err)
                    }`, retval, nowf.Name, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.Struct {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf(`%s
                if err := binary.Write(outbuff, binary.LittleEndian, uint32(len(this.%s)) ); err != nil {
                    panic(err)
                }
                for _,val := range this.%s {
                    if err := binary.Write(outbuff, binary.LittleEndian, val.Pack() ); err != nil {
                        panic(err)
                    }
                }`, retval, nowf.Name, nowf.Name)
			} else {
				retval = fmt.Sprintf(`%s
                if err := binary.Write(outbuff, binary.LittleEndian, this.%s.Pack() ); err != nil {
                    panic(err)
                }`, retval, nowf.Name)
			}
		}
	}

	retval = fmt.Sprintf("%s\n return outbuff.Bytes() \n}", retval)

	//gene parse
	retval = fmt.Sprintf("%s\n func (this *%s) Parse(_data []byte) {\n outbuff := bytes.NewBuffer(_data)", retval, t.Name())
	for i := 0; i < t.NumField(); i++ {
		nowf := t.Field(i)
		if nowf.Type.Kind() == reflect.Uint64 || nowf.Type.Kind() == reflect.Uint32 {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf(`%s
                {
                lenv := 0
                if err := binary.Read(outbuff, binary.LittleEndian, &lenv ); err != nil {
                    panic(err)
                }
                for i:=0;i<lenv;i++ {
                    var _tmp %s
                    if err := binary.Read(outbuff, binary.LittleEndian, &_tmp); err != nil {
                        panic(err)
                    }
                    this.%s = append(this.%s, _tmp)                 
                }
                }`, retval, nowf.Type.Name(), nowf.Name, nowf.Name)
			} else {
				retval = fmt.Sprintf(`%s
                if err := binary.Read(outbuff, binary.LittleEndian, &this.%s); err != nil {
                    panic(err)
                }                   
                `, retval, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.String {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf(`%s
                lenv := 0
                if err := binary.Read(outbuff, binary.LittleEndian, &lenv ); err != nil {
                    panic(err)
                }
                for i:=0;i<lenv;i++{
                    var _strlen uint32
                    if err := binary.Read(outbuff, binary.LittleEndian, &_strlen); err != nil {
                        panic(err)
                    }                   
                    _slc := make([]byte, _strlen)
                    if err := binary.Read(outbuff, binary.LittleEndian, &_slc); err != nil {
                        panic(err)
                    }   
                    this.%s = Bytes2str(_slc)               
                }
                for _,val := range this.%s {
                    if err := binary.Write(outbuff, binary.LittleEndian, uint32(len(val))); err != nil {
                        panic(err)
                    }
                    if err := binary.Write(outbuff, binary.LittleEndian, Str2bytes(val) ); err != nil {
                        panic(err)
                    }
                }`, retval, nowf.Name, nowf.Name)

			} else {
				retval = fmt.Sprintf(`%s
                    if err := binary.Write(outbuff, binary.LittleEndian, uint32(len(this.%s))); err != nil {
                        panic(err)
                    }
                    if err := binary.Write(outbuff, binary.LittleEndian, Str2bytes(this.%s) ); err != nil {
                        panic(err)
                    }`, retval, nowf.Name, nowf.Name)
			}
		} else if nowf.Type.Kind() == reflect.Struct {
			if nowf.Tag == "repeat" {
				retval = fmt.Sprintf(`%s
                if err := binary.Write(outbuff, binary.LittleEndian, uint32(len(this.%s)) ); err != nil {
                    panic(err)
                }
                for _,val := range this.%s {
                    if err := binary.Write(outbuff, binary.LittleEndian, val.Pack() ); err != nil {
                        panic(err)
                    }
                }`, retval, nowf.Name, nowf.Name)
			} else {
				retval = fmt.Sprintf(`%s
                if err := binary.Write(outbuff, binary.LittleEndian, this.%s.Pack() ); err != nil {
                    panic(err)
                }`, retval, nowf.Name)
			}
		}
	}

	retval = fmt.Sprintf("%s\n }", retval)
	return retval
	//fmt.Println(retval)
}

func main() {
	geneGolang()
	geneCpp()
}
