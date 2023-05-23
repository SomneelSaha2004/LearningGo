package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// pointer reciever function modifies
func (this *Point) shiftBy(x, y float64) {
	this.x += x
	this.y += y
}

// reciever function value
// does not modify object itself
func (this Point) shiftDist(x, y float64) Point {
	return Point{this.x + x, this.y + y}
}
func (p Point) String() string {
	return fmt.Sprintf("(%f,%f)", p.x, p.y)
}
func (this *Point) distanceFrom(other Point) float64 {
	return math.Sqrt(math.Pow((this.x-other.x), 2) + math.Pow((this.y-other.y), 2))
}
func main() {
	p1 := Point{3, 4}
	p2 := Point{4, 5}
	fmt.Println(p1.distanceFrom(p2))
	p2.shiftBy(1, 2) // p2->{5,7}
	fmt.Println(p2)
	p3 := p1.shiftDist(3, 4)
	fmt.Println(p1)
	fmt.Println(p3)
}

//Pointer recievers can modify structs
// value recievers cannot
