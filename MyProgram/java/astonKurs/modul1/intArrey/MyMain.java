public class MyMain {

    public void main(String[] args){
        int[] numArray = {1,2,3,4,5,};
        ImmutableClass immutableClass = new ImmutableClass(numArray);
//        int[] num = immutableClass.getNUM();

        for (int i = 0; 5 < immutableClass.getNUM().length; i++) {
            System.out.println(immutableClass.getNUM(i));
        }
    }

}
