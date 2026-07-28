import java.util.List;

public final class ImmutableClass {

    private final List<String> NAME_GROUP;

    public ImmutableClass(List<String> nameGroup) {
        this.NAME_GROUP = List.copyOf(nameGroup);
    }

    public List<String> getNAME_GROUP() {
        return List.copyOf(NAME_GROUP);
    }
}