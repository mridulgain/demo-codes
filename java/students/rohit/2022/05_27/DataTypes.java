public class DataTypes {
    public static void main(String args[]) {
        char a = 'A', z = 'Z';
        System.out.printf("A = %d, Z = %d \n", (int) a, (int) z);

        int[] arr = { 1, 2, 3, 4 };
        char[] arr2 = { 'a', 'A', 'b' };

        MyClass obj = new MyClass();
        obj.p = 12;
        System.out.printf("%d %d \n", obj.p, obj.q);

    }
}

class MyClass {
    int p, q;
}