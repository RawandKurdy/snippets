# is Crypt Solution - Interview Question - Asked by Palantir Technologies
# A Cryptarithm is a mathematical puzzle,
# Here we check if the solution is correct for the crypt puzzle

def isCryptSolution(crypt, solution):
    letterToNumber = {} # map letter > number
    equation = ["", "", ""] # number1, number2, result
    
    for pair in solution:
        letter, number = pair
        letterToNumber[letter] = number
        
    for word_index in range(len(crypt)):
        word = crypt[word_index] # a word in crypt
        zeroExist = False
        
        for index in range(len(word)):
            # number that correspond with the letter
            number = letterToNumber[word[index]]
            
            if (zeroExist): # if zero exist at start then exit
                return False
            if (number == '0' and index == 0):
                zeroExist = True
            
            equation[word_index]+= number

    # Evaluate the equation and see if its valid
    return int(equation[0])+int(equation[1]) == int(equation[2])

crypt= ["SEND", 
        "MORE", 
        "MONEY"]
solution= [["O","0"], 
           ["M","1"], 
           ["Y","2"], 
           ["E","5"], 
           ["N","6"], 
           ["D","7"], 
           ["R","8"], 
           ["S","9"]]

print(isCryptSolution(crypt, solution)) # True