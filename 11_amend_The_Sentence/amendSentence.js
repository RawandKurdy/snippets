// Amend The Sentence - Adobe - Interview Question -
// in JavaScript

function amendTheSentence(text) {
    var new_str = '';
    
    for (char of text) {
        if (char>="A" && char <= "Z") new_str += ' '+ char;
        else new_str += char;
    }
    return new_str.trim().toLowerCase();
}

let input = "YouAreInTheRightPlace";
console.log("input: ", input);
console.log("output: ", amendTheSentence(input));