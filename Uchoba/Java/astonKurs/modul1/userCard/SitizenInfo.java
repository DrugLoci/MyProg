// класс будети принимать данные о пользователе 

public final class SitizenInfo {

    private final String NAME;
    private final String SURNAME;
    private final String СITY_BIRTHS;
    private final String YEAR_BIRTHS;
    //private int REAL_YEAR = 2025;

    public SitizenInfo(String name, String surname, String сityBirths, String yearBirths) {
        this.NAME = name;
        this.SURNAME = surname;
        this.СITY_BIRTHS = сityBirths;
        this.YEAR_BIRTHS = yearBirths;
    }

    public class InnerSitizenInfo {
        private String hom;
        private String adg;
        private String car;

        public InnerSitizenInfo(String hom, String adg, String car) {
            this.hom = hom;
            this.adg = adg;
            this.hom = car;
        }


        public String getName() {
            return NAME;
        }

        public String getSurname() {
            return SURNAME;
        }

        public String getСityBirths() {
            return СITY_BIRTHS;
        }

        public String getYearBirths() {
            return YEAR_BIRTHS;
        }

        public void setAdg(String adg) {
            this.adg = adg;
        }

        public String getAdg() {
            return String.valueOf(adg);
        }

        public void setCar(String car) {
            this.car = car;
        }

        public String getCar() {
            return car;
        }

        public void setHom(String hom) {
            this.hom = hom;
        }

        public String getHom() {
            return hom;
        }
    }
}
