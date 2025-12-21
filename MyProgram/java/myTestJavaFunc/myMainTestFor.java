public class myMainTestFor {

    public static void main(String[] args) {
        long startTime = System.nanoTime();
//        int i = 0;

//        while (i++ < 100_000_000) {
//            System.out.printf("%d\n", i); //  4 cors intel i 5 VM  370138 ms
//        }

//        do {
//            System.out.printf("%d\n", i++);
//        } while (i < 100_000_000);  //  4 cors intel i 5 VM  348000 ms | 114472

//        for (int i = 0; i < 100_000_000; i++) {
//            System.out.printf("%d\n", i); // 4 cors intel i 5 VM 441334 ms | 116540
//        }

        for (int i = 0; i < 100_000_000; ++i) {
            System.out.printf("%d\n", i); // 4 cors intel i 5 VM 647827 ms | 114688
        }


        long endTime = System.nanoTime();
        long duration = (endTime - startTime) / 1_000_000;
        System.out.println(duration);
    }
}