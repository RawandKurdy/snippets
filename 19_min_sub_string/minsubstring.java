package main;

import java.util.HashSet;
import java.util.ArrayList;

public class MinSubString {
    // Minimum Substring With All Chars - Interview Question -
    // Asked by Facebook
    // In Java
    String minSubstringWithAllChars(String s, String t) {
        int size = s.length() + 1; // length of the text

        for (int x = 0; x < size; x++) {
            for (int p = 0; p < size - x; p++) {
                // sub string candidate
                // also gets filtered
                // chars that are not in t get removed for now
                ArrayList c = filter(s.substring(p, p + x), t);

                // if sub string contain all characters that are in t
                // then correct answer is found :D
                if (new HashSet(c).size() == t.length()) {
                    return s.substring(p, p + x);
                }
            }
        }
        return "";
    }

    // Used to filter a string
    ArrayList filter(String s, String pattern) {
        ArrayList list = new ArrayList();

        for (int x = 0; x < s.length(); x++) {
            String letter = s.substring(x, x + 1);

            if (pattern.indexOf(letter) != -1) {
                list.add(letter);
            }
        }
        return list;
    }

    public static void main(String[] args) {
        MinSubString obj = new MinSubString();

        // Example
        String s = "adobecodebanc";
        String t = "abc";
        System.out.println(obj.minSubstringWithAllChars(s, t));// "banc"
    }
}