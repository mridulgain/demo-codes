public class SmallCaseToUpperCase{
    public static void main(String[] args) {
        String str = "APPLE";
        // System.out.println(str.toUpperCase());
        // char diff = 'a' - 'A';
        String ans = "";

        for(int i = 0; i < str.length(); i++){
            // do something with str.charAt(i)
            ans += (char)(str.charAt(i) - 'A' + 'a');
        }
        System.out.println(ans);
    }
}