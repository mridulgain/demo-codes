public class FunctionExample {
    public static void main(String[] args) {
        int[] arr = {10, 20, 30, 40, 50};

        // where is  30 in arr ?
        search(arr, 30);
        // where is 50 ?
        search(arr, 50);
        // where is 60 ?
        search(arr, 60);

        int[] arr2 = {100, 200, 300, 400, 500, 6000};
        search(arr2, 6000);
    }
    
    static void search(int[] arr, int key){
        boolean found = false;
        for(int i = 0; i < arr.length; i++){
            if(arr[i] == key){
                System.out.println(key +" found at index " + i);
                found = true;
                break;
            }
        }
        if(!found){
            System.out.println(key + " not present in arr");
        }
    }
}
