import java.io.*;
import java.util.*;

public class MyMainList1_1 {
    public static void main(String[] args) {
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
//                .limit(3)
//                .forEach(System.out::println);

//        list.stream() // Получить для каждого студента список книг
//                .limit(1)
//                .forEach(System.out::println);
//        list.stream()
//                .skip(3)
//                .forEach(System.out::println);
//
//        list.stream()
//                .limit(2)
//                .skip(1)
//                .forEach(System.out::println);
//        list.stream()
//                .skip(3)
//                .forEach(System.out::println);
//
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

//        list.stream() // Отсортировать книги по количеству страниц (Не забывайте про условия для сравнения объектов)
//                .skip(3)
////                .sorted(Comparator.comparingInt(s -> Integer.parseInt(s.split("pages ")[1].split(" yer")[0]))) // я понимаю как остортировать только так (стрый варинат)
//                .sorted(Comparator.comparing(s -> s.split("pages ")[1].split(" yer")[0]))  // вот более правильный вариант
////                .sorted()  // нет, я мог бы конечно переделать файл так что бы pages шел первым в строке и тогда бы так сработало но мне почему то эта идея не поравилась
//                .forEach(System.out::println);  // по этому прошу сильно помидорами в меня не бросаться :)

//        list.stream() // Оставить только уникальные книги
//                .skip(3)
//                .distinct()
//                .forEach(System.out::println);

//        list.stream() // Отфильтровать книги, оставив только те, которые были выпущены после 2000 года
//                .skip(3)
//                .limit(4)
//                .forEach(System.out::println);

//        list.stream() // Ограничить стрим на 3 элементах
//                .limit(3)
//                .forEach(System.out::println);

//        list.stream() // Получить из книг годы выпуска
//                .skip(3)
//                .forEach(s -> System.out.println(s.split(" yer ")[1].split("]")[0]));

//        System.out.println(  // При помощи методов получения значения из Optional вывести в консоль год выпуска найденной книги, либо запись о том, что такая книга отсутствует
//                list.stream() // При помощи методов короткого замыкания (почитайте самостоятельно что это такое) вернуть Optional от года
//                .filter(s -> s.contains("Book "))
//                .map(s -> Integer.parseInt(s.split(" yer ")[1].split("]")[0]))
//                .filter(s -> s > 2000)
//                .findFirst()
//        );

    }
}
