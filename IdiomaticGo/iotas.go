package main

import (
	"fmt"
)

type Direction byte

func (d Direction) String() string {
	return []string{"North", "South", "West", "East"}[d]
}
func main() {

	//iota helps us make constants
	const (
		zero = iota
		one
		two
		three
		_ //can slip values with _
		_
		six
		//iota starts from zero
		//from then onwards next value is +1
	)
	//iota can start anywhere
	fmt.Println(zero, two, six)
	//iota can start anywhere
	const (
		a = iota + (22.0 / 7)
		b
		c
	)
	fmt.Println(a, b, c)

	const (
		north Direction = iota
		south
		west
		east
	)
	n := north
	fmt.Println(n)
}
