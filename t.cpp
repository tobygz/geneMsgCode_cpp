#include "msg.h"
#include "stdio.h"

int main(){
    C2s_2000::getInst()->pid = 100;
    printf("%d\n", C2s_2000::getInst()->pid );


    return 0;
}
