public class Wrappers{
    public static void main(String[] args){
        int a = 5;
        Integer b = 6, c = 9;
        // Integer v = new Integer("8");
        String num = "123";
        System.out.println(Integer.MAX_VALUE);
        System.out.println(Integer.bitCount(a));
        System.out.println(num + c);
        System.out.println(Integer.parseInt(num) + b);
        char d = 'a';
        System.out.println(Character.isDigit(d));
        
    }
}