// Longest Subarray By Sum in an Array - Interview Question -
// Asked by Facebook & Palantir Technologies
// In Javascript
function findLongestSubarrayBySum(sum, arr) {
    let prefixSumsToIndex = {};
    let fromIndex = -1;
    let toIndex = -1;
    let sumExistAt = -1;
    
    // Calculates the prefixSums
    let previousSum = 0;
    for (let index = 0; index < arr.length; index++){
        let number = arr[index];
        let newSum = previousSum + number;
        previousSum = newSum;
        prefixSumsToIndex[newSum] = index + 1;
        if (number == sum)
            sumExistAt = index + 1;
    }

    // Sees if such sum exists within the array
    // initializes the indices with initial range if exists
    if (prefixSumsToIndex[sum] != undefined){
        fromIndex = 1;
        toIndex = prefixSumsToIndex[sum];
    }

    // Searches the array for a wider range if available
    for (let numSum in prefixSumsToIndex){
        let indexofSum = prefixSumsToIndex[parseInt(numSum)];
        let indexOfSumWithS = prefixSumsToIndex[parseInt(numSum) + sum];
        if (indexOfSumWithS != undefined){
            let tmpFromIndex = indexofSum + 1;
            let tmpToIndex = indexOfSumWithS;
            if ((toIndex - fromIndex) < (tmpToIndex - tmpFromIndex)){
                toIndex = tmpToIndex;
                fromIndex = tmpFromIndex;
            }
        }
    }

    // Returns depending on if indices has been found or not
    if (fromIndex != -1 && toIndex != -1)
        return [fromIndex, toIndex];
    else if (sumExistAt != -1)
        return [sumExistAt, sumExistAt];
    else return [-1];
}

// Example
let sum = 15;
let arr = [6, 7, 8, 1, 2, 3, 0 , 4, 5, 0, 0, 9, 10];
console.log(findLongestSubarrayBySum(sum, arr)) // [4, 11]