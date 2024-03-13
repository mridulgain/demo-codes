class MyAccount{
    private static int counter;
    String accNo, accHoldersName; 
    MyAccount(String name){
        counter++; //++counter;
        accNo = "UBI08" + counter;
        accHoldersName = name;
    }



}
