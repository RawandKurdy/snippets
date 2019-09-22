// First Duplicate Value - Google - Interview Question -
// in JavaScript

function firstDuplicate(array) {
    const arrLength = array.length;
    if (arrLength === 1) return -1;
    if (arrLength === 0) return -1;

    let valuesMap = new Map();

    for (let index = 0; index < arrLength; index += 1) {
        let valueAtIndex = array[index]
        let valueIndexes = valuesMap.get(valueAtIndex);

        if (valueIndexes == undefined || !valueIndexes) {
            valuesMap.set(valueAtIndex, [index]);
        } else {
            valueIndexes.push(index);
            if (valueIndexes.length == 2)
                return valueAtIndex;
        };
    }
    return -1;
}

console.log(firstDuplicate([2, 1, 3, 5, 3, 2])); // 3