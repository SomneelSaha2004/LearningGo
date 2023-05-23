//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part struct {
	name  string
	price float64
}

func printAssembly(a []Part) {
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i].name, " ")
	}
	fmt.Println()
}
func main() {
	Line := []Part{{"screws", 23}, {"metal", 45}, {"summin", 244}}
	printAssembly(Line)
	Line = append(Line, Part{"eggs", 121}, Part{"yeet", 34})
	printAssembly(Line)
	Line = Line[len(Line)-2:]
	printAssembly(Line)
}
