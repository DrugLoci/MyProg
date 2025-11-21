package my;

import java.lang.Math; 

public class TrueFalse {
    public static void main(String[] args) { // в main рандомно подберем значения true false и передадим значения методу booleanExpression
        boolean a = false;
        boolean b = false;
        boolean c = false;
        boolean d = false;

        for (int i = 1; i < 17; i++) { // хочу 16 раз передвать значения 
            double rand = 0;
            //int randInt = 0;

            for (int f = 1; f < 5; f++) { // тут я рандомно подбираю значения 
                rand = Math.random()+1;
                //randInt = (int) rand;
                //System.out.println(rand); // вывод для проверки

                switch (f) {
                    case 1:
                        if (rand < 1.5) {
                            a = false;
                        } else {
                            a = true;
                        }
                        break;
                    case 2:
                        if (rand < 1.5) {
                            b = false;
                        } else {
                            b = true;
                        }
                        break;
                    case 3:
                        if (rand < 1.5) {
                            c = false;
                        } else {
                            c = true;
                        }
                        break;
                    case 4:
                        if (rand < 1.5) {
                            d = false;
                        } else {
                            d = true;
                        }
                        break;
                }
            }

            boolean expression = booleanExpression(a, b, c, d);
            /*System.out.print("проход: ");
            System.out.println(i);
            System.out.println(a); // вывод для проверки 
            System.out.println(b);
            System.out.println(c);
            System.out.println(d);*/
            System.out.println(expression);
        }
    }

    public static boolean booleanExpression(boolean a, boolean b, boolean c, boolean d) { // должен перебрать все перменые и найти значений true. Если 2 true, вернуть true в противном слечае вернуть false 
        int count = 0; // сюда запишем количесво переменых со значением true 

        if (a == true) {
            count++;
        }

        if (b == true) {
            count++;
        }

        if (c == true) {
            count++;
        }

        if (d == true) {
            count++;
        }

        if (count == 2) {
            return true;
        } else {
            return false;
        }
        
    }
}

