# isCryptSolution

A cryptarithm is a mathematical puzzle for which the goal is to find the correspondence between letters and digits, such that the given arithmetic equation consisting of letters holds true when the letters are converted to digits.</br>

You have an array of strings crypt, the cryptarithm, and an an array containing the mapping of letters and digits, solution. The array crypt will contain three non-empty strings that follow the structure: [word1, word2, word3], which should be interpreted as the word1 + word2 = word3 cryptarithm.</br>

If crypt, when it is decoded by replacing all of the letters in the cryptarithm with digits using the mapping in solution, becomes a valid arithmetic equation containing no numbers with leading zeroes, the answer is true. If it does not become a valid arithmetic solution, the answer is false.</br>

Note that number 0 doesn't contain leading zeroes (while for example 00 or 0123 do).</br>

### Example

For crypt = ["SEND", "MORE", "MONEY"] and</br>

solution = [['O', '0'],</br>
            ['M', '1'],</br>
            ['Y', '2'],</br>
            ['E', '5'],</br>
            ['N', '6'],</br>
            ['D', '7'],</br>
            ['R', '8'],</br>
            ['S', '9']]</br>
the output should be</br>
isCryptSolution(crypt, solution) = true.</br>

When you decrypt "SEND", "MORE", and "MONEY" using the mapping given in crypt, you get 9567 + 1085 = 10652 which is correct and a valid arithmetic equation.</br>

For crypt = ["TEN", "TWO", "ONE"] and</br>

solution = [['O', '1'],</br>
            ['T', '0'],</br>
            ['W', '9'],</br>
            ['E', '5'],</br>
            ['N', '4']]</br>
the output should be</br>
isCryptSolution(crypt, solution) = false.</br>

Even though 054 + 091 = 145, 054 and 091 both contain leading zeroes, meaning that this is not a valid solution.</br>

[Source](https://app.codesignal.com/interview-practice/task/yM4uWYeQTHzYewW9H)

## Passed All tests on Codesignal