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

    public static int partOne(int[] codes) {
        int i = 0;
        boolean running = true;
        while (i < codes.length && running) {
            switch (codes[i]) {
                case 1:
                    codes[codes[i + 3]] = codes[codes[i + 1]] + codes[codes[i + 2]];
                    i += 4;
                    break;
                case 2:
                    codes[codes[i + 3]] = codes[codes[i + 1]] * codes[codes[i + 2]];
                    i += 4;
                    break;
                case 99:
                    running = false;
                    i++;
                    break;
                default:
                    running = false;
                    System.out.println("Error: Unknown instruction");
                    break;
            }
        }
        return codes[0];
    }

    public static int partTwo(int[] codes) {
        int output = 19690720;
        for (int i = 0; i < 100; i++) {
            for (int j = 0; j < 100; j++) {
                int[] copy = Arrays.copyOf(codes, codes.length);
                copy[1] = i;
                copy[2] = j;
                if (partOne(copy) == output) {
                    return 100 * i + j;
                }
            }
        }
        return 0;
    }

    public static void main(String[] args) {
        String[] lines = readLines().get(0).split(",");
        int[] codes = new int[lines.length];
        for (int i = 0; i < lines.length; i++) {
            codes[i] = Integer.parseInt(lines[i]);
        }

        System.out.println("1: " + partOne(Arrays.copyOf(codes, codes.length)));
        System.out.println("2: " + partTwo(codes));
    }
}