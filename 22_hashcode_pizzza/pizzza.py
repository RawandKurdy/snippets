import os

# More Pizza
# Google Hash Code 2020
# Practice Problem
# In Python

def fileOperation(fileName, flag):
    path, op = "", ""
    if flag == "a":
        path = f"./output/{fileName}.out"
        op = "Writing To"
    elif flag == "r":
        path = f"./input/{fileName}.in"
        op = "Reading From"

    print(f"{op} {path}")
    return open(path, flag, 1, "UTF8", None, "\n")

# Used to read the file
def readFile(fileName):
    return fileOperation(fileName, "r")

# Used to write the file
def writeFile(fileName, array=[]):
    f = fileOperation(fileName, "a")
    if f.writable():
        f.write(f"{len(array)} \n")  # Header
        f.write(" ".join(array))
    else:
        print("File not writable")
    f.close()


# Gets the file name
fileNameEnv = os.getenv("p_file")
fileName = fileNameEnv if fileNameEnv != None else input("Enter File Name:\n")

# Starts Reading file from source
file = readFile(fileName)
header = file.readline()
data = file.readline()
file.close()

# Problem Variables
maxSlice, numOftypes = [int(x) for x in header.split(" ")]
array = data.split(" ")
numOfImpTypes = len(array)

print(f"""Maximum Number of Slices: {maxSlice}
Number of Available Types: {numOftypes}""")
print(f"Number of Imported Types: {numOfImpTypes}")

usedTypes = []  # Pizza Type Keys (index)
slices = 0  # Points (Slices)
for i in range(0, numOfImpTypes): # Processing the data
    v = int(array[numOfImpTypes-1-i])
    if v <= maxSlice:
        usedTypes.append(str(numOfImpTypes-1-i))
        maxSlice -= v
        slices += v
    print(f"\r Processing Types {i+1}/{numOfImpTypes} ", end="")

print("\n Done!")
noTypesUsed = len(usedTypes)
print(f"Grabbed {noTypesUsed} Pizza Types")
print(f"Which is {slices} slices in total")

# Writing the output (Solution)
writeFile(fileName, sorted(usedTypes))

# Optional (Extra)
# Writes Output Extra Info
writeFile(fileName + "_details",
          [str(noTypesUsed), str(slices)])
