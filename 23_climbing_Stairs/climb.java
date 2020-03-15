package main;

import java.util.Vector;
import java.util.Arrays;

public class Climb {

    // Climbing Stairs - Interview Question -
    // Asked By Apple and Adobe
    // in Java

    public static int climbingStairs(int n) {
        Vector<Integer> r = new Vector<Integer>();
        // Initial Result Set
        r.addAll(Arrays.asList(new Integer[] { 1, 1, 2 }));

        if (n <= 2)
            return r.get(n);

        for (int step = 3; step < n + 1; step++)
            r.add(r.get(step - 1) + r.get(step - 2));
        return r.get(n);
    }

    public static void main(String[] args) {

        int n = 26; // Steps
        int ways = climbingStairs(n);
        System.out.println(ways); // 196418
    }
}