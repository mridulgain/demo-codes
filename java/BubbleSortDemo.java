
public class BubbleSortDemo {
    public static void main(String[] args) {
        int [] arr = {28, 6, 4, 2, 24};
        System.out.println("Before sorting: ");
        printArray(arr);
        // bubble sort logic
        for(int i = 0; i < arr.length - 1; i++){
            for(int j = 0; j < (arr.length - 1) - i; j++){
                System.out.printf("Comparing : %d < %d ?\n", arr[j], arr[i]);
                if(arr[j] > arr[j+1]){
                    System.out.printf("\tyes! Swap <-> %d & %d\n", arr[j], arr[i]);
                    int tmp = arr[j+1];
                    arr[j+1] = arr[j];
                    arr[j] = tmp; 

                    //
                    System.out.print("\tResultant array: ");
                    printArray(arr);
                } else{
                    System.out.println("\tNo!");
                }
            }
            System.out.println();
        }

        System.out.println("After sorting: ");
        printArray(arr);
    } 

    public static void printArray(int[] a){
        for(int i = 0; i < a.length; i++){
            System.out.printf("%d ", a[i]);
        }
        System.out.println();
    } 
}

