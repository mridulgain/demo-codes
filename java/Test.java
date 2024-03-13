import java.util.Scanner;

class Car{
    String name;
    public Car(String n){
        this.name = n;
    }
    public String toString(){
        return "I am " + this.name + "\n";
    }
}
public class Test{
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String name = in.nextLine();
        System.out.println("hi " + name);
        Car obj = new Car("Honda");
        System.out.println(obj);
        hello();
        System.out.println(Math.PI);
    }
    static String hello(){
        return "hi";
    }
}
