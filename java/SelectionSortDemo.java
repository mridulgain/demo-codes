
public class SelectionSortDemo {
    public static void main(String[] args) {
        int [] arr = {28, 6, 4, 2, 24};
        System.out.println("Before sorting: ");
        printArray(arr);
        // selection sort logic
        for(int i = 0; i <= arr.length - 2; i++){
            for(int j = i + 1; j <= arr.length - 1; j++){
                System.out.printf("Comparing : %d < %d ?\n", arr[j], arr[i]);
                if(arr[j] < arr[i]){
                    System.out.printf("\tyes! Swap <-> %d & %d\n", arr[j], arr[i]);
                    int tmp = arr[i];
                    arr[i] = arr[j];
                    arr[j] = tmp;
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

