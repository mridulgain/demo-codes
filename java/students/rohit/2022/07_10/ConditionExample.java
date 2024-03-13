public class ConditionExample {
    public static void main(String[] args) {
        int input = 1000, result;
        result = input % 7;
        if(result == 0){
            System.out.println(input + " is divisible by 7");
        }
        else{
            System.out.println(input + " is not divisible by 7");
        }      
    }
}
