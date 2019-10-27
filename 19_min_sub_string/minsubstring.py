# Minimum Substring With All Chars - Interview Question -
# Asked by Facebook
# In Python
def minSubstringWithAllChars(s, t):
    size = len(s) + 1  # length of the text
    for x in range(size):
        for p in range(size - x):
            # sub string candidate
            # also gets filtered
            # chars that are not in t get removed for now
            c = filter(lambda el: el in t, s[p:p + x])

            # if sub string contain all characters that are in t
            # then correct answer is found :D
            if len(set(c)) == len(t):
                return s[p:p + x]
    return ""


# Example
s = "adobecodebanc"
t = "abc"
print(minSubstringWithAllChars(s, t))  # "banc"
