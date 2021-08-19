package main

import "fmt"

func main() {

	//if
	x := 323
	y := 38

	if x < y {
		fmt.Printf("%d is greater than %d \n", x, y)
	} else {
		fmt.Printf("%d is greater than %d \n", y, x)

	}
	// else condition

	// if-elseif-else condition
	color := "black"
	if color == "red" {
		fmt.Println("this is red color")
	} else if color == "blue" {
		fmt.Println("this color is blue")
	} else {
		fmt.Println("this color is neither red nor blue")
	}
	// switch
	switch color {
	case "red":
		fmt.Println("this is red color")
	case "blue":
		fmt.Println("this is blur color")
	default:
		fmt.Println("this color is neither red nor blue")

	}

}
