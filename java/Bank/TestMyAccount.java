import java.util.Scanner;

public class TestMyAccount{
    public static void main(String[] args){
        Scanner input = new Scanner(System.in);
        MyAccount acc = null;
        boolean flag = true;
        while(flag){
            System.out.println("-------------------------------");
            System.out.println("Enter 1 to create an account");
            System.out.println("Enter 0 to exit");
            System.out.print("Input: ");
            int op = input.nextInt(); 
            System.out.println("-------------------------------");
            switch(op){
                case 1:
                    System.out.print("Name? ");
                    String name = input.next();
                    acc = new MyAccount(name);
                    System.out.println("Account has been created successfully");
                    System.out.println("Account Number: " + acc.accNo);
                    System.out.println("Account Holders name: " + acc.accHoldersName);
                    break;
                case 0:
                    flag = false;
                    break;
                default:
                    System.out.println("Invalid input, try again");
            }
        }
    }
}
