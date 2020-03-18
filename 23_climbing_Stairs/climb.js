
// Climbing Stairs - Interview Question -
// Asked By Apple and Adobe
// in JavaScript

function climbingStairs(n) {
    let r = [1, 1, 2] // Initial Result Set
    if (n <= 2) return r[n];

    for (let step = 3; step < n + 1; step++) {
        r.push(r[step - 1] + r[step - 2])
    }
    return r[n]
}

let n = 26; // Steps
let ways = climbingStairs(n);
console.log(ways); // 196418
