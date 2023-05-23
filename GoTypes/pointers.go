package main

import "fmt"

//function calls in go are pass by value
// can be slow for large data structures
type Counter struct {
	count int
}

func inc(i *int) {
	*i++
}
func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}
func increment(counter *Counter) {
	counter.count++
	fmt.Println(counter)
}
func main() {
	var intptr *int
	s := 23
	intptr = &s
	fmt.Println(intptr)
	fmt.Println(*intptr) // dereferencing
	inc(intptr)
	fmt.Println(s)
	str := "Old"
	c := Counter{}
	replace(&str, "New", &c)
	fmt.Println(str)
}
