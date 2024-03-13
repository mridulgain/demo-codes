class Example{
    int a;
    public int b;
    private int c;
    static int d;
    static void f(){
        System.out.println("I don't need obj");
    }
    public int getC(){
        return c;
    }
    public void setC(int x){
        this.c = x;
    }
}
public class ExampleMain{
    public static void main(String[] args) {
        //System.out.println(args[0]);
        //System.out.println(args[1].getClass());
        Example obj = new Example();
        obj.a = 10;
        //obj.c = 23;
        obj.b = 9;
        obj.setC(23);
        obj.d = 99;
        System.out.println("obj1: " + obj.a + " " + obj.b + " " + obj.getC() + " " + obj.d);//10 9 99
        Example obj2 = new Example();
        obj2.a = 11;
        obj2.b = 100; 
        obj2.d = 199;
        System.out.println("obj2: " + obj2.a + " " + obj2.b + " " + obj2.d);//11 100 99
        System.out.println("obj1, after creating obj2: " + obj.a + " " + obj.b + " " + obj.d);//10 9 198
        Example.f();
    }
}
