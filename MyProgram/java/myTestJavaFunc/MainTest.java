public class MainTest {
    public static void main(String[] args) {
        String test = "Hello World 123";
        boolean test2 = test.matches("\\w+\\s*");
        System.out.println(test2);
    }
}
