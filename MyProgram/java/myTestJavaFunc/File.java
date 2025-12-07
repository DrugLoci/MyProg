import java.io.IOException;

public class File {

        public static void main(String[] args) {
            java.io.File file = new java.io.File("/home/loci/Students.txt");

            try {
                file.createNewFile();
            } catch (IOException e) {
                throw new RuntimeException(e);
            }
        }
}

