import java.util.List;
import java.util.Scanner;

public class ScanClass {

    public int numGroup() {
        System.out.print("Введите количесво групп: ");
        Scanner scanner = new Scanner(System.in);
        TempClass tempClass = new TempClass();
        tempClass.numGroup = scanner.nextInt();
        return tempClass.numGroup;
    }

//    public List<String> nameGroup() {
//
//    }
}
