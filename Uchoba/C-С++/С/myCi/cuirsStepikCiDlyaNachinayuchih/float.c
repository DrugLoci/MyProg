#include <stdio.h>

// int main() { // индекс массы тела человека 

//     float massa, rost, floatIndex;
//     int intIndex;

//     scanf("%f%f", &massa, &rost);

//     floatIndex = massa / (rost * rost);
//     intIndex = floatIndex;

//     printf("%d", intIndex);

//     return 0;
// }

int main() {

    int kmCh;
    float mSec, secVChas, metrVKm; 
    secVChas = 3600;
    metrVKm = 1000;

    scanf("%d", &kmCh);

    mSec = kmCh * metrVKm / secVChas;

    printf("%f", mSec);

    return 0; 

}
