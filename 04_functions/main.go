package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(greeting("Pratik"), getSum(1, 2))
}
