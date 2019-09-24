# Amend The Sentence - Adobe - Interview Question -
# in Python

def amendTheSentence(text):
    new_str = ''

    for char in text:
        if (char.isupper()):
            new_str += ' '+ char
        else: new_str += char

    return new_str.strip().lower()

input = "YouAreInTheRightPlace"
print("input: ", input)
print("output: ",amendTheSentence(input))