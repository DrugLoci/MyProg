// ссылка на задачник https://code.mu/ru/java/tasker/stager/
// Дано число. Проверьте, отрицательное оно или нет. Выведите об этом информацию в консоль.

import myLogicProgram.Or;
import myNum.GeneraciyaNum;
import myString.MyPrint;

public class Main_1_1 {

    public static void main(String[] args) {

        // создаем объекты 
        GeneraciyaNum generaciyaNum = new GeneraciyaNum();
        Or or = new Or();
        MyPrint myPrint = new MyPrint();

        // генерируем число 
        int num = generaciyaNum.generaciya1Num();

        // определяем пооложительное число или отрицательное 
        String PolpOrOtriNum = or.negativeOrPasitive(num);

        // выводим в косоль резултат 
        myPrint.print1StringInt(num);
        myPrint.print1String(PolpOrOtriNum);
    }
}