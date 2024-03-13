import java.util.Scanner;
public class ArrayExample{
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        // # array declaration & initialisation
        int a[] = new int[10];
        int[] b = new int[10];
        // initialisation
        int[] c = {1,2,3,4};
        // accessing
        for(int i = 0; i < c.length; i++){
            System.out.println(i + " " + c[i]);
        }
        //input multiple values
        for(int i = 0; i < 10; i++){
            System.out.print("Enter a number: ");
            a[i] = sc.nextInt();
        }
        // output
        for(int i = 0; i < 10; i++){
            System.out.print(a[i] + " ");
        }

    }
}
// ha