package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// float32 float64

	// Using var
	var name = "Pratik"
	var age = 27
	const isCool = true

	// Shorthand
	name2 := "Mayank"
	size := 1.33

	// Shorthand-2
	name, email := "PRL", "prl@gmail.com"

	fmt.Println(name, age, isCool, name2, size, email)
	fmt.Printf("%T\n", size)
}
