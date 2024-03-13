public class ArrayExample{
    public static void main(String[] args) {
        int[] arr = {1,2,3,4};
        // accessing array element
        System.out.println(arr[1]);
        int sum = arr[0] + arr[1] + arr[2] + arr[3];
        System.out.println(sum);
        // 
        int[] arr2 = new int[5];
        arr2[0] = 10;
        arr2[1] = 20;
        arr2[4] = 30;
        // 
        System.out.println("Contents of arr2: ");
        for(int i: arr2){
            System.out.println(i);
        }
        //
    }
}