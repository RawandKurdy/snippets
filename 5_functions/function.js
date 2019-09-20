// Functions in JavaScript

// 1- Traditional function
// also used to demonstrate how an anonymous func can be useful
// It runs a function passed as a param
function traditionalFunc(anotherFunc){
    anotherFunc()
}

// 2- Anonymous Function
// I passed it as a parameter
traditionalFunc(()=> {
    console.log("Hello World!")
})

// 3- Function Expression
var anotherFunc = function (){
    console.log("Hello Again!")
}

// Exec function expr.
anotherFunc()