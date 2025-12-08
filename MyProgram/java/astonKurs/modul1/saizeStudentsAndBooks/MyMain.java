import java.io.*;
import java.util.*;

public class MyMain {
    public static void main(String[] args) {
        File file = new File("/home/loci/Students.txt");
        List<String> list = new ArrayList<>();

//        try (FileOutputStream outputStream = new FileOutputStream(file)) {
//            outputStream.write(("Name: Loci\nName: Anton\nName: Nikita\n" +
//                    "Book: Java page: 2000 yer: 2019\n" +
//                    "Book: C++ page: 1000 yer: 2009\n" +
//                    "Book: Python page: 300, yer: 2018\n" +
//                    "Book: C# page: 1200 yer: 2017\n" +
//                    "Book: JavaScript page: 505 yer: 2016\n" +
//                    "Book: PHP page: 358 yer: 1999\n").getBytes());
//        } catch (IOException e) {
//            System.out.println("Error: " + e);
//        }



//        list.stream() // Вывести в консоль каждого студента (переопределите toString)
//                .filter(s -> s.contains("Name: "))
//                .map(s -> s.split("Name: ")[1].split("]")[0])
//                .forEach(System.out::println);


//        list.stream()
//                .filter(s -> s.contains("Name: ") || s.contains("Book: "))
//                .map(s -> s.split("Name: ")[1].split("]")[0].split("Book: ")[1])
//                .forEach(System.out::println);
    }
}
