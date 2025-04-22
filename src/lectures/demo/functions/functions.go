package main

import "fmt"

func double(number int) int {
	return number + number
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from Greet Function!")
}

func main() {
	greet()
	dozen := double(6)
	fmt.Println("A Dozen is: ", dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println("A baker's dozen is: ", bakersDozen)

	anotherBakersDozen := add(double(6), 1)
	fmt.Println("Another baker's dozen is: ", anotherBakersDozen)
}
