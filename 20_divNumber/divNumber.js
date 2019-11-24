// DivNumber Challenge
// Given three integers k, l and r
// return the number of integers between l and r inclusive
// that have exactly k divisors.
// In Javascript

function divNumber(k, l, r) {
    // di is the number of Numbers
    // that has 'k' exact divisors
    di = 0
    while (l <= r) {
        d = 0 //d is number of divisors of individual numbers
        c = 1  // c is counter(Number)

        while (c <= l) {
            d += l % c == 0 ? 1 : 0
            c++
        }
        di += d == k ? 1 : 0
        l++
    }
    return di
}

// Example
k = 3 
l = 2
r = 49
console.log(divNumber(k, l, r)) // 4

// We have 4 integers in that range
// that have exactly 3 divisors - 4, 9, 25, 49.