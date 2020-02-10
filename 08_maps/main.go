package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign key-value
	emails["Pratik"] = "pratik@gmail.com"
	emails["Mayank"] = "mayank@gmail.com"
	emails["Chinmay"] = "chinmay@gmail.com"

	// Define and Assign in single step
	mails := map[string]string{"Bob": "bob@gmail.com", "Mike": "mike@gmail.com"}

	fmt.Println(mails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Mayank")
	fmt.Println(emails)
}
