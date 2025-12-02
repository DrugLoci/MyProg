import java.util.Arrays;

public final class ImmutableClass {
    private final int[] NUM;

    public ImmutableClass(int[] numArray) {
        this.NUM = Arrays.copyOf(numArray, numArray.length);
    }

    public int[] getNUM() {
        return Arrays.copyOf(NUM, NUM.length);
    }
}
