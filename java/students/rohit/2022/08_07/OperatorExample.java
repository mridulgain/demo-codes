// package 08_07;

public class OperatorExample {
    public static void main(String[] args) {
        int a = 1 + 10; // arthmetic operator: +, -, *, /
        System.out.println(a);
        int m = 10, n = 100;
        // relation operators: <, >, ==, <=, >=, !=
        boolean c = m > n; // false
        System.out.println(c);
        // m < n true
        // m == n false
        // m != n true
        // 100 <= 100 true
        // logical operators: &&, ||, !
        // marks >= 0 && marks <= 100
        if(m > n){
            System.out.println("I get printed if condition is true");
        }
        else{
            System.out.println("I get printed if condition is false");
        }
        // bit wise operators: & | ~
        System.out.println(3 & 4);
        System.out.println(3 | 4);
        //increment decrement: ++, --
        m++; // m = m + 1
        //shorthand: +=, -=, *=, /= %=
        int b = 10;
        b += 20; // b = b + 20
        System.out.println(b);

        //ternary
        System.out.printf("Highest : %d\n", m > n ? m : n);
    }
}
