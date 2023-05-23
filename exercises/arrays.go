//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type product struct {
	name  string
	price float64
}

func sum(arr [4]product) float64 {
	sum := 0.0
	for i := 0; i < len(arr); i++ {
		sum += arr[i].price
	}
	return sum
}
func main() {
	var arr [4]product
	for i := 0; i < len(arr); i++ {
		fmt.Print("Enter name : ")
		fmt.Scan(&arr[i].name)
		fmt.Print("Enter price : ")
		fmt.Scan(&arr[i].price)
		fmt.Println()
	}
	fmt.Println(arr[len(arr)-1])
	fmt.Println(len(arr))
	fmt.Println(sum(arr))
}
