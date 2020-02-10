package main

import (
	"fmt"
	"strconv"
)

// Define Person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	// //Alternative
	// firstName lastName city gender string
	// age int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Pratik", lastName: "Lodha", city: "Nashik", gender: "F", age: 27}

	// // Alternative
	// person2 := Person{"Pratik", "Lodha", "Nashik", "M", 27}

	// fmt.Println(person1.firstName)
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Lodha")
	fmt.Println(person1.greet())
}
