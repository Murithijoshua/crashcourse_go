package main

import "fmt"

func main() {
	// this is long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}
	//short method javascript style
	for i := 1; i <= 10; i++ {
		fmt.Printf("loop number %d\n", i)

		// ********************

		// loop number 1
		// loop number 2
		// loop number 3
		// loop number 4
		// loop number 5
		// loop number 6
		// loop number 7
		// loop number 8
		// loop number 9
		// loop number 10
		// ********************
	}
	// fizzbuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("fizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Buzz")
		} else if i%5 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}
