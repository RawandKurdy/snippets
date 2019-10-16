# Product Except Self [in O(n)] - Interview Question -
# Asked by Amazon, LinkedIn, Facebook, Microsoft & Apple
# In Python

def productExceptSelf(nums, m):
    s = 0 # Sum
    p = 1 # Product
    
    for num in nums:
        s = (p + s * num) % m
        p = (p * num) % m
    return s

	# Example
nums = [3, 3, 9, 5, 5, 4, 2, 8, 5, 9] # Input Array
m = 17  # Mod value

print(productExceptSelf(nums, m)) # 16