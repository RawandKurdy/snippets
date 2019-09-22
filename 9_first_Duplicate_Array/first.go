package main

import "fmt"

// First Duplicate Value - Google - Interview Question -
// in Golang
func firstDuplicate(array []int) int {
	arrLength := len(array)
	if arrLength == 0 || arrLength == 1 {
		return -1
	}

	var valuesMap = map[int][]int{}

	for index := 0; index < arrLength; index++ {
		valueAtIndex := array[index]
		valueIndexes, mapRefExist := valuesMap[valueAtIndex]

		if !mapRefExist {
			valuesMap[valueAtIndex] = []int{index}
		} else {
			valueIndexes = append(valueIndexes, index)
			if len(valueIndexes) == 2 {
				return valueAtIndex
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(firstDuplicate([]int{2, 1, 3, 5, 3, 2})) //3
}
