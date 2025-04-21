// Summary:
//
//	Print basic information to the terminal using various variable
//	creation techniques. The information may be printed using any
//	formatting you like.
//
// Notes:
//   - Use fmt.Println() to print out information
//   - Basic math operations are:
//     Subtraction    -
//     Addition       +
//     Multiplication *
//     Division       /
package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var myFavColor = "blue"
	fmt.Println("My Favourite Color is: ", myFavColor)

	//* Store your birth year and age (in years) in two variables using compound assignment
	birthYear, age := 1981, 45
	fmt.Println("I was born in the year: ", birthYear, " and I am ", age, " years old.")

	//* Store your first & last initials in two variables using block assignment
	var (
		firstName    = "Olayinka"
		lastName     = "Korede"
		firstInitial = string(firstName[0])
		lastInitial  = string(lastName[0])
	)
	fmt.Println("My name is", firstName, "and my initials are ", firstInitial, ".", lastInitial, ".")

	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	var ageInDays int
	ageInDays = 365 * 45
	fmt.Println("My age in days is: ", ageInDays)
}
