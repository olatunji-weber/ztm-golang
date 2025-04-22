// --Summary:
//
//	Use functions to perform basic operations and print some
//	information to the terminal
package main

import (
	"fmt"
)

// --Requirements:
// * Write a function that accepts a person's name as a function parameter and displays a greeting to that person.
func greeting(name string) {
	fmt.Println("Hello there,", name, "it's nice to meet you.")
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func messageEcho(message string) string {
	return message
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func addNumbers(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

// - Write a function that returns any number
func anyNumber(num int) int {
	return num
}

// - Write a function that returns any two numbers
// - Add three numbers together using any combination of the existing functions.
// - Print the result
func main() {
	//* Call every function at least once
	var username string
	fmt.Println("What is your name? ")
	fmt.Scanln(&username)
	greeting(username)

	myMessage := "This is my message to you"
	fmt.Println(messageEcho(myMessage))

	a, b, c := 500, 100, 400
	fmt.Println("The sum of the 3 numbers is: ", addNumbers(a, b, c))

	fmt.Println("The number to be returned is: ", anyNumber(20))
}
