// Getting the sum of numbers of an array
// in JavaScript

let arr = [1, 22, 3, 44, 5, 66, 7]

// Using the reduce function
let sumShort = arr.reduce((prevVal, curVal, index, array)=>
prevVal + array[index]);

console.log(sumShort)

// Using for loop
let sum = 0;
for (index in arr) {
    sum += arr[index]
}

console.log(sum)