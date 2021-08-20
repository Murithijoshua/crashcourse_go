package main

import "fmt"

func main() {
	a := 12
	b := &a
	// If you use asterisk infront of a variable declaring to pointer, return the actual value for the referred value
	// else it will point out the  memory address.
	fmt.Println(a, *b)

	//changing valu using pointer
	*b = 45
	fmt.Println(a)
	// this mean using pointer you can easily change the previous referred variable content
}
