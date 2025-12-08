import java.io.*;
import java.util.*;

public class Main {

    static void main(String[] args) {
        File myFile = new File("/home/vboxuser/testFile.txt");
        List<String> list = new ArrayList<>();

        try (FileOutputStream fileOutputStream = new FileOutputStream(myFile, true)) {
            String strStudents = "Name: Anton\nName: Mike\nName: Kate\n";
            String strBook = "Book Java, pages 300, yer 2015\n" +
                            "Book C++, pages 200, yer 2016\n" +
                            "Book C#, pages 100, yer 2017\n" +
                            "Book Python, pages 500, yer 2018\n" +
                            "Book C, pages 400, yer 1998\n";
            fileOutputStream.write(strStudents.getBytes());
            fileOutputStream.write(strBook.getBytes());
        } catch (IOException e) {
            throw new RuntimeException(e);
        }


        try (FileInputStream fileInputStream = new FileInputStream(myFile)) {
            list = new ArrayList<>(List.of(new String(fileInputStream.readAllBytes()).split("\n")));
            System.out.println(list);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }


    }
}
