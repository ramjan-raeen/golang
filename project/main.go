package main

// main.go
// This is a simple Go program that prints a greeting message and some personal information to the terminal.

import (
	"fmt"
	// "net/http"
)

/* 
The main function is the entry point of the program.
It initializes some variables and prints a message to the terminal.
The myAge function returns an integer representing the age.
The program demonstrates variable declaration, reassignment, and function usage in Go.
*/

func main() {
	fmt.Println("Hello World...") // This prints in terminal on start
	var myname string = "Ramjan" // Assigning value to myname
	lastname := "Rahman"   // Assigning value to lastname using short variable declaration
	lastname = "Raeen" // Reassigning value to lastname
	myage := myAge() // Calling myAge function to get age
	fmt.Println("My full name is ", myname, lastname, "and I am", myage, "years old.") // Prints my name in terminal

    favNumbers := []int{44, 444, 444, favoriteNum()}    // Initializing a slice of integers with some favorite numbers
    favNumbers = append(favNumbers, 18) // Appending another favorite number to the slice
    favNumbers = append(favNumbers, 4444, 44444) // Appending more favorite numbers to the slice
    fmt.Println("My favorite numbers are:", favNumbers) // Prints the favorite numbers in terminal
}

/* 
myAge is a function that returns an integer representing the age.
It is used to demonstrate how to define and call a function in Go.
In this case, it returns a hardcoded value of 25.
*/
func myAge() int {
	return 25
}

// favoriteNum is a function that returns an integer representing a favorite number.
func favoriteNum() int {
    return 4
}