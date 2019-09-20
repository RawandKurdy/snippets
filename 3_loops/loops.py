# Loops in python

arr = ["New York", "Washington", "California"]


# Foreach loop
for city in arr:
    print(city)

# While loop 
# (continuous till condition is true) 
condition = True
while condition:
    city = arr.pop()
    print(city)
    condition = not (len(arr) == 0)
