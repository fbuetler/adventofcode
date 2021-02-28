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

    public static int requiredFuel(int mass) {
        return ((int) Math.floor(mass / 3)) - 2;
    }

    public static int partOne(List<String> lines) {
        int fuel = 0;
        for (int i = 0; i < lines.size(); i++) {
            int mass = Integer.parseInt(lines.get(i));
            fuel += requiredFuel(mass);
        }
        return fuel;
    }

    public static int partTwo(List<String> lines) {
        int totalFuel = 0;
        for (int i = 0; i < lines.size(); i++) {
            int mass = Integer.parseInt(lines.get(i));
            do {
                int fuel = requiredFuel(mass);
                totalFuel += fuel > 0 ? fuel : 0;
                mass = fuel;
            } while (mass > 0);
        }
        return totalFuel;
    }

    public static void main(String[] args) {
        List<String> lines = readLines();
        System.out.println("1: " + partOne(lines));
        System.out.println("2: " + partTwo(lines));
    }
}