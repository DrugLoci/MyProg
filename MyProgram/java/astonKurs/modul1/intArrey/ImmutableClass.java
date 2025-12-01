import java.util.Arrays;

public final class ImmutableClass {
    private final int[] NUM = new int[4];

    public ImmutableClass(int[] numArray) {
        this.NUM = Arrays.copyOf(numArray, numArray.length);
    }

    public int[] getNUM(){
        return Arrays.copyOf(NUM, NUM.length);
    }
}
