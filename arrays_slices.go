package main

import "fmt"

func main() {
	// arrays
	// var fruitArr [2]string
	// 1. method 1
	// fruitArr[0] = "mangos"
	// fruitArr[1] = "Oranges"
	// 2. method 2
	fruitArr := [2]string{"strawberry", "mangoes"}

	fmt.Println(fruitArr)
	// +++++++++++++++++++
	// slicing
	// +++++++++++++++++++++

	fruitSlice := []string{"james", "ken", "mike", "dan"}

	fmt.Println(len(fruitSlice))
}
