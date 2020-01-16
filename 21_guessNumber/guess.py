import random
import time

# Guess Random Number Game
# In Python

# generated Number won't exceed this value
max_num = 500

# Seeds the random generator
random.seed(time.time())

# Range, From & To, a Random Range Between 1, max_num
range_From = random.randint(1, max_num)
range_To = random.randint(range_From + 1, max_num)

# Random Number that the User have to guess :D
random_Num = random.randint(range_From, range_To)
# Number of allowed guesses` >= 3
number_Of_Guesses = 3 + (range_To - range_From) % 3

user_attempts = 0  # Number of tries by User
user_won = False  # Is True if User enter random number
tip = ""  # Tip msg, changes according
print(f"""
Welcome to Guess the Random Number Game :D
Written in Python o_o
The number is between {range_From} ~ {range_To}\n
""")
while number_Of_Guesses > user_attempts:
    msg = f"\r You have {number_Of_Guesses-user_attempts-1} tries left"
    msg += f",{tip} Enter a number: "
    user_number = int(input(msg))
    user_attempts += 1
    if user_number == random_Num:
        user_won = True
        break
    elif user_number > random_Num:
        tip = "maybe you should try a smaller number :D"
    elif user_number < random_Num:
        tip = "maybe you should try a bigger number :D"

if user_won:
    print(f"\nCongrats!!! YOU WON! with {user_attempts} attempts")
else:
    print(
        f"\n:( You Lost!, you literally had {user_attempts} attempts,",
        f"Number was {random_Num}"
    )
