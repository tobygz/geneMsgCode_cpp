# geneMsgCode_cpp
generate cpp pack &amp; parse code by golang struct reflect

describe proto by golang struct
------------------------------------
type C2s_2000 struct {
    pid      uint32
    roomid   uint32
    data     uint32 `repeat`
    frameidx uint32 `repeat`
    name     string
}

type C2s_2001 struct {
    lst    C2s_2000 `repeat`
    single C2s_2000
}

------------------------------------
use main.go to generate msg.h&msg.cpp file which contain pack&parse function.
