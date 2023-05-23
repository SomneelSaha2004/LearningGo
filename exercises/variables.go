//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var color string = "purple"
	birthYear, ageInYears := 1945, 13
	var (
		fi = "S"
		fl = "L"
	)
	var ageInDays int
	ageInDays = 365 * ageInYears
	fmt.Println("favourite color : ", color, "\nBirth Year : ", birthYear, "\nAge in years : ", ageInYears, "\n Initials ", fi, fl, "\nAge In Days : ", ageInDays)

}
