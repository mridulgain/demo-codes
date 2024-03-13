public class List2 {
    public static void main(String[] args){
        int arr[] = {10, 0, 1, -10, 50, 20};
        double max = Double.NEGATIVE_INFINITY, min = Double.POSITIVE_INFINITY;
        int max_pair[] = {0, 0}, min_pair[] = {0, 0}, diff;
        for(int i = 0; i < arr.length; i++){
            for(int j = i + 1; j < arr.length; j++){
                // System.out.println(arr[i] + "~" + arr[j]);
                diff = Math.abs(arr[i] - arr[j]);
                if(diff > max){
                    max = diff;
                    max_pair[0] = arr[i];
                    max_pair[1] = arr[j];
                }
                if(diff < min){
                    min = diff;
                    min_pair[0] = arr[i];
                    min_pair[1] = arr[j];
                }
            }
        }
        System.out.println("Max pair:" + max_pair[0] + ", " + max_pair[1]);
        System.out.println("Min pair:" + min_pair[0] + ", " + min_pair[1]);
    }
}
