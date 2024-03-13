class Point
{
	int x , y;
	public Point(int x,int y) 
	{
		this.x=x;
		this.y=y;
	}
	public double getDistance(Point p)
	{
		return Math.sqrt(Math.pow(x - p.x, 2) + Math.pow(y - p.y, 2));
	}
	public String toString()
	{
		// String t= "Distance between "+ "("+x+","+y+")"+" and "+"("+a+","+b+") is ";
		return "(" + x + ", " + y + ")";
	}
}
public class TestPoint
{
	public static void main(String args[])
	{
		Point p = new Point(1,1);
		Point q = new Point(3,4);
		System.out.println("distance between "+ p + " & " + q +" is " + p.getDistance(q));
		//main();
	}
}
