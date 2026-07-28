import java.util.*;

public final class ImmutableClass{
    private final String[] NAME;

    public ImmutableClass(String[] name) {
        this.NAME = Arrays.copyOf(name, name.length);
    }
}