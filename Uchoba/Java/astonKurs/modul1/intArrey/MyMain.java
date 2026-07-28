public class MyMain {

    public void main(String[] args) {
        int[] numArray = {1, 2, 3, 4, 5,};
        ImmutableClass immutableClass = new ImmutableClass(numArray);
        int[] num = new int[4];
        num = immutableClass.getNUM();

        for (int i = 0; i < num.length; i++) {
            System.out.println(num[i]);
        }

    }
}
