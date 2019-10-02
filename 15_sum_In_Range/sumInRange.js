// Sum In Range - Interview Question - Asked by Palantir Technologies
// In Javascript
// In O(n) time
function sumInRange(nums, queries) {
	let totalSum = 0
	let prefixSums = {}
	let modulo = 1000000000 + 7

	// calculates the prefix-sums
	for (let index in nums) {
		if (index == 0) {
			prefixSums[index] = nums[index] % modulo
		} else {
			prefixSums[index] = (prefixSums[index-1] + nums[index]) % modulo
		}
	}

	// Now adds up each separate range
	for (let query of queries) {
		let fromIndex = query[0]
		let toIndex = query[1]
		totalSum += prefixSums[toIndex]

		if (fromIndex != 0) {
			totalSum -= prefixSums[fromIndex-1]
		}
		totalSum %= modulo
	}
	// if totalSum is negative we just add the modulo
	if (totalSum < 0) {
		totalSum += modulo
	}
	return totalSum
}

// Example
let nums =  [34, 19, 21, 5, 1, 10, 26, 46, 33, 10]
let queries = [[3,7], 
 [3,4], 
 [3,7], 
 [4,5], 
 [0,5]]

console.log(sumInRange(nums, queries)) // 283

