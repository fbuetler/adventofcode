import java.io.*;
import java.util.*;

public class Solve {

    public static List<String> readLines() {
        List<String> lines = new ArrayList<>();
        File file = new File(System.getProperty("user.dir") + "/input.txt");
        try {
            BufferedReader br = new BufferedReader(new FileReader(file));
            String line;
            while ((line = br.readLine()) != null) {
                lines.add(line);
            }
            br.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
        return lines;
    }

    public static int partOne(List<String> lines) {
        return 0;
    }

    public static int partTwo(List<String> lines) {
        return 0;
    }

    public static void main(String[] args) {
        List<String> lines = readLines();
        System.out.println("1: " + partOne(lines));
        System.out.println("2: " + partTwo(lines));
    }
}