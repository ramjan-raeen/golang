package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start line of the program...") // This prints in terminal on start
	var firstname string = "Ramjan" // Assigning value to firstname
	lastname := "Rahman"   // Assigning value to lastname using short variable declaration
	lastname = "Raeen" // Reassigning value to lastname
	num := 444 // Assigning value to num of dynamic type
	const title = "Mr." // Assigning value to title of constant type
	//title = "Md"	// Constants cannot be reassigned after their declaration. It is immutable.

	fmt.Println("My name is", title,firstname, lastname, "and my favorite number is",num) // Prints my name in terminal
}