public class Rohitashu {
    public static void main(String[] args) {
        Fruit f1 = new Fruit("Orange", "spherical", "orange", "sweet");
        f1.print();
        Fruit f2 = new Fruit("Mango", "oval", "yellow", "sweet");
        f2.print();
    }
}

class Fruit {
    String name, shape, colour, taste;

    // Constructor
    public Fruit(String name, String shape, String colour, String taste) {
        this.shape = shape;
        this.taste = taste;
        this.colour = colour;
        this.name = name;
    }

    void print() {
        String tmp = name + " is " + shape + ",its colour is " + colour + " and it tastes somewhat " + taste;
        System.out.println(tmp);
    }
}