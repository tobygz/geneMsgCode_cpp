#include "msg.h"
#include "stdio.h"

int main(){
    msg_s2c_2002 *p = msg_s2c_2002::getInst();

    for(int i=0;i<100;i++){
        msg_s2c_2001 *pm = new msg_s2c_2001;
        for(int j=0;j<1024;j++){
            pm->Datalist.push_back( j );
            pm->Difflist.push_back( j+j );
        }
        p->FrameList.push_back(pm);
    }

    printf("%d\n", p->size());

    for(int i=0; i<p->FrameList.size();i++){
        msg_s2c_2001 *pm = p->FrameList[i];
        printf("i: %d msg size: %d\n", i, pm->size());

    }


    return 0;
}

/* output below 
820404
i: 0 msg size: 8204
i: 1 msg size: 8204
i: 2 msg size: 8204
i: 3 msg size: 8204
i: 4 msg size: 8204
i: 5 msg size: 8204
i: 6 msg size: 8204
i: 7 msg size: 8204
i: 8 msg size: 8204
i: 9 msg size: 8204
i: 10 msg size: 8204
i: 11 msg size: 8204
i: 12 msg size: 8204
i: 13 msg size: 8204
i: 14 msg size: 8204
i: 15 msg size: 8204
i: 16 msg size: 8204
i: 17 msg size: 8204
i: 18 msg size: 8204
i: 19 msg size: 8204
i: 20 msg size: 8204
i: 21 msg size: 8204
i: 22 msg size: 8204
i: 23 msg size: 8204
i: 24 msg size: 8204
i: 25 msg size: 8204
i: 26 msg size: 8204
i: 27 msg size: 8204
i: 28 msg size: 8204
i: 29 msg size: 8204
i: 30 msg size: 8204
i: 31 msg size: 8204
i: 32 msg size: 8204
i: 33 msg size: 8204
i: 34 msg size: 8204
i: 35 msg size: 8204
i: 36 msg size: 8204
i: 37 msg size: 8204
i: 38 msg size: 8204
i: 39 msg size: 8204
i: 40 msg size: 8204
i: 41 msg size: 8204
i: 42 msg size: 8204
i: 43 msg size: 8204
i: 44 msg size: 8204
i: 45 msg size: 8204
i: 46 msg size: 8204
i: 47 msg size: 8204
i: 48 msg size: 8204
i: 49 msg size: 8204
i: 50 msg size: 8204
i: 51 msg size: 8204
i: 52 msg size: 8204
i: 53 msg size: 8204
i: 54 msg size: 8204
i: 55 msg size: 8204
i: 56 msg size: 8204
i: 57 msg size: 8204
i: 58 msg size: 8204
i: 59 msg size: 8204
i: 60 msg size: 8204
i: 61 msg size: 8204
i: 62 msg size: 8204
i: 63 msg size: 8204
i: 64 msg size: 8204
i: 65 msg size: 8204
i: 66 msg size: 8204
i: 67 msg size: 8204
i: 68 msg size: 8204
i: 69 msg size: 8204
i: 70 msg size: 8204
i: 71 msg size: 8204
i: 72 msg size: 8204
i: 73 msg size: 8204
i: 74 msg size: 8204
i: 75 msg size: 8204
i: 76 msg size: 8204
i: 77 msg size: 8204
i: 78 msg size: 8204
i: 79 msg size: 8204
i: 80 msg size: 8204
i: 81 msg size: 8204
i: 82 msg size: 8204
i: 83 msg size: 8204
i: 84 msg size: 8204
i: 85 msg size: 8204
i: 86 msg size: 8204
i: 87 msg size: 8204
i: 88 msg size: 8204
i: 89 msg size: 8204
i: 90 msg size: 8204
i: 91 msg size: 8204
i: 92 msg size: 8204
i: 93 msg size: 8204
i: 94 msg size: 8204
i: 95 msg size: 8204
i: 96 msg size: 8204
i: 97 msg size: 8204
i: 98 msg size: 8204
i: 99 msg size: 8204
*/
