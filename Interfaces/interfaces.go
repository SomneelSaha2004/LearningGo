package main

import "fmt"

type Integral interface {
	double(x int) int
	triple(x int) int
}

type MyInt int

func (m MyInt) double(x int) int {
	return 2 * x
}

// we should always use values not pointers
func (m MyInt) triple(x int) int {
	return 3 * x
}
func execute(i Integral) int {
	return i.triple(3)
}
func main() {
	i := MyInt(2)
	fmt.Println(i.double(3))
	m := MyInt(2)
	fmt.Println(execute(m))
}
