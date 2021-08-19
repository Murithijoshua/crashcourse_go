package main

//  #types of variable in go

import "fmt"

func main() {
	// string datatype variable
	var name = "Jamesons'"
	var age = 37
	var isCool = true
	// shorthard for assignment which is only workable  inside a function
	names := "Jameson Johns"
	fmt.Println(name, names, age, isCool)
	fmt.Printf("%T\n", name)
}
