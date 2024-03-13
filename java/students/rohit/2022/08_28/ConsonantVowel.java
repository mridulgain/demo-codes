// package 08_28;
import java.util.Scanner;

public class ConsonantVowel {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        System.out.print("Enter a charecter: ");
        char c = input.next().charAt(0);
        if(c >= 'a' && c <= 'z'){
            if(c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u')
                System.out.println("Vowel");
            else
                System.out.println("Consonant");
        }       
        else {
            System.out.println("Not an alphabet");
        }
        input.close();
    }
}
