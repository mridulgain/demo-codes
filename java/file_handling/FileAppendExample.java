import java.io.FileWriter;
import java.io.IOException;

public class FileAppendExample {
    public static void main(String[] args) {
        try {
            FileWriter writer = new FileWriter("example.txt", true); // true enables append mode
            writer.write("\nAppending new content to the file.");
            writer.close();
            System.out.println("Successfully appended to the file.");
        } catch (IOException e) {
            System.out.println("An error occurred.");
            e.printStackTrace();
        }
    }
}
