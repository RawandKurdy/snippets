# Longest Subarray By Sum in an Array - Interview Question -
# Asked by Facebook & Palantir Technologies
# In Python
def findLongestSubarrayBySum(sum, arr):
    prefixSumsToIndex = {}
    fromIndex = -1
    toIndex = -1
    sumExistAt = -1
    
    # Calculates the prefixSums
    previousSum = 0
    for index in range(len(arr)):
        number = arr[index]
        newSum = previousSum + number
        previousSum = newSum
        prefixSumsToIndex[newSum] = index + 1
        if number == sum:
            sumExistAt = index + 1
  
    # Sees if such sum exists within the array
    # initializes the indices with initial range if exists
    if (prefixSumsToIndex.get(sum) is not None):
        fromIndex = 1
        toIndex = prefixSumsToIndex.get(sum)

    # Searches the array for a wider range if available
    for numSum in prefixSumsToIndex:
        indexofSum = prefixSumsToIndex[numSum]
        indexOfSumWithS = prefixSumsToIndex.get(numSum + sum)
        if (indexOfSumWithS is not None):
            tmpFromIndex = indexofSum + 1
            tmpToIndex = indexOfSumWithS
            if (toIndex - fromIndex) < (tmpToIndex - tmpFromIndex):
                toIndex = tmpToIndex
                fromIndex = tmpFromIndex

    # Returns depending on if indices has been found or not
    if fromIndex != -1 and toIndex != -1:
        return [fromIndex, toIndex]
    elif sumExistAt != -1:
        return [sumExistAt, sumExistAt]
    else: return [-1]

# Example
sum = 15
arr = [6, 7, 8, 1, 2, 3, 0 , 4, 5, 0, 0, 9, 10]
print(findLongestSubarrayBySum(sum, arr)) # [4, 11]