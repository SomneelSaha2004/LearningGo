package main

import "fmt"

//variadics allow us to let functions take an uknown number of inputs
func sum(nums ...int) int {
	// nums is actually a slice
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

// ... operator expands a slice into its individual values
func main() {
	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 5}
	c := append(a, b...)   // ... expanding b into 3,4,5
	fmt.Println(sum(c...)) //same here ^
}
