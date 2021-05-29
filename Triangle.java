class Circle{
    double radius;
    public Circle(double radius){
        this.radius = radius;
    }
    double area(){
        return Math.PI * this.radius * this. radius; 
    }
    double perimeter(){
        return 2 * Math.PI * this.radius;
    }
}
