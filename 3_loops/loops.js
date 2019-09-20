// Loops in Javascript

let arr = ["New York", "Washington", "California"]

// Foreach loop
for (key in arr) {
    console.log(arr[key])
}

// Traditional for loop
for (let x = 0; x < arr.length; x +=1) {
    console.log(arr[x])  
}

// While loop 
// (continuous till condition is true)
let condition = true
while (condition) {
    let city = arr.pop()
    console.log(city)
    condition = !(arr.length == 0)   
}
