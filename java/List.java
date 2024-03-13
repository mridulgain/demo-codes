class List{
    public static void main(String[] args){
        int arr[] = {99, -10, 0, -99, 50};
        // int max = arr[0], min = arr[0];
        double max = Double.NEGATIVE_INFINITY, min = Double.POSITIVE_INFINITY;
        for(int i = 0; i < arr.length; i++){
            if(arr[i] > max)
                max = arr[i];
            if(arr[i] < min)
                min = arr[i];
        }
        // for(int i: arr){
        //     System.out.print(i + " ");
        // }
        System.out.println("Max :" + max +"\nMin :" + min);
    }
}