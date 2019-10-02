# Sum In Range - Interview Question - Asked by Palantir Technologies
# In Python
# In O(n) time

def sumInRange(nums, queries):
    totalSum = 0
    prefixSums = []
    modulo = 1000000000 + 7
    
    # calculates the prefix-sums
    for i in range(len(nums)):
        if (i == 0):
            prefixSums.append(nums[i] % modulo)
        else:
            prefixSums.append((prefixSums[i-1] + nums[i]) % modulo)

    # Now adds up each separate range
    for query in queries:
        fromIndex = query[0]
        toIndex = query[1]
        totalSum += prefixSums[toIndex]

        if (fromIndex != 0): totalSum -= prefixSums[fromIndex-1]
        totalSum %= modulo
    # if totalSum is negative we just add the modulo
    if (totalSum < 0): totalSum += modulo
    return totalSum

# Example
nums =  [34, 19, 21, 5, 1, 10, 26, 46, 33, 10]
queries = [[3,7], 
 [3,4], 
 [3,7], 
 [4,5], 
 [0,5]]

print(sumInRange(nums, queries)) # 283