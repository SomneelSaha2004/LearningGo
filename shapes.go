package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
	String() string
}
type Rectangle struct {
	l, b float64
}

func (r Rectangle) area() float64 {
	return r.l * r.b
}
func (r Rectangle) perimeter() float64 {
	return 2 * (r.l + r.b)
}
func (r Rectangle) diagonal() float64 {
	return math.Sqrt(r.l*r.l + r.b*r.b)
}
func (r Rectangle) String() string {
	return fmt.Sprintf("length : %f\nbreadth : %f\nArea : %f", r.l, r.b, r.area())
}

type Circle struct {
	r float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}
func (c Circle) String() string {
	return fmt.Sprintf("Radius : %f\nArea : %f", c.r, c.area())
}
func getArea(s Shape) float64 {
	return s.area()
}
func getPerimeter(s Shape) float64 {
	return s.perimeter()
}
func main() {
	arr := []Shape{Rectangle{5, 6}, Circle{4}, Circle{14}}
	for _, animal := range arr {
		if typeOfShape, ok := animal.(Rectangle); ok {
			fmt.Println(typeOfShape)
		}

	}
}
