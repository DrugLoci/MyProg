import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;

class Book {    // данный класс будет выполнять роль кеша
    String name;
    String pages;
    String yer;

    public void setName(String name) { this.name = name; }

    public void setPages(String pages) { this.pages = pages; }

    public void setYer(String yer) { this.yer = yer; }

    public String getName() { return name; }

    public String getPages() {return pages;}

    public String getYer() {return yer;}
}

class Student { // данный класс будет выполнять роль кеша
    String name;

    public void setName(String name) { this.name = name; }

    public String getName() { return name; }
}

public class MyMainList_2 {
    public static void main(String[] args) {
        String nameFail = "./StudentAndBook.txt";
        List<String> list = null;

        try {
            list = Files.readAllLines(Paths.get(nameFail));
        } catch (Exception e) {
            System.out.println("Ошибка чтения файла");
        }

//        list.stream() // Вывести в консоль каждого студента (переопределите toString)
//                .limit(3)
//                .forEach(System.out::println);

//        list.stream() // Получить для каждого студента список книг
//                .limit(1)
//                .forEach(System.out::println);
//        list.stream()
//                .skip(3)
//                .forEach(System.out::println);
//        list.stream()
//                .limit(2)
//                .skip(1)
//                .forEach(System.out::println);
//        list.stream()
//                .skip(3)
//                .forEach(System.out::println);
//        list.stream()
//                .limit(3)
//                .skip(2)
//                .forEach(System.out::println);
//        list.stream()
//                .skip(3)
//                .forEach(System.out::println);

//        list.stream() // Получить книги
//                .skip(3)
//                .forEach(System.out::println);

        list.stream() // Отсортировать книги по количеству страниц (Не забывайте про условия для сравнения объектов
                .skip(3)
                .sorted()
                .forEach(System.out::println);
    }
}

