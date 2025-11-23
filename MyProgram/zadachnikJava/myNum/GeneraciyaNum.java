// генерация положительных и отрицательных чисел 

package myNum;

import java.util.Random;

public class GeneraciyaNum {

    public int generaciya1Num() { // генерируем случаное целое отрицательно или положительно одно число
        Random random = new Random();
        int num = random.nextInt(); 
        return num;
    }
    
}