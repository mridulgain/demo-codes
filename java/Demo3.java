public class Demo3{
    public static void main(String[] args){
        Circle obj = new Circle(3.1);
        Circle obj2 = new Circle(1);
        System.out.println("Circle1:");
        System.out.println("Area:" + obj.area() + "\nPerimeter:" + obj.perimeter());
        System.out.println("Circle2:");
        System.out.println("Area:" + obj2.area() + "\nPerimeter:" + obj2.perimeter());
    }
}
