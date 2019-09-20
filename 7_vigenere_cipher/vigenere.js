// Vigenere Cipher in JavaScript

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ".split("");
let stringSecret = "Key";  // Secret key used
let secretKeyIndexes = [];  // Indexes of letters in string secret
var secretKeyLength = stringSecret.length;  // length of the string secret


function operation(secretKey, input, decrypt) {
    let allCapsInput = `${input}`.toUpperCase().split("");
    let tempCounter = 0; // Temporary Counter
    if (secretKeyIndexes.length == 0) { // Convert secret key to set of indexes
        for (char of secretKey.toUpperCase()) {
            secretKeyIndexes.push(alphabet.indexOf(char));
        }
    }

    return allCapsInput.map(
        (char) => {
            let index = alphabet.indexOf(char);
            if (index == null || index == -1) return char;
            else {
                let key = secretKeyIndexes[tempCounter % secretKeyLength];
                tempCounter += 1;
                if (decrypt) key = key * -1;

                return alphabet[(26 + index + key) % 26];
            }
        }).join("");
}

// Function expressions to make stuff easier
const encrypt = text => operation(stringSecret, text, false);
const decrypt = text => operation(stringSecret, text, true);

let plaintext = "New York!";
let ciphertext = encrypt(plaintext); // "XIU ISPU!"
let decryptedtext = decrypt(ciphertext); // "NEW YORK!"
console.log(ciphertext, decryptedtext);