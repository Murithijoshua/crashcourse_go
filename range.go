package main

import "fmt"

func main() {
	ids := []int{12, 13, 15, 17, 3, 0, 5, 7, 3, 6, 8}
	// loop using range taking account index
	for i, id := range ids {
		fmt.Printf("%d and number %d\n", i, id)

	}
	for _, id := range ids {
		fmt.Printf("number %d\n", id)

	}
	// key value pairs
	users := map[string]string{"joshua": "joshua@blob.com", "sharon": "sharoo@gmail.com", "jameson": "jameson@yahoo.com"}

	for k, v := range users {
		fmt.Printf("%s: %s\n", k, v)
	}

}
