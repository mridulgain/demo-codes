public class ElectricBill{
    public static void main(String[] args) {
        int totalUnit = 251;
        int unit = totalUnit;
        Double bill = 0.0;
        if(unit > 250){
            bill += (unit - 250) * 1.50;
            unit = 250;
        }
        if(unit > 150){
            bill += (unit - 150) * 1.20;
            unit = 150;
        }
        if(unit > 50){
            bill += (unit - 50) * 0.75;
            unit = 50;
        }
        if(unit >=0)
            bill += unit * 0.5;
        else{
            System.out.println("Invalid input");
        }
        System.out.println("Your bill for " + totalUnit + " units is Rs. " + bill * 1.2);
    }
}
