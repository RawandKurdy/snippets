package main

import "fmt"

// Using Maps in Golang
func main() {
	// Create the Map
	var mapDic = map[string]string{}

	// Set a key-value pair
	mapDic["01"] = "New York"
	fmt.Println(mapDic)

	// Get a value through its key
	value := mapDic["01"]
	fmt.Println(value)

	// Clear the whole map
	mapDic = make(map[string]string)
	fmt.Println(mapDic)
}
