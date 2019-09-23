# First Not Repeating Character - Amazon - Interview Question-
# in Python

def firstNotRepeatingCharacter(text):
    # Holds repeating and non repeating characters
    nonRepeatedChars, RepeatedChars = {}, {}

    for index in range(len(text)):
        charAtIndex = text[index]
        isOnRepeatedChars = charAtIndex in RepeatedChars
        isOnNonRepeatedChars = charAtIndex in nonRepeatedChars

        if (not isOnNonRepeatedChars and not isOnRepeatedChars):
            nonRepeatedChars[charAtIndex] = charAtIndex

        elif (isOnNonRepeatedChars and not isOnRepeatedChars):
            del nonRepeatedChars[charAtIndex]
            RepeatedChars[charAtIndex] = charAtIndex

    if (len(nonRepeatedChars) == 0): return '_'

    return next(iter(nonRepeatedChars))

input = "thebestshoweverisgameofthrones"
char = firstNotRepeatingCharacter(input)
print("input: ", input)
print("result: ", char)