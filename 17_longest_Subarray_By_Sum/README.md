# Find Longest Subarray By Sum
You have an unsorted array arr of non-negative integers and a number s. Find a longest contiguous subarray in arr that has a sum equal to s. Return two integers that represent its inclusive bounds. If there are several possible answers, return the one with the smallest left bound. If there are no answers, return [-1]. </br>

Your answer should be 1-based, meaning that the first position of the array is 1 instead of 0. </br>

### Example

For s = 12 and arr = [1, 2, 3, 7, 5], the output should be
findLongestSubarrayBySum(s, arr) = [2, 4]. </br>

The sum of elements from the 2nd position to the 4th position (1-based) is equal to 12: 2 + 3 + 7. </br>

For s = 15 and arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], the output should be
findLongestSubarrayBySum(s, arr) = [1, 5]. </br>

The sum of elements from the 1st position to the 5th position (1-based) is equal to 15: 1 + 2 + 3 + 4 + 5. </br>

For s = 15 and arr = [1, 2, 3, 4, 5, 0, 0, 0, 6, 7, 8, 9, 10], the output should be
findLongestSubarrayBySum(s, arr) = [1, 8]. </br>

The sum of elements from the 1st position to the 8th position (1-based) is equal to 15: 1 + 2 + 3 + 4 + 5 + 0 + 0 + 0. </br>

[Source](https://app.codesignal.com/interview-practice/task/izLStwkDr5sMS9CEm)

## Passed All tests on Codesignal