package main

import "fmt"

func main() {
	brodiamyzalius@gmail.coms := make(map[string]string)
	// assiging key value pairs method 1

	brodiamyzalius@gmail.coms["bob"] = "bob@yahoo.com"
	brodiamyzalius@gmail.coms["brian"] = "brian@yahoo.com"
	brodiamyzalius@gmail.coms["mary"] = "mary@yahoo.com"
	fmt.Println(brodiamyzalius@gmail.coms)
	fmt.Println(len(brodiamyzalius@gmail.coms))
	fmt.Println(brodiamyzalius@gmail.coms["mary"])

	// deleting map object

	delete(brodiamyzalius@gmail.coms, "Bob")
	fmt.Println(brodiamyzalius@gmail.coms)

	// assinging key valuepairs

	users := map[string]string{"joshua": "joshua@blob.com", "sharon": "sharoo@gmail.com", "jameson": "jameson@yahoo.com"}
	fmt.Println(users)
}
