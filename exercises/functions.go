// --Summary:
//
//	Use functions to perform basic operations and print some
//	information to the terminal.
//
// --Requirements:
//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
//   - Write a function that returns any message, and call it from within
//     fmt.Println()
//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
//   - Write a function that returns any number
//   - Write a function that returns any two numbers
//   - Add three numbers together using any combination of the existing functions.
//   - Print the result
//   - Call every function at least once
package main

import "fmt"

func greet(n string) {
	fmt.Println("Hello", n)
}
func getMessage() string   { return "Hello World" }
func sum3(a, b, c int) int { return a + b + c }
func getNum() int          { return sum3(1, 2, 3) }
func get2Num() (int, int)  { return getNum(), getNum() }
func main() {
	fmt.Println(getMessage())
	a, b := get2Num()
	fmt.Println(a, b)
}
