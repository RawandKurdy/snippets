#include <iostream>
#include <random>

// Guess Random Number Game
// In C++

// generated Number won't exceed this value
int max_num = 500;

int main(){
    // Seeds the random generator
    srand(time(NULL));

    // Range, From & To, a Random Range Between 1, max_num
    unsigned int range_From = rand() % max_num + 1;
    unsigned int range_To = rand() % (max_num - range_From + 1) + range_From;

    // Random Number that the User have to guess :D
    unsigned int random_Num = rand() % (range_To - range_From + 1) + range_From;

    // Number of allowed guesses` >= 3
    unsigned int number_Of_Guesses = 3 + (range_To - range_From) % 3;

    unsigned int user_attempts = 0; // Number of tries by User
    bool user_won = false; // Is True if User enter random number
    std::string tip = ""; // Tip msg, changes according
    
    std::string msg = "Welcome to Guess the Random Number Game :D\n"+
    std::string("Written in C++ o_o\nThe number is between {%d} ~ {%d}\n");
    std::printf(msg.c_str(), range_From, range_To);

    while (number_Of_Guesses > user_attempts)
    {
        std::cout << "\rYou have " << (number_Of_Guesses-user_attempts-1);
        std::cout << " tries left," << tip << " Enter a number: ";
        unsigned int user_number; 
        std::cin >> user_number;

        user_attempts += 1;
        if (user_number == random_Num){
            user_won = true;
            break;
        } else if (user_number > random_Num){
            tip = "maybe you should try a smaller number :D";
        } else if (user_number < random_Num){
            tip = "maybe you should try a bigger number :D";
        }
    }

    if (user_won){
        std::cout << "\nCongrats!!! YOU WON! with " << user_attempts;
        std::cout << " attempts" << std::endl;
    } else {
        std::cout << "\n:( You Lost!, you literally had " << user_attempts;
        std::cout << " attempts, Number was " << random_Num << std::endl;
    }
}
