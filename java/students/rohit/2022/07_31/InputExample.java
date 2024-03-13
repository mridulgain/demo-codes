// package 07_31;
import java.util.Scanner;

public class InputExample {
    public static void main(String[] args) {
        int userInput;
        System.out.print("Please Enter your marks: ");
        Scanner sc = new Scanner(System.in);
        userInput = sc.nextInt();
        System.out.println("You have entered " + userInput);
        sc.close();
    }
}
