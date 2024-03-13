public class BoxExample {
    public static void main(String[] args) {
        Box b1 = new Box(1.2, 2.3, 1);
        System.out.println("Box1: " + b1.volume());

        Box b2 = new Box();
        System.out.println("Box2: " + b2.volume());

        Box cube = new Box(2.5);
        System.out.println("Cube: " + cube.volume());
    }

}

class Box {
    // propertiees
    double length, breadth, height;

    // constructors
    public Box() {
        length = breadth = height = 0;
    }

    public Box(double length, double breadth, double height) {
        this.length = length;
        this.breadth = breadth;
        this.height = height;
    }

    public Box(double side) {
        length = breadth = height = side;
    }

    // behaviours
    double volume() {
        double vol = length * breadth * height;
        return vol;
    }
}
