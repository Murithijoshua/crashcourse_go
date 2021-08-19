package main

import "fmt"

func main() {
	emails := make(map[string]string)
	// assiging key value pairs method 1

	emails["bob"] = "bob@yahoo.com"
	emails["brian"] = "brian@yahoo.com"
	emails["mary"] = "mary@yahoo.com"
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["mary"])

	// deleting map object

	delete(emails, "Bob")
	fmt.Println(emails)

	// assinging key valuepairs

	users := map[string]string{"joshua": "joshua@blob.com", "sharon": "sharoo@gmail.com", "jameson": "jameson@yahoo.com"}
	fmt.Println(users)
}
