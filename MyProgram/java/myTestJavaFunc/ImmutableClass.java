import java.util.Arrays;

public final  class ImmutableClass {

    private final String[] nameUsers;

//    public ImmutableClass(String[] nameUsers) {
//        this.nameUsers = nameUsers; // опасно
//    }

    public  ImmutableClass(String[] nameUsers) {
        this.nameUsers = Arrays.copyOf(nameUsers , nameUsers.length);
    }

//    private final String[] nameUsers = {"A", "B", "C"};

//    public ImmutableClass(String[] nameUsers) {
//        this.nameUsers = nameUsers;
//    }

    public String[] getNameUsers() {
        return nameUsers; // ошбика, из за которой внешний мир получил адрес массива в куще на который ссылалось иммутабельное поле 
    }
}
