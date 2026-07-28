public class TestClass extends AbstractClass implements Interface{

    public TestClass(int testInt, String testString) {
        super(testInt, testString);
    }

    @Override
    public void methodAbstractClass(int num) {

    }

    @Override
    public void test1() {
        System.out.println(testString);
    }

    @Override
    public void test2() {
        System.out.println(testInt);
    }
}
