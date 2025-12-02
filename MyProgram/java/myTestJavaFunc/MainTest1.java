import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;

public class MainTest1 {

    static void main(String[] args) {
        List<String> name = new ArrayList<>();
        Scanner scanner = new Scanner(System.in);

        for (int i = 0; i < 3; i++) {
            String inputUser = scanner.nextLine().trim();
            name.add(inputUser);
        }

        ImmutableClass immutableClass = new ImmutableClass(name);

        System.out.println(immutableClass.getNAME_GROUP());


    }

}
