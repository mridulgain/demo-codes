public class StringDemo {
    public static void main(String args[]) {
        String str = "Hello";
        String str2 = new String("World");
        System.out.println(str.length());
        String msg = str.concat(str2);
        System.out.println(str);
        System.out.println(str2);
        System.out.println(msg);
        String str3 = new String("Hello");
        System.out.println(str + " == " + str3 + " ? " + (str == str3));
        System.out.println(str + " == " + str3 + " ? " + (str.equals(str3)));
    }
}
