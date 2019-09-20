// Caesar Cipher in JavaScript

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ".split("");
let shift_key = 7;

function operation(key, input) {
    let allCapsInput = `${input}`.toUpperCase().split("");

    return allCapsInput.map(
        (char)=> {
            let index = alphabet.indexOf(char);
            if (index == null || index == -1) return char;
            else return alphabet[(26 + index + key) % 26];
        }).join("");
}

// Function expressions to make stuff easier
const encrypt = text => operation(shift_key, text);
const decrypt = text => operation((shift_key * -1), text);

let plaintext = "New York!";
let ciphertext = encrypt(plaintext); // "ULD FVYR!"
let decryptedtext = decrypt(ciphertext) // "NEW YORK!"
console.log(ciphertext, decryptedtext)