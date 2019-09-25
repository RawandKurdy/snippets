package main

import "fmt"

// Rotate the n x n 2D image matrix by 90 degrees (clockwise)
// Asked By Apple, Microsoft, Amazon
// in Golang
func rotateImage(array [][]int) [][]int {
	n := len(array) // n x n Matrix so ROW == COLUMN

	var rotatedImg = make([][]int, n)

	for row := 0; row < n; row++ {
		// Init. the sub Array for each cell
		rotatedImg[row] = make([]int, n)

		for col := 0; col < n; col++ {
			// Enter the value to their new corresponding indices
			oldCol := row         // old Col to grab from
			oldRow := n - 1 - col // old Row to grab from
			rotatedImg[row][col] = array[oldRow][oldCol]
		}
	}
	return rotatedImg
}

func main() {
	matrix := [][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}

	fmt.Println(rotateImage(matrix))
	/* Output
	   [[7, 4, 1],
	    [8, 5, 2],
	    [9, 6, 3]]
	*/
}
