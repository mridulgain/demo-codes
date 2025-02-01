public class ExceptionHandlingExample {
    public static void main(String[] args) {
        int[] numbers = {1, 2, 3};
        try {
            System.out.println(numbers[5]); // ArrayIndexOutOfBoundsException
        } catch (ArrayIndexOutOfBoundsException e) {
            System.out.println(numbers[numbers.length - 1]);
        } catch (Exception e) {
            System.out.println("An error occurred: " + e.getMessage());
        } finally {
            System.out.println("Execution completed.");
        }
    }
}
