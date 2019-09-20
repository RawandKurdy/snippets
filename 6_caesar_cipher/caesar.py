# Caesar Cipher in Python

alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
shift_key = 7

def operation(key, input):
    allCapsInput = input.upper()
    output = ""
    for char in allCapsInput:
        index = alphabet.find(char)
        if index == -1:
            output += char
        else: 
            output += alphabet[(26 + index + key) % 26]
    return output

# Function expressions to make stuff easier
encrypt = lambda text : operation(shift_key, text)
decrypt = lambda text : operation((shift_key * -1), text)

plaintext = "New York!"
ciphertext = encrypt(plaintext) # "ULD FVYR!"
decryptedtext = decrypt(ciphertext) # "NEW YORK!"
print(ciphertext, decryptedtext)