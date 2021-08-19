package main

import "fmt"

func greeting(name string) string {
	return "hello" + name

}
func sum(number1 int, number2 int) int {
	return number1 + number2

}
func main() {
	fmt.Println(greeting("james"))
	fmt.Println(sum(23, 53))

}
