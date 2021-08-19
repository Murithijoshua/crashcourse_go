package main

//important to know that you need to import first then a must to use imported package,else it will throw error
import (
	"fmt"
	"math"

	"./strutils"
)

func main() {
	fmt.Println(math.Floor(2.7) * 2)
	fmt.Println(math.Ceil(2.7) * 2)
	fmt.Println(math.Sqrt(2.7) * 2)
	fmt.Println(strutils.Reverse("olleh"))

}
