# Functions in Python

# 1- Traditional function
# also used to demonstrate how an anonymous func can be useful
# It  runs a function passed as a param then prints its result
def traditionalFunc(anotherFunc):
    text_with_duplicates = anotherFunc("NYC!", 3)
    print(text_with_duplicates)

# 2- Anonymous Function
# I passed it as a parameter
# it duplicated the text as many time as the user wishes
traditionalFunc(lambda text, times : text * times)

# 3- Function Expression
anotherFunc = lambda text: text[len(text)-1] 

# Exec Expr. function
last_char = anotherFunc("NYC")
print(last_char) # C
