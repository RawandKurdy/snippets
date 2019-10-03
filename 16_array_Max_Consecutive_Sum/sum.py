# Array Max Consecutive Sum - Interview Question -
# Asked by Microsoft, LinkedIn, Amazon and Samsung 
# Maximum possible sum you can get from one of contiguous subarrays.
# In Python
def arrayMaxConsecutiveSum2(inputArray):
    sumMax = 0
    sumSteps = 0
    maxNumber = 0

    for index in range(len(inputArray)):
        number = inputArray[index]
        sumSteps += number

        if (maxNumber == 0 or maxNumber < number):
            maxNumber = number
    
        if (sumSteps < 0):
            sumSteps = 0
        if (sumMax < sumSteps):
            sumMax = sumSteps
            
    return sumMax if sumMax != 0 else maxNumber

# Example
arr = [1, -2, 3, -4, 5, -3, 2, 2, 2]
print(arrayMaxConsecutiveSum2(arr)) # 8