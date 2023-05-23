package main

import "fmt"

// fmt - format

/*
Printf - custom format using %f etc
Print - just print
Println - simple print with a newline

We have F and S variants also :
 F prints to data stream Fprintf etc
 S prints to a new string Sprintf
*/
func pad(s string, lhs, rhs rune) string {
	return fmt.Sprintf("%c%v%c", lhs, s, rhs)
}
func main() {
	//printing runes
	r := '😈'
	fmt.Printf("%c\n", r)
	//to print hexcode of a number
	fmt.Printf("%X\n", r)
	//unicode
	fmt.Printf("%U\n", r)
	//printing bool
	fmt.Println(true)
	fmt.Printf("%t\n", true)
	fmt.Println(pad("hello", '🐲', 'm'))
}
