import java.io.*;
import java.util.*;

public class MyMainList {

    static void main(String[] args) {
        File myFile = new File("/home/vboxuser/testFile.txt");
        List<String> list = new ArrayList<>();

//        try (FileOutputStream fileOutputStream = new FileOutputStream(myFile)) {
//            list.add("Name: Anton\n" +
//                    "Name: Mike\n" +
//                    "Name: Kate\n" +
//                    "Book Java pages 300 yer 2015\n" +
//                    "Book C++ pages 200 yer 2016\n" +
//                    "Book C# pages 100 yer 2017\n" +
//                    "Book Python pages 500 yer 2018\n" +
//                    "Book C pages 400 yer 1998\n");
//            fileOutputStream.write(list.get(0).getBytes());
//        } catch (IOException e) {
//            throw new RuntimeException(e);
//        }

        try (Scanner scanner = new Scanner(new FileInputStream(myFile))) {
            while (scanner.hasNextLine()) {
                list.add(Arrays.toString(scanner.nextLine().split("\n")));
            }
//            System.out.println(list);
        } catch (IOException e) {
            System.out.println("Error: " + e);
        }

//        list.stream() // Вывести в консоль каждого студента (переопределите toString)
//                .filter(s -> s.contains("Name"))
//                .map(s -> s.replaceAll("Name: ", ""))
//                .map(s -> s.replaceAll("]", ""))
//                .map(s -> s.substring(1))
//                .forEach(System.out::println);

//        for (int i = 0; i < 3; i++) { // Получить для каждого студента список книг
//
//            switch (i) {
//                case 0 -> {
//                    list.stream()
//                            .filter(s -> s.contains("Anton"))
//                            .map(s -> s.replaceAll("Name: ", ""))
//                            .map(s -> s.replaceAll("]", ""))
//                            .map(s -> s.substring(1))
//                            .forEach(System.out::println);
//                    list.stream()
//                            .filter(s -> s.contains("Java") || s.contains("C++") || s.contains("C#"))
//                            .map(s -> s.split("pages")[0])
//                            .map(s -> s.substring(1))
//                            .forEach(System.out::println);
//                    System.out.println();
//                    break;
//                }
//                case 1 -> {
//                    list.stream()
//                            .filter(s -> s.contains("Mike"))
//                            .map(s -> s.replaceAll("Name: ", ""))
//                            .map(s -> s.replaceAll("]", ""))
//                            .map(s -> s.substring(1))
//                            .forEach(System.out::println);
//                    list.stream()
//                            .filter(s -> s.contains("Java") || s.contains("Python") || s.contains("C#"))
//                            .map(s -> s.split("pages")[0])
//                            .map(s -> s.substring(1))
//                            .forEach(System.out::println);
//                    System.out.println();
//                    break;
//                }
//                case 2 -> {
//                    list.stream()
//                            .filter(s -> s.contains("Kate"))
//                            .map(s -> s.replaceAll("Name: ", ""))
//                            .map(s -> s.replaceAll("]", ""))
//                            .map(s -> s.substring(1))
//                            .forEach(System.out::println);
//                    list.stream()
//                            .filter(s -> s.contains("C") || s.contains("Python") || s.contains("C++"))
//                            .map(s -> s.split("pages")[0])
//                            .map(s -> s.substring(1))
//                            .forEach(System.out::println);
//                    System.out.println();
//                    break;
//                }
//                default -> System.out.println("Error");
//            }
//        }

//        list.stream() // Получить книги
//                .filter(s -> s.contains("Book "))
//                .map(s -> s.split("pages")[0])
//                .map(s -> s.substring(1))
//                .forEach(System.out::println);

//        list.stream() // Отсортировать книги по количеству страниц (Не забывайте про условия для сравнения объектов)
//                .filter(s -> s.contains("Book "))
//                .sorted(Comparator.comparingInt(s -> Integer.parseInt(s.split("pages")[1].split(" ")[1])))  // ?????????????????????????????
//                .map(s -> s.split("]")[0])
//                .map(s -> s.substring(1))
//                .forEach(System.out::println);

//        list.stream() // Получить список книг, которые написал каждый студент
    }
}



