// package 08_07;
import java.util.Scanner;

public class GradeCal {
    public static void main(String[] args) {
        int marks;
        System.out.print("Please Enter your marks: ");
        Scanner sc = new Scanner(System.in);
        marks = sc.nextInt();
        System.out.println("You have entered " + marks);
        sc.close();
        if(marks >= 0 && marks <= 100){
            // valid
            System.out.println(marks + " is inside valid range");
            // calculate grade
        }
        else{
            // invalid
            System.out.println(marks + " is outside valid range");
        }
    }
    
}
