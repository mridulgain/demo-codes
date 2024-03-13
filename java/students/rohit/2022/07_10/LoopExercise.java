public class LoopExercise {
    public static void main(String[] args) {
        for(int i=1; i <= 40; i++){ // 1. Find all 4 digit number 
            int result = i % 7;
            if(result == 0){ // 2. check if div by 7
                System.out.println(i + " is divisible by 7");
            }
            else{
                System.out.println(i + " is not divisible by 7");
            }
        }
    }
}
//Find sum of all 4 digit numbers which are divisible by 7.
