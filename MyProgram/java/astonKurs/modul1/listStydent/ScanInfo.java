import java.util.ArrayList;
import java.util.Scanner;

public class ScanInfo {

    public static int quantityGroup() { // запрос на количесво групп 
        int numGroup;
        Scanner scanner = new Scanner(System.in);

            while (true) {
                System.out.print("Количесво групп: ");
                numGroup = scanner.nextInt();

                if (numGroup != 0) {
                    break;
                }
            }

    }

    public String[] nameGroup() { // запрос на ввод имени группы
        Scanner scanner = new Scanner(System.in);
        ArrayList<String> nameGroupList = new ArrayList<>();

        for (int i = 0; numGroup > i; i++) {
            System.out.print("Введите имя группы №" + (i + 1) + ": ");
            nameGroupList.add(scanner.nextLine().trim());
        }

        String[] nameGroupArrey = nameGroupList.toArray(new String[nameGroupList.size()]);
        return nameGroupArrey;
    }

    public int[] quantityStudentInGroup() {
        Scanner scanner = new Scanner(System.in);
        int[] numStudent;

        for (int i = 0; numGroup > i; i++) {
            System.out.print("Введите количесво сутдентов в гурппе " + nameGroup(n));
        }
    }
}
