# DivNumber Challenge
# Given three integers k, l and r
# return the number of integers between l and r inclusive
# that have exactly k divisors.
# In Python
def divNumber(k, l, r):
    # di is the number of Numbers
    # that has 'k' exact divisors
    di=0
    
    # n is short for Number
    for n in range(l ,r+1):
        
        #d is number of divisors of n
        d = 0
        for i in range(1, n+1):
            if n % i == 0: 
                d += 1
        if d == k:
            di += 1
    return di

# Example
k, l, r = 3, 2, 49
print(divNumber(k, l, r)) # 4

# We have 4 integers in that range
# that have exactly 3 divisors - 4, 9, 25, 49.
