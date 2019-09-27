package main

import "fmt"

// sudoku2 Interview Question Asked by Apple & Uber
// in Golang
func sudoku2(grid [][]string) bool {
	var trackColNumbers = map[int]map[string]int{} // Numbers in each column list
	var trackRowNumbers = map[int]map[string]int{} // Numbers in each row list
	var currentSubMatrix = map[string]int{}        // Numbers in Each Sub Matrix (3x3)
	var subSize = int(len(grid) / 3)               // Sub group sizes, here its 9/3 = 3

	colAdd := 0 // Added to the sub matrix col to get the real column
	rowAdd := 0 // Added to the sub matrix row to get the real row

	for rowAdd != 9 {
		for subCol := 0; subCol < subSize; subCol++ {
			realCol := subCol + colAdd // real column in the grid
			if trackColNumbers[realCol] == nil {
				trackColNumbers[realCol] = map[string]int{}
			}
			for subRow := 0; subRow < subSize; subRow++ {
				realRow := subRow + rowAdd // real row in the grid
				if trackRowNumbers[realRow] == nil {
					trackRowNumbers[realRow] = map[string]int{}
				}
				item := grid[realRow][realCol] // each item in the cells

				if item != "." {
					if _, ok := currentSubMatrix[item]; !ok {
						currentSubMatrix[item] = 0
					} else {
						return false
					}

					if _, ok := trackRowNumbers[realRow][item]; !ok {
						trackRowNumbers[realRow][item] = 0
					} else {
						return false
					}

					if _, ok := trackColNumbers[realCol][item]; !ok {
						trackColNumbers[realCol][item] = 0
					} else {
						return false
					}
				}
			}
		}
		currentSubMatrix = map[string]int{} // Reset Sub Matrix each time
		colAdd += 3                         // moves to the next 3 columns
		if colAdd == 9 {
			colAdd = 0  // reset back columns to start
			rowAdd += 3 // moves to the next 3 rows
		}
	}
	return true
}

func main() {
	grid := [][]string{{".", ".", ".", ".", "2", ".", ".", "9", "."},
		{".", ".", ".", ".", "6", ".", ".", ".", "."},
		{"7", "1", ".", ".", "7", "5", ".", ".", "."},
		{".", "7", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", "8", "3", ".", ".", "."},
		{".", ".", "8", ".", ".", "7", ".", "6", "."},
		{".", ".", ".", ".", ".", "2", ".", ".", "."},
		{".", "1", ".", "2", ".", ".", ".", ".", "."},
		{".", "2", ".", ".", "3", ".", ".", ".", "."}}

	fmt.Println("is Grid Valid?", sudoku2(grid)) // answer :> False
}
