package main;

import java.io.File;
import java.io.FileWriter;
import java.nio.charset.Charset;
import java.util.Scanner;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;
import java.util.List;

public class Pizzza {

    // More Pizza
    // Google Hash Code 2020
    // Practice Problem
    // In Java

    public static void main(String[] args) {
        // Gets the file name
        String fileNameEnv = System.getenv("p_file");
        String fileName = fileNameEnv != null ? 
        fileNameEnv : input("Enter File Name:");

        // Starts Reading file from source
        ArrayList d = readFile(fileName);
        String[] header = ((String) d.get(0)).split(" ");
        String[] array = ((String) d.get(1)).split(" ");

        // Problem Variables
        int maxSlice = Integer.parseInt(header[0]);
        int numOftypes = Integer.parseInt(header[1]);
        int numOfImpTypes = array.length;

        System.out.printf(
                "Maximum Number of Slices: %d\n" +
                "Number of Available Types: %d\n" +
                "Number of Imported Types: %d\n",
                maxSlice, numOftypes, numOfImpTypes);

        // Type Keys (index)
        ArrayList<String> usedTypes = new ArrayList<String>();
        int slices = 0; // Points (Slices)

        // Processing the data
        for (int i = 0; i < numOfImpTypes; i++) {
            int v = Integer.parseInt(array[numOfImpTypes - 1 - i], 10);
            if (v <= maxSlice) {
                usedTypes.add(String.valueOf(numOfImpTypes - 1 - i));
                maxSlice -= v;
                slices += v;
                System.out.printf("\r Processing Types %d/%d", i + 1,
                numOfImpTypes);
            }
        }

        System.out.println("\nDone!");
        long noTypesUsed = usedTypes.size();
        System.out.printf("Grabbed %d Pizza Types\n", noTypesUsed);
        System.out.printf("Which is %d slices in total\n", slices);

        usedTypes.sort(Comparator.naturalOrder());
        writeFile(fileName, usedTypes); // Writing the output (Solution)

        // Optional (Extra)
        // Writes Output Extra Info
        writeFile(fileName + "_details",
                Arrays.asList(new String[] { String.valueOf(noTypesUsed),
                    String.valueOf(slices) }));
    }

    // Used to read the file
    static ArrayList<String> readFile(String fileName) {
        String path = "./input/" + fileName + ".in";
        System.out.printf("Reading From %s \n", path);
        ArrayList r = new ArrayList<String>();
        try {
            File f = new File(path);
            Scanner i = new Scanner(f, "UTF8");
            while (i.hasNextLine())
                r.add(i.nextLine());
            i.close();
        } catch (Exception e) {
            System.out.println(e);
            System.exit(-1);
        }
        return r;
    }

    //  Used to write the file
    static boolean writeFile(String fileName, List a) {
        String path = "./output/" + fileName + ".out";
        System.out.printf("Writing To %s \n", path);
        boolean r = false;
        try {
            File f = new File(path);
            FileWriter fw = new FileWriter(f, Charset.forName("UTF8"));
            fw.write(String.format("%d\n", a.size()));
            fw.write(String.join(" ", a));
            fw.close();
            r = true;
        } catch (Exception e) {
            System.out.println(e);
            System.exit(-1);
        }
        return r;
    }

    // Used to get user input if needed
    static String input(String msg) {
        System.out.println(msg);
        Scanner i = new Scanner(System.in);
        String input = i.nextLine();
        i.close();
        return input;
    }
}