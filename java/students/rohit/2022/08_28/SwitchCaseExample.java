// package 08_28;
import java.util.*;

public class SwitchCaseExample {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        String today = input.next();
        System.out.println("Today's menue is: ");
        switch(today){
            case "Sunday":
                //sunday's menue
            break;
            case "Saturday":
                //sat's menue
            break;
            case "Friday":
                //
            default:
                System.out.println("Wrong input");
        }
        input.close();
    }
}
