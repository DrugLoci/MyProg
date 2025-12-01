import java.util.Arrays;

public class MainTest1 {

    public static void main(String[] args) {

        String names[] = {"Anton", "Kiril", "Oleg"};
        ImmutableClass immutableClass = new ImmutableClass(names);

        System.out.println(Arrays.toString(immutableClass.getNameUsers()));

        names[0] = "Nikita";
        System.out.println(Arrays.toString(immutableClass.getNameUsers()));

        String names1[];
        names1 = immutableClass.getNameUsers();   // ошибка допущена в иммутабельном классе

        names1[0] = "Nikita";
        System.out.println(Arrays.toString(immutableClass.getNameUsers()));
    }

}
