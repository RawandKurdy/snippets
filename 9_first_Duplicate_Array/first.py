# First Duplicate Value - Google - Interview Question -
# in Python

def firstDuplicate(array):
    arrLength = len(array)
    if arrLength == 0 or arrLength == 1:
        return -1

    valuesMap = {}
    for index in range(arrLength):
        valueAtIndex = array[index]
        valueIndexes = valuesMap.get(valueAtIndex)

        if valueIndexes is None:
            valuesMap[valueAtIndex]=[index]
        else:
            valueIndexes.append(index)
            if len(valueIndexes) == 2:
                return valueAtIndex
    return -1

print(firstDuplicate([2, 1, 3, 5, 3, 2])) # 3