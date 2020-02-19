(async function(){
const fs = require("fs");
const rl = require("readline");
const { log, time, timeEnd } = console;
const { stdin, stdout } = process;

// More Pizza
// Google Hash Code 2020
// Practice Problem
// In JavaScript

let reader = rl.createInterface(stdin, stdout);

// Used to read the file
function readFile(fileName) {
    let path = `./input/${fileName}.in`;
    log(`Reading From ${path}`);
    return fs.readFileSync(path, {
        encoding: "UTF8",
        flag: "r"
    });
}

// Used to write the file
function writeFile(fileName, array = []) {
    let path = `./output/${fileName}.out`;
    log(`Writing to ${path}`);

    let f = fs.createWriteStream(path, {
        encoding: "UTF8",
        flags: "a"
    });

    if (f.writable) {
        f.write(`${array.length} \n`);  // Header
        f.end(array.join(" "));
    } else {
        print("File not writable");
    }
}

// Gets the file name
let fileNameEnv = process.env["p_file"];
let fileName = fileNameEnv != null ? fileNameEnv :
    await new Promise((resolve, rej) => {
        reader.question("Enter File Name:\n", (answer) => {
            resolve(answer)
        });
    });

time("Operation"); // Timer (Optional)

// Starts Reading file from source
let file = readFile(fileName).split("\n");
let header = file[0].split(" ").map(v => parseInt(v, 10));
let array = file[1].split(" "); // Data

// Problem Variables
let [maxSlice, numOftypes] = header;
let numOfImpTypes = array.length;

log(`Maximum Number of Slices: ${maxSlice}
Number of Available Types: ${numOftypes}`);
log(`Number of Imported Types: ${numOfImpTypes}`);

let usedTypes = [];  // Pizza Type Keys (index)
let slices = 0;  // Points (Slices)
for (let i = 0; i < numOfImpTypes; i++) { // Processing the data
    let v = parseInt(array[numOfImpTypes - 1 - i], 10);
    if (v <= maxSlice) {
        usedTypes.push(numOfImpTypes - 1 - i);
        maxSlice -= v;
        slices += v;
        stdout.write(`\r Processing Types ${i + 1}/${numOfImpTypes}`);
    }
}
log("\n Done!");

let noTypesUsed = usedTypes.length;
log(`Grabbed ${noTypesUsed} Pizza Types`);
log(`Which is ${slices} slices in total`);

// Writing the output (Solution)
writeFile(fileName, usedTypes.sort());

// Optional (Extra)
// Writes Output Extra Info
writeFile(fileName + "_details",
    [noTypesUsed, slices]);

timeEnd("Operation"); // Stops the Timer
log("You can exit now :)");
})();