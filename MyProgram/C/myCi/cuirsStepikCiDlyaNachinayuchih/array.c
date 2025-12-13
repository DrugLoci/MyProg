#include <stdio.h>

int main() {

    int a = 5;
    int * pa = &a;
    *pa = 6;
    int point = (int)pa;

    printf("%d\n", a);
    printf("%S", pa);
    
    return 0;
}