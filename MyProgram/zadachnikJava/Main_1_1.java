// ссылка на задачник https://code.mu/ru/java/tasker/stager/
// Дано число. Проверьте, отрицательное оно или нет. Выведите об этом информацию в консоль.

import myLogicProgram.Or;
import myNum.GeneraciyaNum;
import myString.MyPrint;

public class Main_1_1 {

    public static void main(String[] args) {

        // создаем объекты 
        GeneraciyaNum generNum = new GeneraciyaNum();
        Or poloOrOtri = new Or();
        MyPrint print = new MyPrint();

        // генерируем число 
        int num = generNum.generaciya1Num();

        // определяем пооложительное число или отрицательное 
        String PolpOrOtriNum = poloOrOtri.negativeOrPasitive(num);

        // выводим в косоль резултат 
        print.print1StringInt(num);
        print.print1String(PolpOrOtriNum);


    }
}