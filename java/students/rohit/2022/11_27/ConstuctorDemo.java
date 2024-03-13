public class ConstuctorDemo {
    public static void main(String[] args) {
        Counter c = new Counter(100);
        System.out.println("Initial value: " + c.getCount());
        c.increment();
        c.increment();
        System.out.println("After increment: " + c.getCount());
        //
        Counter test_counter = new Counter();
        System.out.println("Initial value: " + test_counter.getCount());
        test_counter.increment();
        test_counter.increment();
        System.out.println("After increment: " + test_counter.getCount());
        //
        Counter c2 = new Counter("X");
        System.out.println("Initial value: " + c2.getCount());
        c2.increment();
        c2.increment();
        System.out.println("After increment: " + c2.getCount());
    }
}

class Counter {
    // properties
    private int count;
    //Constructor
    Counter(int i){
        System.out.println("Init user defined counter");
        count = i;
    }
    Counter(){
        System.out.println("Init default counter");
        count = 0;
    }
    Counter(String s){
        if(s == "I"){
            count = 1;
        }else if(s == "II"){
            count = 2;
        } else if(s == "III"){
            count = 3;
        }else if(s == "V"){
            count = 5;
        }else if(s == "X"){
            count = 10;
        }else {
            count = 0;
        }
    }
    // methods
    void increment(){
        count++;
    }
    void decrement(){
        count--;
    }
    void reset(){
        count = 0;
    }
    public int getCount() {
        return count;
    }
}