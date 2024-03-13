class Account{
    static int id;
    String name;
    double balance;
    public int a;
    int b;
    private int c;
    Account(String name){
        this.name = name;
    }
    void credit(double amt){
        balance += amt;
        System.out.println("Your account has been credited with "+ amt);
        showBalance();
    }
    void debit(double amt){
        balance -= amt;
        System.out.println("Your account has been debited with "+ amt);
        showBalance();
    }
    void showBalance(){
        System.err.println("Current balance is "+ balance);
    }
}

public class TestAccount{
    public static void main(String[] args) {
        Account a = new Account("Sayan");
        a.credit(100);
        a.debit(200);
        a.c = 12;
        System.out.println(a.c);
    }
}