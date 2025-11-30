public abstract class AbstractClass {

    public  int testInt;

    public String testString;

    //public abstract int intTest;

    public AbstractClass(int testInt, String testString) {
        this.testInt = testInt;
        this.testString = testString;
    }

    public abstract void methodAbstractClass(int num);
}
