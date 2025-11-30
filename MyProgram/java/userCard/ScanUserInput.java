import java.util.Scanner;

public class ScanUserInput {

    void printHint(){
        System.out.println("Поля со знаком * обязательны для ввода.");
    }

    public String userInputName() {

        Scanner scanner = new Scanner(System.in);

        String name;
        while (true) {
            printHint();
            System.out.print("*Имя: ");
            name = scanner.nextLine().trim();

            if (!name.isEmpty()) {
                break;
            }
        }
        return name;
    }

    public String userInputSurName() {

        Scanner scanner = new Scanner(System.in);

        String surName;
        while (true) {
            printHint();
            System.out.print("*Фамилия: ");
            surName = scanner.nextLine().trim();

            if (!surName.isEmpty()) {
                break;
            }
        }
        return surName;
    }

    public String userСityBirths() {

        Scanner scanner = new Scanner(System.in);

        String cityBirths;
        while (true) {
            printHint();
            System.out.print("*Город рождения: ");
            cityBirths = scanner.nextLine().trim();

            if (!cityBirths.isEmpty()) {
                break;
            }
        }
        return cityBirths;
    }

    public String userInputYearBirths() {

        Scanner scanner = new Scanner(System.in);

        String yearBirths;
        while (true) {
            printHint();
            System.out.print("*Год рождения: ");
            yearBirths = scanner.nextLine().trim();

            if (!yearBirths.isEmpty()) {
                break;
            }
        }
        return yearBirths;
    }

    public String userInputHom() {
        System.out.print("Место жительства: ");
        Scanner scanner = new Scanner(System.in);
        return scanner.nextLine();
    }

    public String userInputAdg() {
        System.out.print("Возраст: ");
        Scanner scanner = new Scanner(System.in);
        return scanner.nextLine();
    }

    public String userInputCar() {
        System.out.print("Марка автомобиля: ");
        Scanner scanner = new Scanner(System.in);
        return scanner.nextLine();
    }
}
