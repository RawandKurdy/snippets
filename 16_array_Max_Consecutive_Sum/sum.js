// Array Max Consecutive Sum - Interview Question -
// Asked by Microsoft, LinkedIn, Amazon and Samsung
// Maximum possible sum you can get from one of contiguous subarrays.
// In Javascript
function arrayMaxConsecutiveSum2(inputArray) {
    let sumMax = 0;
    let sumSteps = 0;
    let maxNumber = 0;

    for (let number of inputArray) {
        sumSteps += number;

        if (maxNumber == 0 || maxNumber < number)
            maxNumber = number;

        if (sumSteps < 0)
            sumSteps = 0;
        if (sumMax < sumSteps)
            sumMax = sumSteps;
    }
    return sumMax != 0 ? sumMax : maxNumber;
}

// Example
arr = [1, -2, 3, -4, 5, -3, 2, 2, 2];
console.log(arrayMaxConsecutiveSum2(arr)); // 8