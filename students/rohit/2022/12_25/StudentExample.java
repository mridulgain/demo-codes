import java.time.LocalDate;
import java.time.Period;

public class StudentExample {
    public static void main(String[] args){
        Student s1 = new Student();
        s1.roll = 1;
        s1.name = "Rohitashu";
        s1.class_section = "IX-A";
        s1.grade = 'A';
        s1.introduction();

        Student s2 = new Student();
        s2.roll = 10;
        s2.name = "Rohit";
        s2.class_section = "IX-B";
        s2.grade = 'A';
        s2.introduction();

        // handeling dates
        String date = "1996-11-15";
        LocalDate ld = LocalDate.parse(date);
        LocalDate curDate = LocalDate.now();
        System.out.println(Period.between(curDate, ld));
    }
}

class Student{
    // properties
    // <data_type> <variable_name>;
    int roll, age;
    String name, class_section;
    char grade;
    //behaviour
    public void introduction() {
        String greet = "Hello, my name is " + name + ". My roll number is " + roll + ". I study in class " + class_section + ".";
        System.out.println(greet);
    }
}
