package MyProg;
import java.util.Random;

class VesakosniyGod { // нужно рандомно выбрать год и опеределить сколько весакосных лет было до этого года 
    public static void main(String[] args) {
        Random random = new Random(); 
        int randNum = random.nextInt(2025) + 1;
        int numVesakoGod = vesakosniyGod(randNum);
        System.out.println(randNum);
        System.out.println(numVesakoGod);
    }

    public static int vesakosniyGod(int num) {
        return num / 4 - num / 100 + num / 400;
    }
}
