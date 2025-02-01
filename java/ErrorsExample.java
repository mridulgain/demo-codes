public class ErrorsExample {
    public static void main(String[] args) {
        try {
            int result = 10 / 0; // Division by zero (ArithmeticException)
            System.out.println(result);
        } catch (ArithmeticException e) {
            System.out.println("undefined");
        }
    }
}

// Syntax Errors (Compile-time Errors)
// Runtime Errors (Exceptions)
// Logical Errors