// is Crypt Solution - Interview Question - Asked by Palantir Technologies
// A Cryptarithm is a mathematical puzzle,
// Here we check if the solution is correct for the crypt puzzle

function isCryptSolution(crypt, solution) {
    let letterToNumber = {}; // map letter > number
    let equation = ["", "", ""]; // number1, number2, result
    
    for (let pair of solution) {
        let letter = pair[0];
        let number = pair[1];
        letterToNumber[letter] = number;
    }
        
    for (let word_index = 0; word_index < crypt.length; word_index+=1){
        word = crypt[word_index]; // a word in crypt
        zeroExist = false;
        
        for (let index = 0; index < word.length; index+=1) {
            // number that correspond with the letter
            let number = letterToNumber[word[index]];
            
            if (zeroExist) return false; // if zero exist at start then exit
            if (number == '0' && index == 0) zeroExist = true;
            
            equation[word_index]+= number;
        }
    }
    // Evaluate the equation and see if its valid
    return parseInt(equation[0])+parseInt(equation[1]) == parseInt(equation[2]);
}

let crypt= ["SEND", 
            "MORE", 
            "MONEY"]
let solution= [["O","0"], 
               ["M","1"], 
               ["Y","2"], 
               ["E","5"], 
               ["N","6"], 
               ["D","7"], 
               ["R","8"], 
               ["S","9"]]

console.log(isCryptSolution(crypt, solution)) // True