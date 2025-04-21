package main

import "fmt"

func main() {
	var myName, age = "Olayinka", 45
	fmt.Println("My name is: ", myName, ", I am ", age, " years old. Thank you.")

	var name string = "Toluwalope"
	fmt.Println("His wife's name is: ", name)

	userName := "YinkaVitula"
	fmt.Println("User name is ", userName)

	var sum int
	fmt.Println("The sum is: ", sum)

	// number1, otherInfo := 50, "random info"
	// fmt.Println("The first number is: ", number1, "and otherinfo is: ", otherInfo)

	// number2, otherInfo := 100, "some other info***"
	// fmt.Println("The second number is: ", number2, "and otherinfo is: ", otherInfo)
	var (
		number1 = 500
		number2 = 600
	)

	sum = number1 + number2
	fmt.Println("The sum of the number is: ", sum)
}
