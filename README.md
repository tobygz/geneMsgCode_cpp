# geneMsgCode_cpp
generate cpp pack &amp; parse code by golang struct reflect

describe proto by golang struct
------------------------------------
type C2s_2000 struct {<br>
    pid      uint32<br>
    roomid   uint32<br>
    data     uint32 `repeat`<br>
    frameidx uint32 `repeat`<br>
    name     string<br>
}<br>
<br>
type C2s_2001 struct {<br>
    lst    C2s_2000 `repeat`<br>
    single C2s_2000<br>
}<br>
<br>
------------------------------------
use main.go to generate msg.h&msg.cpp file which contain pack&parse function.
