package MyProg;
import java.util.Random;

public class PravdIliLoch { // правда или лож, задача передать 3 значение другому методу и други методом выяснить врем или нет 
    public static void main(String[] args) { // тут берем два числа и складывает, в рандомный момент меняем результат сложения на ложный либо оставляем так как есть 
        Random random = new Random();
        float numA = random.nextFloat();
        float numB = random.nextFloat();
        int numInt = random.nextInt(2);
        float resul = numA + numB;

        if (numInt == 1) {
            resul = random.nextFloat(resul);
        }

        System.out.printf("%f + %f = %f\n", numA, numB, resul);
        System.out.println(doubleExpression(numA, numB, resul));

    }

    public static boolean doubleExpression(float numA, float numB, float resul) {
        float resulProver = numA + numB;

        if (resul == resulProver) {
            return true;
        } else {
            return false;
        }
    }
}
