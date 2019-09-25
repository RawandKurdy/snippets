// Rotate the n x n 2D image matrix by 90 degrees (clockwise)
// Asked By Apple, Microsoft, Amazon
// in JavaScript
function rotateImage(array) {
	let n = array.length // n x n Matrix so ROW == COLUMN
	var rotatedImg = []

	for (let row = 0; row < n; row++) {
		// Init. the sub Array for each cell
		rotatedImg[row] = []

		for (let col = 0; col < n; col++) {
			// Enter the value to their new corresponding indices
			let oldCol = row         // old Col to grab from
			let oldRow = n - 1 - col // old Row to grab from
			rotatedImg[row][col] = array[oldRow][oldCol]
		}
	}
	return rotatedImg
}

matrix = [[1, 2, 3],
     [4, 5, 6],
     [7, 8, 9]]

console.log(rotateImage(matrix))
/* Output
     [[7, 4, 1],
      [8, 5, 2],
      [9, 6, 3]] */