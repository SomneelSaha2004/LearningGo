package main

import "fmt"

func sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
func intersection(a []int, b []int) []int {
	result := []int{}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				result = append(result, a[i])

			}
		}
	}
	fmt.Println(result)
	return result
}

func main() {
	arr := [...]int{1, 2, 3, 4, 4, 3}
	fmt.Println(arr[0:4])
	fmt.Println(sum(arr[:]))
	s := arr[:] //entire array as a slice
	fmt.Println(s)
	s = arr[1:]
	fmt.Println(s)
	fmt.Println(arr[:2])
	// slices are dynamic
	arr2 := [...]int{2, 3, 12, 1}
	i := intersection(arr[:], arr2[:])
	fmt.Println(append(i, arr[:]...))
	//to make empty slices
	slice := make([]bool, 3)
	fmt.Println(slice, arr[0:3])
}
