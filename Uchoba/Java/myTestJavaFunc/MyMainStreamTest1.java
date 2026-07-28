//import java.io.*;
//import java.util.*;
//
//public class MyMainStreamTest1 {
//
//    public static void main(String[] args) {
//        File myFile = new File("/home/vboxuser/testFile.txt");
//        List<String> list = new ArrayList<>();
//
//        try (FileInputStream fileInputStream = new FileInputStream(myFile)) {
//            list = new ArrayList<>(List.of(new String(fileInputStream.readAllBytes())));
//            System.out.println(list);
//        } catch (IOException e) {
//            throw new RuntimeException(e);
//        }
//
//        list<String> listResult = list.stream();
//
//    }
//}
