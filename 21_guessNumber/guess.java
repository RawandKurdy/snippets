package main;

import java.util.Random;
import java.util.Scanner;

public class Guess {
    // Guess Random Number Game
    // In Java

    // generated Number won't exceed this value
    static int max_num = 500;

    public static void main(String args[]) {
        Random ran = new Random();

        // Seeds the random generator
        ran.setSeed(System.currentTimeMillis());

        // Range, From & To, a Random Range Between 1, max_num
        int range_From = ran.nextInt(max_num);
        if (range_From == 0) range_From = 1;
        int range_To = ran.nextInt(max_num - range_From + 1) + range_From;

        // Random Number that the User have to guess :D
        int random_Num = ran.nextInt(range_To - range_From + 1) + range_From;

        // Number of allowed guesses` >= 3
        int number_Of_Guesses = 3 + (range_To - range_From) % 3;

        int user_attempts = 0; // Number of tries by User
        boolean user_won = false; // Is True if User enter random number
        String tip = ""; // Tip msg, changes according

        System.out.printf("Welcome to Guess the Random Number Game :D\n" +
            "Written in Java o_o\n The number is between {%d} ~ {%d}\n",
            range_From, range_To);

        // Used to read inputs from user
        Scanner reader = new Scanner(System.in);

        while (number_Of_Guesses > user_attempts) {
            String msg = "\r You have %d tries left,%s Enter a number: ";
            System.out.printf(msg, (number_Of_Guesses - user_attempts - 1),
            tip);

            int user_number = reader.nextInt();
            user_attempts += 1;
            if (user_number == random_Num) {
                user_won = true;
                break;
            } else if (user_number > random_Num)
                tip = "maybe you should try a smaller number :D";
            else if (user_number < random_Num)
                tip = "maybe you should try a bigger number :D";
        }

        if (user_won)
            System.out.printf("\nCongrats!!! YOU WON! with %d attempts\n",
            user_attempts);
        else
            System.out.printf("\n:( You Lost!, you literally had %d attempts," +
                " Number was %d \n", user_attempts, random_Num);
    }
}