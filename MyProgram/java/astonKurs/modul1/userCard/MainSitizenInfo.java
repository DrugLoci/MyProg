import java.util.Scanner;

public class MainSitizenInfo {

    public static void main(String[] args) {

        ScanUserInput scanUserInput = new ScanUserInput();

        // запрос от пользователя ввода своих данных
        String userName = scanUserInput.userInputName();
        String userSurName = scanUserInput.userInputSurName();
        String userCityBirths = scanUserInput.userСityBirths();
        String userYearBirths = scanUserInput.userInputYearBirths();
        String userHom = scanUserInput.userInputHom();
        String userAdg = scanUserInput.userInputAdg();
        String userCar = scanUserInput.userInputCar();
        

        SitizenInfo sitizenInfo = new SitizenInfo(userName, userSurName, userCityBirths, userYearBirths);
        SitizenInfo.InnerSitizenInfo innerSitizenInfo = sitizenInfo.new InnerSitizenInfo(userHom, userAdg, userCar);

        PrintCardUser printCardUser = new PrintCardUser(userName, userSurName, userCityBirths, userYearBirths, userHom, userAdg, userCar);

        printCardUser.printCard();
    }
}