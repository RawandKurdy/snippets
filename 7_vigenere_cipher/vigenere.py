# Vigenere Cipher in Python

alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
stringSecret = "Key"  # Secret key used
secretKeyIndexes = []  # Indexes of letters in string secret
secretKeyLength = len(stringSecret)  # length of the string secret


def operation(secret_key, input, decrypt):
    tempCounter, output, allCapsInput = 0, "", input.upper()
    global secretKeyIndexes, secretKeyLength

    if len(secretKeyIndexes) == 0: # Convert Secret key to a set of indexes
        secretKeyIndexes = [alphabet.find(char) for char in secret_key.upper()]

    for char in allCapsInput:
        index = alphabet.find(char)
        if index == -1:
            output += char
        else:
            key = secretKeyIndexes[tempCounter % secretKeyLength]
            if decrypt:
                key = key * -1
            tempCounter +=1
            output += alphabet[(26 + index + key) % 26]
    return output

# Function expressions to make stuff easier
encrypt = lambda text : operation(stringSecret, text, False)
decrypt = lambda text : operation(stringSecret, text, True)

plaintext = "New York!"
ciphertext = encrypt(plaintext) # "XIU ISPU!"
decryptedtext = decrypt(ciphertext) # "NEW YORK!"
print(ciphertext, decryptedtext)