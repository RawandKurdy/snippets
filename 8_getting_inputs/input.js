// getting an input from the console
// in JavaScript (using NodeJS)

const readline = require("readline")

// Reader Object to read from stdin (input) and write to (stdout)
let reader = readline.createInterface(process.stdin, process.stdout);

reader.question("What is your name?", (answer)=> {
    reader.write(`Hello ${answer} :) \n`); // Writes to stdout (output)
    reader.close(); // Closes the reader
});
