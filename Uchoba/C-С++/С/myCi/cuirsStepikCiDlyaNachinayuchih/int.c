#include <stdio.h>

// int main() {

//     int a = 10;
//     int b = 11;
//     int c = a + b;

//     printf("%d + %d = %d", a, b, c);

//     return 0;
// }

// int main() {

//     int sec, min;
//     sec = 60;
//     min = 60;

//     printf("В 2-ух часах %d минту и %d секунд.", min * 2, min * sec * 2);

//     return 0; 
// }

// int main() {

//     int sec, min, hour;
//     sec = 60;
//     min = 60;
    
//     printf("Введите сколько ехать в часах: ");
//     scanf("%d", &hour);

//     printf("Ехать %d часов2, в минутах это %d мин., в секундах это %d сек.", hour, hour * min, hour * min * sec);

//     return 0;
// }

// int main() {
//     // Напишем программу. Дано сколько часов и минут длился эксперимент. Нужно напечатать:
//     // сколько длился эксперимент в минутах,
//     // сколько длился эксперимент в секундах.

//     int min, hou;

//     // printf("Сколько часов и минут длился эксперимент?\n");
//     // printf("Часы: ");
//     scanf("%d", &hou);
//     // printf("Минуты: ");
//     scanf("%d", &min);

//     int ixspMin = hou * 60 + min;

//     // printf("Экспиремент длялся в минутах %d.\n", ixspMin);
//     // printf("И также в секундах %d.\n", ixspMin * 60);
//     printf("%d\n", ixspMin);
//     printf("%d\n", ixspMin * 60);  

//     return 0;
// }

int main() {

    int k_rice, k_veg, suma;

    scanf("%d", &k_rice);
    scanf("%d", &k_veg);

    suma = k_rice + k_veg * 2;

    printf("%d", suma);

    return 0;
}