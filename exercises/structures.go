//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y float64
}
type Rectangle struct {
	bottomRight, topLeft Coordinate
}

func getWidth(r Rectangle) float64 {
	return r.bottomRight.y - r.topLeft.y
}
func getLength(r Rectangle) float64 {
	return r.bottomRight.x - r.topLeft.x
}
func main() {
	r := Rectangle{Coordinate{2.3, 4.5}, Coordinate{}}
	fmt.Println("Area : ", getLength(r)*getWidth(r))
	fmt.Println(getLength(r), getWidth(r))
	fmt.Println("Perimeter : ", 2*(getLength(r)+getWidth(r)))
	r.topLeft.x *= 2
	r.topLeft.y *= 2
	fmt.Println("Area : ", getLength(r)*getWidth(r))
	fmt.Println("Perimeter : ", 2*(getLength(r)+getWidth(r)))
	fmt.Println(getLength(r), getWidth(r))
}
