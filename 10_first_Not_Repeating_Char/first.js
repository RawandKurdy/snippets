// First Not Repeating Character - Amazon - Interview Question-
// in JavaScript

function firstNotRepeatingCharacter(text) {

    // Holds repeating and non repeating characters
    let nonRepeatedChars = new Map();
    let RepeatedChars = new Map();
    
    for (let index = 0; index < text.length; index += 1) {
        let charAtIndex = text[index];
        let isOnRepeatedChars = RepeatedChars.get(charAtIndex) != undefined;
        let isOnNonRepeatedChars = nonRepeatedChars.get(charAtIndex) != undefined;

        if (!isOnNonRepeatedChars
            && !isOnRepeatedChars) {
            nonRepeatedChars.set(charAtIndex, charAtIndex);

        } else if (isOnNonRepeatedChars
            && !isOnRepeatedChars) {
            nonRepeatedChars.delete(charAtIndex);
            RepeatedChars.set(charAtIndex, charAtIndex);
        }
    }

    if (nonRepeatedChars.size == 0) return '_';

    let firstNonRepeatedChar = nonRepeatedChars.values().next().value;

    return firstNonRepeatedChar == undefined ? '_' : firstNonRepeatedChar;
}

input = "thebestshoweverisgameofthrones"
char = firstNotRepeatingCharacter(input)
console.log("input: ", input)
console.log("result: ", char)