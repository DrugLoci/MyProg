public final class PrintCardUser {

    private final String NAME;
    private final String SURNAME;
    private final String CITY_BIRTHS;
    private final String YEAR_BIRTHS;
    private final String ADG;
    private final String CAR;
    private final String HOM;

    public PrintCardUser(String name, String surname, String cityBirths, String yearBirths, String hom, String adg, String car) {
        this.NAME = name;
        this.SURNAME = surname;
        this.CITY_BIRTHS = cityBirths;
        this.YEAR_BIRTHS = yearBirths;
        this.ADG = adg;
        this.CAR = car;
        this.HOM = hom;
    }

    public void printCard() {
        System.out.println( "\n\n\nКарточка пользователя: \n" +
                            "Имя: " + NAME + "\n" +
                            "Фамилия: " + SURNAME + "\n" +
                            "Год рождения: " + YEAR_BIRTHS + "\n" +
                            "Город рождения: " + CITY_BIRTHS + "\n" +
                            "Место жительсва: " + HOM + "\n" +
                            "Возраст: " + ADG + "\n" +
                            "Автомобиль: " + CAR);
    }
}