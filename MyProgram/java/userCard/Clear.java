public class Clear {

    public void clearConsole() {

        String os = System.getProperty("os.name");

        if (os.contains("Linux")) {
            ProcessBuilder processBuilder = new ProcessBuilder("clear");
        } else {
            ProcessBuilder processBuilder = new ProcessBuilder("cmd", "/c", "cls");
        }
    }
}