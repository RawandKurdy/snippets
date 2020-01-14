const readline = require("readline").createInterface(
    process.stdin, process.stdout
);

// Guess Random Number Game
// In Javascript

(async()=>{

// generated Number won't exceed this value
const max_num = 500;

// Range, From & To, a Random Range Between 1, max_num
let range_From = Math.floor(Math.random() * max_num) + 1;
let range_To = Math.floor(Math.random() * (max_num - range_From+1)) + range_From;

// Random Number that the User have to guess :D
let random_Num = Math.floor(Math.random() * (range_To - range_From+1))+range_From;
// Number of allowed guesses` >= 3
let number_Of_Guesses = 3 + (range_To - range_From) % 3;

let user_attempts = 0;  // Number of tries by User
let user_won = false;  // Is True if User enter random number
let tip = "";  // Tip msg, changes according

console.log("\nWelcome to Guess the Random Number Game :D" +
    "\nWritten in Javascript o_o" +
    `\nThe number is between ${range_From} ~ ${range_To}\n\n`);

while (number_Of_Guesses > user_attempts) {
    msg = `\r You have ${number_Of_Guesses - user_attempts - 1} ` +
    `tries left,${tip} Enter a number: `;

    user_number = Number.parseInt(await new Promise((resolve, reject) => {
        readline.question(msg, (value) => {
            resolve(value);
        });
    }));

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
    console.log(`\nCongrats!!! YOU WON! with ${user_attempts} attempts`);
else
    console.log(
        `\n:( You Lost!, you literally had ${user_attempts} attempts` +
        `, Number was ${random_Num}`
    );

process.exit(0);
})();
