public class LoopExample {
    public static void main(String args[]) {
        // for loop
        for(int i = 1; i < 20; i = i + 2){
            System.out.print(i + " ");
        }
        System.out.println("");
        // while loop
        int j = 20;
        while(j > 1){
            System.out.print(j + " ");
            j = j / 2;
        }
        System.out.println("");
        System.out.println("j = " + j);
    }
}
// 1st 10 odd numbers are: 1 3 5 7 9.. 