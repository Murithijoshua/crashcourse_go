package main

import (
	"fmt"
	"strconv"
)

// define struct, for instance Person

type Person struct {
	firstname string
	lastname  string
	city      string
	gender    string
	age       int
}

//methods
// there are two types of methods
// 1.value recievers---> they dont change anything
// 2. pointer recievers--> changing something
// example greeting method ==> Value reciever

func (p Person) greeting() string {
	return "hi my name is" + " " + p.firstname + " " + p.lastname + " " + "and am from " + p.city + " " + "aged" + " " + strconv.Itoa(p.age)
}

// this is a pointer reciever that will be used to change the  struct above (using asterisk is good for this kind of method)
func (p *Person) aged() {
	p.age++
}

func main() {
	// intialize your struct using property name
	person1 := Person{firstname: "james", lastname: "Walker", city: "Nairobi", gender: "Male", age: 34}
	fmt.Println(person1)
	//alternative
	person2 := Person{"Dennis", "Smith", "Nakuru", " male", 23}
	fmt.Println(person2)

	// calling the pointer reciever in order to increase their age a couple of times before calling it
	person2.aged()
	person2.aged()
	person2.aged()
	person2.aged()
	person2.aged()
	person2.aged()
	person2.aged()
	person2.aged()
	// calling the method greeting above(value reciever)
	fmt.Println(person2.greeting())

}
