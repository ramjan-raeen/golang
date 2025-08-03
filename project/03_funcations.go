package main

import (
	"fmt"
)

func add_num(a int, b int) int {
	return a + b
}

func concat_strings(firstname string, lastname string) string {
	return firstname + " " + lastname
}

func main() {
	fmt.Println("The sum of 5 and 10 is:", add_num(5, 10)) // This prints the sum of two numbers
	fmt.Println("Concatenated string is:", concat_strings("Ramjan", "Raeen")) // This prints the concatenated string
}